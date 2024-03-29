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

package array

import "context"

// mockable wrapper around the full REST API that is exposed to the terraform resources
type FAClientAPI interface {
	Deactivate(ctx context.Context) error
}

func NewFAClient(host string, secretPayload string) (FAClientAPI, error) {
	return buildFAClient(host, secretPayload)
}
