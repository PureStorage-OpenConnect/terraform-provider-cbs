// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DriveOAIGenAllOf1 Information about a flash, NVRAM, or cache module.
//
// swagger:model driveOaiGenAllOf1
type DriveOAIGenAllOf1 struct {

	// Physical storage capacity of the module in bytes.
	// Read Only: true
	Capacity int64 `json:"capacity,omitempty"`

	// Details about the status of the module if not healthy.
	// Read Only: true
	Details string `json:"details,omitempty"`

	// Storage protocol of the module. Valid values are `NVMe` and `SAS`.
	// Read Only: true
	Protocol string `json:"protocol,omitempty"`

	// Current status of the module. Valid values are `empty`, `failed`, `healthy`, `identifying`, `missing`, `recovering`, `unadmitted`, `unhealthy`, `unrecognized`, and `updating`.
	// Read Only: true
	Status string `json:"status,omitempty"`

	// The type of the module. Valid values are `cache`, `NVRAM`, `SSD`, and `virtual`.
	// Read Only: true
	Type string `json:"type,omitempty"`
}

// Validate validates this drive Oai gen all of1
func (m *DriveOAIGenAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DriveOAIGenAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DriveOAIGenAllOf1) UnmarshalBinary(b []byte) error {
	var res DriveOAIGenAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}