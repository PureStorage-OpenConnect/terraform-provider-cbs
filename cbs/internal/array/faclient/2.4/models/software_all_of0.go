// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SoftwareAllOf0 software all of0
//
// swagger:model softwareAllOf0
type SoftwareAllOf0 struct {
	ResourceFixedNonUniqueName

	SoftwareOAIGen1AllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SoftwareAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceFixedNonUniqueName
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceFixedNonUniqueName = aO0

	// AO1
	var aO1 SoftwareOAIGen1AllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SoftwareOAIGen1AllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SoftwareAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ResourceFixedNonUniqueName)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SoftwareOAIGen1AllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this software all of0
func (m *SoftwareAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceFixedNonUniqueName
	if err := m.ResourceFixedNonUniqueName.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SoftwareOAIGen1AllOf1
	if err := m.SoftwareOAIGen1AllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareAllOf0) UnmarshalBinary(b []byte) error {
	var res SoftwareAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}