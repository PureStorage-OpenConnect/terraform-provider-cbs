// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirectoryPolicyExportPostPoliciesItems directory policy export post policies items
//
// swagger:model directoryPolicyExportPostPoliciesItems
type DirectoryPolicyExportPostPoliciesItems struct {

	// The name of the export to create when applying the export policy to the directory.
	ExportName string `json:"export_name,omitempty"`

	// policy
	Policy *DirectoryPolicyExportPostPoliciesItemsPolicy `json:"policy,omitempty"`
}

// Validate validates this directory policy export post policies items
func (m *DirectoryPolicyExportPostPoliciesItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DirectoryPolicyExportPostPoliciesItems) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	if m.Policy != nil {
		if err := m.Policy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DirectoryPolicyExportPostPoliciesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectoryPolicyExportPostPoliciesItems) UnmarshalBinary(b []byte) error {
	var res DirectoryPolicyExportPostPoliciesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
