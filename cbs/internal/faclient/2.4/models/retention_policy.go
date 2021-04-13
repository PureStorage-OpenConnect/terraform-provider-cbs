// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RetentionPolicy The snapshot retention policy.
//
// swagger:model retentionPolicy
type RetentionPolicy struct {

	// The length of time to keep the specified snapshots. Measured in seconds.
	AllForSec int32 `json:"all_for_sec,omitempty"`

	// The number of days to keep the snapshots after the `all_for_sec` period has passed.
	Days int32 `json:"days,omitempty"`

	// The number of snapshots to keep per day after the `all_for_sec` period has passed.
	PerDay int32 `json:"per_day,omitempty"`
}

// Validate validates this retention policy
func (m *RetentionPolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RetentionPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionPolicy) UnmarshalBinary(b []byte) error {
	var res RetentionPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
