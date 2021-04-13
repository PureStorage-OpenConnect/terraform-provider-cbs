// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicySmbPostAllOf1 policy smb post all of1
//
// swagger:model policySmbPostAllOf1
type PolicySmbPostAllOf1 struct {

	// If set to `true`, enables access based enumeration on the policy. When access based enumeration is enabled on a policy, files and folders within exports that are attached to the policy will be hidden from users who do not have permission to view them. If not specified, defaults to `false`.
	AccessBasedEnumerationEnabled bool `json:"access_based_enumeration_enabled,omitempty"`
}

// Validate validates this policy smb post all of1
func (m *PolicySmbPostAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicySmbPostAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicySmbPostAllOf1) UnmarshalBinary(b []byte) error {
	var res PolicySmbPostAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}