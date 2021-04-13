// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertWatcherPatch alert watcher patch
//
// swagger:model alertWatcherPatch
type AlertWatcherPatch struct {

	// If set to `true`, email notifications will be sent to this watcher for alerts. If set to `false`, email notifications are disabled.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this alert watcher patch
func (m *AlertWatcherPatch) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertWatcherPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertWatcherPatch) UnmarshalBinary(b []byte) error {
	var res AlertWatcherPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
