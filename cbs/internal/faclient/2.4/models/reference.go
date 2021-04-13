// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Reference reference
//
// swagger:model reference
type Reference struct {

	// A globally unique, system-generated ID. The ID cannot be modified.
	ID string `json:"id,omitempty"`

	// The resource name, such as volume name, pod name, snapshot name, and so on.
	Name string `json:"name,omitempty"`
}

// Validate validates this reference
func (m *Reference) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Reference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Reference) UnmarshalBinary(b []byte) error {
	var res Reference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
