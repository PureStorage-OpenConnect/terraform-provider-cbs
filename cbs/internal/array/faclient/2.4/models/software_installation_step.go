// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SoftwareInstallationStep software installation step
//
// swagger:model softwareInstallationStep
type SoftwareInstallationStep struct {
	ResourceFixedNonUniqueName

	StartEndTime

	SoftwareInstallationStepAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SoftwareInstallationStep) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ResourceFixedNonUniqueName
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ResourceFixedNonUniqueName = aO0

	// AO1
	var aO1 StartEndTime
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.StartEndTime = aO1

	// AO2
	var aO2 SoftwareInstallationStepAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.SoftwareInstallationStepAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SoftwareInstallationStep) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.ResourceFixedNonUniqueName)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.StartEndTime)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.SoftwareInstallationStepAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this software installation step
func (m *SoftwareInstallationStep) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ResourceFixedNonUniqueName
	if err := m.ResourceFixedNonUniqueName.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with StartEndTime
	if err := m.StartEndTime.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SoftwareInstallationStepAllOf2
	if err := m.SoftwareInstallationStepAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareInstallationStep) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareInstallationStep) UnmarshalBinary(b []byte) error {
	var res SoftwareInstallationStep
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
