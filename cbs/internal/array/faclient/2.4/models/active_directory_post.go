// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActiveDirectoryPost active directory post
//
// swagger:model activeDirectoryPost
type ActiveDirectoryPost struct {
	ActiveDirectoryPostAllOf0

	ActiveDirectoryPostAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ActiveDirectoryPost) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ActiveDirectoryPostAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ActiveDirectoryPostAllOf0 = aO0

	// AO1
	var aO1 ActiveDirectoryPostAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ActiveDirectoryPostAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ActiveDirectoryPost) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ActiveDirectoryPostAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ActiveDirectoryPostAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this active directory post
func (m *ActiveDirectoryPost) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ActiveDirectoryPostAllOf0
	// validation for a type composition with ActiveDirectoryPostAllOf1
	if err := m.ActiveDirectoryPostAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectoryPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryPost) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ActiveDirectoryPostAllOf0 active directory post all of0
//
// swagger:model ActiveDirectoryPostAllOf0
type ActiveDirectoryPostAllOf0 interface{}
