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

package auth

import (
	"encoding/json"
)

func NewBootstrapService() Bootstrapper {
	return &mockBootstrapService{}
}

type mockBootstrapService struct{}

func (b *mockBootstrapService) GenerateSecretPayload(host string, pureuserPrivateKey []byte) ([]byte, error) {
	mockPayload := SecretPayload{
		UserName:       "mockUser",
		Issuer:         "mockIssuer",
		ClientID:       "mockClientID",
		KeyID:          "mockKeyID",
		RestPrivateKey: "mockRestPrivateKey",
	}
	credentials, _ := json.MarshalIndent(mockPayload, "", " ")

	return credentials, nil
}
