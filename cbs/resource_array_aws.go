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
	"log"
	"strings"
	"time"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/cloud"
	mapset "github.com/deckarep/golang-set"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

var cftParams = []interface{}{
	"ArrayName",
	"LogSenderDomain",
	"PurityInstanceType",
	"LicenseKey",
	"KeyName",
	"SystemSubnet",
	"ReplicationSubnet",
	"ISCSISubnet",
	"ManagementSubnet",
	"ReplicationSecurityGroup",
	"ISCSISecurityGroup",
	"ManagementSecurityGroup",
}

// By default, the terraform parameter is just the CloudFormation template converted to snake case.
// This map holds exceptions to that rule.
var renamedParams = map[string]string{
	"PurityInstanceType": "array_model",
	"KeyName":            "pureuser_key_pair_name",
}

var templateOutputs = []interface{}{
	"ReplicationEndpointCT0",
	"iSCSIEndpointCT0",
	"ReplicationEndpointCT1",
	"iSCSIEndpointCT1",
	"GUIEndpoint",
	"ManagementEndpoint",
	"ResumeLambda",
}

func resourceArrayAWS() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceArrayAWSCreate,
		ReadContext:   resourceArrayAWSRead,
		UpdateContext: resourceArrayAWSUpdate,
		DeleteContext: resourceArrayAWSDelete,
		Schema: map[string]*schema.Schema{
			"deployment_template_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"array_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"deployment_role_arn": {
				Type:     schema.TypeString,
				Required: true,
			},
			"log_sender_domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"alert_recipients": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"array_model": {
				Type:     schema.TypeString,
				Required: true,
			},
			"license_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pureuser_key_pair_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"pureuser_private_key_path": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"pureuser_private_key_path", "pureuser_private_key"},
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"pureuser_private_key": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"pureuser_private_key_path", "pureuser_private_key"},
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"system_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replication_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iscsi_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replication_security_group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iscsi_security_group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"management_security_group": {
				Type:     schema.TypeString,
				Required: true,
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"stack_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_endpoint_ct0": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iscsi_endpoint_ct0": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_endpoint_ct1": {

				Type:     schema.TypeString,
				Computed: true,
			},
			"iscsi_endpoint_ct1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gui_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resume_lambda": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			// Timeouts are slightly longer than the timeouts on the waiters to ensure the waiter
			// timeouts trigger first.
			Create: schema.DefaultTimeout(130 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(50 * time.Minute),
		},
	}
}

func resourceArrayAWSCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsClient, diags := m.(*CbsService).awsClientService()
	if diags.HasError() {
		return diags
	}

	arrayName := d.Get("array_name").(string)
	deploymentArn := d.Get("deployment_role_arn").(string)
	templateURL := d.Get("deployment_template_url").(string)

	var params []*cloudformation.Parameter

	// From 6.3.5 the AvailabilityZone parameter is required for the CFT. This parameter can be computed
	// on the fly so users does not have to provide it. In order to backward compatibility with older CBS
	// releases we want to pass AZ only if the template requires it.
	template, err := awsClient.ValidateTemplate(&cloudformation.ValidateTemplateInput{
		TemplateURL: aws.String(templateURL),
	})
	if err != nil {
		return diag.FromErr(err)
	}

	for _, templateParam := range template.Parameters {
		// Only calculate and pass the AvailabilityZone parameter if required by the CFT
		if *templateParam.ParameterKey == "AvailabilityZone" {
			// Derive the availability_zone from the provided subnet parameter
			// All of the subnet ids provided by users must be from the same AZ so we picked system_subnet
			subnets, subnetErr := awsClient.DescribeSubnets(&ec2.DescribeSubnetsInput{
				SubnetIds: aws.StringSlice([]string{d.Get("system_subnet").(string)}),
			})
			if subnetErr != nil {
				return diag.FromErr(subnetErr)
			}

			params = append(params, &cloudformation.Parameter{
				ParameterKey:   aws.String("AvailabilityZone"),
				ParameterValue: subnets.Subnets[0].AvailabilityZone,
			})
			break
		}
	}

	if v, ok := d.GetOk("alert_recipients"); ok {
		recips := convertToStringSlice(v.([]interface{}))
		params = append(params, &cloudformation.Parameter{
			ParameterKey:   aws.String("AlertRecipients"),
			ParameterValue: aws.String(strings.Join(recips, ",")),
		})
	}

	for _, cftParam := range cftParams {
		tfParam := templateToTFParam(cftParam.(string), renamedParams)
		val := d.Get(tfParam).(string)

		params = append(params, &cloudformation.Parameter{
			ParameterKey:   aws.String(cftParam.(string)),
			ParameterValue: aws.String(val),
		})
	}

	// Get Caller Identity on behalf of provider role
	res, err := awsClient.GetCallerIdentity(&sts.GetCallerIdentityInput{})
	if err != nil {
		return diag.FromErr(err)
	}
	params = append(params, &cloudformation.Parameter{
		ParameterKey:   aws.String("AdminArn"),
		ParameterValue: aws.String(*res.Arn),
	})

	var tags []*cloudformation.Tag
	if v, ok := d.GetOk("tags"); ok {
		tags = expandTags(v.(map[string]interface{}))
	}

	input := &cloudformation.CreateStackInput{
		Capabilities:    aws.StringSlice([]string{"CAPABILITY_IAM"}),
		DisableRollback: aws.Bool(true),
		RoleARN:         aws.String(deploymentArn),
		StackName:       aws.String(arrayName),
		Tags:            tags,
		TemplateURL:     aws.String(templateURL),
		Parameters:      params,
	}

	output, stackErr := awsClient.CreateStack(input)
	if stackErr != nil {
		return diag.FromErr(stackErr)
	}

	d.SetId(*output.StackId)

	waiterInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(*output.StackId),
	}

	if err := awsClient.WaitUntilStackCreateCompleteWithContext(ctx, waiterInput); err != nil {
		return diag.FromErr(err)
	}

	diags = resourceArrayAWSRead(ctx, d, m)
	if diags.HasError() {
		return diags
	}

	adminSecretsManagerArn, err := getSecretsManagerArn(awsClient, waiterInput)
	if err != nil {
		return diag.FromErr(err)
	}

	// Bootstrap the array
	credentials, err := generateSecretPayload(ctx, d)
	if err != nil {
		return diag.Errorf("failed to bootstrap the CloudFormation stack with Id %s. Please contact "+
			"Pure Storage support to deactivate the instance: %+v.", d.Id(), err)
	}

	// Set secret payload to secrets manager
	secretInput := &secretsmanager.PutSecretValueInput{
		SecretId:     aws.String(string(adminSecretsManagerArn)),
		SecretString: aws.String(string(credentials)),
	}
	if _, err := awsClient.PutSecretValue(secretInput); err != nil {
		return diag.Errorf("failed to store credentials on the CloudFormation stack with Id %s. Please contact "+
			"Pure Storage support to deactivate the instance: %+v.", d.Id(), err)
	}

	return nil
}

func resourceArrayAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsClient, diags := m.(*CbsService).awsClientService()
	if diags.HasError() {
		return diags
	}

	input := &cloudformation.DescribeStacksInput{
		StackName: aws.String(d.Id()),
	}

	res, err := awsClient.DescribeStacks(input)
	if err != nil {
		return diag.FromErr(err)
	}

	stacks := res.Stacks
	if len(stacks) < 1 {
		log.Printf("[WARN] No CloudFormation stack found with Id %s, removing resource from state.", d.Id())
		d.SetId("")
		return nil
	}

	stack := stacks[0]
	if *stack.StackStatus == cloudformation.StackStatusDeleteComplete {
		log.Printf("[WARN] CloudFormation stack with Id %s has already been deleted, removing resource from state.", d.Id())
		d.SetId("")
		return nil
	}

	d.Set("stack_name", stack.StackName)

	params := stack.Parameters

	cftParamSet := mapset.NewSetFromSlice(cftParams)
	for _, param := range params {
		if *param.ParameterKey == "AlertRecipients" {
			recips := strings.Split(*param.ParameterValue, ",")
			d.Set("alert_recipients", recips)
		}
		if cftParamSet.Contains(*param.ParameterKey) {
			tfParam := templateToTFParam(*param.ParameterKey, renamedParams)
			d.Set(tfParam, param.ParameterValue)
		}
	}

	outputs := stack.Outputs
	templateOutputSet := mapset.NewSetFromSlice(templateOutputs)
	for _, output := range outputs {
		if templateOutputSet.Contains(*output.OutputKey) {
			d.Set(toSnake(*output.OutputKey), output.OutputValue)
		}
	}

	if err := d.Set("tags", flattenTags(stack.Tags)); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func resourceArrayAWSUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Perform a read to overwrite any attempted updates
	err := resourceArrayAWSRead(ctx, d, m)
	if err != nil {
		return err
	}
	return diag.Errorf("Updates are not supported.")
}

func resourceArrayAWSDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	awsClient, diags := m.(*CbsService).awsClientService()
	if diags.HasError() {
		return diags
	}

	waiterInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(d.Id()),
	}

	adminSecretsManagerArn, err := getSecretsManagerArn(awsClient, waiterInput)
	if err != nil {
		return diag.FromErr(err)
	}

	faClient, err := awsClient.NewFAClient(ctx, d.Get("management_endpoint").(string), adminSecretsManagerArn)
	if err != nil {
		return diag.Errorf("failed to create FA client on the CloudFormation stack with Id %s. Please contact "+
			"Pure Storage support to deactivate the instance: %+v.", d.Id(), err)
	}

	if err := faClient.Deactivate(ctx); err != nil {
		return diag.Errorf("failed to deactivate the instance with Id %s. Please contact "+
			"Pure Storage support: %+v.", d.Id(), err)
	}

	if err := awsClient.WaitUntilStackDeleteCompleteWithContext(ctx, waiterInput); err != nil {
		return diag.FromErr(err)
	}

	return nil
}

func expandTags(m map[string]interface{}) []*cloudformation.Tag {
	var tags []*cloudformation.Tag
	for k, v := range m {
		tags = append(tags, &cloudformation.Tag{
			Key:   aws.String(k),
			Value: aws.String(v.(string)),
		})
	}
	return tags
}

func flattenTags(tags []*cloudformation.Tag) map[string]string {
	m := make(map[string]string)
	for _, t := range tags {
		m[aws.StringValue(t.Key)] = aws.StringValue(t.Value)
	}
	return m
}

func getSecretsManagerArn(awsClient cloud.AWSClientAPI, waiterInput *cloudformation.DescribeStacksInput) (string, error) {
	stacksOutput, err := awsClient.DescribeStacks(waiterInput)
	if err != nil {
		return "", err
	}
	outputs := stacksOutput.Stacks[0].Outputs
	var adminSecretsManagerArn string
	for _, output := range outputs {
		if *output.OutputKey == "AdminSecretsManagerArn" {
			adminSecretsManagerArn = *output.OutputValue
			break
		}
	}
	return adminSecretsManagerArn, nil
}
