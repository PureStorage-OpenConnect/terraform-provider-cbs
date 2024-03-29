// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirectorySnapshotPatchAllOf1Policy The snapshot policy that manages this snapshot. Set to `name` or `id` to `""` to clear the policy.
//
// swagger:model directorySnapshotPatchAllOf1Policy
type DirectorySnapshotPatchAllOf1Policy struct {
	Reference
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *DirectorySnapshotPatchAllOf1Policy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Reference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Reference = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m DirectorySnapshotPatchAllOf1Policy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.Reference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this directory snapshot patch all of1 policy
func (m *DirectorySnapshotPatchAllOf1Policy) Validate(formats strfmt.Registry) error {
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
func (m *DirectorySnapshotPatchAllOf1Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectorySnapshotPatchAllOf1Policy) UnmarshalBinary(b []byte) error {
	var res DirectorySnapshotPatchAllOf1Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
