// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicySmbPatchAllOf1 policy smb patch all of1
//
// swagger:model policySmbPatchAllOf1
type PolicySmbPatchAllOf1 struct {

	// If set to `true`, enables access based enumeration on the policy. When access based enumeration is enabled on a policy, files and folders within exports that are attached to the policy will be hidden from users who do not have permission to view them.
	AccessBasedEnumerationEnabled bool `json:"access_based_enumeration_enabled,omitempty"`
}

// Validate validates this policy smb patch all of1
func (m *PolicySmbPatchAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicySmbPatchAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicySmbPatchAllOf1) UnmarshalBinary(b []byte) error {
	var res PolicySmbPatchAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
