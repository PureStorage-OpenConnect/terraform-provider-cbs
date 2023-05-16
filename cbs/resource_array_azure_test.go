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

var cbsAzureParam acceptance.AccTestCbsAzureParams
var azureParamsConfigure sync.Once

func TestAccArrayAzure_basic(t *testing.T) {
	loadAccAzureParams(t)

	if os.Getenv(acceptance.EnvTfAccAzureSkipMarketplace) != "" {
		t.Skipf("Skipping acc test due to env variable '%s'", acceptance.EnvTfAccAzureSkipMarketplace)
	}

	arrayName := acctest.RandomWithPrefix(cbsAzureParam.ArrayName)
	resourceName := "cbs_array_azure.test_array_azure"
	orgDomain := "example.com"
	orgDomain2 := "example-invalid-update.com"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckArrayAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureConfig(arrayName, orgDomain, false),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureBasicCheck(resourceName, arrayName, orgDomain),
					testAccArrayAzureOptOutDefaultProtectionPolicy(resourceName),
				),
			},
			{
				Config:      testAccAzureConfig(arrayName, orgDomain2, false),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureBasicCheck(resourceName, arrayName, orgDomain),
				),
			},
		},
	})
}

func TestAccArrayAzure_basicFusion(t *testing.T) {
	loadAccAzureParams(t)

	if os.Getenv(acceptance.EnvTfAccAzureSkipMarketplace) != "" {
		t.Skipf("Skipping acc tests due to env variable '%s'", acceptance.EnvTfAccAzureSkipMarketplace)
	}

	if cbsAzureParam.FusionSECIdentity == "" {
		t.Skip("Skipping acc test as fusion_sec_identity is not set in the param file")
	}

	arrayName := acctest.RandomWithPrefix(cbsAzureParam.ArrayName)
	resourceName := "cbs_array_azure.test_array_azure"
	orgDomain := "example.com"
	orgDomain2 := "example-invalid-update.com"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckArrayAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureConfig(arrayName, orgDomain, true),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureBasicCheck(resourceName, arrayName, orgDomain),
					testAccArrayAzureOptOutDefaultProtectionPolicy(resourceName),
				),
			},
			{
				Config:      testAccAzureConfig(arrayName, orgDomain2, true),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureBasicCheck(resourceName, arrayName, orgDomain),
				),
			},
		},
	})
}

func TestAccArrayAzure_basicAppId(t *testing.T) {
	loadAccAzureParams(t)

	if cbsAzureParam.AppDefinitionId == "" {
		t.Skip("Skipping acc test due to app_definition_id is not set in the param file")
	}

	arrayName := acctest.RandomWithPrefix(cbsAzureParam.ArrayName)
	resourceName := "cbs_array_azure.test_array_azure"
	orgDomain := "example.com"
	orgDomain2 := "example-invalid-update.com"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckArrayAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureConfigAppId(arrayName, orgDomain, false),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAppIdCheck(resourceName, arrayName, orgDomain),
					testAccArrayAzureOptOutDefaultProtectionPolicy(resourceName),
				),
			},
			{
				Config:      testAccAzureConfigAppId(arrayName, orgDomain2, false),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAppIdCheck(resourceName, arrayName, orgDomain),
				),
			},
		},
	})
}

func TestAccArrayAzure_basicAppIdFusion(t *testing.T) {
	loadAccAzureParams(t)

	if cbsAzureParam.AppDefinitionId == "" {
		t.Skip("Skipping acc test due to app_definition_id is not set in the param file")
	}

	if os.Getenv(acceptance.EnvTfAccAzureSkipFusionAppId) != "" {
		t.Skipf("Skipping acc tests due to env variable '%s'", acceptance.EnvTfAccAzureSkipFusionAppId)
	}

	if cbsAzureParam.FusionSECIdentity == "" {
		t.Skip("Skipping acc test as fusion_sec_identity is not set in the param file")
	}

	arrayName := acctest.RandomWithPrefix(cbsAzureParam.ArrayName)
	resourceName := "cbs_array_azure.test_array_azure"
	orgDomain := "example.com"
	orgDomain2 := "example-invalid-update.com"

	resource.ParallelTest(t, resource.TestCase{
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckArrayAzureDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccAzureConfigAppId(arrayName, orgDomain, true),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAppIdCheck(resourceName, arrayName, orgDomain),
					testAccArrayAzureOptOutDefaultProtectionPolicy(resourceName),
				),
			},
			{
				Config:      testAccAzureConfigAppId(arrayName, orgDomain2, true),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAppIdCheck(resourceName, arrayName, orgDomain),
				),
			},
		},
	})
}

func testAccAzureConfig(name string, orgDomain string, fusionArray bool) string {
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

	fusionHCL := ""
	if fusionArray {
		fusionHCL = fmt.Sprintf(`fusion_sec_identity = "%s"`, cbsAzureParam.FusionSECIdentity)
	}

	return fmt.Sprintf(`
	resource "cbs_array_azure" "test_array_azure" {
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
		array_model = "%[13]s"
		zone = 3

		%[14]s

		jit_approval_group_object_ids = [
				"%[15]s",
		]

		%[16]s

		tags = {
			foo = "bar"
			test = "value"
		}
	}`, name, orgDomain, cbsAzureParam.ResourceGroupName, cbsAzureParam.LicenseKey, cbsAzureParam.PureuserPrivateKeyPath, cbsAzureParam.SystemSubnet,
		cbsAzureParam.ReplicationSubnet, cbsAzureParam.ISCSISubnet, cbsAzureParam.ManagementSubnet, cbsAzureParam.VirtualNetworkId,
		cbsAzureParam.Location, cbsAzureParam.KeyvaultId, cbsAzureParam.ArrayModel, fusionHCL, cbsAzureParam.JitGroupID, planHCL)
}

func testAccAzureConfigAppId(name string, orgDomain string, fusionArray bool) string {
	fusionHCL := ""
	if fusionArray {
		fusionHCL = fmt.Sprintf(`fusion_sec_identity = "%s"`, cbsAzureParam.FusionSECIdentity)
	}

	return fmt.Sprintf(`
	resource "cbs_array_azure" "test_array_azure" {
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
		array_model = "%[13]s"
		zone = 3

		%[14]s

		app_definition_id = "%[15]s"

		tags = {
			foo = "bar"
			test = "value"
		}
	}`, name, orgDomain, cbsAzureParam.ResourceGroupName, cbsAzureParam.LicenseKey, cbsAzureParam.PureuserPrivateKeyPath, cbsAzureParam.SystemSubnet,
		cbsAzureParam.ReplicationSubnet, cbsAzureParam.ISCSISubnet, cbsAzureParam.ManagementSubnet, cbsAzureParam.VirtualNetworkId,
		cbsAzureParam.Location, cbsAzureParam.KeyvaultId, cbsAzureParam.ArrayModel, fusionHCL, cbsAzureParam.AppDefinitionId)
}

// Lazy load the Azure param values from the json file specified at TEST_ACC_AZURE_PARAMS_PATH.
func loadAccAzureParams(t *testing.T) {
	azureParamsConfigure.Do(func() {
		cfgPath := os.Getenv(acceptance.EnvTfAccAzureParamsPath)
		if cfgPath == "" {
			t.Fatalf("%s environment variable must be set for acceptance testing", acceptance.EnvTfAccAzureParamsPath)
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

func testAccCheckAzureBasicCheck(resourceName string, arrayName string, orgDomain string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		testAccCheckAzureCommonAttrsCheck(resourceName, arrayName, orgDomain),
		resource.TestCheckResourceAttr(resourceName, "jit_approval_group_object_ids.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "jit_approval_group_object_ids.0", cbsAzureParam.JitGroupID),
		testAccCheckAzureOutputSetCheck(resourceName),
	)
}

func testAccCheckAzureAppIdCheck(resourceName string, arrayName string, orgDomain string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		testAccCheckAzureCommonAttrsCheck(resourceName, arrayName, orgDomain),
		resource.TestCheckResourceAttr(resourceName, "jit_approval_group_object_ids.#", "0"),
		testAccCheckAzureOutputSetCheck(resourceName),
	)
}

// Check for common parameters that are same in all of the test cases
func testAccCheckAzureCommonAttrsCheck(resourceName string, arrayName string, orgDomain string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "array_name", arrayName),
		resource.TestCheckResourceAttr(resourceName, "array_model", cbsAzureParam.ArrayModel),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.0", "user@example.com"),
		resource.TestCheckResourceAttr(resourceName, "key_vault_id", cbsAzureParam.KeyvaultId),
		resource.TestCheckResourceAttr(resourceName, "location", cbsAzureParam.Location),
		resource.TestCheckResourceAttr(resourceName, "log_sender_domain", orgDomain),
		resource.TestCheckResourceAttr(resourceName, "pureuser_private_key_path", cbsAzureParam.PureuserPrivateKeyPath),
		resource.TestCheckResourceAttr(resourceName, "resource_group_name", cbsAzureParam.ResourceGroupName),
		resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
		resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
		resource.TestCheckResourceAttr(resourceName, "tags.test", "value"),
		resource.TestCheckResourceAttr(resourceName, "virtual_network_id", cbsAzureParam.VirtualNetworkId),
		resource.TestCheckResourceAttr(resourceName, "zone", "3"),
		resource.TestCheckResourceAttr(resourceName, "system_subnet", cbsAzureParam.SystemSubnet),
		resource.TestCheckResourceAttr(resourceName, "replication_subnet", cbsAzureParam.ReplicationSubnet),
		resource.TestCheckResourceAttr(resourceName, "iscsi_subnet", cbsAzureParam.ISCSISubnet),
		resource.TestCheckResourceAttr(resourceName, "management_subnet", cbsAzureParam.ManagementSubnet),
	)
}

// Check that all the output parameters were set
func testAccCheckAzureOutputSetCheck(resourceName string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
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

// testAccArrayAzureOptOutDefaultProtectionPolicy opts out of the automatic safemode.
//
// Starting with 6.3 we have a default pgroup, "pgroup-auto",  which is created
// automatically and that will contain all the volumes on the array.
// This is normal and due to the always on safemode feature.
//
// In Terraform Acceptance testing, however, we don't want that pgroup and would
// like to remove it. The presence of the pgroup would've caused the destroy to fail.
func testAccArrayAzureOptOutDefaultProtectionPolicy(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ctx := context.Background()

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Managed Application not found: %s", resourceName)
		}

		return optOutDefaultProtectionPolicy(ctx, rs.Primary.Attributes)
	}
}
