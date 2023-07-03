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

	"os"
	"regexp"
	"strings"
	"time"

	vaultSecret "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/auth"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/cloud"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/internal/tfazurerm"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/version"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managedapplications"
	"github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/go-autorest/autorest/to"
	mapset "github.com/deckarep/golang-set"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

// Default managed application plan
const (
	defaultPlanName      = "cbs_azure_6_3_8"
	defaultPlanProduct   = "pure_storage_cloud_block_store_deployment"
	defaultPlanPublisher = "purestoragemarketplaceadmin"
	defaultPlanVersion   = "1.0.0"
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
	"orgDomain": "log_sender_domain",
	"sku":       "array_model",
}

var azureTFOutputs = []string{
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
	sch := &schema.Resource{
		CreateContext: resourceArrayAzureCreate,
		ReadContext:   resourceArrayAzureRead,
		UpdateContext: resourceArrayAzureUpdate,
		DeleteContext: resourceArrayAzureDelete,
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
			"array_name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateAzureManagedApplicationName,
			},

			"alert_recipients": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"license_key": {
				Type:     schema.TypeString,
				Required: true,
			},

			"log_sender_domain": {
				Type:     schema.TypeString,
				Required: true,
			},

			"pureuser_private_key_path": {
				Type:         schema.TypeString,
				Optional:     true,
				ExactlyOneOf: []string{"pureuser_private_key_path", "pureuser_private_key"},
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"pureuser_private_key": {
				Type:         schema.TypeString,
				Optional:     true,
				Sensitive:    true,
				ExactlyOneOf: []string{"pureuser_private_key_path", "pureuser_private_key"},
				ValidateFunc: validation.StringIsNotEmpty,
			},
			"key_vault_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateKeyVaultId,
			},

			"array_model": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"V10MUR1",
					"V20MUR1",
					"V20MP2R2",
				}, false),
			},

			"management_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"system_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"iscsi_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replication_subnet": {
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_network_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringIsNotEmpty,
			},

			"zone": {
				Type:     schema.TypeInt,
				Required: true,
				ValidateFunc: validation.IntInSlice([]int{
					1,
					2,
					3,
				}),
			},

			"user_assigned_identity": {
				Type:        schema.TypeString,
				Description: "A required input that denotes the identity of the customer User Assigned identity.",
				Required:    true,
			},

			"fusion_sec_identity": {
				Type:        schema.TypeString,
				Description: "Optional input that denotes the identity of a Fusion Storage Endpoint Collection, obtained during Azure Portal GUI or CLI deployment",
				Optional:    true,
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
			"ct0_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ct1_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint_ct0": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_endpoint_ct1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_endpoint_ct0": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"replication_endpoint_ct1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iscsi_endpoint_ct0": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"iscsi_endpoint_ct1": {
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

	// Allow for more extensive testing when in development mode
	if version.IsDevelopmentMode() {
		sch.Schema["app_definition_id"] = &schema.Schema{
			Type:          schema.TypeString,
			Optional:      true,
			ConflictsWith: []string{"plan"},
		}

		sch.Schema["plan"].ConflictsWith = []string{"app_definition_id"}
		sch.Schema["array_model"].ValidateFunc = nil
		sch.Schema["jit_approval_group_object_ids"].Required = false
		sch.Schema["jit_approval_group_object_ids"].Optional = true
	}
	return sch
}

func resourceArrayAzureCreate(ctx context.Context, d *schema.ResourceData, m interface{}) (returnedDiags diag.Diagnostics) {
	tflog.Trace(ctx, "resourceArrayAzureCreate")
	azureClient, diags := m.(*CbsService).azureClientService(ctx)
	if diags.HasError() {
		return diags
	}

	name := d.Get("array_name").(string)
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
			return diag.Errorf("A resource with the ID %q already exists - to be managed via Terraform this resource needs to be imported into the State.", *existing.ID)
		}
	}

	parameters := managedapplications.Application{
		Location: to.StringPtr(d.Get("location").(string)),
	}

	targetResourceGroupId := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", azureClient.SubscriptionID(), managedResourceGroup)
	parameters.ApplicationProperties = &managedapplications.ApplicationProperties{
		ManagedResourceGroupID: to.StringPtr(targetResourceGroupId),
	}

	if appDefinitionId, ok := d.GetOk("app_definition_id"); ok {
		parameters.Kind = to.StringPtr("ServiceCatalog")
		parameters.ApplicationDefinitionID = to.StringPtr(appDefinitionId.(string))
	} else {
		parameters.Kind = to.StringPtr("MarketPlace")
		if v, ok1 := d.GetOk("plan"); ok1 && len(v.([]interface{})) > 0 {
			parameters.Plan = expandPlan(v.([]interface{}))
		} else {
			parameters.Plan = &managedapplications.Plan{
				Name:      to.StringPtr(defaultPlanName),
				Product:   to.StringPtr(defaultPlanProduct),
				Publisher: to.StringPtr(defaultPlanPublisher),
				Version:   to.StringPtr(defaultPlanVersion),
			}
		}
	}

	vnetName, vnetRGName, err := tfazurerm.ParseNameRGFromID(d.Get("virtual_network_id").(string), "virtualNetworks")
	if err != nil {
		return diag.FromErr(err)
	}

	parameters.Parameters = make(map[string]interface{})
	setAppParameter := func(key string, value interface{}) {
		(parameters.Parameters.(map[string]interface{}))[key] = map[string]interface{}{"value": value}
	}
	for _, value := range azureParams {
		valueStr := value.(string)
		if strings.HasSuffix(valueStr, "Vnet") {
			setAppParameter(valueStr, vnetName)
		} else if strings.HasSuffix(valueStr, "ResourceGroup") {
			setAppParameter(valueStr, vnetRGName)
		} else {
			setAppParameter(valueStr, d.Get(templateToTFParam(valueStr, renamedAzureParams)))
		}
	}

	returnedDiags = setAzureJitAccessPolicy(&parameters, d)

	if v, ok := d.GetOk("alert_recipients"); ok {
		newRecips := convertToStringSlice(v.([]interface{}))
		setAppParameter("alertRecipients", strings.Join(newRecips, ","))
	} else { // Deployment template has validation check on 'alertRecipients'. If not set, it should be "" instead of null.
		setAppParameter("alertRecipients", "")
	}

	var identities = []string{}
	if v, ok := d.GetOk("user_assigned_identity"); ok {
		identities = append(identities, v.(string))
	} else {
		return diag.Errorf("failed to retrieve user_assigned_identity")
	}

	if v, ok := d.GetOk("fusion_sec_identity"); ok {
		identities = append(identities, v.(string))
		setAppParameter("fusionSECIdentity", expandIdentityObject(identities[1:]))
	}

	parameters.Identity = expandIdentityObject(identities)

	if v, ok := d.GetOk("tags"); ok {
		tags := v.(map[string]interface{})
		tagsMap := make(map[string]interface{})
		for _, tag := range templateTags {
			tagsMap[tag] = tags
		}
		setAppParameter("tagsByResource", tagsMap)
	}

	err = prevalidateKeyVaultId(ctx, d, azureClient)
	if err != nil {
		return diag.FromErr(err)
	}
	defer func() {
		if returnedDiags.HasError() {
			vaultId, secretName := vaultIdSecretName(d)
			azureClient.SecretDelete(ctx, vaultId, secretName)
		}
	}()

	pvtKeyBytes, err := getSSHPrivateKeyBytesFromResourceData(d)
	if err != nil {
		return diag.FromErr(err)
	}

	sshPublicKey, err := auth.PrivateKeyDerivePublicKey(pvtKeyBytes)
	if err != nil {
		return diag.FromErr(err)
	}
	setAppParameter("pureuserPublicKey", string(sshPublicKey))
	// Error out now, before we create resources
	if returnedDiags.HasError() {
		return returnedDiags
	}

	tflog.Trace(ctx, "resourceArrayAzureCreate AppsCreateOrUpdate")
	err = azureClient.AppsCreateOrUpdate(ctx, resourceGroupName, name, parameters)
	defer func() {
		if returnedDiags.HasError() {
			azureClient.AppsDelete(ctx, resourceGroupName, name)
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

	diags = resourceArrayAzureRead(ctx, d, m)
	if diags.HasError() {
		return diags
	}

	err = generateAndSetSecret(ctx, d, azureClient)
	if err != nil {
		return diag.FromErr(err)
	}

	return returnedDiags
}

func resourceArrayAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "resourceArrayAzureRead")
	azureClient, diags := m.(*CbsService).azureClientService(ctx)
	if diags.HasError() {
		return diags
	}

	v, ok := d.GetOk("array_name")
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

	d.Set("application_name", appName)
	d.Set("managed_resource_group_name", managedResourceGroup)
	d.Set("resource_group_name", resourceGroup)
	d.Set("location", resp.Location)

	if props := resp.ApplicationProperties; props != nil {
		params := formatAzureParameters(props.Parameters)
		azureParamSet := mapset.NewSetFromSlice(azureParams)
		var vnetName, vnetRGName string
		for k, v := range params {
			// SecureString parameters will always have a null value, so ignore them
			if v.valType != "SecureString" {
				if k == "AlertRecipients" {
					recips := strings.Split(v.value.(string), ",")
					d.Set("alert_recipients", recips)
				}
				if k == "tagsByResource" {
					maps := v.value.(map[string]interface{})
					for _, tagValue := range maps {
						d.Set("tags", tagValue)
						break
					}
				}
				if azureParamSet.Contains(k) {
					if strings.HasSuffix(k, "Vnet") {
						vnetName = v.value.(string)
					} else if strings.HasSuffix(k, "ResourceGroup") {
						vnetRGName = v.value.(string)
					} else {
						d.Set(templateToTFParam(k, renamedAzureParams), v.value)
					}

				}
			}
		}

		// set virtual_network_id using VirtualNetwork Name and ResourceGroup name
		if vnetName == "" || vnetRGName == "" {
			return diag.Errorf("failed to read VirtualNetwork ID of Managed Application %q", appName)
		}
		vnetId := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Network/virtualNetworks/%s",
			azureClient.SubscriptionID(), vnetRGName, vnetName)
		d.Set("virtual_network_id", vnetId)

		if err := d.Set("jit_approval_group_object_ids", flattenAzureJitApprovalGroupIds(props.JitAccessPolicy)); err != nil {
			return diag.FromErr(err)
		}

		outputs := props.Outputs.(map[string]interface{})

		azureTFOutputSet := mapset.NewSet()
		for _, s := range azureTFOutputs {
			azureTFOutputSet.Add(s)
		}

		for k, v := range outputs {
			if v != nil {
				v := v.(map[string]interface{})
				if k == "floatingManagementIP" {
					d.Set("management_endpoint", v["value"])
				}
				if azureTFOutputSet.Contains(k) {
					d.Set(toSnake(k), v["value"])
				}
			}
		}

	}
	return nil
}

func resourceArrayAzureUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "resourceArrayAzureUpdate")
	diags := resourceArrayAzureRead(ctx, d, m)
	if diags.HasError() {
		return diags
	}
	return diag.Errorf("Updates are not supported.")
}

func resourceArrayAzureDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	tflog.Trace(ctx, "resourceArrayAzureDelete")
	azureClient, diags := m.(*CbsService).azureClientService(ctx)
	if diags.HasError() {
		return diags
	}

	vaultId, secretName := vaultIdSecretName(d)

	faClient, err := azureClient.NewFAClient(ctx, d.Get("management_endpoint").(string), vaultId, secretName)
	if err != nil {
		return diag.Errorf("failed to create FA client with managed application ID: %s. "+
			"Please contact Pure Storage support to deactivate the instance: %+v.", d.Id(), err)
	}

	if err := faClient.Deactivate(ctx); err != nil {
		return diag.Errorf("failed to deactivate the instance with Id %s. Please contact "+
			"Pure Storage support: %+v.", d.Id(), err)
	}

	azureClient.DeactivateWait()

	resourceGroup := d.Get("resource_group_name").(string)
	appName := d.Get("array_name").(string)

	err = azureClient.AppsDelete(ctx, resourceGroup, appName)
	if err != nil {
		return diag.FromErr(err)
	}

	azureClient.SecretDelete(ctx, vaultId, secretName)

	return nil
}

func validateKeyVaultId(v interface{}, k string) (warnings []string, errors []error) {

	_, _, err := tfazurerm.ParseNameRGFromID(v.(string), "vaults")
	if err != nil {
		errors = append(errors, err)
	}

	return
}

func groupGetByDisplayName(ctx context.Context, client cloud.AzureClientAPI, displayName string) (*graphrbac.ADGroup, error) {

	filter := fmt.Sprintf("displayName eq '%s'", displayName)

	values, err := client.GroupsListComplete(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("listing Groups for filter %q: %+v", filter, err)
	}
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

func expandIdentityObject(identities []string) *managedapplications.Identity {
	var userIdentities = make(map[string]*managedapplications.UserAssignedResourceIdentity)
	for _, identity := range identities {
		userIdentities[identity] = new(managedapplications.UserAssignedResourceIdentity)
	}
	return &managedapplications.Identity{
		Type:                   managedapplications.ResourceIdentityTypeUserAssigned,
		UserAssignedIdentities: userIdentities,
	}
}

var invalidSecretCharacters = regexp.MustCompile("[^a-zA-Z0-9-]+")

func vaultIdSecretName(d *schema.ResourceData) (vaultId string, secretName string) {
	secretName = fmt.Sprintf("cbs-%s-%s", d.Get("resource_group_name").(string), d.Get("array_name").(string))
	secretName = invalidSecretCharacters.ReplaceAllString(secretName, "")
	vaultId = d.Get("key_vault_id").(string)
	return
}

// Here we just write/read from the secret, just to make sure that we have a properly configured key_vault_id
func prevalidateKeyVaultId(ctx context.Context, d *schema.ResourceData, azureClient cloud.AzureClientAPI) error {
	tflog.Trace(ctx, "prevalidateKeyVaultId")
	vaultId, secretName := vaultIdSecretName(d)
	const placeholderText = "PLACEHOLDER"

	existing, err := azureClient.SecretGet(ctx, vaultId, secretName, "")
	if !(utils.ResponseWasNotFound(existing.Response) || (existing.Value != nil && *existing.Value == placeholderText)) {
		if err != nil {
			return fmt.Errorf("failed to check for existing secret: %+v Secret name: %s key_vault_id: %s", err, secretName, vaultId)
		} else if existing.Value != nil {
			return fmt.Errorf("secret already exists, please check for existing deployment, or change resource group or name. Secret name: %s key_vault_id: %s", secretName, vaultId)
		} else {
			return fmt.Errorf("unhandled exceptional case while checking existing secret: Secret name: %s key_vault_id: %s  Response Status: %s Body: %s",
				secretName, vaultId, existing.Request.Response.Status, existing.Response.Body)
		}
	}

	tflog.Trace(ctx, "prevalidateKeyVaultId azureClient.SecretSet")
	setSecret, err := azureClient.SecretSet(ctx, vaultId, secretName, vaultSecret.SecretSetParameters{Value: to.StringPtr(placeholderText)})
	if utils.ResponseWasStatusCode(setSecret.Response, 409 /* Conflict: Deleted but not purged */) {
		err = azureClient.SecretRecover(ctx, vaultId, secretName)
		if err != nil {
			return fmt.Errorf("failed to recover previously deleted secret: %+v Secret name: %s key_vault_id: %s", err, secretName, vaultId)
		}
		for {
			setSecret, err = azureClient.SecretSet(ctx, vaultId, secretName, vaultSecret.SecretSetParameters{Value: to.StringPtr(placeholderText)})
			if !utils.ResponseWasStatusCode(setSecret.Response, 409) {
				break
			}
			fmt.Fprintf(os.Stderr, "[WARN] An existing deleted secret exists, we are recovering it so that we may overwrite it: %+v\n", err)
			time.Sleep(10 * time.Second)
		}
	}
	if err != nil {
		return fmt.Errorf("prevalidation failure: Please check value for: %+v Secret name: %s key_vault_id: %s", err, secretName, vaultId)
	}

	secret, err := azureClient.SecretGet(ctx, vaultId, secretName, "")
	if err != nil {
		return fmt.Errorf("failed to get secret: %+v Secret name: %s key_vault_id: %s", err, secretName, vaultId)
	}
	if *secret.Value != placeholderText {
		return fmt.Errorf("secret did not match secret that we set Secret name: %s key_vault_id: %s", secretName, vaultId)
	}
	return nil
}

func generateAndSetSecret(ctx context.Context, d *schema.ResourceData, azureClient cloud.AzureClientAPI) error {
	vaultId, secretName := vaultIdSecretName(d)

	restCredentials, err := generateSecretPayload(ctx, d)
	if err != nil {
		return err
	}

	tflog.Trace(ctx, "generateAndSetSecret azureClient.SecretSet")
	_, err = azureClient.SecretSet(ctx, vaultId, secretName, vaultSecret.SecretSetParameters{Value: to.StringPtr((string(restCredentials)))})
	if err != nil {
		return err
	}
	return nil
}
