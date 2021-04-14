/*

	Copyright 2021, Pure Storage Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/

package cbs

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/managedapplications"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type cloudformationAPI interface {
	CreateStack(input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)
	DescribeStacks(input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)
	DeleteStack(input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)
	WaitUntilStackCreateCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput, opts ...request.WaiterOption) error
	WaitUntilStackDeleteCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput, opts ...request.WaiterOption) error
}

type CbsService struct {
	CloudFormation cloudformationAPI
	AzureClient    AzureClientAPI
	awsRegionStr   string
	azureConfig    azureUserConfig
}

func (m *CbsService) CloudFormationService() (cloudformationAPI, diag.Diagnostics) {
	if m.CloudFormation == nil {
		cftSvc, diags := buildAWSSession(m.awsRegionStr)
		if diags.HasError() {
			return nil, diags
		}
		m.CloudFormation = cftSvc
	}

	return m.CloudFormation, nil
}

func (m *CbsService) AzureClientService() (AzureClientAPI, diag.Diagnostics) {
	if m.AzureClient == nil {
		azureClient, diags := buildAzureClient(m.azureConfig)
		if diags.HasError() {
			return nil, diags
		}
		m.AzureClient = azureClient
	}

	return m.AzureClient, nil
}

func buildAWSSessionPreCheck(region string) diag.Diagnostics {
	var diags diag.Diagnostics
	if region == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary: fmt.Sprintf("No AWS region specified. The AWS region must be provided either in "+
				"the provider configuration block or with the %s environment variable.", awsRegionVar),
		})
		return diags
	}
	return nil
}

type AzureClientAPI interface {
	SubscriptionID() string
	groupsListComplete(ctx context.Context, filter string) (*[]graphrbac.ADGroup, error)
	appsCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters managedapplications.Application) error
	appsGet(ctx context.Context, resourceGroupName string, applicationName string) (managedapplications.Application, error)
	appsDelete(ctx context.Context, resourceGroupName string, applicationName string) error
}

type AzureClient struct {
	ApplicationsClient *managedapplications.ApplicationsClient
	GroupsClient       *graphrbac.GroupsClient
}
