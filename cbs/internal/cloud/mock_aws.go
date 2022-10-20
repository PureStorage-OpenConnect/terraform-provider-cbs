//go:build mock

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

package cloud

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/internal/mockdb"
)

// Hash the stack name, use it for stack id
func awsMockStackID(stackName string) string {
	nh := sha1.Sum([]byte(stackName))
	return fmt.Sprintf("arn:aws:cloudformation:mock:%012d:stack/%s/%x-%x-%x-%x-%x",
		binary.BigEndian.Uint32(nh[16:20]), stackName, nh[0:4], nh[4:6], nh[6:8], nh[8:10], nh[10:16])
}

func buildAWSClient(region string) (AWSClientAPI, error) {
	return &mockAWSClient{}, nil
}

type mockAWSClient struct{}

func (m *mockAWSClient) CreateStack(input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error) {
	stackId := awsMockStackID(*input.StackName)

	if mockdb.AWSGetStack(*input.StackName) != nil {
		return nil, fmt.Errorf("Stack with name `%s` already exists", *input.StackName)
	}
	if mockdb.AWSGetStack(stackId) != nil {
		return nil, fmt.Errorf("Stack with StackId `%s` (from name `%s`) already exists", stackId, *input.StackName)
	}

	outputs := make([]*cloudformation.Output, 0)
	for _, key := range []string{
		"ReplicationEndpointCT0", "ReplicationEndpointCT1",
		"iSCSIEndpointCT0", "iSCSIEndpointCT1",
		"GUIEndpoint", "ResumeLambda",
	} {
		outputs = append(outputs, &cloudformation.Output{
			OutputKey:   aws.String(key),
			OutputValue: aws.String(fmt.Sprintf("%s mocked value", key)),
		})
	}

	outputs = append(outputs, &cloudformation.Output{
		OutputKey:   aws.String("AdminSecretsManagerArn"),
		OutputValue: aws.String(awsMockSecretsManagerArn()),
	})

	host := awsMockManagementEndpoint()
	outputs = append(outputs, &cloudformation.Output{
		OutputKey:   aws.String("ManagementEndpoint"),
		OutputValue: aws.String(host),
	})

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

	mockdb.AWSSetStack(stackId, host, stack)
	mockdb.AWSSetStack(*input.StackName, host, stack)
	return &cloudformation.CreateStackOutput{StackId: &stackId}, nil
}

func (m *mockAWSClient) DescribeStacks(input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {
	stack := mockdb.AWSGetStack(*input.StackName)
	if stack == nil {
		return nil, fmt.Errorf("Stack with name or Id `%s` does not exist", *input.StackName)
	}
	return &cloudformation.DescribeStacksOutput{Stacks: []*cloudformation.Stack{stack}}, nil
}

func (m *mockAWSClient) DeleteStack(input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
	stack := mockdb.AWSGetStack(*input.StackName)
	if stack == nil {
		return nil, fmt.Errorf("Stack with name or Id `%s` does not exist", *input.StackName)
	}
	mockdb.AWSDelStack(*stack.StackId)
	mockdb.AWSDelStack(*stack.StackName)

	stack.StackStatus = aws.String("DELETE_COMPLETE")
	mockdb.AWSSetStack(*stack.StackId, "", *stack)
	mockdb.AWSSetStack(*stack.StackName, "", *stack)

	return &cloudformation.DeleteStackOutput{}, nil
}

func (m *mockAWSClient) WaitUntilStackCreateCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput) error {
	return nil
}

func (m *mockAWSClient) WaitUntilStackDeleteCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput) error {
	return nil
}

func (m *mockAWSClient) GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	mockUserArn := "arn:aws:iam::123456789012:user/mock"
	return &sts.GetCallerIdentityOutput{Arn: &mockUserArn}, nil
}

func (m *mockAWSClient) PutSecretValue(input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
	mockdb.AWSSetSecret(*input.SecretId, *input.SecretString)
	return &secretsmanager.PutSecretValueOutput{ARN: input.SecretId}, nil
}

func awsMockSecretsManagerArn() string {
	b := make([]byte, 19)
	rand.Read(b)
	return fmt.Sprintf("arn:aws:secretsmanager:mock:123456789012:secret:pure/cbs/%x-%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:16], b[16:19])
}

func (m *mockAWSClient) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	secret := mockdb.AWSGetSecret(*input.SecretId)
	if secret == nil {
		return nil, fmt.Errorf("Secret with Id `%s` does not exist", *input.SecretId)
	}
	return &secretsmanager.GetSecretValueOutput{SecretString: secret}, nil
}

func (m *mockAWSClient) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	return &ec2.DescribeSubnetsOutput{
		Subnets: []*ec2.Subnet{
			&ec2.Subnet{
				AvailabilityZone: aws.String("mock-az"),
			},
		},
	}, nil
}

func (m *mockAWSClient) ValidateTemplate(input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error) {
	return &cloudformation.ValidateTemplateOutput{
		Parameters: []*cloudformation.TemplateParameter{
			&cloudformation.TemplateParameter{
				ParameterKey: aws.String("AvailabilityZone"),
			},
		},
	}, nil
}

func (m *mockAWSClient) NewFAClient(host string, adminSecretsManagerArn string) (array.FAClientAPI, error) {
	return &array.MockFAClient{Host: host, Kind: array.FAClientKindAWS}, nil
}

func awsMockManagementEndpoint() string {
	bg := big.NewInt(255)
	blocks := []string{}
	for i := 0; i < 4; i++ {
		number, _ := rand.Int(rand.Reader, bg)
		blocks = append(blocks, strconv.Itoa(int(number.Int64())))
	}
	return strings.Join(blocks, ".")
}
