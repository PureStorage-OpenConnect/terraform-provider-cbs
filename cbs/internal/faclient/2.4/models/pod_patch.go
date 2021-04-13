// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodPatch pod patch
//
// swagger:model podPatch
type PodPatch struct {
	PodPatchAllOf0

	PodPatchAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PodPatch) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PodPatchAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PodPatchAllOf0 = aO0

	// AO1
	var aO1 PodPatchAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PodPatchAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PodPatch) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PodPatchAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PodPatchAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pod patch
func (m *PodPatch) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PodPatchAllOf0
	if err := m.PodPatchAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PodPatchAllOf1
	if err := m.PodPatchAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PodPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodPatch) UnmarshalBinary(b []byte) error {
	var res PodPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}