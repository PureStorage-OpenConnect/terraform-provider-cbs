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

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type cbsAzureParams struct {
	PlanName               string `json:"plan_name"`
	PlanProduct            string `json:"plan_product"`
	PlanPublisher          string `json:"plan_publisher"`
	PlanVersion            string `json:"plan_version"`
	ResourceGroupName      string `json:"resource_group_name"`
	Location               string `json:"location"`
	LicenseKey             string `json:"license_key"`
	PureuserPrivateKeyPath string `json:"pureuser_private_key_path"`
	KeyvaultId             string `json:"keyvault_id"`
	ManagementSubnet       string `json:"management_subnet"`
	ISCSISubnet            string `json:"iscsi_subnet"`
	ReplicationSubnet      string `json:"replication_subnet"`
	SystemSubnet           string `json:"system_subnet"`
	VirtualNetworkId       string `json:"virtual_network_id"`
	JitGroup               string `json:"jit_group"`
	JitGroupID             string `json:"jit_group_id"`
}

const azureParamsPathVar = "TEST_ACC_AZURE_PARAMS_PATH"

var cbsAzureParam cbsAzureParams
var azureParamsConfigure sync.Once

func TestAccArrayAzure_basic(t *testing.T) {

	arrayName := acctest.RandomWithPrefix("tf-test-array-azure")
	resourceName := "cbs_array_azure.test_array_azure"
	orgDomain := "example.com"
	orgDomain2 := "example-invalid-update.com"
	loadAccAzureParams(t)

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckArrayAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureConfig(arrayName, orgDomain, false),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAllAttrs(resourceName, arrayName, orgDomain),
				),
			},
			{
				Config:      testAccAzureConfig(arrayName, orgDomain2, false),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAllAttrs(resourceName, arrayName, orgDomain),
				),
			},
		},
	})
}

func testAccAzureConfig(name string, orgDomain string, oldJit bool) string {
	planHCL := ""
	if cbsAzureParam.PlanName != "" ||
		cbsAzureParam.PlanProduct != "" ||
		cbsAzureParam.PlanPublisher != "" ||
		cbsAzureParam.PlanVersion != "" {
		planHCL = fmt.Sprintf(`
		plan {
			name = "%s"
			product = "%s"
			publisher = "%s"
			version = "%s"
		}
		`,
			cbsAzureParam.PlanName,
			cbsAzureParam.PlanProduct,
			cbsAzureParam.PlanPublisher,
			cbsAzureParam.PlanVersion,
		)
	}

	jitHCL := fmt.Sprintf(`
		jit_approval_group_object_ids = [
				"%s",
		]
	`, cbsAzureParam.JitGroupID)
	if oldJit {
		jitHCL = fmt.Sprintf(`
			jit_approval {
				activation_maximum_duration = "PT8H"
				approvers {
					groups = [
						"%s",
					]
				}
			}
		`, cbsAzureParam.JitGroup)
	}

	return fmt.Sprintf(`
	resource "cbs_array_azure" "test_array_azure" {
		%[14]s
		array_name = "%[1]s"
		log_sender_domain = "%[2]s"
		resource_group_name = "%[3]s"
		license_key = "%[4]s"
		pureuser_private_key_path = "%[5]s"
		system_subnet = "%[6]s"
		replication_subnet = "%[7]s"
		iscsi_subnet = "%[8]s"
		management_subnet = "%[9]s"
		virtual_network_id = "%[10]s"
		location = "%[11]s"
		key_vault_id = "%[12]s"

		alert_recipients = ["user@example.com"]
		array_model = "V10MUR1"
		zone = 3

		%[13]s

		tags = {
			foo = "bar"
			test = "value"
		}
	}`, name, orgDomain, cbsAzureParam.ResourceGroupName, cbsAzureParam.LicenseKey, cbsAzureParam.PureuserPrivateKeyPath, cbsAzureParam.SystemSubnet,
		cbsAzureParam.ReplicationSubnet, cbsAzureParam.ISCSISubnet, cbsAzureParam.ManagementSubnet, cbsAzureParam.VirtualNetworkId,
		cbsAzureParam.Location, cbsAzureParam.KeyvaultId, jitHCL, planHCL)
}

// Lazy load the Azure param values from the json file specified at TEST_ACC_AZURE_PARAMS_PATH.
func loadAccAzureParams(t *testing.T) {
	azureParamsConfigure.Do(func() {
		cfgPath := os.Getenv(azureParamsPathVar)
		if cfgPath == "" {
			t.Fatalf("%s environment variable must be set for acceptance testing", azureParamsPathVar)
		}
		cfgData, err := ioutil.ReadFile(cfgPath)
		if err != nil {
			t.Fatalf("Error reading azure config JSON file: %s", err)
		}
		if err := json.Unmarshal(cfgData, &cbsAzureParam); err != nil {
			t.Fatalf("Error unmarshaling JSON file: %s", err)
		}
	})
}

func testAccCheckArrayAzureDestroy(s *terraform.State) error {
	azureClient, diags := testAccProvider.Meta().(*CbsService).azureClientService(context.TODO())
	if diags.HasError() {
		return fmt.Errorf("err: %+v", diags)
	}
	ctx := context.Background()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cbs_array_azure" {
			continue
		}

		appName := rs.Primary.Attributes["array_name"]
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

func testAccArrayAzureExists(resourceName string) resource.TestCheckFunc {
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

		appName := rs.Primary.Attributes["array_name"]
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

func testAccCheckAzureAllAttrs(resourceName string, arrayName string, orgDomain string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "array_name", arrayName),
		resource.TestCheckResourceAttr(resourceName, "location", cbsAzureParam.Location),
		resource.TestCheckResourceAttr(resourceName, "resource_group_name", cbsAzureParam.ResourceGroupName),
		resource.TestCheckResourceAttr(resourceName, "zone", "3"),
		resource.TestCheckResourceAttr(resourceName, "log_sender_domain", orgDomain),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.0", "user@example.com"),
		resource.TestCheckResourceAttr(resourceName, "array_model", "V10MUR1"),
		resource.TestCheckResourceAttr(resourceName, "pureuser_private_key_path", cbsAzureParam.PureuserPrivateKeyPath),
		resource.TestCheckResourceAttr(resourceName, "system_subnet", cbsAzureParam.SystemSubnet),
		resource.TestCheckResourceAttr(resourceName, "replication_subnet", cbsAzureParam.ReplicationSubnet),
		resource.TestCheckResourceAttr(resourceName, "iscsi_subnet", cbsAzureParam.ISCSISubnet),
		resource.TestCheckResourceAttr(resourceName, "management_subnet", cbsAzureParam.ManagementSubnet),
		resource.TestCheckResourceAttr(resourceName, "virtual_network_id", cbsAzureParam.VirtualNetworkId),
		resource.TestCheckResourceAttr(resourceName, "key_vault_id", cbsAzureParam.KeyvaultId),
		resource.TestCheckResourceAttr(resourceName, "jit_approval_group_object_ids.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval_group_object_ids.0", cbsAzureParam.JitGroupID),
		resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
		resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
		resource.TestCheckResourceAttr(resourceName, "tags.test", "value"),
		resource.TestCheckResourceAttrSet(resourceName, "application_name"),
		resource.TestCheckResourceAttrSet(resourceName, "managed_resource_group_name"),
		resource.TestCheckResourceAttrSet(resourceName, "ct0_name"),
		resource.TestCheckResourceAttrSet(resourceName, "ct1_name"),
		resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
		resource.TestCheckResourceAttrSet(resourceName, "management_endpoint_ct0"),
		resource.TestCheckResourceAttrSet(resourceName, "management_endpoint_ct1"),
		resource.TestCheckResourceAttrSet(resourceName, "replication_endpoint_ct0"),
		resource.TestCheckResourceAttrSet(resourceName, "replication_endpoint_ct1"),
		resource.TestCheckResourceAttrSet(resourceName, "iscsi_endpoint_ct0"),
		resource.TestCheckResourceAttrSet(resourceName, "iscsi_endpoint_ct1"),
	)
}

func testAccCheckAzureBasicOldJitApprovalAttrs(resourceName string, arrayName string, orgDomain string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "array_name", arrayName),
		resource.TestCheckResourceAttr(resourceName, "location", cbsAzureParam.Location),
		resource.TestCheckResourceAttr(resourceName, "resource_group_name", cbsAzureParam.ResourceGroupName),
		resource.TestCheckResourceAttr(resourceName, "zone", "3"),
		resource.TestCheckResourceAttr(resourceName, "log_sender_domain", orgDomain),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.0", "user@example.com"),
		resource.TestCheckResourceAttr(resourceName, "array_model", "V10MUR1"),
		resource.TestCheckResourceAttr(resourceName, "pureuser_private_key_path", cbsAzureParam.PureuserPrivateKeyPath),
		resource.TestCheckResourceAttr(resourceName, "system_subnet", cbsAzureParam.SystemSubnet),
		resource.TestCheckResourceAttr(resourceName, "replication_subnet", cbsAzureParam.ReplicationSubnet),
		resource.TestCheckResourceAttr(resourceName, "iscsi_subnet", cbsAzureParam.ISCSISubnet),
		resource.TestCheckResourceAttr(resourceName, "management_subnet", cbsAzureParam.ManagementSubnet),
		resource.TestCheckResourceAttr(resourceName, "virtual_network_id", cbsAzureParam.VirtualNetworkId),
		resource.TestCheckResourceAttr(resourceName, "key_vault_id", cbsAzureParam.KeyvaultId),
		resource.TestCheckResourceAttr(resourceName, "jit_approval.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval.0.approvers.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval.0.activation_maximum_duration", "PT8H"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval.0.approvers.0.groups.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval.0.approvers.0.groups.0", cbsAzureParam.JitGroup),
		resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
		resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
		resource.TestCheckResourceAttr(resourceName, "tags.test", "value"),
		resource.TestCheckResourceAttrSet(resourceName, "application_name"),
		resource.TestCheckResourceAttrSet(resourceName, "managed_resource_group_name"),
		resource.TestCheckResourceAttrSet(resourceName, "ct0_name"),
		resource.TestCheckResourceAttrSet(resourceName, "ct1_name"),
		resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
		resource.TestCheckResourceAttrSet(resourceName, "management_endpoint_ct0"),
		resource.TestCheckResourceAttrSet(resourceName, "management_endpoint_ct1"),
		resource.TestCheckResourceAttrSet(resourceName, "replication_endpoint_ct0"),
		resource.TestCheckResourceAttrSet(resourceName, "replication_endpoint_ct1"),
		resource.TestCheckResourceAttrSet(resourceName, "iscsi_endpoint_ct0"),
		resource.TestCheckResourceAttrSet(resourceName, "iscsi_endpoint_ct1"),
	)
}

func TestAccArrayAzure_OldJitApproval(t *testing.T) {

	arrayName := acctest.RandomWithPrefix("tf-test-array-azure")
	resourceName := "cbs_array_azure.test_array_azure"
	orgDomain := "example.com"
	orgDomain2 := "example-invalid-update.com"
	loadAccAzureParams(t)

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckArrayAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureConfig(arrayName, orgDomain, true),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureBasicOldJitApprovalAttrs(resourceName, arrayName, orgDomain),
				),
			},
			{
				Config:      testAccAzureConfig(arrayName, orgDomain2, true),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureBasicOldJitApprovalAttrs(resourceName, arrayName, orgDomain),
				),
			},
		},
	})
}
