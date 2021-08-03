// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyMemberExportPostMembersItems policy member export post members items
//
// swagger:model policyMemberExportPostMembersItems
type PolicyMemberExportPostMembersItems struct {

	// The name of the export to create when applying the export policy to the directory.
	ExportName string `json:"export_name,omitempty"`

	// member
	Member *PolicyMemberExportPostMembersItemsMember `json:"member,omitempty"`
}

// Validate validates this policy member export post members items
func (m *PolicyMemberExportPostMembersItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMember(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyMemberExportPostMembersItems) validateMember(formats strfmt.Registry) error {

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
func (m *PolicyMemberExportPostMembersItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyMemberExportPostMembersItems) UnmarshalBinary(b []byte) error {
	var res PolicyMemberExportPostMembersItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}