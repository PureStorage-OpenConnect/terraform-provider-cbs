//go:build mock

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


	This file contains a mock implementation of the Azure cloud service
	calls.

*/

package cloud

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	vaultSecret "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/managedapplications"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/internal/mockdb"
)

var azureTFOutputs = []string{
	"applicationName",
	"managedResourceGroupName",
	"ct0Name",
	"ct1Name",
	"managementEndpoint",
	"managementEndpointCT0",
	"managementEndpointCT1",
	"replicationEndpointCT0",
	"replicationEndpointCT1",
	"iSCSIEndpointCT0",
	"iSCSIEndpointCT1",
}

func buildAzureClient(ctx context.Context, config AzureConfig) (AzureClientAPI, error) {
	return &mockAzureClient{config}, nil
}

type mockAzureClient struct {
	config AzureConfig
}

func (m *mockAzureClient) SubscriptionID() string {
	return m.config.SubscriptionID

}

var mockGroupsListCompleteDisplayNameRE = regexp.MustCompile(`displayName eq '(\w+)'`)

// Expect filter to be something like `displayName eq 'a_GroupName'`
func (m *mockAzureClient) GroupsListComplete(ctx context.Context, filter string) (*[]graphrbac.ADGroup, error) {
	filterMatches := mockGroupsListCompleteDisplayNameRE.FindStringSubmatch(filter)

	if len(filterMatches) != 2 {
		return nil, fmt.Errorf("Mock didn't understand your groups list filter")
	}

	return &[]graphrbac.ADGroup{
		{DisplayName: &filterMatches[1], ObjectID: to.StringPtr("mock-group-object-id")},
	}, nil
}

func (m *mockAzureClient) AppsCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters managedapplications.Application) error {
	mockdb.AzureAppsDel(resourceGroupName, applicationName)

	outputs := make(map[string]interface{})
	for _, value := range append(azureTFOutputs, "floatingManagementIP") {
		outputs[value] = make(map[string]interface{})
		outputs[value].(map[string]interface{})["value"] = fmt.Sprintf("mock value for %s", value)
		outputs[value].(map[string]interface{})["type"] = "mock type"
	}

	appParams := parameters.ApplicationProperties.Parameters.(map[string]interface{})
	appParams["_artifactsLocationSasToken"] = map[string]interface{}{"type": "SecureString"}
	appParams["licenseKey"].(map[string]interface{})["type"] = "SecureString"
	appParams["licenseKey"].(map[string]interface{})["value"] = nil
	appParams["pureuserPublicKey"].(map[string]interface{})["type"] = "SecureString"
	appParams["pureuserPublicKey"].(map[string]interface{})["value"] = nil

	for _, v := range appParams {
		t := v.(map[string]interface{})["type"]
		if t == nil || t.(string) != "SecureString" {
			v.(map[string]interface{})["type"] = "mock type"
		}
	}

	parameters.ApplicationProperties.Outputs = outputs
	parameters.Response = autorest.Response{
		Response: &http.Response{StatusCode: http.StatusOK, Request: &http.Request{}},
	}
	parameters.Location = to.StringPtr("mock Location")
	parameters.ID = to.StringPtr(fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Solutions/applications/%s",
		m.SubscriptionID(), resourceGroupName, applicationName))

	mockdb.AzureAppsSet(resourceGroupName, applicationName, parameters)

	return nil
}

func (m *mockAzureClient) AppsGet(ctx context.Context, resourceGroupName string, applicationName string) (managedapplications.Application, error) {
	app := mockdb.AzureAppsGet(resourceGroupName, applicationName)
	if app == nil {
		return managedapplications.Application{Response: autorest.Response{
			Response: &http.Response{StatusCode: http.StatusNotFound, Request: &http.Request{}},
		}}, fmt.Errorf("mock appsGet not found")
	}
	return *app, nil
}

func (m *mockAzureClient) AppsDelete(ctx context.Context, resourceGroupName string, applicationName string) error {
	mockdb.AzureAppsDel(resourceGroupName, applicationName)
	return nil
}

const deletedMagicValue = "__DELETED"
const recoveredMagicValue = "__RECOVERED"

func (azureClient *mockAzureClient) SecretSet(ctx context.Context, vaultId string, secretName string, parameters vaultSecret.SecretSetParameters) (vaultSecret.SecretBundle, error) {
	wasDeleted := false
	existing := mockdb.AzureSecretGet(vaultId, secretName)
	if existing != nil {
		if *existing == deletedMagicValue {
			wasDeleted = true
		}
	}
	mockdb.AzureSecretDel(vaultId, secretName)
	// Random chance this was previously deleted

	if (!wasDeleted && existing != nil) || (!wasDeleted && rand.Intn(2) == 0) { // Not deleted
		mockdb.AzureSecretSet(vaultId, secretName, *parameters.Value)
		return vaultSecret.SecretBundle{Value: parameters.Value}, nil
	} else { // Was deleted
		mockdb.AzureSecretSet(vaultId, secretName, deletedMagicValue)
		return vaultSecret.SecretBundle{Response: autorest.Response{Response: &http.Response{StatusCode: 409}}}, fmt.Errorf("secret is deleted")
	}
}

func (azureClient *mockAzureClient) SecretGet(ctx context.Context, vaultId string, secretName string, version string) (vaultSecret.SecretBundle, error) {
	existing := mockdb.AzureSecretGet(vaultId, secretName)
	if existing == nil || *existing == deletedMagicValue {
		return vaultSecret.SecretBundle{Response: autorest.Response{Response: &http.Response{StatusCode: 404}}}, fmt.Errorf("secret not found")
	}
	return vaultSecret.SecretBundle{Value: existing}, nil
}

func (azureClient *mockAzureClient) SecretDelete(ctx context.Context, vaultId string, secretName string) (vaultSecret.DeletedSecretBundle, error) {
	existing := mockdb.AzureSecretGet(vaultId, secretName)
	if existing == nil || *existing == deletedMagicValue {
		return vaultSecret.DeletedSecretBundle{Response: autorest.Response{Response: &http.Response{StatusCode: 404}}}, fmt.Errorf("secret not found")
	}
	mockdb.AzureSecretDel(vaultId, secretName)
	mockdb.AzureSecretSet(vaultId, secretName, deletedMagicValue)
	return vaultSecret.DeletedSecretBundle{}, nil
}

func (azureClient *mockAzureClient) SecretRecover(ctx context.Context, vaultId string, secretName string) error {
	existing := mockdb.AzureSecretGet(vaultId, secretName)
	if existing == nil {
		return fmt.Errorf("secret not found")
	}
	if *existing != deletedMagicValue {
		return fmt.Errorf("secret not deleted")
	}
	mockdb.AzureSecretDel(vaultId, secretName)
	mockdb.AzureSecretSet(vaultId, secretName, recoveredMagicValue)
	return nil
}

func (azureClient *mockAzureClient) NewFAClient(host string, vaultId string, secretName string) (array.FAClientAPI, error) {
	return &array.MockFAClient{Host: host, Kind: array.FAClientKindAzure}, nil
}

func (azureClient *mockAzureClient) DeactivateWait() {}
