//go:build !mock

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

/*

	This file contains wrappers for "real" cloud service calls, as opposed to alternative
	implementations that are mocked.  Some of these wrappers also do a small amount of
	"flattening" in order to simplify some of the cloud APIs.  Simplifications includes
	reducing the number of objects, and putting multiple calls together.  This helps serve
	as a single file for tracking all cloud accesses, and it also makes it clearer what
	interfaces need to be mocked.

*/

package cloud

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
	"time"

	vaultSecret "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	vaultManagement "github.com/Azure/azure-sdk-for-go/services/preview/keyvault/mgmt/2020-04-01-preview/keyvault"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managedapplications"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/tracing"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/internal/tfazurerm"
	"github.com/hashicorp/go-azure-helpers/authentication"
	"github.com/hashicorp/go-azure-helpers/sender"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

const azureEnvironment = "public"

// Used to trace Azure API requests
type tracerTransport struct {
	tflogCtx context.Context
	r        http.RoundTripper
}

func (d *tracerTransport) RoundTrip(h *http.Request) (*http.Response, error) {
	requestDump, _ := httputil.DumpRequestOut(h, true)
	tflog.Trace(d.tflogCtx, "Azure HTTP Trace Round Trip", map[string]interface{}{
		"request": string(requestDump),
	})

	resp, err := d.r.RoundTrip(h)
	responseDump, _ := httputil.DumpResponse(resp, true)
	tflog.Trace(d.tflogCtx, "Azure HTTP Trace Round Trip", map[string]interface{}{
		"response": string(responseDump),
		"err":      err,
	})
	return resp, err
}

// Used to trace Azure API requests
type tracer struct {
	tflogCtx context.Context
}

func (m *tracer) NewTransport(base *http.Transport) http.RoundTripper {
	return &tracerTransport{m.tflogCtx, http.DefaultTransport}
}

func (m *tracer) StartSpan(ctx context.Context, name string) context.Context {
	tflog.Trace(ctx, "Azure HTTP Trace Start Span", map[string]interface{}{"name": name})
	m.tflogCtx = ctx
	return ctx
}

func (m *tracer) EndSpan(ctx context.Context, httpStatusCode int, err error) {
	tflog.Trace(ctx, "Azure HTTP Trace End Span", map[string]interface{}{
		"status_code": httpStatusCode,
		"err":         err,
	})
}

type azureClient struct {
	applicationsClient    managedapplications.ApplicationsClient
	groupsClient          graphrbac.GroupsClient
	vaultManagementClient vaultManagement.VaultsClient
	vaultSecretClient     vaultSecret.BaseClient
}

func buildAzureClient(ctx context.Context, userConfig AzureConfig) (AzureClientAPI, error) {

	if os.Getenv("PS_HTTP_TRACE_LOGGING") != "" {
		tracing.Register(&tracer{ctx})
	}

	// Heavily inspired by https://github.com/terraform-providers/terraform-provider-azurerm/blob/dfba957d737fa474870eebafedb9326edf899858/azurerm/internal/clients/builder.go
	// and https://github.com/terraform-providers/terraform-provider-azurerm/blob/d155f299d12e6e2440f42d7c8695569f8256b9c6/azurerm/internal/services/keyvault/client/client.go

	builder := &authentication.Builder{
		SubscriptionID:           userConfig.SubscriptionID,
		ClientID:                 userConfig.ClientID,
		ClientSecret:             userConfig.ClientSecret,
		TenantID:                 userConfig.TenantID,
		Environment:              azureEnvironment,
		SupportsClientSecretAuth: true,
		SupportsAzureCliToken:    true,
	}

	config, err := builder.Build()
	if err != nil {
		return nil, err
	}

	env, err := authentication.DetermineEnvironment(config.Environment)
	if err != nil {
		return nil, err
	}

	// This indicates that a 429 response should not be included as a retry attempt
	// so that we continue to retry until it succeeds. Set this behavior to keep
	// consistent with azurerm provider.
	autorest.Count429AsRetry = false

	oauthConfig, err := config.BuildOAuthConfig(env.ActiveDirectoryEndpoint)
	if err != nil {
		return nil, err
	}

	sender := sender.BuildSender("cbs")
	auth, err := config.GetADALToken(ctx, sender, oauthConfig, env.TokenAudience)
	if err != nil {
		return nil, err
	}
	graphAuth, err := config.GetADALToken(ctx, sender, oauthConfig, env.GraphEndpoint)
	if err != nil {
		return nil, err
	}
	vaultAuthorizer := config.ADALBearerAuthorizerCallback(ctx, sender, oauthConfig)

	vaultSecretClient := vaultSecret.New()
	vaultSecretClient.Authorizer = vaultAuthorizer

	vaultManagementClient := vaultManagement.NewVaultsClientWithBaseURI(
		env.ResourceManagerEndpoint, config.SubscriptionID)
	vaultManagementClient.Authorizer = auth

	// Create applications client
	applicationsClient := managedapplications.NewApplicationsClient(config.SubscriptionID)
	applicationsClient.SubscriptionID = config.SubscriptionID
	applicationsClient.Authorizer = auth
	// Create groups client
	groupClient := graphrbac.NewGroupsClient(config.TenantID)
	groupClient.Authorizer = graphAuth

	return &azureClient{
		applicationsClient:    applicationsClient,
		groupsClient:          groupClient,
		vaultManagementClient: vaultManagementClient,
		vaultSecretClient:     vaultSecretClient,
	}, nil

}

func (client *azureClient) SubscriptionID() string {
	return client.applicationsClient.SubscriptionID
}

func (client *azureClient) GroupsListComplete(ctx context.Context, filter string) (*[]graphrbac.ADGroup, error) {
	resp, err := client.groupsClient.ListComplete(ctx, filter)
	if err != nil {
		return nil, err
	}
	return resp.Response().Value, nil
}

func (azureClient *azureClient) AppsCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters managedapplications.Application) error {
	future, err := azureClient.applicationsClient.CreateOrUpdate(ctx, resourceGroupName, applicationName, parameters)
	if err != nil {
		return fmt.Errorf("failed to create Managed Application %q (Resource Group %q): %+v", applicationName, resourceGroupName, err)
	}
	err = future.WaitForCompletionRef(ctx, azureClient.applicationsClient.Client)
	if err != nil {
		return fmt.Errorf("failed to wait for creation of Managed Application %q (Resource Group %q): %+v", applicationName, resourceGroupName, err)
	}
	return nil
}

func (azureClient *azureClient) AppsGet(ctx context.Context, resourceGroupName string, applicationName string) (managedapplications.Application, error) {
	return azureClient.applicationsClient.Get(ctx, resourceGroupName, applicationName)
}

func (azureClient *azureClient) AppsDelete(ctx context.Context, resourceGroupName string, applicationName string) error {
	future, err := azureClient.applicationsClient.Delete(ctx, resourceGroupName, applicationName)
	if err != nil {
		return fmt.Errorf("failed to delete Managed Application %q (Resource Group %q): %+v", applicationName, resourceGroupName, err)
	}
	err = future.WaitForCompletionRef(ctx, azureClient.applicationsClient.Client)
	if err != nil {
		return fmt.Errorf("failed to wait for deleting Managed Application (Managed Application Name %q / Resource Group %q): %+v", applicationName, resourceGroupName, err)
	}
	return nil
}

func (azureClient *azureClient) SecretSet(ctx context.Context, vaultId string, secretName string, parameters vaultSecret.SecretSetParameters) (vaultSecret.SecretBundle, error) {
	vaultUri, err := azureClient.baseUriForKeyVault(ctx, vaultId)
	if err != nil {
		return vaultSecret.SecretBundle{}, err
	}

	return azureClient.vaultSecretClient.SetSecret(ctx, *vaultUri, secretName, parameters)
}

func (azureClient *azureClient) SecretGet(ctx context.Context, vaultId string, secretName string, version string) (vaultSecret.SecretBundle, error) {
	vaultUri, err := azureClient.baseUriForKeyVault(ctx, vaultId)
	if err != nil {
		return vaultSecret.SecretBundle{}, err
	}
	return azureClient.vaultSecretClient.GetSecret(ctx, *vaultUri, secretName, version)
}

func (azureClient *azureClient) SecretDelete(ctx context.Context, vaultId string, secretName string) (vaultSecret.DeletedSecretBundle, error) {
	vaultUri, err := azureClient.baseUriForKeyVault(ctx, vaultId)
	if err != nil {
		return vaultSecret.DeletedSecretBundle{}, err
	}
	return azureClient.vaultSecretClient.DeleteSecret(ctx, *vaultUri, secretName)
}

func (azureClient *azureClient) SecretRecover(ctx context.Context, vaultId string, secretName string) error {
	vaultUri, err := azureClient.baseUriForKeyVault(ctx, vaultId)
	if err != nil {
		return err
	}
	_, err = azureClient.vaultSecretClient.RecoverDeletedSecret(ctx, *vaultUri, secretName)
	return err
}

// Inpsired by: https://github.com/terraform-providers/terraform-provider-azurerm/blob/8a924c35f7cc5a977877677dfb0312597aedfa5d/azurerm/internal/services/keyvault/client/helpers.go#L36 but simplified
func (azureClient *azureClient) baseUriForKeyVault(ctx context.Context, azureId string) (*string, error) {

	vaultName, vaultRGName, err := tfazurerm.ParseNameRGFromID(azureId, "vaults")
	if err != nil {
		return nil, err
	}

	resp, err := azureClient.vaultManagementClient.Get(ctx, vaultRGName, vaultName)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			return nil, fmt.Errorf("%s was not found", azureId)
		}
		return nil, fmt.Errorf("retrieving %s: %+v", azureId, err)
	}

	if resp.Properties == nil || resp.Properties.VaultURI == nil {
		return nil, fmt.Errorf("`properties` was nil for %s", azureId)
	}

	return resp.Properties.VaultURI, nil
}

func (azureClient *azureClient) NewFAClient(ctx context.Context, host string, vaultId string, secretName string) (array.FAClientAPI, error) {

	secret, err := azureClient.SecretGet(ctx, vaultId, secretName, "")
	if err != nil {
		return nil, err
	}
	return array.NewFAClient(host, *secret.Value)
}

func (azureClient *azureClient) DeactivateWait() {
	// Docs say that we should wait 5 minutes after deactivation on Azure before we delete the managed app
	// We have pulled this out into function so that the mocked version doesn't have to wait.
	time.Sleep(5 * time.Minute)
}
