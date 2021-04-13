// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodPost pod post
//
// swagger:model podPost
type PodPost struct {
	BuiltIn

	PodPostAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PodPost) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuiltIn
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuiltIn = aO0

	// AO1
	var aO1 PodPostAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PodPostAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PodPost) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.BuiltIn)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PodPostAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pod post
func (m *PodPost) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuiltIn
	if err := m.BuiltIn.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PodPostAllOf1
	if err := m.PodPostAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PodPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodPost) UnmarshalBinary(b []byte) error {
	var res PodPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
