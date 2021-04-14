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
	ResourceGroupName        string `json:"resource_group_name"`
	Location                 string `json:"location"`
	LicenseKey               string `json:"license_key"`
	PureuserPublicKey        string `json:"pureuser_public_key"`
	ManagementSubnet         string `json:"management_subnet"`
	ISCSISubnet              string `json:"iscsi_subnet"`
	ReplicationSubnet        string `json:"replication_subnet"`
	SystemSubnet             string `json:"system_subnet"`
	VirtualNetwork           string `json:"virtual_network"`
	ManagementResourceGroup  string `json:"management_resource_group"`
	ISCSIResourceGroup       string `json:"iscsi_resource_group"`
	ReplicationResourceGroup string `json:"replication_resource_group"`
	SystemResourceGroup      string `json:"system_resource_group"`
	JitGroup                 string `json:"jit_group"`
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
		PreCheck:          func() { testAccAzureConfigPreCheck(t) },
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckArrayAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureBasicConfig(arrayName, orgDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAllAttrs(resourceName, arrayName, orgDomain),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config:      testAccAzureBasicConfig(arrayName, orgDomain2),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAllAttrs(resourceName, arrayName, orgDomain),
				),
			},
		},
	})
}

func testAccAzureBasicConfig(name string, orgDomain string) string {
	return fmt.Sprintf(`
	resource "cbs_array_azure" "test_array_azure" {
		array_name = "%[1]s"
		log_sender_domain = "%[2]s"
		resource_group_name = "%[3]s"
		license_key = "%[4]s"
		pureuser_public_key = "%[5]s"
		system_subnet = "%[6]s"
		replication_subnet = "%[7]s"
		iscsi_subnet = "%[8]s"
		management_subnet = "%[9]s"
		management_resource_group = "%[10]s"
		system_resource_group = "%[11]s"
		iscsi_resource_group = "%[12]s"
		replication_resource_group = "%[13]s"
		virtual_network = "%[14]s"
		location = "%[15]s"
		alert_recipients = ["user@example.com"]
		array_model = "V10MUR1"
		zone = 1

		jit_approval {
			activation_maximum_duration = "PT2H"
			approvers {
				groups = [
					"%[16]s",
				]
			}
		}

		tags = {
			foo = "bar"
			test = "value"
		}
	}`, name, orgDomain, cbsAzureParam.ResourceGroupName, cbsAzureParam.LicenseKey, cbsAzureParam.PureuserPublicKey, cbsAzureParam.SystemSubnet,
		cbsAzureParam.ReplicationSubnet, cbsAzureParam.ISCSISubnet, cbsAzureParam.ManagementSubnet, cbsAzureParam.ManagementResourceGroup,
		cbsAzureParam.SystemResourceGroup, cbsAzureParam.ISCSIResourceGroup, cbsAzureParam.ReplicationResourceGroup, cbsAzureParam.VirtualNetwork,
		cbsAzureParam.Location, cbsAzureParam.JitGroup)
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
	azureClient, diags := testAccProvider.Meta().(*CbsService).AzureClientService()
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
		resp, err := azureClient.appsGet(ctx, resourceGroup, appName)
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
		azureClient, diags := testAccProvider.Meta().(*CbsService).AzureClientService()
		if diags.HasError() {
			return fmt.Errorf("err: %+v", diags)
		}

		appName := rs.Primary.Attributes["array_name"]
		resourceGroup := rs.Primary.Attributes["resource_group_name"]

		if resp, err := azureClient.appsGet(ctx, resourceGroup, appName); err != nil {
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
		resource.TestCheckResourceAttr(resourceName, "zone", "1"),
		resource.TestCheckResourceAttr(resourceName, "log_sender_domain", orgDomain),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.0", "user@example.com"),
		resource.TestCheckResourceAttr(resourceName, "array_model", "V10MUR1"),
		resource.TestCheckResourceAttr(resourceName, "system_subnet", cbsAzureParam.SystemSubnet),
		resource.TestCheckResourceAttr(resourceName, "replication_subnet", cbsAzureParam.ReplicationSubnet),
		resource.TestCheckResourceAttr(resourceName, "iscsi_subnet", cbsAzureParam.ISCSISubnet),
		resource.TestCheckResourceAttr(resourceName, "management_subnet", cbsAzureParam.ManagementSubnet),
		resource.TestCheckResourceAttr(resourceName, "virtual_network", cbsAzureParam.VirtualNetwork),
		resource.TestCheckResourceAttr(resourceName, "management_resource_group", cbsAzureParam.ManagementResourceGroup),
		resource.TestCheckResourceAttr(resourceName, "system_resource_group", cbsAzureParam.SystemResourceGroup),
		resource.TestCheckResourceAttr(resourceName, "iscsi_resource_group", cbsAzureParam.ISCSIResourceGroup),
		resource.TestCheckResourceAttr(resourceName, "replication_resource_group", cbsAzureParam.ReplicationResourceGroup),
		resource.TestCheckResourceAttr(resourceName, "jit_approval.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval.0.approvers.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval.0.activation_maximum_duration", "PT2H"),
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
