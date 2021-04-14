// +build mock

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

package cbs

import (
	"context"
	"fmt"
	"net/http"
	"regexp"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/managedapplications"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

func buildAzureClient(config azureUserConfig) (AzureClientAPI, diag.Diagnostics) {
	return &mockAzureClient{}, nil
}

type mockAzureClient struct {
	config azureUserConfig
}

func (m *mockAzureClient) SubscriptionID() string {
	return m.config.SubscriptionID

}

var mockGroupsListCompleteDisplayNameRE = regexp.MustCompile(`displayName eq '(\w+)'`)

// Expect filter to be something like `displayName eq 'a_GroupName'`
func (m *mockAzureClient) groupsListComplete(ctx context.Context, filter string) (*[]graphrbac.ADGroup, error) {
	filterMatches := mockGroupsListCompleteDisplayNameRE.FindStringSubmatch(filter)

	if len(filterMatches) != 2 {
		return nil, fmt.Errorf("Mock didn't understand your groups list filter")
	}

	return &[]graphrbac.ADGroup{
		{DisplayName: &filterMatches[1], ObjectID: to.StringPtr("mock-group-object-id")},
	}, nil
}

func (m *mockAzureClient) appsCreateOrUpdate(ctx context.Context, resourceGroupName string, applicationName string, parameters managedapplications.Application) error {
	mockDbAzureAppsDel(resourceGroupName, applicationName)

	outputs := make(map[string]interface{})
	for _, value := range append(azureTFOutputs, "floatingManagementIP") {
		outputs[value] = make(map[string]interface{})
		outputs[value].(map[string]interface{})["value"] = fmt.Sprintf("mock value for %s", value)
	}

	appParams := parameters.ApplicationProperties.Parameters.(*map[string]interface{})
	(*appParams)["_artifactsLocationSasToken"] = map[string]string{"value": "SecureString"}
	(*appParams)["licenseKey"].(map[string]interface{})["value"] = "SecureString"
	(*appParams)["pureuserPublicKey"].(map[string]interface{})["value"] = "SecureString"

	parameters.ApplicationProperties.Outputs = outputs
	parameters.Response = autorest.Response{
		Response: &http.Response{StatusCode: http.StatusOK, Request: &http.Request{}},
	}
	parameters.Location = to.StringPtr("mock Location")
	parameters.ID = to.StringPtr(fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Solutions/applications/%s",
		m.SubscriptionID(), resourceGroupName, applicationName))

	mockDbAzureAppsSet(resourceGroupName, applicationName, parameters)

	return nil
}

func (m *mockAzureClient) appsGet(ctx context.Context, resourceGroupName string, applicationName string) (managedapplications.Application, error) {
	app := mockDbAzureAppsGet(resourceGroupName, applicationName)
	if app == nil {
		return managedapplications.Application{Response: autorest.Response{
			Response: &http.Response{StatusCode: http.StatusNotFound, Request: &http.Request{}},
		}}, fmt.Errorf("mock appsGet not found")
	}
	return *app, nil
}

func (m *mockAzureClient) appsDelete(ctx context.Context, resourceGroupName string, applicationName string) error {
	mockDbAzureAppsDel(resourceGroupName, applicationName)
	return nil
}
