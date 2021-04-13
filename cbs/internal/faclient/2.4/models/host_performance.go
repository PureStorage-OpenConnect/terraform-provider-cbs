// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostPerformance host performance
//
// swagger:model HostPerformance
type HostPerformance struct {
	ResourcePerformance
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HostPerformance) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourcePerformance
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourcePerformance = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HostPerformance) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ResourcePerformance)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this host performance
func (m *HostPerformance) Validate(formats strfmt.Registry) error {
	return nil
}
