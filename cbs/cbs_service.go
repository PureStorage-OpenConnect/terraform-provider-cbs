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
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/managedapplications"
	"github.com/Azure/go-autorest/autorest"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

type CbsService struct {
	CloudFormation *cloudformation.CloudFormation
	AzureClient    *AzureClient
}

type AzureClient struct {
	ApplicationsClient *managedapplications.ApplicationsClient
	GroupsClient       *graphrbac.GroupsClient
}

func (m *CbsService) CloudFormationService() (*cloudformation.CloudFormation, diag.Diagnostics) {
	if m.CloudFormation == nil {
		cftSvc, diags := buildSession(awsRegionStr)
		if diags.HasError() {
			return nil, diags
		}
		m.CloudFormation = cftSvc
	}

	return m.CloudFormation, nil
}

func (m *CbsService) AzureClientService() (*AzureClient, diag.Diagnostics) {
	if m.AzureClient == nil {
		azureClient, diags := buildAzureClient(azureBuilder)
		if diags.HasError() {
			return nil, diags
		}
		m.AzureClient = azureClient
	}

	return m.AzureClient, nil
}

func buildSession(region string) (*cloudformation.CloudFormation, diag.Diagnostics) {
	var diags diag.Diagnostics
	if region == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary: fmt.Sprintf("No AWS region specified. The AWS region must be provided either in "+
				"the provider configuration block or with the %s environment variable.", awsRegionVar),
		})
		return nil, diags
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	cftSvc := cloudformation.New(sess)
	return cftSvc, nil
}

func buildAzureClient(builder *authentication.Builder) (*AzureClient, diag.Diagnostics) {
	var azureClient AzureClient
	config, err := builder.Build()
	if err != nil {
		return nil, diag.FromErr(err)
	}

	env, err := authentication.DetermineEnvironment(config.Environment)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	// This indicates that a 429 response should not be included as a retry attempt
	// so that we continue to retry until it succeeds. Set this behavior to keep
	// consistent with azurerm provider.
	autorest.Count429AsRetry = false

	oauthConfig, err := config.BuildOAuthConfig(env.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	sender := sender.BuildSender("cbs")
	auth, err := config.GetAuthorizationToken(sender, oauthConfig, env.TokenAudience)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	graphAuth, err := config.GetAuthorizationToken(sender, oauthConfig, env.GraphEndpoint)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	// Create applications client
	client := managedapplications.NewApplicationsClient(config.SubscriptionID)
	client.SubscriptionID = config.SubscriptionID
	client.Client.Authorizer = auth
	// Create groups client
	groupClient := graphrbac.NewGroupsClient(config.TenantID)
	groupClient.Client.Authorizer = graphAuth

	azureClient.ApplicationsClient = &client
	azureClient.GroupsClient = &groupClient
	return &azureClient, nil
}
