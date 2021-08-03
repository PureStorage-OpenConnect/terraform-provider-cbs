// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HardwarePatchAllOf1 hardware patch all of1
//
// swagger:model hardwarePatchAllOf1
type HardwarePatchAllOf1 struct {

	// State of an LED used to visually identify the component.
	IdentifyEnabled bool `json:"identify_enabled,omitempty"`

	// Number that identifies the relative position of a hardware component within the array.
	Index int32 `json:"index,omitempty"`
}

// Validate validates this hardware patch all of1
func (m *HardwarePatchAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HardwarePatchAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HardwarePatchAllOf1) UnmarshalBinary(b []byte) error {
	var res HardwarePatchAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}