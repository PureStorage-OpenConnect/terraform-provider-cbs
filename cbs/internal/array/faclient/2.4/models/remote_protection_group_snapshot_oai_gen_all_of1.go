// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoteProtectionGroupSnapshotOAIGenAllOf1 A remote protection group snapshot is a collection of point-in-time consistent volume snapshots, for volumes associated with the remote protection group when the snapshot was taken.
//
// swagger:model remoteProtectionGroupSnapshotOaiGenAllOf1
type RemoteProtectionGroupSnapshotOAIGenAllOf1 struct {

	// Creation time of the snapshot on the original source of the snapshot. Measured in milliseconds since the UNIX epoch.
	// Read Only: true
	Created int64 `json:"created,omitempty"`

	// Destroyed and pending eradication? If not specified, defaults to false.
	Destroyed bool `json:"destroyed,omitempty"`

	// Whether or not this remote protection group snapshot is replicated from the current array.
	// Read Only: true
	IsLocal *bool `json:"is_local,omitempty"`

	// remote
	Remote *RemoteProtectionGroupSnapshotOAIGenAllOf1Remote `json:"remote,omitempty"`

	// source
	Source *RemoteProtectionGroupSnapshotOAIGenAllOf1Source `json:"source,omitempty"`

	// The suffix that is appended to the `source_name` value to generate the full remote protection group snapshot name in the form `PGROUP.SUFFIX`. If the suffix is not specified, the system constructs the snapshot name in the form `PGROUP.NNN`, where `PGROUP` is the protection group name, and `NNN` is a monotonically increasing number.
	Suffix string `json:"suffix,omitempty"`

	// Milliseconds remaining until eradication, if the snapshot has been destroyed.
	// Read Only: true
	TimeRemaining int64 `json:"time_remaining,omitempty"`
}

// Validate validates this remote protection group snapshot Oai gen all of1
func (m *RemoteProtectionGroupSnapshotOAIGenAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemoteProtectionGroupSnapshotOAIGenAllOf1) validateRemote(formats strfmt.Registry) error {

	if swag.IsZero(m.Remote) { // not required
		return nil
	}

	if m.Remote != nil {
		if err := m.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("remote")
			}
			return err
		}
	}

	return nil
}

func (m *RemoteProtectionGroupSnapshotOAIGenAllOf1) validateSource(formats strfmt.Registry) error {

	if swag.IsZero(m.Source) { // not required
		return nil
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProtectionGroupSnapshotOAIGenAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProtectionGroupSnapshotOAIGenAllOf1) UnmarshalBinary(b []byte) error {
	var res RemoteProtectionGroupSnapshotOAIGenAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}