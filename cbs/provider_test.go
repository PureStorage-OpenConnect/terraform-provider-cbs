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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/cloud"

	"github.com/stretchr/testify/require"
)

var testAccProvider *schema.Provider
var testAccProvidersFactory map[string]func() (*schema.Provider, error)

func init() {
	testAccProvider = Provider()

	testAccProvidersFactory = map[string]func() (*schema.Provider, error){
		"cbs": func() (*schema.Provider, error) { return testAccProvider, nil },
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ *schema.Provider = Provider()
}

func TestAccProvider_emptyConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccUnsetAWSPreCheck(t) },
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config:      testProviderEmptyConfig(),
				ExpectError: regexp.MustCompile("No AWS region specified."),
			},
		},
	})
}

func TestAccProvider_emptyAWSConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccUnsetAWSPreCheck(t) },
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config:      testProviderEmptyAWSConfig(),
				ExpectError: regexp.MustCompile("No AWS region specified."),
			},
		},
	})
}

func TestAccProvider_unmatchedConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccCBSPreCheck(t) },
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config:      testProviderUnmatchedConfig(),
				ExpectError: regexp.MustCompile("No AWS region specified."),
			},
		},
	})
}

func TestAccProvider_allConfig(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { testAccCBSPreCheck(t) },
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      nil,
		Steps: []resource.TestStep{
			{
				Config:      testProviderAllConfig(),
				ExpectError: regexp.MustCompile("No AWS region specified."),
			},
		},
	})
}

func testProviderEmptyConfig() string {
	return `
	provider "cbs" {}
	` + testInvalidAWSArrayConfig()
}

func testProviderEmptyAWSConfig() string {
	return `
	provider "cbs" {
		aws {}
	}
	` + testInvalidAWSArrayConfig()
}

func testProviderUnmatchedConfig() string {
	return `
    provider "cbs" {
        azure {}
    }
    ` + testInvalidAWSArrayConfig()
}

func testProviderAllConfig() string {
	return `
    provider "cbs" {
		aws {}
		azure {}
    }
    ` + testInvalidAWSArrayConfig()
}

// For the test to initialize the provider, the configuration must contain a resource. We add this
// resource to the test configurations so that we can test the provider setup. The values don't matter
// since we can verify what we need to in the provider before we check the resource values.
func testInvalidAWSArrayConfig() string {
	return `
	resource "cbs_array_aws" "test_array_aws" {

		array_name = "invalid-array"

		log_sender_domain = "example.org"

		deployment_template_url = "fake-template"

		deployment_role_arn = "arn:aws:iam::foo"

		array_model = "V10AR1"
		license_key = "foo"

		pureuser_key_pair_name = "foo"

		system_subnet = "subnet-foo"
		replication_subnet = "subnet-foo"
		iscsi_subnet = "subnet-foo"
		management_subnet = "subnet-foo"

		replication_security_group = "sg-foo"
		iscsi_security_group = "sg-foo"
		management_security_group = "sg-foo"
	}`
}

func testAccUnsetAWSPreCheck(t *testing.T) {
	// Make sure the AWS_DEFAULT_REGION var is unset, so that we can test explicit provider configurations
	os.Unsetenv(awsRegionVar)
}

func TestAccProvider_azureCLIAuth(t *testing.T) {
	if os.Getenv("MOCK_OUT_DIR") != "" {
		t.Skip() // TODO:This requires Azure credentials in container, haven't set that up yet
	}

	if os.Getenv("TF_ACC") == "" {
		return
	}

	// Support only Azure CLI authentication
	_, err := cloud.NewAzureClient(cloud.AzureConfig{})
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
}

func TestAccProvider_azureServicePrincipalAuth(t *testing.T) {
	if os.Getenv("TF_ACC") == "" {
		return
	}

	// Support only Service Principal authentication
	testAccAzureConfigPreCheck(t)

	config := cloud.AzureConfig{
		SubscriptionID: os.Getenv(azureSubscriptionID),
		ClientID:       os.Getenv(azureClientID),
		ClientSecret:   os.Getenv(azureClientSecret),
		TenantID:       os.Getenv(azureTenantID),
	}

	_, err := cloud.NewAzureClient(config)
	if err != nil {
		t.Fatalf("err: %+v", err)
	}
}

func testAccAzureConfigPreCheck(t *testing.T) {
	variables := []string{
		azureSubscriptionID,
		azureClientID,
		azureTenantID,
		azureClientSecret,
	}

	for _, variable := range variables {
		value := os.Getenv(variable)
		if value == "" {
			t.Fatalf("`%s` must be set for azure acceptance tests!", variable)
		}
	}
}

func testAccCBSPreCheck(t *testing.T) {
	testAccUnsetAWSPreCheck(t)
	testAccAzureConfigPreCheck(t)
}

func TestProvider_HasExpectedResources(t *testing.T) {
	expectedResources := []string{
		"cbs_array_aws",
		"cbs_array_azure",
	}

	resources := testAccProvider.ResourcesMap
	require.Equal(t, len(expectedResources), len(resources), "Got an unexpected number of resources.")

	for _, resource := range expectedResources {
		require.Contains(t, resources, resource, "An expected resource was not registered.")
		require.NotNil(t, resources[resource], "A resource schema is nil.")
	}
}
