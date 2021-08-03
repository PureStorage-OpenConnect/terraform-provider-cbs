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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sync"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

type cbsAwsParams struct {
	DeploymentRoleArn      string `json:"deployment_role_arn"`
	LicenseKey             string `json:"license_key"`
	PureuserKeyPairName    string `json:"pureuser_key_pair_name"`
	PureuserPrivateKeyPath string `json:"pureuser_private_key_path"`
	PureuserPrivateKey     string `json:"pureuser_private_key"`
	Subnet                 string `json:"subnet"`
	SecurityGroup          string `json:"security_group"`
}

var awsParams cbsAwsParams
var awsParamsConfigure sync.Once
var awsRegionConfigure sync.Once

const testDefaultRegion = "us-west-2"
const awsParamsPathVar = "TEST_ACC_AWS_PARAMS_PATH"

// 6.1.7
const deploymentTemplateURL = "https://s3.amazonaws.com/awsmp-fulfillment-cf-templates-prod/4ea2905b-7939-4ee0-a521-d5c2fcb41214.6b728728-d8fa-4eb7-b92d-22d9aee3684c.template"

// Basic test of an AWS array. Spins up a new instance, makes sure it exists, and tests that
// the parameter and output values are all correctly stored in the Terraform state. Then try
// an update, confirm that it failed, and make sure the state is still correct
func TestAccArrayAws_basic(t *testing.T) {

	arrayName := acctest.RandomWithPrefix("tf-test-array")
	senderDomain := "example.com"
	senderDomain2 := "example-invalid-update.com"
	loadAccAwsParams(t)
	resourceName := "cbs_array_aws.test_array_aws"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccAWSPreCheck(t) },
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccBasicConfig(arrayName, senderDomain),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAwsExists(resourceName),
					testAccCheckAllAttrs(resourceName, arrayName, senderDomain),
					//resource.TestCheckResourceAttr(resourceName, "pureuser_private_key_path", awsParams.PureuserPrivateKeyPath),
				),
			},
			{
				Config:      testAccBasicConfig(arrayName, senderDomain2),
				ExpectError: regexp.MustCompile("Updates are not supported."),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAwsExists(resourceName),
					testAccCheckAllAttrs(resourceName, arrayName, senderDomain),
				),
			},
		},
	})
}

// Create an array with tags
func TestAccArrayAws_tags(t *testing.T) {

	arrayName := acctest.RandomWithPrefix("tf-test-array")
	senderDomain := "example.com"
	loadAccAwsParams(t)
	resourceName := "cbs_array_aws.test_array_aws"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccAWSPreCheck(t) },
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTagsConfig(arrayName),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAwsExists(resourceName),
					testAccCheckAllAttrs(resourceName, arrayName, senderDomain),
					//resource.TestCheckResourceAttr(resourceName, "pureuser_private_key_path", awsParams.PureuserPrivateKeyPath),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.foo", "bar"),
					resource.TestCheckResourceAttr(resourceName, "tags.test", "value"),
				),
			},
		},
	})
}

func TestAccArrayAws_pureuserPrivateKey(t *testing.T) {
	t.Skip("Deactivation in AWS is disabled")

	arrayName := acctest.RandomWithPrefix("tf-test-array")
	senderDomain := "example.com"
	loadAccAwsParams(t)
	resourceName := "cbs_array_aws.test_array_aws"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { testAccAWSPreCheck(t) },
		ProviderFactories: testAccProvidersFactory,
		CheckDestroy:      testAccCheckDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccPureuserPrivateKeyConfig(arrayName),
				Check: resource.ComposeTestCheckFunc(
					testAccArrayAwsExists(resourceName),
					testAccCheckAllAttrs(resourceName, arrayName, senderDomain),
					resource.TestCheckNoResourceAttr(resourceName, "pureuser_private_key_path"),
					resource.TestCheckResourceAttrSet(resourceName, "pureuser_private_key"),
				),
			},
		},
	})
}

func testAccBasicConfig(name string, senderDomain string) string {
	return fmt.Sprintf(`
	resource "cbs_array_aws" "test_array_aws" {

		array_name = "%[1]s"

		log_sender_domain = "%[2]s"

		deployment_template_url = "%[3]s"

		deployment_role_arn = "%[4]s"

		alert_recipients = ["user@example.com"]
		array_model = "V10AR1"
		license_key = "%[5]s"

		pureuser_key_pair_name = "%[6]s"

		system_subnet = "%[8]s"
		replication_subnet = "%[8]s"
		iscsi_subnet = "%[8]s"
		management_subnet = "%[8]s"

		replication_security_group = "%[9]s"
		iscsi_security_group = "%[9]s"
		management_security_group = "%[9]s"
	}`, name, senderDomain, deploymentTemplateURL, awsParams.DeploymentRoleArn, awsParams.LicenseKey,
		awsParams.PureuserKeyPairName, awsParams.PureuserPrivateKeyPath, awsParams.Subnet, awsParams.SecurityGroup)
}

func testAccTagsConfig(name string) string {
	return fmt.Sprintf(`
	resource "cbs_array_aws" "test_array_aws" {

		array_name = "%[1]s"

		log_sender_domain = "example.com"

		deployment_template_url = "%[2]s"

		deployment_role_arn = "%[3]s"

		alert_recipients = ["user@example.com"]
		array_model = "V10AR1"
		license_key = "%[4]s"

		pureuser_key_pair_name = "%[5]s"

		system_subnet = "%[7]s"
		replication_subnet = "%[7]s"
		iscsi_subnet = "%[7]s"
		management_subnet = "%[7]s"

		replication_security_group = "%[8]s"
		iscsi_security_group = "%[8]s"
		management_security_group = "%[8]s"

		tags = {
			foo = "bar"
			test = "value"
		}
	}`, name, deploymentTemplateURL, awsParams.DeploymentRoleArn, awsParams.LicenseKey,
		awsParams.PureuserKeyPairName, awsParams.PureuserPrivateKeyPath, awsParams.Subnet, awsParams.SecurityGroup)
}

func testAccPureuserPrivateKeyConfig(name string) string {
	return fmt.Sprintf(`
	resource "cbs_array_aws" "test_array_aws" {

		array_name = "%[1]s"

		log_sender_domain = "example.com"

		deployment_template_url = "%[2]s"

		deployment_role_arn = "%[3]s"

		alert_recipients = ["user@example.com"]
		array_model = "V10AR1"
		license_key = "%[4]s"

		pureuser_key_pair_name = "%[5]s"

		system_subnet = "%[7]s"
		replication_subnet = "%[7]s"
		iscsi_subnet = "%[7]s"
		management_subnet = "%[7]s"

		replication_security_group = "%[8]s"
		iscsi_security_group = "%[8]s"
		management_security_group = "%[8]s"
	}`, name, deploymentTemplateURL, awsParams.DeploymentRoleArn, awsParams.LicenseKey,
		awsParams.PureuserKeyPairName, awsParams.PureuserPrivateKey, awsParams.Subnet, awsParams.SecurityGroup)
}

func testAccAWSPreCheck(t *testing.T) {
	awsRegionConfigure.Do(func() {
		region := awsTestRegion()
		os.Setenv(awsRegionVar, region)
	})
}

func awsTestRegion() string {
	if region := os.Getenv("AWS_TEST_REGION"); region != "" {
		return region
	}
	return testDefaultRegion
}

// Lazy load the AWS param values from the json file specified at TEST_ACC_AWS_PARAMS_PATH.
func loadAccAwsParams(t *testing.T) {
	awsParamsConfigure.Do(func() {
		cfgPath := os.Getenv(awsParamsPathVar)
		if cfgPath == "" {
			t.Fatalf("%s environment variable must be set for acceptance testing", awsParamsPathVar)
		}
		cfgData, err := ioutil.ReadFile(cfgPath)
		if err != nil {
			t.Fatalf("Error reading aws config JSON file: %s", err)
		}
		if err := json.Unmarshal(cfgData, &awsParams); err != nil {
			t.Fatalf("Error unmarshaling JSON file: %s", err)
		}
	})
}

func testAccCheckAllAttrs(resourceName string, arrayName string, senderDomain string) resource.TestCheckFunc {
	return resource.ComposeAggregateTestCheckFunc(
		resource.TestCheckResourceAttr(resourceName, "array_name", arrayName),
		resource.TestCheckResourceAttr(resourceName, "stack_name", arrayName),
		resource.TestCheckResourceAttr(resourceName, "deployment_template_url", deploymentTemplateURL),
		resource.TestCheckResourceAttr(resourceName, "deployment_role_arn", awsParams.DeploymentRoleArn),
		resource.TestCheckResourceAttr(resourceName, "log_sender_domain", senderDomain),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.#", "1"),
		resource.TestCheckResourceAttr(resourceName, "alert_recipients.0", "user@example.com"),
		resource.TestCheckResourceAttr(resourceName, "array_model", "V10AR1"),
		resource.TestCheckResourceAttr(resourceName, "license_key", awsParams.LicenseKey),
		resource.TestCheckResourceAttr(resourceName, "pureuser_key_pair_name", awsParams.PureuserKeyPairName),
		resource.TestCheckResourceAttr(resourceName, "system_subnet", awsParams.Subnet),
		resource.TestCheckResourceAttr(resourceName, "replication_subnet", awsParams.Subnet),
		resource.TestCheckResourceAttr(resourceName, "iscsi_subnet", awsParams.Subnet),
		resource.TestCheckResourceAttr(resourceName, "management_subnet", awsParams.Subnet),
		resource.TestCheckResourceAttr(resourceName, "replication_security_group", awsParams.SecurityGroup),
		resource.TestCheckResourceAttr(resourceName, "iscsi_security_group", awsParams.SecurityGroup),
		resource.TestCheckResourceAttr(resourceName, "management_security_group", awsParams.SecurityGroup),
		resource.TestCheckResourceAttrSet(resourceName, "replication_endpoint_ct0"),
		resource.TestCheckResourceAttrSet(resourceName, "iscsi_endpoint_ct0"),
		resource.TestCheckResourceAttrSet(resourceName, "replication_endpoint_ct1"),
		resource.TestCheckResourceAttrSet(resourceName, "iscsi_endpoint_ct1"),
		resource.TestCheckResourceAttrSet(resourceName, "gui_endpoint"),
		resource.TestCheckResourceAttrSet(resourceName, "management_endpoint"),
		resource.TestCheckResourceAttrSet(resourceName, "resume_lambda"),
	)
}

func testAccCheckDestroy(s *terraform.State) error {
	cftSvc, diags := testAccProvider.Meta().(*CbsService).awsClientService()
	if diags.HasError() {
		return fmt.Errorf("err: %+v", diags)
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cbs_array_aws" {
			continue
		}

		stackID := rs.Primary.ID
		input := &cloudformation.DescribeStacksInput{
			StackName: aws.String(stackID),
		}

		res, err := cftSvc.DescribeStacks(input)
		if err != nil {
			return err
		}

		stack := res.Stacks[0]
		if *stack.StackId == stackID && *stack.StackStatus != cloudformation.StackStatusDeleteComplete {
			return fmt.Errorf("Array still exists. Stack ID: %s", stackID)
		}

	}
	return nil
}

func testAccArrayAwsExists(resourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		cftSvc, diags := testAccProvider.Meta().(*CbsService).awsClientService()
		if diags.HasError() {
			return fmt.Errorf("err: %+v", diags)
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource not found: %s", resourceName)
		}
		stackName := rs.Primary.Attributes["stack_name"]

		input := &cloudformation.DescribeStacksInput{
			StackName: aws.String(stackName),
		}

		res, err := cftSvc.DescribeStacks(input)
		if err != nil {
			return err
		}

		if len(res.Stacks) == 0 || res.Stacks[0] == nil {
			return fmt.Errorf("Stack not found. Name: %s, Resource Name: %s", stackName, resourceName)
		}

		return nil
	}
}
