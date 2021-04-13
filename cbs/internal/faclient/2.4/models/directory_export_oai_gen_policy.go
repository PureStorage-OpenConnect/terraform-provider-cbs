// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirectoryExportOAIGenPolicy The export policy that manages this export. An export can be managed by at most one export policy.
//
// swagger:model directoryExportOaiGenPolicy
type DirectoryExportOAIGenPolicy struct {
	FixedReferenceWithType
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DirectoryExportOAIGenPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 FixedReferenceWithType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.FixedReferenceWithType = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DirectoryExportOAIGenPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.FixedReferenceWithType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this directory export Oai gen policy
func (m *DirectoryExportOAIGenPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FixedReferenceWithType
	if err := m.FixedReferenceWithType.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DirectoryExportOAIGenPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectoryExportOAIGenPolicy) UnmarshalBinary(b []byte) error {
	var res DirectoryExportOAIGenPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
