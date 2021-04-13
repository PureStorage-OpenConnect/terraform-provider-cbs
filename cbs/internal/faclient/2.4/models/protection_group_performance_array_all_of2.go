// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtectionGroupPerformanceArrayAllOf2 protection group performance array all of2
//
// swagger:model protectionGroupPerformanceArrayAllOf2
type ProtectionGroupPerformanceArrayAllOf2 struct {

	// The source array from where the data is replicated.
	Source string `json:"source,omitempty"`

	// The target to where the data is replicated.
	Target string `json:"target,omitempty"`

	// The time when the sample performance data was taken. Measured in milliseconds since the UNIX epoch.
	Time int64 `json:"time,omitempty"`
}

// Validate validates this protection group performance array all of2
func (m *ProtectionGroupPerformanceArrayAllOf2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionGroupPerformanceArrayAllOf2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionGroupPerformanceArrayAllOf2) UnmarshalBinary(b []byte) error {
	var res ProtectionGroupPerformanceArrayAllOf2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
