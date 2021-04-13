// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Session session
//
// swagger:model Session
type Session struct {
	BuiltIn

	SessionOAIGenAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Session) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuiltIn
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuiltIn = aO0

	// AO1
	var aO1 SessionOAIGenAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SessionOAIGenAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Session) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BuiltIn)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SessionOAIGenAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this session
func (m *Session) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuiltIn
	if err := m.BuiltIn.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SessionOAIGenAllOf1
	if err := m.SessionOAIGenAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Session) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Session) UnmarshalBinary(b []byte) error {
	var res Session
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
