// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyRuleSmbClientPostRulesItems policy rule smb client post rules items
//
// swagger:model policyRuleSmbClientPostRulesItems
type PolicyRuleSmbClientPostRulesItems struct {

	// Specifies whether access to information is allowed for anonymous users. If not specified, defaults to `false`.
	AnonymousAccessAllowed bool `json:"anonymous_access_allowed,omitempty"`

	// Specifies which clients are given access. Accepted notation, IP, IP mask, or hostname. The default is `*` if not specified.
	Client string `json:"client,omitempty"`

	// Specifies whether the remote client is required to use SMB encryption. If not specified, defaults to `false`.
	SmbEncryptionRequired bool `json:"smb_encryption_required,omitempty"`
}

// Validate validates this policy rule smb client post rules items
func (m *PolicyRuleSmbClientPostRulesItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRuleSmbClientPostRulesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRuleSmbClientPostRulesItems) UnmarshalBinary(b []byte) error {
	var res PolicyRuleSmbClientPostRulesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
