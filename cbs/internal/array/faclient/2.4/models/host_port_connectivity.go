// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostPortConnectivity The connectivity status between the host and the ports on each controller.
//
// swagger:model hostPortConnectivity
type HostPortConnectivity struct {

	// The host connection status. Values include `Redundant`, `N/A`, `Redundant-Failover`, `Uneven`, `Unused Port`, `Single Controller`, `Single Controller-Failover`, and `None`.
	Details string `json:"details,omitempty"`

	// The host connection health status. Values include `healthy`, `unhealthy`, and `critical`.
	Status string `json:"status,omitempty"`
}

// Validate validates this host port connectivity
func (m *HostPortConnectivity) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HostPortConnectivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostPortConnectivity) UnmarshalBinary(b []byte) error {
	var res HostPortConnectivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
