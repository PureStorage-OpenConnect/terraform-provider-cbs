// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SoftwareInstallationAllOf2Software Referenced `software` to which the upgrade belongs.
//
// swagger:model softwareInstallationAllOf2Software
type SoftwareInstallationAllOf2Software struct {
	Reference
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SoftwareInstallationAllOf2Software) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Reference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Reference = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SoftwareInstallationAllOf2Software) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Reference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this software installation all of2 software
func (m *SoftwareInstallationAllOf2Software) Validate(formats strfmt.Registry) error {
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
func (m *SoftwareInstallationAllOf2Software) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareInstallationAllOf2Software) UnmarshalBinary(b []byte) error {
	var res SoftwareInstallationAllOf2Software
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
