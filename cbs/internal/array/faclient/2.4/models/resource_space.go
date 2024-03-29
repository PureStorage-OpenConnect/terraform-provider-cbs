// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourceSpace resource space
//
// swagger:model resourceSpace
type ResourceSpace struct {
	BuiltIn

	ResourceSpaceAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourceSpace) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuiltIn
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuiltIn = aO0

	// AO1
	var aO1 ResourceSpaceAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ResourceSpaceAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourceSpace) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BuiltIn)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ResourceSpaceAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource space
func (m *ResourceSpace) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuiltIn
	if err := m.BuiltIn.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ResourceSpaceAllOf1
	if err := m.ResourceSpaceAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ResourceSpace) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceSpace) UnmarshalBinary(b []byte) error {
	var res ResourceSpace
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
