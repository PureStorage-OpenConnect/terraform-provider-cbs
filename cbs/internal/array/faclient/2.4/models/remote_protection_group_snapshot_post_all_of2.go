// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoteProtectionGroupSnapshotPostAllOf2 remote protection group snapshot post all of2
//
// swagger:model remoteProtectionGroupSnapshotPostAllOf2
type RemoteProtectionGroupSnapshotPostAllOf2 struct {

	// destroyed
	// Read Only: true
	Destroyed *bool `json:"destroyed,omitempty"`

	// Specifies a name suffix for the snapshots created. The snapshot is created on the FlashArray specified by the `on` parameter. The `on` parameter cannot refer to an offload target. Snapshots with suffixes specified have names in the form `PGROUP.SUFFIX` instead of the default `PGROUP.NNN` form. The names of all snapshots created by a single command that specifies this option have the same suffix.
	Suffix string `json:"suffix,omitempty"`
}

// Validate validates this remote protection group snapshot post all of2
func (m *RemoteProtectionGroupSnapshotPostAllOf2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProtectionGroupSnapshotPostAllOf2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProtectionGroupSnapshotPostAllOf2) UnmarshalBinary(b []byte) error {
	var res RemoteProtectionGroupSnapshotPostAllOf2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}