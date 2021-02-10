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
	"os"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		PreCheck:          func() { testAccPreCheck(t) },
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
		PreCheck:          func() { testAccPreCheck(t) },
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

func testProviderEmptyConfig() string {
	return fmt.Sprintf(`
	provider "cbs" {}
	`) + testInvalidAWSArrayConfig()
}

func testProviderEmptyAWSConfig() string {
	return fmt.Sprintf(`
	provider "cbs" {
		aws {}
	}
	`) + testInvalidAWSArrayConfig()
}

// For the test to initialize the provider, the configuration must contain a resource. We add this
// resource to the test configurations so that we can test the provider setup. The values don't matter
// since we can verify what we need to in the provider before we check the resource values.
func testInvalidAWSArrayConfig() string {
	return fmt.Sprintf(`
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
	}`)
}

func testAccPreCheck(t *testing.T) {
	// Make sure the AWS_DEFAULT_REGION var is unset, so that we can test explicit provider configurations
	os.Unsetenv(awsRegionVar)
}
