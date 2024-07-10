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

package cloud

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managedapplications"
	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	vaultSecret "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
)

type AWSClientAPI interface {
	CreateStack(input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)
	DescribeStacks(input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)
	WaitUntilStackCreateCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput) error
	WaitUntilStackDeleteCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput) error
	GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error)
	PutSecretValue(input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error)
	GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error)
	DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error)
	ValidateTemplate(input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)
	NewFAClient(ctx context.Context, host string, adminSecretsManagerArn string) (array.FAClientAPI, error)
}

func NewAWSClient(region string) (AWSClientAPI, error) {
	return buildAWSClient(region)
}

type AzureConfig struct {
	SubscriptionID string
	ClientID       string
	ClientSecret   string
	TenantID       string
}

type AzureClientAPI interface {
	SubscriptionID() string
	GroupsListComplete(ctx context.Context, filter string) (*[]graphrbac.ADGroup, error)
	AppsCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters managedapplications.Application) error
	AppsGet(ctx context.Context, resourceGroupName string, applicationName string) (managedapplications.Application, error)
	ResourcesGetByType(ctx context.Context, resourceType string, managedResourceGroup string) ([]resources.GenericResourceExpanded, error)
	ResourceGet(ctx context.Context, resourceID string) (resources.GenericResource, error)
	AppsDelete(ctx context.Context, resourceGroupName string, applicationName string) error
	SecretSet(ctx context.Context, vaultId string, secretName string, parameters vaultSecret.SecretSetParameters) (vaultSecret.SecretBundle, error)
	SecretGet(ctx context.Context, vaultId string, secretName string, version string) (vaultSecret.SecretBundle, error)
	SecretDelete(ctx context.Context, vaultId string, secretName string) (vaultSecret.DeletedSecretBundle, error)
	SecretRecover(ctx context.Context, vaultId string, secretName string) error
	DeactivateWait()
	NewFAClient(ctx context.Context, host string, vaultId string, secretName string) (array.FAClientAPI, error)
}

func NewAzureClient(ctx context.Context, config AzureConfig) (AzureClientAPI, error) {
	return buildAzureClient(ctx, config)
}
