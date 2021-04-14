// +build mock

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


	This file contains a mock implementation of the AWS cloud service
	calls.

*/

// TODO: Would be nice to have AZ, AWS prefixes on things

package cbs

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

// Hash the stack name, use it for stack id
func awsMockStackID(stackName string) string {
	nh := sha1.Sum([]byte(stackName))
	return fmt.Sprintf("arn:aws:cloudformation:mock:%012d:stack/%s/%x-%x-%x-%x-%x",
		binary.BigEndian.Uint32(nh[16:20]), stackName, nh[0:4], nh[4:6], nh[6:8], nh[8:10], nh[10:16])
}

func buildAWSSession(region string) (cloudformationAPI, diag.Diagnostics) {
	var diags = buildAWSSessionPreCheck(region)
	if diags != nil {
		return nil, diags
	}
	return &mockCloudformation{}, diags
}

type mockCloudformation struct{}

func (m *mockCloudformation) CreateStack(input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error) {
	stackId := awsMockStackID(*input.StackName)

	if mockDbAWSGetStack(*input.StackName) != nil {
		return nil, fmt.Errorf("Stack with name `%s` already exists", *input.StackName)
	}
	if mockDbAWSGetStack(stackId) != nil {
		return nil, fmt.Errorf("Stack with StackId `%s` (from name `%s`) already exists", stackId, *input.StackName)
	}

	outputs := make([]*cloudformation.Output, 0)
	for _, key := range []string{
		"ReplicationEndpointCT0", "ReplicationEndpointCT1",
		"iSCSIEndpointCT0", "iSCSIEndpointCT1",
		"GUIEndpoint", "ManagementEndpoint", "ResumeLambda",
	} {
		outputs = append(outputs, &cloudformation.Output{
			OutputKey:   aws.String(key),
			OutputValue: aws.String(fmt.Sprintf("%s mocked value", key)),
		})
	}

	stack := cloudformation.Stack{
		Capabilities:    input.Capabilities,
		DisableRollback: input.DisableRollback,
		Parameters:      input.Parameters,
		RoleARN:         input.RoleARN,
		StackName:       input.StackName,
		StackId:         &stackId,
		StackStatus:     aws.String("CREATE_COMPLETE"),
		Outputs:         outputs,
		Tags:            input.Tags,
	}

	mockDbAWSSetStack(stackId, stack)
	mockDbAWSSetStack(*input.StackName, stack)
	return &cloudformation.CreateStackOutput{StackId: &stackId}, nil
}

func (m *mockCloudformation) DescribeStacks(input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {
	stack := mockDbAWSGetStack(*input.StackName)
	if stack == nil {
		return nil, fmt.Errorf("Stack with name or Id `%s` does not exist", *input.StackName)
	}
	return &cloudformation.DescribeStacksOutput{Stacks: []*cloudformation.Stack{stack}}, nil
}

func (m *mockCloudformation) DeleteStack(input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
	stack := mockDbAWSGetStack(*input.StackName)
	if stack == nil {
		return nil, fmt.Errorf("Stack with name or Id `%s` does not exist", *input.StackName)
	}
	mockDbAWSDelStack(*stack.StackId)
	mockDbAWSDelStack(*stack.StackName)

	stack.StackStatus = aws.String("DELETE_COMPLETE")
	mockDbAWSSetStack(*stack.StackId, *stack)
	mockDbAWSSetStack(*stack.StackName, *stack)

	return &cloudformation.DeleteStackOutput{}, nil
}

func (m *mockCloudformation) WaitUntilStackCreateCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput, opts ...request.WaiterOption) error {
	return nil
}

func (m *mockCloudformation) WaitUntilStackDeleteCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput, opts ...request.WaiterOption) error {
	return nil
}
