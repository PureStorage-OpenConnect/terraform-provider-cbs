//go:build !mock
// +build !mock

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

package array

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	bootstrapauth "github.com/PureStorage-OpenConnect/terraform-provider-cbs/auth"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/client/arrays"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/auth"
)

const defaultTimeout = 30 * time.Second

type faClient struct {
	*client.Flasharray
	host string
}

func buildFAClient(host string, secretPayload string) (FAClientAPI, error) {
	credentials := bootstrapauth.SecretPayload{}
	if err := json.Unmarshal([]byte(secretPayload), &credentials); err != nil {
		return nil, err
	}
	tokenConfig := auth.TokenConfig{
		Issuer:     credentials.Issuer,
		ClientID:   credentials.ClientID,
		KeyID:      credentials.KeyID,
		User:       credentials.UserName,
		PrivateKey: credentials.RestPrivateKey,
	}

	lowerClient, err := faclient.New(host, tokenConfig)
	if err != nil {
		return nil, err
	}
	return &faClient{lowerClient, host}, nil
}

func (c *faClient) Deactivate(ctx context.Context) error {
	tempCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	postFactoryResetTokenParams := arrays.NewPostAPI24ArraysFactoryResetTokenParamsWithContext(tempCtx)
	resp, err := c.Arrays.PostAPI24ArraysFactoryResetToken(postFactoryResetTokenParams)
	if err != nil {
		return fmt.Errorf("failed to create factory reset token on management endpoint %q : %+v", c.host, err)
	}

	token := resp.GetPayload().Items[0].ArrayFactoryResetTokenOAIGenAllOf1.Token
	eradicate := true
	deleteArrayParams := arrays.NewDeleteAPI24ArraysParamsWithContext(tempCtx).WithEradicateAllData(&eradicate).WithFactoryResetToken(&token)
	if _, err = c.Arrays.DeleteAPI24Arrays(deleteArrayParams); err != nil {
		return fmt.Errorf("failed to deactivate management endpoint %q : %+v", c.host, err)
	}
	return nil
}
