// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeOAIGen1AllOf1Pod A reference to the pod.
//
// swagger:model volumeOaiGen1AllOf1Pod
type VolumeOAIGen1AllOf1Pod struct {
	Reference
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VolumeOAIGen1AllOf1Pod) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Reference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Reference = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VolumeOAIGen1AllOf1Pod) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Reference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this volume Oai gen1 all of1 pod
func (m *VolumeOAIGen1AllOf1Pod) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Reference
	if err := m.Reference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeOAIGen1AllOf1Pod) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeOAIGen1AllOf1Pod) UnmarshalBinary(b []byte) error {
	var res VolumeOAIGen1AllOf1Pod
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
