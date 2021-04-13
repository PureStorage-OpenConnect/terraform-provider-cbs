// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtectionGroupSnapshotOAIGenAllOf1Space Returns provisioned (virtual) size and physical storage consumption data for each protection group.
//
// swagger:model protectionGroupSnapshotOaiGenAllOf1Space
type ProtectionGroupSnapshotOAIGenAllOf1Space struct {
	ArraySpaceOAIGenAllOf1SpaceAllOf0
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectionGroupSnapshotOAIGenAllOf1Space) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ArraySpaceOAIGenAllOf1SpaceAllOf0
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ArraySpaceOAIGenAllOf1SpaceAllOf0 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectionGroupSnapshotOAIGenAllOf1Space) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ArraySpaceOAIGenAllOf1SpaceAllOf0)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protection group snapshot Oai gen all of1 space
func (m *ProtectionGroupSnapshotOAIGenAllOf1Space) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ArraySpaceOAIGenAllOf1SpaceAllOf0
	if err := m.ArraySpaceOAIGenAllOf1SpaceAllOf0.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionGroupSnapshotOAIGenAllOf1Space) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionGroupSnapshotOAIGenAllOf1Space) UnmarshalBinary(b []byte) error {
	var res ProtectionGroupSnapshotOAIGenAllOf1Space
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
