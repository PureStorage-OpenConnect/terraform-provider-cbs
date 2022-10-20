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

package array

import (
	"context"
	"fmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/internal/mockdb"
	"github.com/aws/aws-sdk-go/aws"
)

const (
	FAClientKindAWS = iota
	FAClientKindAzure
)

type MockFAClient struct {
	Kind int
	Host string
}

func buildFAClient(host string, secretPayload string) (FAClientAPI, error) {
	// the mock aws and mock azure clients don't actually use this, instead they build the
	// MockFAClient directly so that they can set the Kind enum.
	// TODO we should try and use this instead for the mocks
	return nil, nil
}

func (self *MockFAClient) Deactivate(ctx context.Context) error {
	if self.Kind == FAClientKindAWS {
		stack := mockdb.AWSGetHostStack(self.Host)
		if stack == nil {
			return fmt.Errorf("Stack for management endpoint address `%s` does not exist", self.Host)
		}
		mockdb.AWSDelStack(*stack.StackId)
		mockdb.AWSDelStack(*stack.StackName)

		stack.StackStatus = aws.String("DELETE_COMPLETE")
		mockdb.AWSSetStack(*stack.StackId, self.Host, *stack)
		mockdb.AWSSetStack(*stack.StackName, self.Host, *stack)
	} else if self.Kind == FAClientKindAzure {
		// Dont do anything for now, as in Azure the array doesn't delete itself on deactivation,
		// but we should follow up with some better mocks later though...
	} else {
		return fmt.Errorf("this error should not happen")
	}

	return nil
}
