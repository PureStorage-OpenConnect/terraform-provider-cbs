// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtectionGroupSnapshotPostAllOf2 protection group snapshot post all of2
//
// swagger:model protectionGroupSnapshotPostAllOf2
type ProtectionGroupSnapshotPostAllOf2 struct {

	// destroyed
	// Read Only: true
	Destroyed *bool `json:"destroyed,omitempty"`
}

// Validate validates this protection group snapshot post all of2
func (m *ProtectionGroupSnapshotPostAllOf2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionGroupSnapshotPostAllOf2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionGroupSnapshotPostAllOf2) UnmarshalBinary(b []byte) error {
	var res ProtectionGroupSnapshotPostAllOf2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
