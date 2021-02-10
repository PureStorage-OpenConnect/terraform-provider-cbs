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
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	awsRegionVar = "AWS_DEFAULT_REGION"
)

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
		},
		ResourcesMap: map[string]*schema.Resource{
			"cbs_array_aws": resourceArrayAWS(),
		},
		ConfigureContextFunc: configureProvider,
	}
}

func configureProvider(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	region, diags := awsRegion(d)

	if diags.HasError() {
		return nil, diags
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	cftSvc := cloudformation.New(sess)

	return cftSvc, diags
}

func awsRegion(d *schema.ResourceData) (string, diag.Diagnostics) {
	var diags diag.Diagnostics

	if awsL, ok := d.Get("aws").([]interface{}); ok && len(awsL) > 0 && awsL[0] != nil {
		awsM := awsL[0].(map[string]interface{})
		if region, ok := awsM["region"]; ok {
			return region.(string), diags
		}
	}

	if region := os.Getenv(awsRegionVar); region != "" {
		return region, diags
	}

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary: fmt.Sprintf("No AWS region specified. The AWS region must be provided either in "+
			"the provider configuration block or with the %s environment variable.", awsRegionVar),
	})
	return "", diags
}
