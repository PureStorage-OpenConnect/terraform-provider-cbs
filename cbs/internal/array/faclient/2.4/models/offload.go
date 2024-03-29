// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Offload offload
//
// swagger:model Offload
type Offload struct {
	OffloadPost

	ResourceNoID

	OffloadOAIGenAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *Offload) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 OffloadPost
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.OffloadPost = aO0

	// AO1
	var aO1 ResourceNoID
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ResourceNoID = aO1

	// AO2
	var aO2 OffloadOAIGenAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.OffloadOAIGenAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m Offload) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.OffloadPost)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ResourceNoID)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.OffloadOAIGenAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this offload
func (m *Offload) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OffloadPost
	if err := m.OffloadPost.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ResourceNoID
	if err := m.ResourceNoID.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OffloadOAIGenAllOf2
	if err := m.OffloadOAIGenAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Offload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Offload) UnmarshalBinary(b []byte) error {
	var res Offload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
