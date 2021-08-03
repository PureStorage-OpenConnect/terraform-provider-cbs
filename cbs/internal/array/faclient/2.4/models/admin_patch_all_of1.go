// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminPatchAllOf1 admin patch all of1
//
// swagger:model adminPatchAllOf1
type AdminPatchAllOf1 struct {

	// The current password.
	OldPassword string `json:"old_password,omitempty"`
}

// Validate validates this admin patch all of1
func (m *AdminPatchAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AdminPatchAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminPatchAllOf1) UnmarshalBinary(b []byte) error {
	var res AdminPatchAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}