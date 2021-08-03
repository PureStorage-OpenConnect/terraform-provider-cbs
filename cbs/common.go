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
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/iancoleman/strcase"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/auth"
)

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

func getSSHPrivateKeyBytes(data *schema.ResourceData) ([]byte, error) {
	if v, ok := data.GetOk("pureuser_private_key"); ok {
		return []byte(v.(string)), nil
	}
	return ioutil.ReadFile(data.Get("pureuser_private_key_path").(string))
}

func generateSecretPayload(data *schema.ResourceData) ([]byte, error) {
	keyContent, err := getSSHPrivateKeyBytes(data)
	if err != nil {
		return nil, err
	}
	bootstrap := auth.NewBootstrapService()
	return bootstrap.GenerateSecretPayload(data.Get("management_endpoint").(string), keyContent)
}
