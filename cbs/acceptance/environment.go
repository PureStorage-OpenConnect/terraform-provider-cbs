/*

	Copyright 2022, Pure Storage Inc.

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

package acceptance

const (
	// Environment variable for controlling the Azure acceptance tests
	// related to deploying an app from the Azure Marketplace
	EnvTfAccAzureSkipMarketplace = "TC_ACC_SKIP_AZURE_MARKETPLACE"

	// Environment variable controlling if the Azure acceptance tests
	// for user az cli login should be run. This testing is not available
	// when using Service Principal auth.
	EnvTfAccSkipUserPrincipalAuth = "TF_ACC_SKIP_USER_PRINCIPAL_AUTH"

	// Environment variable for controlling the Azure acceptance tests
	// related to deploying an CBS Fusion app from an App Definition
	EnvTfAccAzureSkipFusionAppId = "TC_ACC_SKIP_AZURE_FUSION_APP_ID"

	// Environment variable controlling if the Fusion Storage Endpoint
	// Collection Azure acceptance tests should be run.
	EnvTfAccSkipFusionSECAzure = "TF_ACC_SKIP_FUSION_SEC_AZURE"

	// Enviromment variable with path to the Azure acceptance tests
	// parameters file in json format
	EnvTfAccAzureParamsPath = "TEST_ACC_AZURE_PARAMS_PATH"

	// Enviromment variable with path to the Fusion Storage Endpoint
	// Collection Azure acceptance tests parameters file in json format
	EnvTfAccFusionSECAzureParamsPath = "TEST_ACC_FUSION_SEC_AZURE_PARAMS_PATH"

	// Environment variable with path to the AWS acceptance tests
	// parameters file in json format
	EnvTfAccAwsParamsPath = "TEST_ACC_AWS_PARAMS_PATH"
)
