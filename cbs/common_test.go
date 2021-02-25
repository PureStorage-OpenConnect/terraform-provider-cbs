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
	"testing"
)

func TestCbs_toSnake(t *testing.T) {

	m := map[string]string{
		"ArrayName":              "array_name",
		"ISCSISubnet":            "iscsi_subnet",
		"ReplicationEndpointCT0": "replication_endpoint_ct0",
		"iSCSIEndpointCT1":       "iscsi_endpoint_ct1",
	}

	for camel, snakeExp := range m {
		snakeStr := toSnake(camel)
		if snakeStr != snakeExp {
			t.Errorf("error converting camel case to snake case, expected: %s, actual: %s", snakeExp, snakeStr)
		}
	}
}
