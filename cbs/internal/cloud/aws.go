//go:build !mock

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

/*

	This file contains wrappers for "real" cloud service calls, as opposed to alternative
	implementations that are mocked.  Some of these wrappers also do a small amount of
	"flattening" in order to simplify some of the cloud APIs.  Simplifications includes
	reducing the number of objects, and putting multiple calls together.  This helps serve
	as a single file for tracking all cloud accesses, and it also makes it clearer what
	interfaces need to be mocked.

*/

package cloud

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array"
)

type awsClient struct {
	cloudFormation *cloudformation.CloudFormation
	sts            *sts.STS
	secretsManager *secretsmanager.SecretsManager
	ec2Service     *ec2.EC2
}

func buildAWSClient(region string) (AWSClientAPI, error) {
	var awsClient awsClient
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, err
	}

	awsClient.cloudFormation = cloudformation.New(sess)
	awsClient.sts = sts.New(sess)
	awsClient.secretsManager = secretsmanager.New(sess)
	awsClient.ec2Service = ec2.New(sess)
	return &awsClient, nil
}

func (client *awsClient) CreateStack(input *cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error) {
	return client.cloudFormation.CreateStack(input)
}

func (client *awsClient) DescribeStacks(input *cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error) {
	return client.cloudFormation.DescribeStacks(input)
}

func (client *awsClient) DeleteStack(input *cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error) {
	return client.cloudFormation.DeleteStack(input)
}

func (client *awsClient) WaitUntilStackCreateCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput) error {
	return client.cloudFormation.WaitUntilStackCreateCompleteWithContext(ctx, input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
		request.WithWaiterMaxAttempts(240),
	)
}

func (client *awsClient) WaitUntilStackDeleteCompleteWithContext(ctx aws.Context, input *cloudformation.DescribeStacksInput) error {
	return client.cloudFormation.WaitUntilStackDeleteCompleteWithContext(ctx, input,
		request.WithWaiterDelay(request.ConstantWaiterDelay(30*time.Second)),
		request.WithWaiterMaxAttempts(60),
	)
}

func (client *awsClient) GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	return client.sts.GetCallerIdentity(input)
}

func (client *awsClient) PutSecretValue(input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
	return client.secretsManager.PutSecretValue(input)
}

func (client *awsClient) GetSecretValue(input *secretsmanager.GetSecretValueInput) (*secretsmanager.GetSecretValueOutput, error) {
	return client.secretsManager.GetSecretValue(input)
}

func (client *awsClient) DescribeSubnets(input *ec2.DescribeSubnetsInput) (*ec2.DescribeSubnetsOutput, error) {
	return client.ec2Service.DescribeSubnets(input)
}

func (client *awsClient) ValidateTemplate(input *cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error) {
	return client.cloudFormation.ValidateTemplate(input)
}

func (client *awsClient) NewFAClient(host string, adminSecretsManagerArn string) (array.FAClientAPI, error) {
	secretInput := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(string(adminSecretsManagerArn)),
	}
	result, err := client.GetSecretValue(secretInput)
	if err != nil {
		return nil, err
	}
	return array.NewFAClient(host, *result.SecretString)
}
