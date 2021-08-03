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

package service

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/version"
)

const testConfigDir = "../../../testing"
const testConfigURL = "http://localhost:%d/test_version_config.json"

// Make sure that we can retrieve the real config file
func TestVersionMapper_retrieveRealConfig(t *testing.T) {
	mapper := versionMapperImpl{
		providerVersion: version.ProviderVersion,
		configURL:       versionConfigURL,
	}
	if _, err := mapper.retrieveVersionDetailsFile(); err != nil {
		t.Errorf("failed to retrieve config file from %s: %s", versionConfigURL, err)
	}
}

func launchServer(port int) *http.Server {

	server := http.Server{
		Handler: http.FileServer(http.Dir(testConfigDir)),
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	go server.Serve(l)

	return &server
}

func TestVersionMapper_badRequest(t *testing.T) {
	port := 8080
	server := launchServer(port)
	defer server.Shutdown(context.Background())

	tests := map[string]struct {
		purityVersion   string
		providerVersion string
		expectedError   string
	}{
		"purityVersionTooEarly": {
			purityVersion:   "6.1.0",
			providerVersion: "0.5.0",
			expectedError:   "invalid version 6.1.0. The \"version\" parameter only supports purity versions 6.1.6 and later",
		},
		"purityVersionNotFound": {
			purityVersion:   "6.2.0",
			providerVersion: "0.5.0",
			expectedError:   "invalid version 6.2.0",
		},
		"purityVersionInvalid": {
			purityVersion:   "foobar",
			providerVersion: "0.5.0",
			expectedError:   "invalid version foobar",
		},
		"providerVersionTooEarly": {
			purityVersion:   "6.1.7",
			providerVersion: "0.5.0",
			expectedError:   "purity version 6.1.7 is only supported by provider versions 0.6.0 and later. Current provider version: 0.5.0",
		},
		"providerVersionTooLate": {
			purityVersion:   "6.1.6",
			providerVersion: "0.7.0",
			expectedError:   "purity version 6.1.6 is only supported by provider versions 0.6.0 and earlier. Current provider version: 0.7.0",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mapper := versionMapperImpl{
				providerVersion: test.providerVersion,
				configURL:       fmt.Sprintf(testConfigURL, port),
			}
			_, err := mapper.GetCFTForVersion(test.purityVersion)
			assert.NotNil(t, err)

			if err.Error() != test.expectedError {
				t.Errorf("expected error: %s, actual error: %s", test.expectedError, err.Error())
			}
		})
	}
}

func TestVersionMapper_badURL(t *testing.T) {
	port := 8081
	server := launchServer(port)
	defer server.Shutdown(context.Background())

	invalidConfigURL := fmt.Sprintf("http://localhost:%d/test_version_config_invalid.json", port)

	mapper := versionMapperImpl{
		providerVersion: "0.5.0",
		configURL:       invalidConfigURL,
	}

	_, err := mapper.GetCFTForVersion("6.1.6")
	assert.NotNil(t, err)

	expectedError := "received invalid status code 404 when attempting to retrieve purity version configuration"
	assert.EqualError(t, err, expectedError)
}

func TestVersionMapper_success(t *testing.T) {
	port := 8082
	server := launchServer(port)
	defer server.Shutdown(context.Background())

	tests := map[string]struct {
		purityVersion    string
		providerVersion  string
		azurePlanVersion string
	}{
		"test_616": {
			purityVersion:    "6.1.6",
			providerVersion:  "0.5.0",
			azurePlanVersion: "1.0.0",
		},
		"test_617": {
			purityVersion:    "6.1.7",
			providerVersion:  "0.7.0",
			azurePlanVersion: "1.0.9",
		},
		"test_618": {
			purityVersion:    "6.1.8",
			providerVersion:  "foo", //no min or max provider version for 6.1.8 test config, this should be able to be anything
			azurePlanVersion: "1.0.5",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			mapper := versionMapperImpl{
				providerVersion: test.providerVersion,
				configURL:       fmt.Sprintf(testConfigURL, port),
			}

			// Test AWS
			awsURL, err := mapper.GetCFTForVersion(test.purityVersion)
			assert.Nil(t, err)

			expectedURL := fmt.Sprintf("aws_url_%s", test.purityVersion)
			assert.Equal(t, expectedURL, awsURL)

			// Test Azure
			plan, err := mapper.GetManagedAppPlanForVersion(test.purityVersion)
			assert.Nil(t, err)

			assert.Equal(t, fmt.Sprintf("azure_plan_%s", test.purityVersion), *plan.Name)
			assert.Equal(t, fmt.Sprintf("azure_product_%s", test.purityVersion), *plan.Product)
			assert.Equal(t, fmt.Sprintf("azure_publisher_%s", test.purityVersion), *plan.Publisher)
			assert.Equal(t, test.azurePlanVersion, *plan.Version)

		})
	}
}
