// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MemberNoIDGroupOAIGenAllOf0 member no Id group Oai gen all of0
//
// swagger:model memberNoIdGroupOaiGenAllOf0
type MemberNoIDGroupOAIGenAllOf0 struct {

	// group
	Group *ReferenceNoID `json:"group,omitempty"`

	// member
	Member *Reference `json:"member,omitempty"`
}

// Validate validates this member no Id group Oai gen all of0
func (m *MemberNoIDGroupOAIGenAllOf0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMember(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MemberNoIDGroupOAIGenAllOf0) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *MemberNoIDGroupOAIGenAllOf0) validateMember(formats strfmt.Registry) error {

	if swag.IsZero(m.Member) { // not required
		return nil
	}

	if m.Member != nil {
		if err := m.Member.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemberNoIDGroupOAIGenAllOf0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberNoIDGroupOAIGenAllOf0) UnmarshalBinary(b []byte) error {
	var res MemberNoIDGroupOAIGenAllOf0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
