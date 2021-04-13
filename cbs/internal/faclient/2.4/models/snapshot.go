// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Snapshot snapshot
//
// swagger:model snapshot
type Snapshot struct {

	// The snapshot creation time. Measured in milliseconds since the UNIX epoch.
	// Read Only: true
	Created int64 `json:"created,omitempty"`

	// Returns a value of `true` if the snapshot has been destroyed and is pending eradication. The `time_remaining` value displays the amount of time left until the destroyed volume snapshot is permanently eradicated. Before the `time_remaining` period has elapsed, the destroyed volume snapshot can be recovered by setting `destroyed=false`. Once the `time_remaining` period has elapsed, the volume snapshot is permanently eradicated and can no longer be recovered.
	Destroyed bool `json:"destroyed,omitempty"`

	// pod
	Pod *SnapshotPod `json:"pod,omitempty"`

	// The provisioned space of the snapshot. Measured in bytes.
	// Read Only: true
	// Maximum: 4.503599627370496e+15
	// Minimum: 1.048576e+06
	Provisioned int64 `json:"provisioned,omitempty"`

	// source
	Source *SnapshotSource `json:"source,omitempty"`

	// The suffix that is appended to the `source_name` value to generate the full volume snapshot name in the form `VOL.SUFFIX`. If the suffix is not specified, the system constructs the snapshot name in the form `VOL.NNN`, where `VOL` is the volume name, and `NNN` is a monotonically increasing number.
	Suffix string `json:"suffix,omitempty"`

	// The amount of time left until the destroyed snapshot is permanently eradicated. Measured in milliseconds. Before the `time_remaining` period has elapsed, the destroyed snapshot can be recovered by setting `destroyed=false`.
	// Read Only: true
	TimeRemaining int64 `json:"time_remaining,omitempty"`
}

// Validate validates this snapshot
func (m *Snapshot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProvisioned(formats); err != nil {
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

func (m *Snapshot) validatePod(formats strfmt.Registry) error {

	if swag.IsZero(m.Pod) { // not required
		return nil
	}

	if m.Pod != nil {
		if err := m.Pod.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pod")
			}
			return err
		}
	}

	return nil
}

func (m *Snapshot) validateProvisioned(formats strfmt.Registry) error {

	if swag.IsZero(m.Provisioned) { // not required
		return nil
	}

	if err := validate.MinimumInt("provisioned", "body", int64(m.Provisioned), 1.048576e+06, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("provisioned", "body", int64(m.Provisioned), 4.503599627370496e+15, false); err != nil {
		return err
	}

	return nil
}

func (m *Snapshot) validateSource(formats strfmt.Registry) error {

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
func (m *Snapshot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Snapshot) UnmarshalBinary(b []byte) error {
	var res Snapshot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
