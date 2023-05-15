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
	"strconv"
	"testing"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/acceptance"
	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataAzurePlans(t *testing.T) {

	if os.Getenv(acceptance.EnvTfAccAzureSkipMarketplace) != "" {
		t.Skipf("Skipping acc test due to env variable '%s'", acceptance.EnvTfAccAzureSkipMarketplace)
	}

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProvidersFactory,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureDataPlansConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccDataAzurePlans(),
				),
			},
		},
	})
}

func testAccAzureDataPlansConfig() string {
	return `data "cbs_azure_plans" "azure_plans" {}`
}

func testAccDataAzurePlans() resource.TestCheckFunc {
	return func(s *terraform.State) error {
		data_resource := s.RootModule().Resources["data.cbs_azure_plans.azure_plans"]

		// We expect to have at least two plans released at any given time (current + LTS)
		plans_size, err := strconv.Atoi(data_resource.Primary.Attributes["plans.#"])
		if err != nil {
			return err
		}
		if plans_size < 2 {
			return fmt.Errorf("Unexpected plans size: %d", plans_size)
		}
		for i := 0; i < plans_size; i++ {
			// Check product version derivation from plan name makes sense.
			plan_name := data_resource.Primary.Attributes["plans."+strconv.Itoa(i)+".name"]
			name_match := plan_name_regexp.FindStringSubmatch(plan_name)
			name_version, err := version.NewVersion(fmt.Sprintf("%s.%s.%s", name_match[1], name_match[2], name_match[3]))
			if err != nil {
				return err
			}
			reference_version, _ := version.NewVersion("6.0.0")
			if name_version.LessThanOrEqual(reference_version) {
				return fmt.Errorf("Incorrect product version: %s", name_version.String())
			}

			// We expect that product and publisher is a well-known string.
			plan_product := data_resource.Primary.Attributes["plans."+strconv.Itoa(i)+".product"]
			if plan_product != "pure_storage_cloud_block_store_deployment" {
				return fmt.Errorf("Incorrect plan product: %s", plan_product)
			}
			plan_publisher := data_resource.Primary.Attributes["plans."+strconv.Itoa(i)+".publisher"]
			if plan_publisher != "purestoragemarketplaceadmin" {
				return fmt.Errorf("Incorrect plan publisher: %s", plan_publisher)
			}

			// The plan version string should be sane.
			plan_version := data_resource.Primary.Attributes["plans."+strconv.Itoa(i)+".version"]
			parsed_version, err := version.NewVersion(plan_version)
			if err != nil {
				return err
			}
			if plan_version != parsed_version.String() {
				return fmt.Errorf("Incorrect plan version: %s", plan_version)
			}
		}
		return nil
	}
}
