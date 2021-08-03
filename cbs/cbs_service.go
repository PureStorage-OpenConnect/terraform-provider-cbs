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
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/cloud"
)

type CbsService struct {
	AWSClient    cloud.AWSClientAPI
	AzureClient  cloud.AzureClientAPI
	awsRegionStr string
	azureConfig  cloud.AzureConfig
}

func (m *CbsService) awsClientService() (cloud.AWSClientAPI, diag.Diagnostics) {
	if m.AWSClient == nil {
		diags := buildAWSClientPreCheck(m.awsRegionStr)
		if diags != nil {
			return nil, diags
		}
		awsClient, err := cloud.NewAWSClient(m.awsRegionStr)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		m.AWSClient = awsClient
	}

	return m.AWSClient, nil
}

func (m *CbsService) azureClientService() (cloud.AzureClientAPI, diag.Diagnostics) {
	if m.AzureClient == nil {
		azureClient, err := cloud.NewAzureClient(m.azureConfig)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		m.AzureClient = azureClient
	}

	return m.AzureClient, nil
}

func buildAWSClientPreCheck(region string) diag.Diagnostics {
	var diags diag.Diagnostics
	if region == "" {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary: fmt.Sprintf("No AWS region specified. The AWS region must be provided either in "+
				"the provider configuration block or with the %s environment variable.", awsRegionVar),
		})
		return diags
	}
	return nil
}
