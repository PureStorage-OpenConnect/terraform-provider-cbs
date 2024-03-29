// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyRuleNfsClientPostRulesItems policy rule nfs client post rules items
//
// swagger:model policyRuleNfsClientPostRulesItems
type PolicyRuleNfsClientPostRulesItems struct {

	// Specifies access control for the export. Valid values are `root-squash` and `no-root-squash`. The default is `root-squash` if not specified.
	Access string `json:"access,omitempty"`

	// Specifies which clients are given access. Accepted notation includes IP, IP mask, or hostname. The default is `*` if not specified.
	Client string `json:"client,omitempty"`

	// Specifies which read-write client access permissions are allowed for the export. Valid values are `rw` and `ro`. The default is `rw` if not specified.
	Permission string `json:"permission,omitempty"`
}

// Validate validates this policy rule nfs client post rules items
func (m *PolicyRuleNfsClientPostRulesItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRuleNfsClientPostRulesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRuleNfsClientPostRulesItems) UnmarshalBinary(b []byte) error {
	var res PolicyRuleNfsClientPostRulesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
