// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumePost volume post
//
// swagger:model volumePost
type VolumePost struct {
	DestroyedPatchPost

	VolumePostAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VolumePost) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 DestroyedPatchPost
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.DestroyedPatchPost = aO0

	// AO1
	var aO1 VolumePostAllOf1
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VolumePostAllOf1 = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VolumePost) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.DestroyedPatchPost)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VolumePostAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this volume post
func (m *VolumePost) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with DestroyedPatchPost
	if err := m.DestroyedPatchPost.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VolumePostAllOf1
	if err := m.VolumePostAllOf1.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VolumePost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumePost) UnmarshalBinary(b []byte) error {
	var res VolumePost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
