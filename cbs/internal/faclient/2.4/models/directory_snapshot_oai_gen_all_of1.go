// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirectorySnapshotOAIGenAllOf1 A point-in-time image of the contents of the directory.
//
// swagger:model directorySnapshotOaiGenAllOf1
type DirectorySnapshotOAIGenAllOf1 struct {

	// The customizable portion of the client visible snapshot name. A full snapshot name is constructed in the form of `DIR.CLIENT_NAME.SUFFIX` where `DIR` is the full managed directory name, `CLIENT_NAME` is the client name, and `SUFFIX` is the suffix. The client visible snapshot name is `CLIENT_NAME.SUFFIX`.
	ClientName string `json:"client_name,omitempty"`

	// The snapshot creation time, measured in milliseconds since the UNIX epoch.
	// Read Only: true
	Created int64 `json:"created,omitempty"`

	// Returns a value of `true` if the snapshot has been destroyed and is pending eradication. The `time_remaining` value displays the amount of time left until the destroyed directory snapshot is permanently eradicated. Before the `time_remaining` period has elapsed, the destroyed directory snapshot can be recovered by setting `destroyed=false`. Once the `time_remaining` period has elapsed, the directory snapshot is permanently eradicated and can no longer be recovered.
	Destroyed bool `json:"destroyed,omitempty"`

	// policy
	Policy *DirectorySnapshotOAIGenAllOf1Policy `json:"policy,omitempty"`

	// source
	Source *DirectorySnapshotOAIGenAllOf1Source `json:"source,omitempty"`

	// space
	Space *DirectorySnapshotOAIGenAllOf1Space `json:"space,omitempty"`

	// The suffix that is appended to the `source_name` value and the `client_name` value to generate the full directory snapshot name in the form of `DIR.CLIENT_NAME.SUFFIX` where `DIR` is the managed directory name, `CLIENT_NAME` is the client name, and `SUFFIX` is the suffix. If the suffix is a string, this field returns `null`. See the `name` value for the full snapshot name including the suffix.
	Suffix int64 `json:"suffix,omitempty"`

	// The amount of time left until the directory snapshot is permanently eradicated. Measured in milliseconds. Before the `time_remaining` period has elapsed, the snapshot can be recovered by setting `destroyed=false` if it is destroyed, by setting `policy=""` if it is managed by a snapshot policy, or by setting `keep_for=""` if it is a manual snapshot.
	// Read Only: true
	TimeRemaining int64 `json:"time_remaining,omitempty"`
}

// Validate validates this directory snapshot Oai gen all of1
func (m *DirectorySnapshotOAIGenAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectorySnapshotOAIGenAllOf1) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

func (m *DirectorySnapshotOAIGenAllOf1) validateSource(formats strfmt.Registry) error {

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

func (m *DirectorySnapshotOAIGenAllOf1) validateSpace(formats strfmt.Registry) error {

	if swag.IsZero(m.Space) { // not required
		return nil
	}

	if m.Space != nil {
		if err := m.Space.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectorySnapshotOAIGenAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectorySnapshotOAIGenAllOf1) UnmarshalBinary(b []byte) error {
	var res DirectorySnapshotOAIGenAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
