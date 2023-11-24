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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/acceptance"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/cloud"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var cbsAzureParam acceptance.AccTestCbsAzureParams
var azureParamsConfigure sync.Once

func TestAccArrayAzure_basic(t *testing.T) {
	loadAccAzureParams(t)
	plans, err := QueryMarketplaceForPlans(context.TODO())
	if err != nil {
		t.Error(err)
	}

	// when tags are provided plan is needed to get plan specific
	// list of resources to assign all tags to each resource
	cbsAzureParam.PlanName = plans[0].Plan.Name
	cbsAzureParam.PlanProduct = "pure_storage_cloud_block_store_deployment"
	cbsAzureParam.PlanPublisher = "purestoragemarketplaceadmin"
	cbsAzureParam.PlanVersion = plans[0].Plan.Version

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
				Config: testAccAzureConfig(arrayName, orgDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccArrayTags(resourceName),
					testAccCheckAzureBasicCheck(resourceName, arrayName, orgDomain),
					testAccArrayAzureOptOutDefaultProtectionPolicy(resourceName),
				),
			},
			{
				Config:      testAccAzureConfig(arrayName, orgDomain2),
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
				Config: testAccAzureConfigAppId(arrayName, orgDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAppIdCheck(resourceName, arrayName, orgDomain),
					testAccArrayAzureOptOutDefaultProtectionPolicy(resourceName),
				),
			},
			{
				Config:      testAccAzureConfigAppId(arrayName, orgDomain2),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAzureExists(resourceName),
					testAccCheckAzureAppIdCheck(resourceName, arrayName, orgDomain),
				),
			},
		},
	})
}

func TestGetResourcesFromTemplateJsonMalformedJson(t *testing.T) {
	// case for incorrect json file
	_, err := GetResourcesFromTemplateJson([]byte{'x'})
	if err == nil {
		t.Errorf("incorrect json parsing must have failed")
	}
}

func TestGetResourcesFromTemplateJsonEmptyResources(t *testing.T) {
	resource_json := []byte(`{
		"parameters": {
			"steps": [
				{
					"name": "network",
				},
				{
					"name": "tags_name",
					"elements": [
						{
							"name": "tags_name",
							"type": "Microsoft.Common.TagsByResource",
							"resources": [
							],
						}
					]
				},
			],
		}
	}`)

	// resources not present
	_, err := GetResourcesFromTemplateJson(resource_json)
	if err == nil {
		t.Errorf("Should have failed as resources are not found")
	}
}

func TestGetResourcesFromTemplateIncorrectField(t *testing.T) {
	var err error
	resource_json := []byte(`{
		"parameters": {
			"steps": [
				{
					"name": "network",
				},
				{
					"name": "tags_name",
					"elements": [
						{
							"name": "tags_name",
							"type": "Microsoft.Common.TagsByResource",
							"resources": [
								"Microsoft.Compute/virtualMachines",
								"Microsoft.Network/networkInterfaces",
							],
						}
					]
				},
			],
		}
	}`)

	//tags parameter not found
	_, err = GetResourcesFromTemplateJson(resource_json)
	if err == nil {
		t.Errorf("should have failed as tags not found %+v", err)
		return
	}

	if err.Error() != "invalid character '}' looking for beginning of object key string" {
		t.Errorf("should have failed for the error - 'invalid character '}' looking for beginning of object key string'")
	}
}

func TestSetResourceTagsResourceSuccess(t *testing.T) {
	tagsMap := make(map[string]interface{})
	resource := "someResourceType"
	resources := []string{"someResourceType", "someOtherResource"}
	resourceTagList := []interface{}{map[string]interface{}{"name": "someTag", "value": "someTagValue"}}

	err := SetResourceTags(resource, resourceTagList, resources, &tagsMap)
	if err != nil {
		t.Errorf("must have succeeded but have error %s", err.Error())
	}

	if _, ok := tagsMap[resource]; !ok {
		t.Errorf("resource tag is not set")
	}

	if val, ok := tagsMap[resource].(map[string]interface{})["someTag"]; !ok {
		t.Errorf("tag not set for the resource")
	} else {
		if val.(string) != "someTagValue" {
			t.Errorf("tag value not correct")
		}
	}
}

func TestSetResourceTagsResourceNotFound(t *testing.T) {
	tagsMap := make(map[string]interface{})
	resource := "someResourceType1"
	resources := []string{"someResourceType", "someOtherResource"}
	resourceTagList := []interface{}{map[string]string{"tag_name": "someTagName", "tag_value": "someTagValue"}}

	err := SetResourceTags(resource, resourceTagList, resources, &tagsMap)
	if err == nil {
		t.Errorf("must have failed with unknown resource type mssage")
	}

	if err.Error() != "provided resource type someResourceType1 not found" {
		t.Errorf("error must be: %s but was %s", "provided resource type someResourceType1 not found", err.Error())
	}
}

func TestGetResourcesFromTemplateSuccess(t *testing.T) {
	var resources []string
	var err error
	resource_json := []byte(`{
		"parameters": {
			"steps": [
				{
					"name": "network"
				},
				{
					"name": "tags",
					"elements": [
						{
							"name": "tags",
							"type": "Microsoft.Common.TagsByResource",
							"resources": [
								"Microsoft.Compute/virtualMachines",
								"Microsoft.Network/networkInterfaces"
							]
						}
					]
				}
			]
		}
	}`)

	resources, err = GetResourcesFromTemplateJson(resource_json)
	if err != nil {
		t.Errorf("json parsing must have been successful %s", err.Error())
	}

	if len(resources) != 2 {
		t.Errorf("unexpected number of resources %d", len(resources))
	}
}

func testAccAzureConfig(name string, orgDomain string) string {
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

	resourceTags := "[]"
	if cbsAzureParam.ResourceTags != "" {
		resourceTags = cbsAzureParam.ResourceTags
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

		jit_approval_group_object_ids = [
				"%[14]s",
		]

		%[15]s

		user_assigned_identity = "%[16]s"
		tags = {
			foo = "bar"
			dada = "dudu"
			some_tag_2 = "some_tag_2_value"
		}

		dynamic "resource_tags" {
			for_each = %[17]s
			content {
				resource = resource_tags.value.resource
				dynamic "tag" {
					for_each = resource_tags.value.tags
					content {
						name = tag.value.tag_name
						value = tag.value.tag_value
					}
				}
			}
		}
	}`, name, orgDomain, cbsAzureParam.ResourceGroupName, cbsAzureParam.LicenseKey, cbsAzureParam.PureuserPrivateKeyPath, cbsAzureParam.SystemSubnet,
		cbsAzureParam.ReplicationSubnet, cbsAzureParam.ISCSISubnet, cbsAzureParam.ManagementSubnet, cbsAzureParam.VirtualNetworkId,
		cbsAzureParam.Location, cbsAzureParam.KeyvaultId, cbsAzureParam.ArrayModel, cbsAzureParam.JitGroupID, planHCL,
		cbsAzureParam.UserAssignedIdentity, resourceTags)
}

func testAccAzureConfigAppId(name string, orgDomain string) string {
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

		app_definition_id = "%[14]s"
		user_assigned_identity = "%[15]s"
	}`, name, orgDomain, cbsAzureParam.ResourceGroupName, cbsAzureParam.LicenseKey, cbsAzureParam.PureuserPrivateKeyPath, cbsAzureParam.SystemSubnet,
		cbsAzureParam.ReplicationSubnet, cbsAzureParam.ISCSISubnet, cbsAzureParam.ManagementSubnet, cbsAzureParam.VirtualNetworkId,
		cbsAzureParam.Location, cbsAzureParam.KeyvaultId, cbsAzureParam.ArrayModel, cbsAzureParam.AppDefinitionId,
		cbsAzureParam.UserAssignedIdentity)
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

func checkVirtualMachineTags(ctx context.Context, api cloud.AzureClientAPI, managedResourceGroup string) error {
	virtualMachines, err := api.ResourcesGetByType(ctx, "Microsoft.Compute/virtualMachines", managedResourceGroup)
	if err != nil {
		return err
	}

	for _, virtualMachine := range virtualMachines {
		if _, ok := virtualMachine.Tags["foo"]; !ok {
			return errors.New("foo was not set in tags for Microsoft.Compute/virtualMachines")
		}

		if _, ok := virtualMachine.Tags["some_tag_1"]; !ok {
			return errors.New("some_tag_1 was not set in tags for Microsoft.Compute/virtualMachines")
		}

		if _, ok := virtualMachine.Tags["some_tag_2"]; !ok {
			return errors.New("some_tag_2 was not set in tags for Microsoft.Compute/virtualMachines")
		}

		if *virtualMachine.Tags["some_tag_1"] != "some_tag_1_value" ||
			*virtualMachine.Tags["some_tag_2"] != "some_tag_2_value" ||
			*virtualMachine.Tags["foo"] != "bar" {
			return errors.New("tags for Microsoft.Compute/virtualMachines not correct")
		}
	}
	return nil
}

func checkNetworkInterfacesTags(ctx context.Context, api cloud.AzureClientAPI, managedResourceGroup string) error {
	disks, err := api.ResourcesGetByType(ctx, "Microsoft.Network/networkInterfaces", managedResourceGroup)
	if err != nil {
		return err
	}

	for _, disk := range disks {
		if _, ok := disk.Tags["foo"]; !ok {
			return errors.New("foo was not set in tags for Microsoft.Network/networkInterfaces")
		}

		if _, ok := disk.Tags["some_tag_2"]; !ok {
			return errors.New("some_tag_2 was not set in tags for Microsoft.Network/networkInterfaces")
		}

		if _, ok := disk.Tags["some_tag_3"]; !ok {
			return errors.New("some_tag_3 was not set in tags for Microsoft.Network/networkInterfaces")
		}

		if _, ok := disk.Tags["some_tag_4"]; !ok {
			return errors.New("some_tag_4 was not set in tags for Microsoft.Network/networkInterfaces")
		}

		if *disk.Tags["some_tag_3"] != "some_tag_3_value" ||
			*disk.Tags["some_tag_4"] != "some_tag_4_value" ||
			*disk.Tags["some_tag_2"] != "some_tag_2_value" ||
			*disk.Tags["foo"] != "bar" {
			return errors.New("tags for Microsoft.Network/networkInterfaces not correct")
		}
	}
	return nil
}

// we want to verify that the global tags and the tags for specific resource types and are set
func testAccArrayTags(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Managed Application not found: %s", resourceName)
		}

		ctx := context.Background()
		azureClient, diags := testAccProvider.Meta().(*CbsService).azureClientService(ctx)
		if diags.HasError() {
			return fmt.Errorf("err: %+v", diags)
		}

		if err := checkVirtualMachineTags(ctx, azureClient, rs.Primary.Attributes["managed_resource_group_name"]); err != nil {
			return err
		}

		if err := checkNetworkInterfacesTags(ctx, azureClient, rs.Primary.Attributes["managed_resource_group_name"]); err != nil {
			return err
		}

		return nil
	}
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
