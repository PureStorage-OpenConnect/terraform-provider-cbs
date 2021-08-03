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
	"encoding/json"
	"fmt"
	"net/http"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/version"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/managedapplications"
	"github.com/Azure/go-autorest/autorest/to"
	goversion "github.com/hashicorp/go-version"
)

const (
	// The version mapper only supports purity release versions starting at 6.1.6
	minPurityVersionString = "6.1.6"

	// TODO finalize public location and URL
	versionConfigURL = "https://github.dev.purestorage.com/raw/FlashArray/terraform-provider-cbs-utils/master/version_map.json"
)

var minPurityVersion *goversion.Version

type VersionMapper interface {
	GetCFTForVersion(purityVersionString string) (string, error)
	GetManagedAppPlanForVersion(purityVersionString string) (*managedapplications.Plan, error)
}

type awsDetails struct {
	TemplateURL string `json:"template_url"`
}

type azureDetails struct {
	PlanName      string `json:"plan_name"`
	PlanProduct   string `json:"plan_product"`
	PlanPublisher string `json:"plan_publisher"`
	PlanVersion   string `json:"plan_version"`
}

type versionDetails struct {
	AWS                awsDetails   `json:"aws"`
	Azure              azureDetails `json:"azure"`
	MinProviderVersion string       `json:"min_provider_version"`
	MaxProviderVersion string       `json:"max_provider_version"`
}

type versionDetailsMap map[string]versionDetails

type versionConfig struct {
	Versions versionDetailsMap `json:"purity_versions"`
}

type versionMapperImpl struct {
	providerVersion string
	configURL       string
}

func init() {
	var err error
	if minPurityVersion, err = goversion.NewVersion(minPurityVersionString); err != nil {
		panic(err)
	}
}

func NewVersionMapper() VersionMapper {
	return &versionMapperImpl{
		providerVersion: version.ProviderVersion,
		configURL:       versionConfigURL,
	}
}

func (vm *versionMapperImpl) GetCFTForVersion(purityVersionString string) (string, error) {
	details, err := vm.getAndValidateVersionDetails(purityVersionString)
	if err != nil {
		return "", err
	}
	return details.AWS.TemplateURL, nil
}

func (vm *versionMapperImpl) GetManagedAppPlanForVersion(purityVersionString string) (*managedapplications.Plan, error) {
	details, err := vm.getAndValidateVersionDetails(purityVersionString)
	if err != nil {
		return nil, err
	}
	plan := &managedapplications.Plan{
		Name:      to.StringPtr(details.Azure.PlanName),
		Product:   to.StringPtr(details.Azure.PlanProduct),
		Publisher: to.StringPtr(details.Azure.PlanPublisher),
		Version:   to.StringPtr(details.Azure.PlanVersion),
	}
	return plan, nil
}

func (vm *versionMapperImpl) getAndValidateVersionDetails(purityVersionString string) (*versionDetails, error) {

	config, err := vm.retrieveVersionDetailsFile()
	if err != nil {
		return nil, err
	}

	versionDetails, ok := config.Versions[purityVersionString]
	if !ok {
		purityVersion, err := goversion.NewVersion(purityVersionString)

		if err == nil && purityVersion.LessThan(minPurityVersion) {
			return nil, fmt.Errorf("invalid version %s. The \"version\" parameter only supports purity versions %s and later",
				purityVersion, minPurityVersion)
		} else {
			return nil, fmt.Errorf("invalid version %s", purityVersionString)
		}

	} else if versionDetails.MaxProviderVersion != "" || versionDetails.MinProviderVersion != "" {
		maxProviderVersionString := versionDetails.MaxProviderVersion
		minProviderVersionString := versionDetails.MinProviderVersion

		providerVersion := goversion.Must(goversion.NewVersion(vm.providerVersion))

		if maxProviderVersionString != "" {
			maxProviderVersion := goversion.Must(goversion.NewVersion(maxProviderVersionString))

			if providerVersion.GreaterThan(maxProviderVersion) {
				return nil, fmt.Errorf("purity version %s is only supported by provider versions %s and earlier. Current provider version: %s",
					purityVersionString, maxProviderVersionString, vm.providerVersion)
			}
		}
		if minProviderVersionString != "" {
			minProviderVersion := goversion.Must(goversion.NewVersion(minProviderVersionString))

			if providerVersion.LessThan(minProviderVersion) {
				return nil, fmt.Errorf("purity version %s is only supported by provider versions %s and later. Current provider version: %s",
					purityVersionString, minProviderVersionString, vm.providerVersion)
			}
		}
	}

	return &versionDetails, nil
}

func (vm *versionMapperImpl) retrieveVersionDetailsFile() (*versionConfig, error) {

	res, err := http.Get(vm.configURL)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve purity version configuration: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received invalid status code %d when attempting to retrieve purity version configuration", res.StatusCode)
	}

	var config versionConfig
	if err := json.NewDecoder(res.Body).Decode(&config); err != nil {
		panic(err)
	}
	return &config, nil
}
