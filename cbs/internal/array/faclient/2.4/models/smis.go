// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SMIS smis
//
// swagger:model smis
type SMIS struct {

	// If set to `true`, the Service Location Protocol (SLP) and its ports, TCP 427 and UDP 427, are enabled.
	SlpEnabled bool `json:"slp_enabled,omitempty"`

	// If set to `true`, the SMI-S provider and its port, TCP 5989 is enabled.
	WbemHTTPSEnabled bool `json:"wbem_https_enabled,omitempty"`
}

// Validate validates this smis
func (m *SMIS) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SMIS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SMIS) UnmarshalBinary(b []byte) error {
	var res SMIS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
