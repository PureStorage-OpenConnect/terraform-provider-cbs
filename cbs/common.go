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
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managedapplications"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/auth"
	"github.com/avast/retry-go/v4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iancoleman/strcase"
)

type azureParameterValue struct {
	valType string
	value   interface{}
}

func toSnake(s string) string {
	snakeStr := strcase.ToSnake(s)
	// Post process the string to correct some edge cases
	if strings.Contains(snakeStr, "i_scsi") {
		snakeStr = strings.ReplaceAll(snakeStr, "i_scsi", "iscsi")
	}
	exp := regexp.MustCompile("ct_([0,1])")
	if exp.MatchString(snakeStr) {
		snakeStr = exp.ReplaceAllString(snakeStr, "ct${1}")
	}
	return snakeStr
}

// returns the corresponding terraform schema parameter name given a template
// parameter name and renamed string map.
func templateToTFParam(cftParam string, renamedMap map[string]string) string {
	var tfParam string
	if renamedParam, ok := renamedMap[cftParam]; ok {
		tfParam = renamedParam
	} else {
		tfParam = toSnake(cftParam)
	}
	return tfParam
}

func convertToStringSlice(vals []interface{}) []string {
	strs := make([]string, len(vals))
	for i := range vals {
		strs[i] = vals[i].(string)
	}
	return strs
}

func getSSHPrivateKeyBytesFromResourceData(data *schema.ResourceData) ([]byte, error) {
	if v, ok := data.GetOk("pureuser_private_key"); ok {
		return []byte(v.(string)), nil
	}
	return ioutil.ReadFile(data.Get("pureuser_private_key_path").(string))
}

func getSSHPrivateKeyBytesFromMap(m map[string]string) ([]byte, error) {
	if v, ok := m["pureuser_private_key"]; ok {
		return []byte(v), nil
	}
	return ioutil.ReadFile(m["pureuser_private_key_path"])
}

func generateSecretPayload(ctx context.Context, data *schema.ResourceData) ([]byte, error) {
	keyContent, err := getSSHPrivateKeyBytesFromResourceData(data)
	if err != nil {
		return nil, err
	}
	bootstrap := auth.NewBootstrapService()
	return bootstrap.GenerateSecretPayload(ctx, data.Get("management_endpoint").(string), keyContent)
}

func optOutDefaultProtectionPolicy(ctx context.Context, attributes map[string]string) error {
	keyContent, err := getSSHPrivateKeyBytesFromMap(attributes)
	if err != nil {
		return err
	}
	bootstrap := auth.NewBootstrapService()
	return bootstrap.OptOutDefaultProtectionPolicy(ctx, attributes["management_endpoint"], keyContent)
}

// Download a file as a byte buffer
func downloadToBuffer(url string) ([]byte, error) {
	var body []byte

	err := retry.Do(func() error {
		response, err := http.Get(url)
		if err != nil {
			return err
		}
		defer response.Body.Close()

		body, err = ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		return nil
	}, retry.Attempts(5), retry.Delay(500*time.Millisecond))

	return body, err
}

func validateAzureManagedApplicationName(v interface{}, k string) (warnings []string, errors []error) {
	value := v.(string)

	if !regexp.MustCompile(`^[-\da-zA-Z]{3,64}$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("%q must be between 3 and 64 characters in length and contains only letters, numbers or hyphens", k))
	}

	return warnings, errors
}

func validateAzureResourceGroupName(v interface{}, k string) (warnings []string, errors []error) {
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

func toAzureManagedResourceGroup(name string) string {
	result := name + "-cbs-mrg"
	return result
}

func expandPlan(input []interface{}) *managedapplications.Plan {
	plan := input[0].(map[string]interface{})

	return &managedapplications.Plan{
		Name:      to.StringPtr(plan["name"].(string)),
		Product:   to.StringPtr(plan["product"].(string)),
		Publisher: to.StringPtr(plan["publisher"].(string)),
		Version:   to.StringPtr(plan["version"].(string)),
	}
}

func setAzureJitAccessPolicy(parameters *managedapplications.Application, d *schema.ResourceData) diag.Diagnostics {
	var (
		approvers     []managedapplications.JitApproverDefinition
		returnedDiags diag.Diagnostics
	)

	if objectIds, ok := d.GetOk("jit_approval_group_object_ids"); ok {
		ids := convertToStringSlice(objectIds.([]interface{}))
		if len(ids) == 0 {
			returnedDiags = diag.Errorf("jit_approval_group_object_ids must not be empty")
		}
		for _, id := range ids {
			approvers = append(approvers, managedapplications.JitApproverDefinition{
				ID:   to.StringPtr(id),
				Type: managedapplications.Group,
			})
		}
	}

	if len(approvers) > 0 {
		parameters.ApplicationProperties.JitAccessPolicy = &managedapplications.ApplicationJitAccessPolicy{
			JitAccessEnabled:         to.BoolPtr(true),
			JitApprovalMode:          managedapplications.JitApprovalModeManualApprove,
			MaximumJitAccessDuration: to.StringPtr("PT8H"),
			JitApprovers:             &approvers,
		}
	}

	return returnedDiags
}

func flattenAzureJitApprovalGroupIds(policy *managedapplications.ApplicationJitAccessPolicy) []string {
	var groupList []string
	if policy != nil {
		for _, g := range *policy.JitApprovers {
			if g.Type == managedapplications.Group {
				groupList = append(groupList, *g.ID)
			}
		}
	}
	return groupList
}

func formatAzureParameters(val interface{}) map[string]azureParameterValue {
	rawParams := val.(map[string]interface{})
	params := make(map[string]azureParameterValue)
	for k, v := range rawParams {
		if v != nil {
			params[k] = azureParameterValue{
				valType: v.(map[string]interface{})["type"].(string),
				value:   v.(map[string]interface{})["value"],
			}
		}
	}
	return params
}
