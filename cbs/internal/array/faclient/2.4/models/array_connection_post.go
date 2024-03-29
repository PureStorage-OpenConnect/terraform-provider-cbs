// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArrayConnectionPost array connection post
//
// swagger:model arrayConnectionPost
type ArrayConnectionPost struct {
	ArrayConnectionPostAllOf0

	ArrayConnectionPostAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ArrayConnectionPost) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ArrayConnectionPostAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ArrayConnectionPostAllOf0 = aO0

	// AO1
	var aO1 ArrayConnectionPostAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ArrayConnectionPostAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ArrayConnectionPost) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.ArrayConnectionPostAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ArrayConnectionPostAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this array connection post
func (m *ArrayConnectionPost) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ArrayConnectionPostAllOf0
	// validation for a type composition with ArrayConnectionPostAllOf1
	if err := m.ArrayConnectionPostAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ArrayConnectionPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArrayConnectionPost) UnmarshalBinary(b []byte) error {
	var res ArrayConnectionPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ArrayConnectionPostAllOf0 array connection post all of0
//
// swagger:model ArrayConnectionPostAllOf0
type ArrayConnectionPostAllOf0 interface{}
