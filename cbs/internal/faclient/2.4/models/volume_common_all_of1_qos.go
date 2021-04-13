// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeCommonAllOf1Qos Displays QoS limit information.
//
// swagger:model volumeCommonAllOf1Qos
type VolumeCommonAllOf1Qos struct {
	Qos
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VolumeCommonAllOf1Qos) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Qos
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Qos = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VolumeCommonAllOf1Qos) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Qos)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this volume common all of1 qos
func (m *VolumeCommonAllOf1Qos) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Qos
	if err := m.Qos.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeCommonAllOf1Qos) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeCommonAllOf1Qos) UnmarshalBinary(b []byte) error {
	var res VolumeCommonAllOf1Qos
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
