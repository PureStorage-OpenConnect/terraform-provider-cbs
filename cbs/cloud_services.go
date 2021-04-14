// +build !mock

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


	This file contains wrappers for "real" cloud service calls, as opposed to alternative
	implementations that are mocked.  Some of these wrappers also do a small amount of
	"flattening" in order to simplify some of the cloud APIs.  Simplifications includes
	reducing the number of objects, and putting multiple calls together.  This helps serve
	as a single file for tracking all cloud accesses, and it also makes it clearer what
	interfaces need to be mocked.

*/

package cbs

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managedapplications"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// Aws things:

func buildAWSSession(region string) (*cloudformation.CloudFormation, diag.Diagnostics) {
	var diags = buildAWSSessionPreCheck(region)
	if diags != nil {
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

// Azure things:

func buildAzureClient(userConfig azureUserConfig) (AzureClientAPI, diag.Diagnostics) {

	builder := &authentication.Builder{
		SubscriptionID:           userConfig.SubscriptionID,
		ClientID:                 userConfig.ClientID,
		ClientSecret:             userConfig.ClientSecret,
		TenantID:                 userConfig.TenantID,
		Environment:              azureEnvironment,
		SupportsClientSecretAuth: true,
		SupportsAzureCliToken:    true,
	}

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
	applicationsClient := managedapplications.NewApplicationsClient(config.SubscriptionID)
	applicationsClient.SubscriptionID = config.SubscriptionID
	applicationsClient.Client.Authorizer = auth
	// Create groups client
	groupClient := graphrbac.NewGroupsClient(config.TenantID)
	groupClient.Client.Authorizer = graphAuth

	azureClient.ApplicationsClient = &applicationsClient
	azureClient.GroupsClient = &groupClient
	return &azureClient, nil
}

func (client *AzureClient) SubscriptionID() string {
	return client.ApplicationsClient.SubscriptionID
}

func (client *AzureClient) groupsListComplete(ctx context.Context, filter string) (*[]graphrbac.ADGroup, error) {
	resp, err := client.GroupsClient.ListComplete(ctx, filter)
	if err != nil {
		return nil, err
	}
	return resp.Response().Value, nil
}

func (azureClient *AzureClient) appsCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters managedapplications.Application) error {
	future, err := azureClient.ApplicationsClient.CreateOrUpdate(ctx, resourceGroupName, applicationName, parameters)
	if err != nil {
		return fmt.Errorf("failed to create Managed Application %q (Resource Group %q): %+v", applicationName, resourceGroupName, err)
	}
	err = future.WaitForCompletionRef(ctx, azureClient.ApplicationsClient.Client)
	if err != nil {
		return fmt.Errorf("failed to wait for creation of Managed Application %q (Resource Group %q): %+v", applicationName, resourceGroupName, err)
	}
	return nil
}

func (azureClient *AzureClient) appsGet(ctx context.Context, resourceGroupName string, applicationName string) (managedapplications.Application, error) {
	return azureClient.ApplicationsClient.Get(ctx, resourceGroupName, applicationName)
}

func (azureClient *AzureClient) appsDelete(ctx context.Context, resourceGroupName string, applicationName string) error {
	future, err := azureClient.ApplicationsClient.Delete(ctx, resourceGroupName, applicationName)
	if err != nil {
		return fmt.Errorf("failed to delete Managed Application %q (Resource Group %q): %+v", applicationName, resourceGroupName, err)
	}
	err = future.WaitForCompletionRef(ctx, azureClient.ApplicationsClient.Client)
	if err != nil {
		return fmt.Errorf("failed to wait for deleting Managed Application (Managed Application Name %q / Resource Group %q): %+v", applicationName, resourceGroupName, err)
	}
	return nil
}
