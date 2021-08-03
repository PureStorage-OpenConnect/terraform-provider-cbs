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
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/cloud"
)

const (
	awsRegionVar = "AWS_DEFAULT_REGION"
)

const (
	azureSubscriptionID = "ARM_SUBSCRIPTION_ID"
	azureClientID       = "ARM_CLIENT_ID"
	azureTenantID       = "ARM_TENANT_ID"
	azureClientSecret   = "ARM_CLIENT_SECRET"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"aws": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": {
							Type:        schema.TypeString,
							Optional:    true,
							DefaultFunc: schema.EnvDefaultFunc(awsRegionVar, nil),
						},
					},
				},
			},

			"azure": {
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

	cbsService.awsRegionStr = awsRegion(d)

	if _, ok := d.GetOk("aws"); ok {
		if _, diags := cbsService.awsClientService(); diags.HasError() {
			return nil, diags
		}
	}

	cbsService.azureConfig = azureMakeConfig(d)

	if _, ok := d.GetOk("azure"); ok {
		azureClient, err := cloud.NewAzureClient(cbsService.azureConfig)
		if err != nil {
			return nil, diag.FromErr(err)
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

func azureMakeConfig(d *schema.ResourceData) (config cloud.AzureConfig) {
	config.SubscriptionID = os.Getenv(azureSubscriptionID)
	config.ClientID = os.Getenv(azureClientID)
	config.ClientSecret = os.Getenv(azureClientSecret)
	config.TenantID = os.Getenv(azureTenantID)

	if azureL, ok := d.Get("azure").([]interface{}); ok && len(azureL) > 0 && azureL[0] != nil {
		azureM := azureL[0].(map[string]interface{})
		config.SubscriptionID = azureM["subscription_id"].(string)
		config.ClientID = azureM["client_id"].(string)
		config.ClientSecret = azureM["client_secret"].(string)
		config.TenantID = azureM["tenant_id"].(string)
	}

	return
}
