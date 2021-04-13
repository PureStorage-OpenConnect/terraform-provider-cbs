// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MemberNoIDAllOAIGenAllOf0Group The resource in which the host, volume, or other item in the environment is a member.
//
// swagger:model memberNoIdAllOaiGenAllOf0Group
type MemberNoIDAllOAIGenAllOf0Group struct {
	ReferenceNoID
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *MemberNoIDAllOAIGenAllOf0Group) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ReferenceNoID
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ReferenceNoID = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m MemberNoIDAllOAIGenAllOf0Group) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ReferenceNoID)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this member no Id all Oai gen all of0 group
func (m *MemberNoIDAllOAIGenAllOf0Group) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ReferenceNoID
	if err := m.ReferenceNoID.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *MemberNoIDAllOAIGenAllOf0Group) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberNoIDAllOAIGenAllOf0Group) UnmarshalBinary(b []byte) error {
	var res MemberNoIDAllOAIGenAllOf0Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
