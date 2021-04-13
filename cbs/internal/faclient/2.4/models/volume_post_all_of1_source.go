// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumePostAllOf1Source The source volume of a volume copy.
//
// swagger:model volumePostAllOf1Source
type VolumePostAllOf1Source struct {
	Reference
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VolumePostAllOf1Source) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Reference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Reference = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VolumePostAllOf1Source) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Reference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this volume post all of1 source
func (m *VolumePostAllOf1Source) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Reference
	if err := m.Reference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VolumePostAllOf1Source) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumePostAllOf1Source) UnmarshalBinary(b []byte) error {
	var res VolumePostAllOf1Source
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
