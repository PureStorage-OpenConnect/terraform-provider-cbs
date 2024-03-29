// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuiltIn A built-in resource. Many are singletons predefined by Purity (e.g., support settings). Some correspond to a piece of software, like an app, or hardware, like a controller. Others are created by the system in response to some event (e.g., alerts, audit records).
// Typically, a user can't create, delete or rename a built-in resource. A few can be created or deleted, but not renamed because the names are meaningful to Purity (e.g., VIFs).
//
// swagger:model builtIn
type BuiltIn struct {

	// A globally unique, system-generated ID. The ID cannot be modified and cannot refer to another resource.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// A locally unique, system-generated name. The name cannot be modified.
	// Read Only: true
	Name string `json:"name,omitempty"`
}

// Validate validates this built in
func (m *BuiltIn) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuiltIn) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuiltIn) UnmarshalBinary(b []byte) error {
	var res BuiltIn
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
