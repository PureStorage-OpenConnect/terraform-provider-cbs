// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeGroupPost volume group post
//
// swagger:model volumeGroupPost
type VolumeGroupPost struct {
	BuiltIn

	VolumeGroup

	VolumeGroupPostAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VolumeGroupPost) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuiltIn
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuiltIn = aO0

	// AO1
	var aO1 VolumeGroup
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VolumeGroup = aO1

	// AO2
	var aO2 VolumeGroupPostAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.VolumeGroupPostAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VolumeGroupPost) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.BuiltIn)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VolumeGroup)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.VolumeGroupPostAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this volume group post
func (m *VolumeGroupPost) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuiltIn
	if err := m.BuiltIn.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VolumeGroup
	if err := m.VolumeGroup.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VolumeGroupPostAllOf2
	if err := m.VolumeGroupPostAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeGroupPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeGroupPost) UnmarshalBinary(b []byte) error {
	var res VolumeGroupPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
