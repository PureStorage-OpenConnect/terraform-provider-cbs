// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostAllOf0 host all of0
//
// swagger:model hostAllOf0
type HostAllOf0 struct {
	ResourceNoID

	HostOAIGen1AllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *HostAllOf0) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceNoID
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceNoID = aO0

	// AO1
	var aO1 HostOAIGen1AllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.HostOAIGen1AllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m HostAllOf0) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ResourceNoID)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.HostOAIGen1AllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this host all of0
func (m *HostAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceNoID
	if err := m.ResourceNoID.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with HostOAIGen1AllOf1
	if err := m.HostOAIGen1AllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *HostAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostAllOf0) UnmarshalBinary(b []byte) error {
	var res HostAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
