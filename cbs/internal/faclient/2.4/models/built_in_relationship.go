// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuiltInRelationship built in relationship
//
// swagger:model builtInRelationship
type BuiltInRelationship struct {

	// A non-modifiable, globally unique ID chosen by the system.
	// Read Only: true
	ID string `json:"id,omitempty"`
}

// Validate validates this built in relationship
func (m *BuiltInRelationship) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BuiltInRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuiltInRelationship) UnmarshalBinary(b []byte) error {
	var res BuiltInRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}