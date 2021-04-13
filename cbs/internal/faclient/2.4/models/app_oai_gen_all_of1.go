// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AppOAIGenAllOf1 An app that is installed on Purity to extend its functionality. The app is an add-on service that is integrated into the system to extend array functionality.
//
// swagger:model appOaiGenAllOf1
type AppOAIGenAllOf1 struct {

	// A description of the app.
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Details of the status of the app.
	// Read Only: true
	Details string `json:"details,omitempty"`

	// If set to `true`, the app is enabled. By default, apps are disabled.
	Enabled bool `json:"enabled,omitempty"`

	// The status of the app. Values include `healthy` and `unhealthy`. For cluster apps, this represents the aggregate status of the app. The aggregate status is only `healthy` if all nodes are `healthy`&#59; otherwise, it is `unhealthy`.
	// Read Only: true
	Status string `json:"status,omitempty"`

	// The app version. For cluster apps, this represents the node version if all nodes are of the same version. If the node versions differ, a value of `null` is returned.
	// Read Only: true
	Version string `json:"version,omitempty"`

	// If set to `true`, VNC access is enabled. By default, VNC access is disabled.
	VncEnabled bool `json:"vnc_enabled,omitempty"`
}

// Validate validates this app Oai gen all of1
func (m *AppOAIGenAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AppOAIGenAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppOAIGenAllOf1) UnmarshalBinary(b []byte) error {
	var res AppOAIGenAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}