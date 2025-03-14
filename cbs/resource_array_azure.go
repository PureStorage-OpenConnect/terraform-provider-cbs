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
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"os"
	"regexp"
	"strings"
	"time"

	vaultSecret "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.1/keyvault"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/auth"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/appcatalog"
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
	defaultPlanName      = "cbs_azure_6_5_x"
	defaultPlanProduct   = "pure_cloud_block_store_product_deployment"
	defaultPlanPublisher = "purestoragemarketplaceadmin"
	defaultPlanVersion   = "1.0.0"
)

const marketplacePlanIdPrefix = "purestoragemarketplaceadmin.pure_cloud_block_store_product_deployment"

var staticResourceList = []string{
	"Microsoft.Solutions/applications",
	"Microsoft.Resources/tags",
	"Microsoft.Compute/virtualMachines",
	"Microsoft.Network/networkInterfaces",
	"Microsoft.DocumentDB/databaseAccounts",
	"Microsoft.Storage/storageAccounts",
	"Microsoft.ManagedIdentity/userAssignedIdentities",
	"Microsoft.KeyVault/vaults",
	"Microsoft.Network/loadBalancers",
	"Microsoft.Network/publicIPAddresses",
	"Microsoft.Compute/disks",
	"Microsoft.Compute/virtualMachines/extensions",
	"Microsoft.Network/networkSecurityGroups",
	"Microsoft.Network/applicationSecurityGroups",
	"Microsoft.Compute/capacityReservationGroups",
	"Microsoft.Compute/capacityReservationGroups/capacityReservations",
	"Microsoft.Resources/deploymentScripts",
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
					"V10MP2R2",
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

			"resource_tags": {
				Description: "Optional field that defines specific tags for specific resource types",
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"resource": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validation.StringIsNotEmpty,
						},
						"tag": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},
									"value": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringIsNotEmpty,
									},
								},
							},
						},
					},
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

type JSONCreateUiDefinitionTemplate struct {
	Parameters struct {
		Steps []struct {
			Name     string
			Elements []struct {
				Name      string
				Resources []string
			}
		}
	}
}

// Retrieve resource list from Azure createUiDefinition artifact
func GetResourcesFromTemplateJson(data []byte) ([]string, error) {
	// Parse the createUiDefinition template
	var unmarshalled_data JSONCreateUiDefinitionTemplate
	err := json.Unmarshal(data, &unmarshalled_data)
	if err != nil {
		return nil, err
	}

	if len(unmarshalled_data.Parameters.Steps) == 0 {
		return nil, fmt.Errorf("createUiDefinition resources is of unexpected size %d, it must have at least one element", len(unmarshalled_data.Parameters.Steps))
	}

	var resources []string
	for _, v := range unmarshalled_data.Parameters.Steps {
		if v.Name == "tags" {
			if len(v.Elements) != 1 {
				return nil, fmt.Errorf("createUiDefinition resources array is of unexpected size %d, it must have the size of 1", len(v.Elements))
			}

			if v.Elements[0].Name != "tags" {
				return nil, fmt.Errorf("createUiDefinition resources has incorrect format to retrieve resource list")
			}

			if len(v.Elements[0].Resources) == 0 {
				return nil, fmt.Errorf("createUiDefinition resources must not be empty")
			}

			resources = v.Elements[0].Resources
		}
	}

	return resources, nil
}

func GetPlanArtifacts(ctx context.Context, plan Plan) (map[string]*appcatalog.Artifact, error) {
	productSummary, err := GetProductSummary(ctx)
	if err != nil {
		return nil, err
	}

	for _, response_result := range productSummary.Results {
		// Filter out non Cloud Block Store offers, unfortunatelly the API does not have offer/product ID available so we need
		// to take a look at the inner plans to filter out offers we do not want
		if len(response_result.Plans) == 0 || !strings.HasPrefix(*response_result.Plans[0].UniquePlanID, marketplacePlanIdPrefix) {
			tflog.Debug(ctx, fmt.Sprintf("Skipping non Cloud Block Store offer %s", *response_result.DisplayName))
			continue
		}
		tflog.Debug(ctx, fmt.Sprintf("Processing plans under offer %s", *response_result.DisplayName))
		for _, response_plan := range response_result.Plans {
			if plan.Name != *response_plan.PlanID {
				tflog.Debug(ctx, fmt.Sprintf("Skipping non matching plan %s", *response_plan.PlanID))
				continue
			}
			tflog.Debug(ctx, fmt.Sprintf("Found matching plan %s", *response_plan.PlanID))

			artifacts := make(map[string]*appcatalog.Artifact)
			var templatePlan *Plan
			for _, response_artifact := range response_plan.Artifacts {
				artifacts[*response_artifact.Name] = response_artifact
				// Get the plan properties from the DefaultTemplate artifact to verify integrity
				if *response_artifact.Name == "DefaultTemplate" {
					template_data, err := downloadToBuffer(*response_artifact.URI)
					if err != nil {
						return nil, err
					}

					templatePlan, err = GetPlanFromTemplateJson(template_data)
					if err != nil {
						return nil, err
					}
				}
			}
			// Verify the integrity of the artifact and that it matches our expectations
			if templatePlan.Name != plan.Name {
				return nil, fmt.Errorf("mismatch between planID in marketplace DefaultTemplate")
			}
			if templatePlan.Product != plan.Product {
				return nil, fmt.Errorf("mismatch between product in marketplace response and DefaultTemplate")
			}
			if templatePlan.Publisher != plan.Publisher {
				return nil, fmt.Errorf("mismatch between publisher in marketplace response and DefaultTemplate")
			}
			if templatePlan.Version != plan.Version {
				return nil, fmt.Errorf("mismatch between plan version in marketplace response and DefaultTemplate")
			}
			return artifacts, nil
		}
	}
	return nil, fmt.Errorf("could not find plan %s in marketplace in order to get Cloud Block Store artifacts", plan.Name)
}

func GetProductSummary(ctx context.Context) (appcatalog.SearchClientGetResponse, error) {
	search_client := appcatalog.NewSearchClient()
	return search_client.Get(ctx, "en", "US", "terraform-cbs-provider", []appcatalog.SearchV2FieldName{"All"}, []string{"purestoragemarketplaceadmin"})
}

func GetResourceListFromUiDefinitionUrl(url string) ([]string, error) {
	template_data, err := downloadToBuffer(url)
	if err != nil {
		return nil, err
	}

	return GetResourcesFromTemplateJson(template_data)
}

func getPlanResources(ctx context.Context, plan Plan) ([]string, error) {
	artifacts, err := GetPlanArtifacts(ctx, plan)
	if err != nil {
		return nil, err
	}

	return GetResourceListFromUiDefinitionUrl(*artifacts["createuidefinition"].URI)
}

func getResourcesFromAppDefinitionId(ctx context.Context, azureClient cloud.AzureClientAPI, appDefinitionID string) ([]string, error) {
	resource, err := azureClient.ResourceGet(ctx, appDefinitionID)
	if err != nil {
		return nil, err
	}

	properties := resource.Properties.(map[string]interface{})
	artifacts := properties["artifacts"].([]interface{})
	var resources []string
	for _, artifact := range artifacts {
		details := artifact.(map[string]interface{})
		name := details["name"].(string)
		if name == "ApplicationResourceTemplate" {
			// to get createUiDefinition.json file we simply construct url by
			// removing last resource number (32 bytes long) and replace it with createUiDefinition.json
			uri := details["uri"].(string)
			uri = uri[:len(uri)-32] + "createUiDefinition.json"
			resources, err = GetResourceListFromUiDefinitionUrl(uri)
			if err != nil {
				return nil, errors.New("cannot get resource list from uidefinition url")
			}
			break
		}
	}

	if len(resources) == 0 {
		return nil, errors.New("Resource list must not be empty")
	}

	return resources, err
}

func getResourcesForCurrentDeployment(ctx context.Context, azureClient cloud.AzureClientAPI, d *schema.ResourceData) ([]string, error) {
	var resources []string
	var err error
	if appDefinitionId, ok := d.GetOk("app_definition_id"); ok {
		resources, err = getResourcesFromAppDefinitionId(ctx, azureClient, appDefinitionId.(string))
		if err != nil {
			return nil, fmt.Errorf("cannot get resource list from app_definition_id %s %+v", appDefinitionId.(string), err)
		}
	}

	if planParam, ok := d.GetOk("plan"); ok {
		// We rely on Terraform framework here that this will always be convertible, otherwise we would panic
		planDecoded := planParam.([]interface{})[0].(map[string]interface{})
		plan := Plan{
			Name:      planDecoded["name"].(string),
			Product:   planDecoded["product"].(string),
			Publisher: planDecoded["publisher"].(string),
			Version:   planDecoded["version"].(string),
		}
		resources, err = getPlanResources(ctx, plan)
		if err != nil {
			return nil, fmt.Errorf("could not get resource list from plan %s: %+v", plan.Name, err)
		}
	}

	if len(resources) == 0 {
		return nil, fmt.Errorf("resource list cannot be empty")
	}

	return resources, nil
}

func SetResourceTags(resourceType string, resourceTagList []interface{}, resources []string, tagsMap *map[string]interface{}) error {
	resourceExists := false
	for _, v := range resources {
		if v == resourceType {
			resourceExists = true
		}
	}
	if !resourceExists {
		return fmt.Errorf("provided resource type %s not found", resourceType)
	}
	for _, tagPair := range resourceTagList {
		tagMap := tagPair.(map[string]interface{})
		key := tagMap["name"].(string)
		value := tagMap["value"].(string)
		if _, ok := (*tagsMap)[resourceType]; !ok {
			(*tagsMap)[resourceType] = make(map[string]interface{})
		}
		(*tagsMap)[resourceType].(map[string]interface{})[key] = value
	}

	return nil
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

	parameters.Identity = expandIdentityObject(identities)

	// Fetch the list of resources which will be deployed. We rely on createUiDefinition.json artifact which should
	// be available through the marketplace api in case of public plans. In case that the plan is not available in
	// the marketplace API for any reason we fall back to a statically defined list.
	resources, err := getResourcesForCurrentDeployment(ctx, azureClient, d)
	if err != nil {
		tflog.Warn(ctx, fmt.Sprintf("Could not get resource list for current deployment, falling back to static resource list: %s", err))
		resources = staticResourceList
	}

	tagsMap := make(map[string]interface{})
	if v, ok := d.GetOk("tags"); ok {
		tags := v.(map[string]interface{})
		for _, tag := range resources {
			copyMap := make(map[string]interface{})
			for key, value := range tags {
				copyMap[key] = value
			}
			tagsMap[tag] = copyMap
		}
	}

	if v, ok := d.GetOk("resource_tags"); ok {
		resource_tags := v.([]interface{})
		for _, resource_tag := range resource_tags {
			resource := resource_tag.(map[string]interface{})["resource"].(string)
			resourceTagList := resource_tag.(map[string]interface{})["tag"].([]interface{})

			err = SetResourceTags(resource, resourceTagList, resources, &tagsMap)
			if err != nil {
				return diag.Errorf("cannot set resource tags %+v", err)
			}
		}
	}

	setAppParameter("tagsByResource", tagsMap)

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
