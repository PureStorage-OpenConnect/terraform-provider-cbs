// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Chap Challenge-Handshake Authentication Protocol (CHAP).
//
// swagger:model chap
type Chap struct {

	// The host password for CHAP authentication. The password must be between 12 and 255 characters (inclusive) and cannot be the same as the target password.
	HostPassword string `json:"host_password,omitempty"`

	// The host username for CHAP authentication.
	HostUser string `json:"host_user,omitempty"`

	// The target password for CHAP authentication. The password must be between 12 and 255 characters (inclusive) and cannot be the same as the host password.
	TargetPassword string `json:"target_password,omitempty"`

	// The target username for CHAP authentication.
	TargetUser string `json:"target_user,omitempty"`
}

// Validate validates this chap
func (m *Chap) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Chap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Chap) UnmarshalBinary(b []byte) error {
	var res Chap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
