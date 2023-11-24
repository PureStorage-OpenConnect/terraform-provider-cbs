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
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-version"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAzurePlans() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceAzurePlansRead,
		Schema: map[string]*schema.Schema{
			"plans": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"product": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"publisher": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCbsPlanAzure() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceCbsPlanAzureRead,
		Schema: map[string]*schema.Schema{
			"plan_version": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validateVersionPrefixTag,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"publisher": {
				Type:     schema.TypeString,
				Computed: true,
				Elem:     schema.TypeString,
			},
			"product": {
				Type:     schema.TypeString,
				Computed: true,
				Elem:     schema.TypeString,
			},
		},
	}
}

type Plan struct {
	Name      string
	Product   string
	Publisher string
	Version   string
}

type JSONDefaultTemplateResources struct {
	Plan Plan
}

type JSONDefaultTemplate struct {
	Resources []JSONDefaultTemplateResources
}

type VersionedPlan struct {
	Version version.Version
	Plan    Plan
}

type PlanByVersion []VersionedPlan

func (a PlanByVersion) Len() int           { return len(a) }
func (a PlanByVersion) Less(i, j int) bool { return a[i].Version.LessThan(&a[j].Version) }
func (a PlanByVersion) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

var plan_name_regexp = regexp.MustCompile(`^[\w]+_([\d]+)_([\d]+)_([\d]+)$`)

// Retrieve a plan information from Azure DefaultTemplate artifact
func GetPlanFromTemplateJson(data []byte) (*Plan, error) {
	// Parse the default template
	var unmarshalled_data JSONDefaultTemplate
	err := json.Unmarshal(data, &unmarshalled_data)
	if err != nil {
		return nil, err
	}
	if len(unmarshalled_data.Resources) != 1 {
		return nil, fmt.Errorf("DefaultTemplate resources is of unexpected size %d", len(unmarshalled_data.Resources))
	}

	// Validate a plan
	unmarshalled_plan := unmarshalled_data.Resources[0].Plan
	if len(unmarshalled_plan.Name) == 0 || len(unmarshalled_plan.Product) == 0 || len(unmarshalled_plan.Publisher) == 0 || len(unmarshalled_plan.Version) == 0 {
		return nil, errors.New("DefaultTemplate plan has unknown format")
	}

	return &unmarshalled_plan, nil
}

func versionPlans(plans []Plan) ([]VersionedPlan, error) {
	var versioned_plans []VersionedPlan
	for _, plan := range plans {
		match := plan_name_regexp.FindStringSubmatch(plan.Name)
		version, err := version.NewVersion(fmt.Sprintf("%s.%s.%s", match[1], match[2], match[3]))
		if err != nil {
			return nil, fmt.Errorf("Unable to parse version string in plan name: %s", plan.Name)
		}
		versioned_plans = append(versioned_plans, VersionedPlan{*version, plan})
	}

	sort.Sort(sort.Reverse(PlanByVersion(versioned_plans)))
	return versioned_plans, nil
}

func QueryMarketplaceForPlans(ctx context.Context) ([]VersionedPlan, error) {
	productSummary, err := GetProductSummary(ctx)
	if err != nil {
		return nil, err
	}

	var template_plans []Plan
	for _, response_result := range productSummary.Results {
		for _, response_plan := range response_result.Plans {
			if !strings.HasPrefix(*response_plan.PlanID, "cbs_azure") {
				// Exclude any plans which aren't the Pure Cloud Block Store on Azure offering.
				// E.g. the Pure Fusion Storage Endpoint Collection offerings.
				continue
			}

			for _, response_artifact := range response_plan.Artifacts {
				if *response_artifact.Name != "DefaultTemplate" {
					continue
				}

				template_data, err := downloadToBuffer(*response_artifact.URI)
				if err != nil {
					continue
				}

				template_plan, err := GetPlanFromTemplateJson(template_data)
				if err != nil {
					continue
				}

				template_plans = append(template_plans, *template_plan)
			}
		}
	}

	return versionPlans(template_plans)
}

func dataSourceAzurePlansRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	versioned_plans, err := QueryMarketplaceForPlans(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	var plans []map[string]string
	for _, versioned_plan := range versioned_plans {
		plan := map[string]string{}
		plan["name"] = versioned_plan.Plan.Name
		plan["product"] = versioned_plan.Plan.Product
		plan["publisher"] = versioned_plan.Plan.Publisher
		plan["version"] = versioned_plan.Plan.Version
		plans = append(plans, plan)
	}

	if len(plans) == 0 {
		d.SetId("")
	} else {
		d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
		err = d.Set("plans", plans)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return diags
}

func dataSourceCbsPlanAzureRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	versioned_plans, err := QueryMarketplaceForPlans(ctx)
	if err != nil {
		return diag.FromErr(err)
	}

	version_prefix_tag := d.Get("plan_version").(string)
	if version_prefix_tag[len(version_prefix_tag)-1] == 'x' {
		version_prefix_tag = version_prefix_tag[0 : len(version_prefix_tag)-1]
	}

	set := false
	for _, versioned_plan := range versioned_plans {
		match := plan_name_regexp.FindStringSubmatch(versioned_plan.Plan.Name)
		version_tag := fmt.Sprintf("%s.%s.%s", match[1], match[2], match[3])
		if strings.HasPrefix(version_tag, version_prefix_tag) {
			d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
			err = d.Set("name", versioned_plan.Plan.Name)
			if err != nil {
				return diag.FromErr(err)
			}
			err = d.Set("product", versioned_plan.Plan.Product)
			if err != nil {
				return diag.FromErr(err)
			}
			err = d.Set("publisher", versioned_plan.Plan.Publisher)
			if err != nil {
				return diag.FromErr(err)
			}
			err = d.Set("version", versioned_plan.Plan.Version)
			if err != nil {
				return diag.FromErr(err)
			}
			set = true
			break
		}
	}

	if !set {
		return diag.FromErr(errors.New("Specific plan for provided version tag not found"))
	}

	return diags
}
