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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/acceptance"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var cbsFusionSECAzureParam acceptance.AccTestCbsFusionSECAzureParams
var fusionSECAzureParamsConfigure sync.Once

func TestAccFusionSECAzure_basic(t *testing.T) {
	loadAccFusionSECAzureParams(t)

	if os.Getenv(acceptance.EnvTfAccSkipFusionSECAzure) != "" {
		t.Skipf("Skipping acc test due to env variable '%s'", acceptance.EnvTfAccSkipFusionSECAzure)
	}

	fusionSECName := acctest.RandomWithPrefix(cbsFusionSECAzureParam.FusionSECName)
	resourceName := "cbs_fusion_sec_azure.test_fusion_sec_azure"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckFusionSECAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccFusionSECAzureConfig(fusionSECName, false),
				Check: resource.ComposeTestCheckFunc(
					testAccFusionSECAzureExists(resourceName),
					testAccCheckFusionSECAzureBasicCheck(resourceName, fusionSECName),
				),
			},
			{
				Config:      testAccFusionSECAzureConfig(fusionSECName, true),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccFusionSECAzureExists(resourceName),
					testAccCheckFusionSECAzureBasicCheck(resourceName, fusionSECName),
				),
			},
		},
	})
}

func testAccFusionSECAzureConfig(name string, update bool) string {
	planHCL := ""
	if cbsFusionSECAzureParam.PlanName != "" ||
		cbsFusionSECAzureParam.PlanProduct != "" ||
		cbsFusionSECAzureParam.PlanPublisher != "" ||
		cbsFusionSECAzureParam.PlanVersion != "" {
		planHCL = fmt.Sprintf(`
		plan {
			name = "%s"
			product = "%s"
			publisher = "%s"
			version = "%s"
		}
		`,
			cbsFusionSECAzureParam.PlanName,
			cbsFusionSECAzureParam.PlanProduct,
			cbsFusionSECAzureParam.PlanPublisher,
			cbsFusionSECAzureParam.PlanVersion,
		)
	}

	if update {
		name = fmt.Sprintf("%s-update", name)
	}

	return fmt.Sprintf(`
	resource "cbs_fusion_sec_azure" "test_fusion_sec_azure" {
		fusion_sec_name = "%[1]s"
		resource_group_name = "%[2]s"
		load_balancer_network_rg = "%[3]s"
		load_balancer_network_name = "%[4]s"
		load_balancer_subnet = "%[5]s"
		location = "%[6]s"

		jit_approval_group_object_ids = [
				"%[7]s",
		]

		%[8]s

		tags = {
			foo = "bar"
			test = "value"
		}
	}`, name, cbsFusionSECAzureParam.ResourceGroupName,
		cbsFusionSECAzureParam.LoadBalancerNetworkRg, cbsFusionSECAzureParam.LoadBalancerNetworkName, cbsFusionSECAzureParam.LoadBalancerSubnet,
		cbsFusionSECAzureParam.Location, cbsFusionSECAzureParam.JitGroupID, planHCL)
}

// Lazy load the Fusion SEC Azure param values from the json file specified at TEST_ACC_FUSION_SEC_AZURE_PARAMS_PATH.
func loadAccFusionSECAzureParams(t *testing.T) {
	fusionSECAzureParamsConfigure.Do(func() {
		cfgPath := os.Getenv(acceptance.EnvTfAccFusionSECAzureParamsPath)
		if cfgPath == "" {
			t.Fatalf("%s environment variable must be set for acceptance testing", acceptance.EnvTfAccFusionSECAzureParamsPath)
		}
		cfgData, err := ioutil.ReadFile(cfgPath)
		if err != nil {
			t.Fatalf("Error reading azure config JSON file: %s", err)
		}
		if err := json.Unmarshal(cfgData, &cbsFusionSECAzureParam); err != nil {
			t.Fatalf("Error unmarshaling JSON file: %s", err)
		}
	})
}

func testAccCheckFusionSECAzureDestroy(s *terraform.State) error {
	azureClient, diags := testAccProvider.Meta().(*CbsService).azureClientService(context.TODO())
	if diags.HasError() {
		return fmt.Errorf("err: %+v", diags)
	}
	ctx := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cbs_fusion_sec_azure" {
			continue
		}

		appName := rs.Primary.Attributes["fusion_sec_name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]
		resp, err := azureClient.AppsGet(ctx, resourceGroup, appName)
		if err != nil {
			if responseWasNotFound(resp.Response) {
				return nil
			}
			return err
		}
		return fmt.Errorf("Managed Application still exists:\n%#v", resp)
	}

	return nil
}

func testAccFusionSECAzureExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Managed Application not found: %s", resourceName)
		}

		ctx := context.Background()
		azureClient, diags := testAccProvider.Meta().(*CbsService).azureClientService(context.TODO())
		if diags.HasError() {
			return fmt.Errorf("err: %+v", diags)
		}

		appName := rs.Primary.Attributes["fusion_sec_name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		if resp, err := azureClient.AppsGet(ctx, resourceGroup, appName); err != nil {
			if responseWasNotFound(resp.Response) {
				return fmt.Errorf("Managed Application %q (Resource Group %q) does not exist", appName, resourceGroup)
			}
			return err
		}

		return nil
	}
}

func testAccCheckFusionSECAzureBasicCheck(resourceName string, fusionSECName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		testAccCheckFusionSECAzureCommonAttrsCheck(resourceName, fusionSECName),
		resource.TestCheckResourceAttr(resourceName, "jit_approval_group_object_ids.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval_group_object_ids.0", cbsFusionSECAzureParam.JitGroupID),
		testAccCheckFusionSECAzureOutputSetCheck(resourceName),
	)
}

func testAccCheckFusionSECAzureCommonAttrsCheck(resourceName string, fusionSECName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "fusion_sec_name", fusionSECName),
		resource.TestCheckResourceAttr(resourceName, "location", cbsFusionSECAzureParam.Location),
		resource.TestCheckResourceAttr(resourceName, "resource_group_name", cbsFusionSECAzureParam.ResourceGroupName),
		resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
		resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
		resource.TestCheckResourceAttr(resourceName, "tags.test", "value"),
		resource.TestCheckResourceAttr(resourceName, "load_balancer_network_rg", cbsFusionSECAzureParam.LoadBalancerNetworkRg),
		resource.TestCheckResourceAttr(resourceName, "load_balancer_network_name", cbsFusionSECAzureParam.LoadBalancerNetworkName),
		resource.TestCheckResourceAttr(resourceName, "load_balancer_subnet", cbsFusionSECAzureParam.LoadBalancerSubnet),
	)
}

// Check that all the output parameters were set
func testAccCheckFusionSECAzureOutputSetCheck(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttrSet(resourceName, "application_name"),
		resource.TestCheckResourceAttrSet(resourceName, "managed_resource_group_name"),
		resource.TestCheckResourceAttrSet(resourceName, "hmvip0"),
		resource.TestCheckResourceAttrSet(resourceName, "hmvip1"),
		resource.TestCheckResourceAttrSet(resourceName, "load_balancer_full_identity_id"),
	)
}
