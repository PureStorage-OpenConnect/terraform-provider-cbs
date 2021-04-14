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

	mapset "github.com/deckarep/golang-set"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},
	}
}

func resourceArrayAWSCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cftSvc, diags := m.(*CbsService).CloudFormationService()
	if diags.HasError() {
		return diags
	}

	arrayName := d.Get("array_name").(string)
	deploymentArn := d.Get("deployment_role_arn").(string)
	templateURL := d.Get("deployment_template_url").(string)

	var params []*cloudformation.Parameter

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

	output, err := cftSvc.CreateStack(input)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(*output.StackId)

	waiterInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(*output.StackId),
	}

	if err := cftSvc.WaitUntilStackCreateCompleteWithContext(ctx, waiterInput,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
		request.WithWaiterMaxAttempts(240),
	); err != nil {
		return diag.FromErr(err)
	}

	return resourceArrayAWSRead(ctx, d, m)
}

func resourceArrayAWSRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	cftSvc, diags := m.(*CbsService).CloudFormationService()
	if diags.HasError() {
		return diags
	}

	input := &cloudformation.DescribeStacksInput{
		StackName: aws.String(d.Id()),
	}

	res, err := cftSvc.DescribeStacks(input)
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
	cftSvc, diags := m.(*CbsService).CloudFormationService()
	if diags.HasError() {
		return diags
	}

	input := &cloudformation.DeleteStackInput{
		RoleARN:   aws.String(d.Get("deployment_role_arn").(string)),
		StackName: aws.String(d.Id()),
	}

	if _, err := cftSvc.DeleteStack(input); err != nil {
		return diag.FromErr(err)
	}

	waiterInput := &cloudformation.DescribeStacksInput{
		StackName: aws.String(d.Id()),
	}

	if err := cftSvc.WaitUntilStackDeleteCompleteWithContext(ctx, waiterInput,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
		request.WithWaiterMaxAttempts(20),
	); err != nil {
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
