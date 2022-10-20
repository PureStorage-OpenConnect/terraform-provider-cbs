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

package version

const InternalProviderVersion = "dev"

// The ProviderVersion will stay 'dev' for internal builds and will be set to
// proper version on releases
var ProviderVersion = InternalProviderVersion

// This is used to identify the logs with the git commit that the provider was
// built from
var ProviderCommit = ""

// In development we might need more testing options to be enabled so this will
// define if we are running in development mode or not. The ProviderVersion will
// be overwritten by the goreleaser to the release tag for dev builds this will
// stay the default.
func IsDevelopmentMode() bool {
	return ProviderVersion == InternalProviderVersion
}
