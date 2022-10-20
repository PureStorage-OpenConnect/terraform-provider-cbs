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
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managedapplications"
	"github.com/Azure/go-autorest/autorest/to"
	mapset "github.com/deckarep/golang-set"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Default managed application plan
const (
	defaultFusionSECPlanName      = "pure_sec_1_0_0"
	defaultFusionSECPlanProduct   = "pure_fusion_storage_endpoint_collection"
	defaultFusionSECPlanPublisher = "purestoragemarketplaceadmin"
	defaultFusionSECPlanVersion   = "1.0.3"
)

var fusionSECAzureTemplateTags = []string{
	"Microsoft.Network/loadBalancers",
	"Microsoft.ManagedIdentity/userAssignedIdentities",
}

var fusionSECAzureParams = []interface{}{
	"fusionSECName",
	"location",
	"loadBalancerNetworkRg",
	"loadBalancerNetworkName",
	"loadBalancerSubnet",
}

var renamedFusionSECAzureParams = map[string]string{}

var fusionSECAzureTFOutputs = []string{
	"applicationName",
	"managedResourceGroupName",
	"hmvip0",
	"hmvip1",
	"loadBalancerFullIdentityId",
}

func resourceFusionSECAzure() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFusionSECAzureCreate,
		ReadContext:   resourceFusionSECAzureRead,
		UpdateContext: resourceFusionSECAzureUpdate,
		DeleteContext: resourceFusionSECAzureDelete,
		Schema: map[string]*schema.Schema{
			"resource_group_name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateAzureResourceGroupName,
			},

			"location": {
				Type:     schema.TypeString,
				Required: true,
			},

			// parameters
			"fusion_sec_name": {
				Description:  "The name of the Fusion Storage Endpoint Collection (SEC). 0-59 alphanumeric characters only.",
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAzureManagedApplicationName,
			},

			"load_balancer_network_rg": {
				Type:     schema.TypeString,
				Required: true,
			},

			"load_balancer_network_name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"load_balancer_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},

			"jit_approval_group_object_ids": {
				Description: "This is a list of Azure group object IDs for people who are allowed to approve JIT requests",
				Required:    true,
				Type:        schema.TypeList,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.IsUUID,
				},
			},

			"plan": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"product": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"publisher": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"version": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
					},
				},
			},

			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type:         schema.TypeString,
					ValidateFunc: validation.StringIsNotEmpty,
				},
			},

			// Outputs
			"application_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_resource_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hmvip0": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hmvip1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"load_balancer_full_identity_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},
	}
}

func resourceFusionSECAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (returnedDiags diag.Diagnostics) {
	tflog.Trace(ctx, "resourceFusionSECAzurereate")
	azureClient, diags := m.(*CbsService).azureClientService(ctx)
	if diags.HasError() {
		return diags
	}

	name := d.Get("fusion_sec_name").(string)
	managedResourceGroup := toAzureManagedResourceGroup(name)
	resourceGroupName := d.Get("resource_group_name").(string)

	if d.IsNewResource() {
		existing, err := azureClient.AppsGet(ctx, resourceGroupName, name)
		if err != nil {
			if !responseWasNotFound(existing.Response) {
				return diag.Errorf("failed to check for presence of existing Managed Application Name %q (Resource Group %q): %+v", name, resourceGroupName, err)
			}
		}
		if existing.ID != nil && *existing.ID != "" {
			return diag.Errorf(
				"A resource with the name %q, Resource Group %q and ID %q already exists - to be managed via Terraform this resource needs to be imported into the State.",
				name,
				resourceGroupName,
				*existing.ID,
			)
		}
	}

	parameters := managedapplications.Application{
		Location: to.StringPtr(d.Get("location").(string)),
	}

	targetResourceGroupId := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", azureClient.SubscriptionID(), managedResourceGroup)
	parameters.ApplicationProperties = &managedapplications.ApplicationProperties{
		ManagedResourceGroupID: to.StringPtr(targetResourceGroupId),
	}

	parameters.Kind = to.StringPtr("MarketPlace")
	if v, ok1 := d.GetOk("plan"); ok1 && len(v.([]interface{})) > 0 {
		parameters.Plan = expandPlan(v.([]interface{}))
	} else {
		parameters.Plan = &managedapplications.Plan{
			Name:      to.StringPtr(defaultFusionSECPlanName),
			Product:   to.StringPtr(defaultFusionSECPlanPublisher),
			Publisher: to.StringPtr(defaultFusionSECPlanPublisher),
			Version:   to.StringPtr(defaultFusionSECPlanVersion),
		}
	}

	parameters.Parameters = make(map[string]interface{})
	setAppParameter := func(key string, value interface{}) {
		(parameters.Parameters.(map[string]interface{}))[key] = map[string]interface{}{"value": value}
	}
	for _, value := range fusionSECAzureParams {
		valueStr := value.(string)
		setAppParameter(valueStr, d.Get(templateToTFParam(valueStr, renamedFusionSECAzureParams)))
	}

	returnedDiags = setAzureJitAccessPolicy(&parameters, d)

	if v, ok := d.GetOk("tags"); ok {
		tags := v.(map[string]interface{})
		tagsMap := make(map[string]interface{})
		for _, tag := range fusionSECAzureTemplateTags {
			tagsMap[tag] = tags
		}
		setAppParameter("tagsByResource", tagsMap)
	}

	// Error out now, before we create resources
	if returnedDiags.HasError() {
		return returnedDiags
	}

	tflog.Trace(ctx, "resourceFusionSECAzureCreate AppsCreateOrUpdate")
	err := azureClient.AppsCreateOrUpdate(ctx, resourceGroupName, name, parameters)
	defer func() {
		if returnedDiags.HasError() {
			if err = azureClient.AppsDelete(ctx, resourceGroupName, name); err != nil {
				tflog.Error(
					ctx,
					fmt.Sprintf(
						"failed to delete Managed Application %q (Resource Group %q) after failed CreateOrUpdate operation: %+v",
						name,
						resourceGroupName,
						err,
					),
				)
			}
		}
	}()
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := azureClient.AppsGet(ctx, resourceGroupName, name)
	if err != nil {
		return diag.Errorf("failed to retrieve Managed Application %q (Resource Group %q): %+v", name, resourceGroupName, err)
	}
	if resp.ID == nil || *resp.ID == "" {
		return diag.Errorf("cannot read Managed Application %q (Resource Group %q) ID", name, resourceGroupName)
	}
	d.SetId(*resp.ID)

	diags = resourceFusionSECAzureRead(ctx, d, m)
	if diags.HasError() {
		returnedDiags = append(returnedDiags, diags...)
	}

	return returnedDiags
}

func resourceFusionSECAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "resourceFusionSECAzureRead")
	azureClient, diags := m.(*CbsService).azureClientService(ctx)
	if diags.HasError() {
		return diags
	}

	v, ok := d.GetOk("fusion_sec_name")
	if !ok {
		log.Printf("[WARN] No Managed Application found with Id %q, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	appName := v.(string)
	managedResourceGroup := toAzureManagedResourceGroup(appName)

	resourceGroup := d.Get("resource_group_name").(string)
	resp, err := azureClient.AppsGet(ctx, resourceGroup, appName)
	if err != nil {
		if responseWasNotFound(resp.Response) {
			log.Printf("[WARN] Managed Application %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to read Managed Application %q (Resource Group %q): %+v", appName, resourceGroup, err)
	}

	if err := d.Set("application_name", appName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("managed_resource_group_name", managedResourceGroup); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("resource_group_name", resourceGroup); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("location", resp.Location); err != nil {
		return diag.FromErr(err)
	}

	if props := resp.ApplicationProperties; props != nil {
		params := formatAzureParameters(props.Parameters)
		fusionSECParamSet := mapset.NewSetFromSlice(fusionSECAzureParams)
		for k, v := range params {
			// SecureString parameters will always have a null value, so ignore them
			if v.valType != "SecureString" {
				if k == "tagsByResource" {
					maps := v.value.(map[string]interface{})
					for _, tagValue := range maps {
						if err := d.Set("tags", tagValue); err != nil {
							return diag.FromErr(err)
						}
						break
					}
				}
				if fusionSECParamSet.Contains(k) {
					if err := d.Set(templateToTFParam(k, renamedFusionSECAzureParams), v.value); err != nil {
						return diag.FromErr(err)
					}
				}
			}
		}

		if err := d.Set("jit_approval_group_object_ids", flattenAzureJitApprovalGroupIds(props.JitAccessPolicy)); err != nil {
			return diag.FromErr(err)
		}

		outputs := props.Outputs.(map[string]interface{})

		fusionSECAzureTFOutputSet := mapset.NewSet()
		for _, s := range fusionSECAzureTFOutputs {
			fusionSECAzureTFOutputSet.Add(s)
		}

		for k, v := range outputs {
			if v != nil {
				v := v.(map[string]interface{})
				if fusionSECAzureTFOutputSet.Contains(k) {
					if !strings.HasPrefix(k, "hmvip") {
						k = toSnake(k)
					}
					if err := d.Set(k, v["value"]); err != nil {
						return diag.FromErr(err)
					}
				}
			}
		}

	}

	return nil
}

func resourceFusionSECAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "resourceFusionSECAzureUpdate")
	diags := resourceArrayAzureRead(ctx, d, m)
	if diags.HasError() {
		return diags
	}
	return diag.Errorf("Updates are not supported.")
}

func resourceFusionSECAzureDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "resourceFusionSECAzureDelete")
	azureClient, diags := m.(*CbsService).azureClientService(ctx)
	if diags.HasError() {
		return diags
	}

	resourceGroup := d.Get("resource_group_name").(string)
	appName := d.Get("fusion_sec_name").(string)

	err := azureClient.AppsDelete(ctx, resourceGroup, appName)
	return diag.FromErr(err)
}
