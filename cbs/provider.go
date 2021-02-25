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
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/go-azure-helpers/authentication"
)

const (
	awsRegionVar = "AWS_DEFAULT_REGION"
)

const (
	azureEnvironment    = "public"
	azureSubscriptionID = "ARM_SUBSCRIPTION_ID"
	azureClientID       = "ARM_CLIENT_ID"
	azureTenantID       = "ARM_TENANT_ID"
	azureClientSecret   = "ARM_CLIENT_SECRET"
)

var awsRegionStr string
var azureBuilder *authentication.Builder

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"aws": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc(awsRegionVar, nil),
						},
					},
				},
			},

			"azure": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subscription_id": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc(azureSubscriptionID, ""),
							Description: "The Subscription ID which should be used.",
						},

						"client_id": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc(azureClientID, ""),
							Description: "The Client ID which should be used for service principal authentication.",
						},

						"tenant_id": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc(azureTenantID, ""),
							Description: "The Tenant ID which should be used for service principal authentication.",
						},

						"client_secret": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc(azureClientSecret, ""),
							Description: "The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.",
						},
					},
				},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cbs_array_aws":   resourceArrayAWS(),
			"cbs_array_azure": resourceArrayAzure(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var cbsService CbsService
	var diags diag.Diagnostics

	awsRegionStr = awsRegion(d)
	azureBuilder = azureConfig(d)

	if _, ok := d.GetOk("aws"); ok {
		cftSvc, diags := buildSession(awsRegionStr)
		if diags.HasError() {
			return nil, diags
		}
		cbsService.CloudFormation = cftSvc
	}

	if _, ok := d.GetOk("azure"); ok {
		azureClient, diags := buildAzureClient(azureBuilder)
		if diags.HasError() {
			return nil, diags
		}
		cbsService.AzureClient = azureClient
	}

	return &cbsService, diags
}

func awsRegion(d *schema.ResourceData) string {
	if awsL, ok := d.Get("aws").([]interface{}); ok && len(awsL) > 0 && awsL[0] != nil {
		awsM := awsL[0].(map[string]interface{})
		if region, ok := awsM["region"]; ok {
			return region.(string)
		}
	}

	if region := os.Getenv(awsRegionVar); region != "" {
		return region
	}

	return ""
}

func azureConfig(d *schema.ResourceData) *authentication.Builder {
	if azureL, ok := d.Get("azure").([]interface{}); ok && len(azureL) > 0 && azureL[0] != nil {
		azureM := azureL[0].(map[string]interface{})
		builder := &authentication.Builder{
			SubscriptionID:           azureM["subscription_id"].(string),
			ClientID:                 azureM["client_id"].(string),
			ClientSecret:             azureM["client_secret"].(string),
			TenantID:                 azureM["tenant_id"].(string),
			Environment:              azureEnvironment,
			SupportsClientSecretAuth: true,
			SupportsAzureCliToken:    true,
		}
		return builder
	}

	builder := &authentication.Builder{
		SubscriptionID:           os.Getenv(azureSubscriptionID),
		ClientID:                 os.Getenv(azureClientID),
		ClientSecret:             os.Getenv(azureClientSecret),
		TenantID:                 os.Getenv(azureTenantID),
		Environment:              azureEnvironment,
		SupportsClientSecretAuth: true,
		SupportsAzureCliToken:    true,
	}

	return builder
}
