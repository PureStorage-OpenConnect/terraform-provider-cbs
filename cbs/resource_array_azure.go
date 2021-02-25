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
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-07-01/managedapplications"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	mapset "github.com/deckarep/golang-set"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

const kind = "MarketPlace"

// Plan block. We will expose plan as input param in future versions.
const (
	planName    = "cbs_azure_6_1_0"
	product     = "pure_storage_cloud_block_store_deployment"
	publisher   = "purestoragemarketplaceadmin"
	planVersion = "1.0.0"
)

var templateTags = []string{
	"Microsoft.Network/applicationSecurityGroups",
	"Microsoft.DocumentDB/databaseAccounts",
	"Microsoft.Compute/disks",
	"Microsoft.KeyVault/vaults",
	"Microsoft.Network/loadBalancers",
	"Microsoft.ManagedIdentity/userAssignedIdentities",
	"Microsoft.Compute/virtualMachines/extensions",
	"Microsoft.Network/networkInterfaces",
	"Microsoft.Network/networkSecurityGroups",
	"Microsoft.Network/publicIPAddresses",
	"Microsoft.Compute/virtualMachines",
}

var azureParams = []interface{}{
	"arrayName",
	"licenseKey",
	"location",
	"orgDomain",
	"pureuserPublicKey",
	"sku",
	"managementSubnet",
	"systemSubnet",
	"iSCSISubnet",
	"replicationSubnet",
	"managementVnet",
	"systemVnet",
	"iSCSIVnet",
	"replicationVnet",
	"managementResourceGroup",
	"systemResourceGroup",
	"iSCSIResourceGroup",
	"replicationResourceGroup",
	"zone",
}

var renamedAzureParams = map[string]string{
	"orgDomain":       "log_sender_domain",
	"sku":             "array_model",
	"managementVnet":  "virtual_network",
	"systemVnet":      "virtual_network",
	"iSCSIVnet":       "virtual_network",
	"replicationVnet": "virtual_network",
}

var azureTFOutputs = []interface{}{
	"applicationName",
	"managedResourceGroupName",
	"ct0Name",
	"ct1Name",
	"managementEndpoint",
	"managementEndpointCT0",
	"managementEndpointCT1",
	"replicationEndpointCT0",
	"replicationEndpointCT1",
	"iSCSIEndpointCT0",
	"iSCSIEndpointCT1",
}

func resourceArrayAzure() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceArrayAzureCreate,
		ReadContext:   resourceArrayAzureRead,
		UpdateContext: resourceArrayAzureUpdate,
		DeleteContext: resourceArrayAzureDelete,
		Schema: map[string]*schema.Schema{
			"resource_group_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateResourceGroupName,
			},

			"location": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"australiaeast",
					"centralus",
					"eastus",
					"eastus2",
					"northeurope",
					"westeurope",
					"westus2",
				}, false),
			},

			// parameters
			"array_name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateManagedApplicationName,
			},

			"alert_recipients": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"license_key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"log_sender_domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"pureuser_public_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},

			"array_model": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"V10MUR1",
					"V20MUR1",
				}, false),
			},

			"management_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"system_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"iscsi_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"replication_subnet": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"management_resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"system_resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"iscsi_resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"replication_resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"zone": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
				ValidateFunc: validation.IntInSlice([]int{
					1,
					2,
					3,
				}),
			},

			"jit_approval": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"activation_maximum_duration": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "PT8H",
							ValidateFunc: validation.StringInSlice([]string{
								"PT1H",
								"PT2H",
								"PT3H",
								"PT4H",
								"PT5H",
								"PT6H",
								"PT7H",
								"PT8H",
							}, false),
						},
						"approvers": {
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"groups": {
										Type:     schema.TypeList,
										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
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
			"application_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"managed_resource_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ct0_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ct1_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint_ct0": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint_ct1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_endpoint_ct0": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_endpoint_ct1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"iscsi_endpoint_ct0": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"iscsi_endpoint_ct1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(50 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(50 * time.Minute),
		},
	}
}

func resourceArrayAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	azureClient, diags := m.(*CbsService).AzureClientService()
	if diags.HasError() {
		return diags
	}
	client := azureClient.ApplicationsClient
	groupClient := azureClient.GroupsClient

	name := d.Get("array_name").(string)
	managedResourceGroup := toManagedResourceGroup(name)
	resourceGroupName := d.Get("resource_group_name").(string)

	if d.IsNewResource() {
		existing, err := client.Get(ctx, resourceGroupName, name)
		if err != nil {
			if !responseWasNotFound(existing.Response) {
				return diag.Errorf("failed to check for present of existing Managed Application Name %q (Resource Group %q): %+v", name, resourceGroupName, err)
			}
		}
		if existing.ID != nil && *existing.ID != "" {
			return diag.Errorf("A resource with the ID %q already exists - to be managed via Terraform this resource needs to be imported into the State.", *existing.ID)
		}
	}

	parameters := managedapplications.Application{
		Location: to.StringPtr(d.Get("location").(string)),
		Kind:     to.StringPtr(kind),
	}

	targetResourceGroupId := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", client.SubscriptionID, managedResourceGroup)
	parameters.ApplicationProperties = &managedapplications.ApplicationProperties{
		ManagedResourceGroupID: to.StringPtr(targetResourceGroupId),
	}

	parameters.Plan = &managedapplications.Plan{
		Name:      to.StringPtr(planName),
		Product:   to.StringPtr(product),
		Publisher: to.StringPtr(publisher),
		Version:   to.StringPtr(planVersion),
	}

	params := make(map[string]interface{})
	for _, value := range azureParams {
		valueStr := value.(string)
		params[valueStr] = struct {
			Val interface{} `json:"value"`
		}{
			Val: d.Get(templateToTFParam(valueStr, renamedAzureParams)),
		}
	}

	approval := d.Get("jit_approval").([]interface{})[0].(map[string]interface{})
	approver := approval["approvers"].([]interface{})[0].(map[string]interface{})
	var approvers []managedapplications.JitApproverDefinition
	displayNameList := convertToStringSlice(approver["groups"].([]interface{}))
	for _, displayName := range displayNameList {
		group, err := groupGetByDisplayName(ctx, groupClient, displayName)
		if err != nil {
			return diag.Errorf("No group found matching specified name %q: %+v", displayName, err)
		}
		if group.ObjectID == nil {
			return diag.Errorf("Group returned with nil object ID: %+v", err)
		}
		newApprover := managedapplications.JitApproverDefinition{
			DisplayName: to.StringPtr(displayName),
			ID:          group.ObjectID,
			Type:        managedapplications.Group,
		}
		approvers = append(approvers, newApprover)
	}

	parameters.ApplicationProperties.JitAccessPolicy = &managedapplications.ApplicationJitAccessPolicy{
		JitAccessEnabled:         to.BoolPtr(true),
		JitApprovalMode:          managedapplications.JitApprovalModeManualApprove,
		MaximumJitAccessDuration: to.StringPtr(approval["activation_maximum_duration"].(string)),
		JitApprovers:             &approvers,
	}

	if v, ok := d.GetOk("alert_recipients"); ok {
		newRecips := convertToStringSlice(v.([]interface{}))
		params["alertRecipients"] = struct {
			Val interface{} `json:"value"`
		}{
			Val: strings.Join(newRecips, ","),
		}
	} else { // Deployment template has validation check on 'alertRecipients'. If not set, it should be "" instead of null.
		params["alertRecipients"] = struct {
			Val interface{} `json:"value"`
		}{
			Val: "",
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		tags := v.(map[string]interface{})
		tagsMap := make(map[string]interface{})
		for _, tag := range templateTags {
			tagsMap[tag] = tags
		}
		params["tagsByResource"] = struct {
			Val interface{} `json:"value"`
		}{
			Val: tagsMap,
		}
	}

	parameters.Parameters = &params

	future, err := client.CreateOrUpdate(ctx, resourceGroupName, name, parameters)
	if err != nil {
		return diag.Errorf("failed to create Managed Application %q (Resource Group %q): %+v", name, resourceGroupName, err)
	}
	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return diag.Errorf("failed to wait for creation of Managed Application %q (Resource Group %q): %+v", name, resourceGroupName, err)
	}

	resp, err := client.Get(ctx, resourceGroupName, name)
	if err != nil {
		return diag.Errorf("failed to retrieve Managed Application %q (Resource Group %q): %+v", name, resourceGroupName, err)
	}
	if resp.ID == nil || *resp.ID == "" {
		return diag.Errorf("cannot read Managed Application %q (Resource Group %q) ID", name, resourceGroupName)
	}
	d.SetId(*resp.ID)

	return resourceArrayAzureRead(ctx, d, m)
}

func resourceArrayAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	azureClient, diags := m.(*CbsService).AzureClientService()
	if diags.HasError() {
		return diags
	}
	client := azureClient.ApplicationsClient

	v, ok := d.GetOk("array_name")
	if !ok {
		log.Printf("[WARN] No Managed Application found with Id %q, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	appName := v.(string)
	managedResourceGroup := toManagedResourceGroup(appName)

	resourceGroup := d.Get("resource_group_name").(string)
	resp, err := client.Get(ctx, resourceGroup, appName)
	if err != nil {
		if responseWasNotFound(resp.Response) {
			log.Printf("[WARN] Managed Application %q does not exist - removing from state", d.Id())
			d.SetId("")
			return nil
		}
		return diag.Errorf("failed to read Managed Application %q (Resource Group %q): %+v", appName, resourceGroup, err)
	}

	d.Set("application_name", appName)
	d.Set("managed_resource_group_name", managedResourceGroup)
	d.Set("resource_group_name", resourceGroup)
	d.Set("location", resp.Location)

	if props := resp.ApplicationProperties; props != nil {
		params := props.Parameters.(map[string]interface{})
		azureParamSet := mapset.NewSetFromSlice(azureParams)
		for k, v := range params {
			if v != nil {
				if k == "AlertRecipients" {
					recips := strings.Split(v.(map[string]interface{})["value"].(string), ",")
					d.Set("alert_recipients", recips)
				}
				if k == "tagsByResource" {
					maps := v.(map[string]interface{})["value"].(map[string]interface{})
					for _, tagValue := range maps {
						d.Set("tags", tagValue)
						break
					}
				}
				if azureParamSet.Contains(k) {
					d.Set(templateToTFParam(k, renamedAzureParams), v.(map[string]interface{})["value"])
				}
			}
		}

		if err = d.Set("jit_approval", flattenJitApproval(props.JitAccessPolicy)); err != nil {
			return diag.FromErr(err)
		}

		outputs := props.Outputs.(map[string]interface{})
		azureTFOutputSet := mapset.NewSetFromSlice(azureTFOutputs)
		for k, v := range outputs {
			if v != nil {
				if k == "floatingManagementIP" {
					d.Set("management_endpoint", v.(map[string]interface{})["value"])
				}
				if azureTFOutputSet.Contains(k) {
					d.Set(toSnake(k), v.(map[string]interface{})["value"])
				}
			}
		}
	}
	return nil
}

func resourceArrayAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	diags := resourceArrayAzureRead(ctx, d, m)
	if diags.HasError() {
		return diags
	}
	return diag.Errorf("Updates are not supported.")
}

func resourceArrayAzureDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	azureClient, diags := m.(*CbsService).AzureClientService()
	if diags.HasError() {
		return diags
	}
	client := azureClient.ApplicationsClient

	resourceGroup := d.Get("resource_group_name").(string)
	appName := d.Get("array_name").(string)
	future, err := client.Delete(ctx, resourceGroup, appName)
	if err != nil {
		return diag.Errorf("failed to delete Managed Application %q (Resource Group %q): %+v", appName, resourceGroup, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return diag.Errorf("failed to wait for deleting Managed Application (Managed Application Name %q / Resource Group %q): %+v", appName, resourceGroup, err)
	}
	return nil
}

func validateManagedApplicationName(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)

	if !regexp.MustCompile(`^[-\da-zA-Z]{3,64}$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("%q must be between 3 and 64 characters in length and contains only letters, numbers or hyphens.", k))
	}

	return warnings, errors
}

func validateResourceGroupName(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)

	if len(value) > 90 {
		errors = append(errors, fmt.Errorf("%q may not exceed 90 characters in length", k))
	}

	if strings.HasSuffix(value, ".") {
		errors = append(errors, fmt.Errorf("%q may not end with a period", k))
	}

	// regex pulled from https://docs.microsoft.com/en-us/rest/api/resources/resourcegroups/createorupdate
	if matched := regexp.MustCompile(`^[-\w._()]+$`).Match([]byte(value)); !matched {
		errors = append(errors, fmt.Errorf("%q may only contain alphanumeric characters, dash, underscores, parentheses and periods", k))
	}

	return warnings, errors
}

func responseWasNotFound(resp autorest.Response) bool {
	if r := resp.Response; r != nil {
		if r.StatusCode == http.StatusNotFound {
			return true
		}
	}

	return false
}

func toManagedResourceGroup(name string) string {
	result := name + "-cbs-mrg"
	return result
}

func groupGetByDisplayName(ctx context.Context, client *graphrbac.GroupsClient, displayName string) (*graphrbac.ADGroup, error) {
	filter := fmt.Sprintf("displayName eq '%s'", displayName)
	resp, err := client.ListComplete(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("listing Groups for filter %q: %+v", filter, err)
	}

	values := resp.Response().Value
	if values == nil {
		return nil, fmt.Errorf("nil values for Groups matching %q", filter)
	}
	if len(*values) == 0 {
		return nil, fmt.Errorf("found no Groups matching %q", filter)
	}
	if len(*values) > 2 {
		return nil, fmt.Errorf("found multiple Groups matching %q", filter)
	}

	group := (*values)[0]
	if group.DisplayName == nil {
		return nil, fmt.Errorf("nil DisplayName for Group matching %q", filter)
	}
	if !strings.EqualFold(*group.DisplayName, displayName) {
		return nil, fmt.Errorf("displayname for Group matching %q does not match (%q!=%q)", filter, *group.DisplayName, displayName)
	}

	return &group, nil
}

func flattenJitApproval(policy *managedapplications.ApplicationJitAccessPolicy) []map[string]interface{} {
	results := make([]map[string]interface{}, 1)

	approver := make(map[string]interface{})
	result := make(map[string]interface{})
	var groupList []string
	for _, g := range *policy.JitApprovers {
		groupList = append(groupList, *g.DisplayName)
	}
	approver["groups"] = groupList
	approvers := make([]map[string]interface{}, 1)
	approvers[0] = approver
	result["approvers"] = approvers
	result["activation_maximum_duration"] = *policy.MaximumJitAccessDuration
	results[0] = result

	return results
}
