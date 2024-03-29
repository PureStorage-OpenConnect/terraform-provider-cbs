//go:build go1.18
// +build go1.18

// Code generated by Microsoft (R) AutoRest Code Generator (autorest: 3.9.2, generator: @autorest/go@4.0.0-preview.44)
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package appcatalog

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type Artifact.
func (a Artifact) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifactType", a.ArtifactType)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "uri", a.URI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Artifact.
func (a *Artifact) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifactType":
			err = unpopulate(val, "ArtifactType", &a.ArtifactType)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "uri":
			err = unpopulate(val, "URI", &a.URI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "error", e.Error)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponse.
func (e *ErrorResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "error":
			err = unpopulate(val, "Error", &e.Error)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponseDetails.
func (e ErrorResponseDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "code", e.Code)
	populate(objectMap, "message", e.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorResponseDetails.
func (e *ErrorResponseDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &e.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FacetValue.
func (f FacetValue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "count", f.Count)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetValue.
func (f *FacetValue) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "count":
			err = unpopulate(val, "Count", &f.Count)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type FacetsItem.
func (f FacetsItem) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "facetValues", f.FacetValues)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type FacetsItem.
func (f *FacetsItem) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", f, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "facetValues":
			err = unpopulate(val, "FacetValues", &f.FacetValues)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", f, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LeadGeneration.
func (l LeadGeneration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "productId", l.ProductID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LeadGeneration.
func (l *LeadGeneration) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "productId":
			err = unpopulate(val, "ProductID", &l.ProductID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Link.
func (l Link) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", l.DisplayName)
	populate(objectMap, "id", l.ID)
	populate(objectMap, "uri", l.URI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Link.
func (l *Link) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &l.DisplayName)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "ID", &l.ID)
			delete(rawMsg, key)
		case "uri":
			err = unpopulate(val, "URI", &l.URI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MarketStartPrice.
func (m MarketStartPrice) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "currency", m.Currency)
	populate(objectMap, "market", m.Market)
	populate(objectMap, "meterUnits", m.MeterUnits)
	populate(objectMap, "minMeterPrice", m.MinMeterPrice)
	populate(objectMap, "minTermPrice", m.MinTermPrice)
	populate(objectMap, "termUnits", m.TermUnits)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MarketStartPrice.
func (m *MarketStartPrice) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "currency":
			err = unpopulate(val, "Currency", &m.Currency)
			delete(rawMsg, key)
		case "market":
			err = unpopulate(val, "Market", &m.Market)
			delete(rawMsg, key)
		case "meterUnits":
			err = unpopulate(val, "MeterUnits", &m.MeterUnits)
			delete(rawMsg, key)
		case "minMeterPrice":
			err = unpopulate(val, "MinMeterPrice", &m.MinMeterPrice)
			delete(rawMsg, key)
		case "minTermPrice":
			err = unpopulate(val, "MinTermPrice", &m.MinTermPrice)
			delete(rawMsg, key)
		case "termUnits":
			err = unpopulate(val, "TermUnits", &m.TermUnits)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PlanMetadata.
func (p PlanMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "altStackReference", p.AltStackReference)
	populate(objectMap, "generation", p.Generation)
	populate(objectMap, "relatedSkus", p.RelatedSKUs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PlanMetadata.
func (p *PlanMetadata) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "altStackReference":
			err = unpopulate(val, "AltStackReference", &p.AltStackReference)
			delete(rawMsg, key)
		case "generation":
			err = unpopulate(val, "Generation", &p.Generation)
			delete(rawMsg, key)
		case "relatedSkus":
			err = unpopulate(val, "RelatedSKUs", &p.RelatedSKUs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PlanSKURelation.
func (p PlanSKURelation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "relationType", p.RelationType)
	populate(objectMap, "sku", p.SKU)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PlanSKURelation.
func (p *PlanSKURelation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "relationType":
			err = unpopulate(val, "RelationType", &p.RelationType)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &p.SKU)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PlanSKURelationSKU.
func (p PlanSKURelationSKU) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "generation", p.Generation)
	populate(objectMap, "identity", p.Identity)
	populate(objectMap, "name", p.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PlanSKURelationSKU.
func (p *PlanSKURelationSKU) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "generation":
			err = unpopulate(val, "Generation", &p.Generation)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &p.Identity)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &p.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PlanSummary.
func (p PlanSummary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "artifacts", p.Artifacts)
	populate(objectMap, "cspState", p.CspState)
	populate(objectMap, "description", p.Description)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "displayRank", p.DisplayRank)
	populate(objectMap, "isPrivate", p.IsPrivate)
	populate(objectMap, "isQuantifiable", p.IsQuantifiable)
	populate(objectMap, "leadGeneration", p.LeadGeneration)
	populate(objectMap, "metadata", p.Metadata)
	populate(objectMap, "planId", p.PlanID)
	populate(objectMap, "pricingTypes", p.PricingTypes)
	populate(objectMap, "purchaseDurationDiscounts", p.PurchaseDurationDiscounts)
	populate(objectMap, "skuId", p.SKUID)
	populate(objectMap, "summary", p.Summary)
	populate(objectMap, "type", p.Type)
	populate(objectMap, "uiDefinitionUri", p.UIDefinitionURI)
	populate(objectMap, "uniquePlanId", p.UniquePlanID)
	populate(objectMap, "vmArchitectureType", p.VMArchitectureType)
	populate(objectMap, "vmSecurityTypes", p.VMSecurityTypes)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PlanSummary.
func (p *PlanSummary) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "artifacts":
			err = unpopulate(val, "Artifacts", &p.Artifacts)
			delete(rawMsg, key)
		case "cspState":
			err = unpopulate(val, "CspState", &p.CspState)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &p.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "displayRank":
			err = unpopulate(val, "DisplayRank", &p.DisplayRank)
			delete(rawMsg, key)
		case "isPrivate":
			err = unpopulate(val, "IsPrivate", &p.IsPrivate)
			delete(rawMsg, key)
		case "isQuantifiable":
			err = unpopulate(val, "IsQuantifiable", &p.IsQuantifiable)
			delete(rawMsg, key)
		case "leadGeneration":
			err = unpopulate(val, "LeadGeneration", &p.LeadGeneration)
			delete(rawMsg, key)
		case "metadata":
			err = unpopulate(val, "Metadata", &p.Metadata)
			delete(rawMsg, key)
		case "planId":
			err = unpopulate(val, "PlanID", &p.PlanID)
			delete(rawMsg, key)
		case "pricingTypes":
			err = unpopulate(val, "PricingTypes", &p.PricingTypes)
			delete(rawMsg, key)
		case "purchaseDurationDiscounts":
			err = unpopulate(val, "PurchaseDurationDiscounts", &p.PurchaseDurationDiscounts)
			delete(rawMsg, key)
		case "skuId":
			err = unpopulate(val, "SKUID", &p.SKUID)
			delete(rawMsg, key)
		case "summary":
			err = unpopulate(val, "Summary", &p.Summary)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &p.Type)
			delete(rawMsg, key)
		case "uiDefinitionUri":
			err = unpopulate(val, "UIDefinitionURI", &p.UIDefinitionURI)
			delete(rawMsg, key)
		case "uniquePlanId":
			err = unpopulate(val, "UniquePlanID", &p.UniquePlanID)
			delete(rawMsg, key)
		case "vmArchitectureType":
			err = unpopulate(val, "VMArchitectureType", &p.VMArchitectureType)
			delete(rawMsg, key)
		case "vmSecurityTypes":
			err = unpopulate(val, "VMSecurityTypes", &p.VMSecurityTypes)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PlanSummaryLeadGeneration.
func (p PlanSummaryLeadGeneration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "productId", p.ProductID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PlanSummaryLeadGeneration.
func (p *PlanSummaryLeadGeneration) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "productId":
			err = unpopulate(val, "ProductID", &p.ProductID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PlanSummaryMetadata.
func (p PlanSummaryMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "altStackReference", p.AltStackReference)
	populate(objectMap, "generation", p.Generation)
	populate(objectMap, "relatedSkus", p.RelatedSKUs)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PlanSummaryMetadata.
func (p *PlanSummaryMetadata) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "altStackReference":
			err = unpopulate(val, "AltStackReference", &p.AltStackReference)
			delete(rawMsg, key)
		case "generation":
			err = unpopulate(val, "Generation", &p.Generation)
			delete(rawMsg, key)
		case "relatedSkus":
			err = unpopulate(val, "RelatedSKUs", &p.RelatedSKUs)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ProductSummary.
func (p ProductSummary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "azureBenefit", p.AzureBenefit)
	populate(objectMap, "badges", p.Badges)
	populate(objectMap, "categoryIds", p.CategoryIDs)
	populate(objectMap, "cspLegalTermsUri", p.CspLegalTermsURI)
	populate(objectMap, "description", p.Description)
	populate(objectMap, "displayName", p.DisplayName)
	populate(objectMap, "largeIconUri", p.LargeIconURI)
	populate(objectMap, "legalTermsType", p.LegalTermsType)
	populate(objectMap, "legalTermsUri", p.LegalTermsURI)
	populate(objectMap, "linkedAddIns", p.LinkedAddIns)
	populate(objectMap, "links", p.Links)
	populate(objectMap, "mediumIconUri", p.MediumIconURI)
	populate(objectMap, "operatingSystems", p.OperatingSystems)
	populate(objectMap, "plans", p.Plans)
	populate(objectMap, "popularity", p.Popularity)
	populate(objectMap, "pricingTypes", p.PricingTypes)
	populate(objectMap, "privacyPolicyUri", p.PrivacyPolicyURI)
	populate(objectMap, "productType", p.ProductType)
	populate(objectMap, "publisherDisplayName", p.PublisherDisplayName)
	populate(objectMap, "publisherId", p.PublisherID)
	populate(objectMap, "publisherType", p.PublisherType)
	populate(objectMap, "publishingStage", p.PublishingStage)
	populate(objectMap, "ratingAverage", p.RatingAverage)
	populate(objectMap, "ratingBuckets", p.RatingBuckets)
	populate(objectMap, "screenshots", p.Screenshots)
	populate(objectMap, "smallIconUri", p.SmallIconURI)
	populate(objectMap, "startingPrice", p.StartingPrice)
	populate(objectMap, "summary", p.Summary)
	populate(objectMap, "supportUri", p.SupportURI)
	populate(objectMap, "uniqueProductId", p.UniqueProductID)
	populate(objectMap, "vmArchitectureTypes", p.VMArchitectureTypes)
	populate(objectMap, "vmImageGenerations", p.VMImageGenerations)
	populate(objectMap, "vmSecurityTypes", p.VMSecurityTypes)
	populate(objectMap, "version", p.Version)
	populate(objectMap, "wideIconUri", p.WideIconURI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ProductSummary.
func (p *ProductSummary) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "azureBenefit":
			err = unpopulate(val, "AzureBenefit", &p.AzureBenefit)
			delete(rawMsg, key)
		case "badges":
			err = unpopulate(val, "Badges", &p.Badges)
			delete(rawMsg, key)
		case "categoryIds":
			err = unpopulate(val, "CategoryIDs", &p.CategoryIDs)
			delete(rawMsg, key)
		case "cspLegalTermsUri":
			err = unpopulate(val, "CspLegalTermsURI", &p.CspLegalTermsURI)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &p.Description)
			delete(rawMsg, key)
		case "displayName":
			err = unpopulate(val, "DisplayName", &p.DisplayName)
			delete(rawMsg, key)
		case "largeIconUri":
			err = unpopulate(val, "LargeIconURI", &p.LargeIconURI)
			delete(rawMsg, key)
		case "legalTermsType":
			err = unpopulate(val, "LegalTermsType", &p.LegalTermsType)
			delete(rawMsg, key)
		case "legalTermsUri":
			err = unpopulate(val, "LegalTermsURI", &p.LegalTermsURI)
			delete(rawMsg, key)
		case "linkedAddIns":
			err = unpopulate(val, "LinkedAddIns", &p.LinkedAddIns)
			delete(rawMsg, key)
		case "links":
			err = unpopulate(val, "Links", &p.Links)
			delete(rawMsg, key)
		case "mediumIconUri":
			err = unpopulate(val, "MediumIconURI", &p.MediumIconURI)
			delete(rawMsg, key)
		case "operatingSystems":
			err = unpopulate(val, "OperatingSystems", &p.OperatingSystems)
			delete(rawMsg, key)
		case "plans":
			err = unpopulate(val, "Plans", &p.Plans)
			delete(rawMsg, key)
		case "popularity":
			err = unpopulate(val, "Popularity", &p.Popularity)
			delete(rawMsg, key)
		case "pricingTypes":
			err = unpopulate(val, "PricingTypes", &p.PricingTypes)
			delete(rawMsg, key)
		case "privacyPolicyUri":
			err = unpopulate(val, "PrivacyPolicyURI", &p.PrivacyPolicyURI)
			delete(rawMsg, key)
		case "productType":
			err = unpopulate(val, "ProductType", &p.ProductType)
			delete(rawMsg, key)
		case "publisherDisplayName":
			err = unpopulate(val, "PublisherDisplayName", &p.PublisherDisplayName)
			delete(rawMsg, key)
		case "publisherId":
			err = unpopulate(val, "PublisherID", &p.PublisherID)
			delete(rawMsg, key)
		case "publisherType":
			err = unpopulate(val, "PublisherType", &p.PublisherType)
			delete(rawMsg, key)
		case "publishingStage":
			err = unpopulate(val, "PublishingStage", &p.PublishingStage)
			delete(rawMsg, key)
		case "ratingAverage":
			err = unpopulate(val, "RatingAverage", &p.RatingAverage)
			delete(rawMsg, key)
		case "ratingBuckets":
			err = unpopulate(val, "RatingBuckets", &p.RatingBuckets)
			delete(rawMsg, key)
		case "screenshots":
			err = unpopulate(val, "Screenshots", &p.Screenshots)
			delete(rawMsg, key)
		case "smallIconUri":
			err = unpopulate(val, "SmallIconURI", &p.SmallIconURI)
			delete(rawMsg, key)
		case "startingPrice":
			err = unpopulate(val, "StartingPrice", &p.StartingPrice)
			delete(rawMsg, key)
		case "summary":
			err = unpopulate(val, "Summary", &p.Summary)
			delete(rawMsg, key)
		case "supportUri":
			err = unpopulate(val, "SupportURI", &p.SupportURI)
			delete(rawMsg, key)
		case "uniqueProductId":
			err = unpopulate(val, "UniqueProductID", &p.UniqueProductID)
			delete(rawMsg, key)
		case "vmArchitectureTypes":
			err = unpopulate(val, "VMArchitectureTypes", &p.VMArchitectureTypes)
			delete(rawMsg, key)
		case "vmImageGenerations":
			err = unpopulate(val, "VMImageGenerations", &p.VMImageGenerations)
			delete(rawMsg, key)
		case "vmSecurityTypes":
			err = unpopulate(val, "VMSecurityTypes", &p.VMSecurityTypes)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &p.Version)
			delete(rawMsg, key)
		case "wideIconUri":
			err = unpopulate(val, "WideIconURI", &p.WideIconURI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PurchaseDurationDiscount.
func (p PurchaseDurationDiscount) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "discountPercentage", p.DiscountPercentage)
	populate(objectMap, "duration", p.Duration)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PurchaseDurationDiscount.
func (p *PurchaseDurationDiscount) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "discountPercentage":
			err = unpopulate(val, "DiscountPercentage", &p.DiscountPercentage)
			delete(rawMsg, key)
		case "duration":
			err = unpopulate(val, "Duration", &p.Duration)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RelatedSKU.
func (r RelatedSKU) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "generation", r.Generation)
	populate(objectMap, "identity", r.Identity)
	populate(objectMap, "name", r.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RelatedSKU.
func (r *RelatedSKU) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "generation":
			err = unpopulate(val, "Generation", &r.Generation)
			delete(rawMsg, key)
		case "identity":
			err = unpopulate(val, "Identity", &r.Identity)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SearchResponse.
func (s SearchResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "facets", s.Facets)
	populate(objectMap, "nextLink", s.NextLink)
	populate(objectMap, "results", s.Results)
	populate(objectMap, "showingResultsFor", s.ShowingResultsFor)
	populate(objectMap, "totalCount", s.TotalCount)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SearchResponse.
func (s *SearchResponse) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "facets":
			err = unpopulate(val, "Facets", &s.Facets)
			delete(rawMsg, key)
		case "nextLink":
			err = unpopulate(val, "NextLink", &s.NextLink)
			delete(rawMsg, key)
		case "results":
			err = unpopulate(val, "Results", &s.Results)
			delete(rawMsg, key)
		case "showingResultsFor":
			err = unpopulate(val, "ShowingResultsFor", &s.ShowingResultsFor)
			delete(rawMsg, key)
		case "totalCount":
			err = unpopulate(val, "TotalCount", &s.TotalCount)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
