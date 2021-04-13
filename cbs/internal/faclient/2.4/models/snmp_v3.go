// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SNMPV3 The v3 configurations of SNMP.
//
// swagger:model snmpV3
type SNMPV3 struct {

	// Passphrase used by Purity//FA to authenticate the array with the specified managers.
	AuthPassphrase string `json:"auth_passphrase,omitempty"`

	// Hash algorithm used to validate the authentication passphrase. Valid values are `MD5` and `SHA`.
	AuthProtocol string `json:"auth_protocol,omitempty"`

	// Passphrase used to encrypt SNMP messages.
	PrivacyPassphrase string `json:"privacy_passphrase,omitempty"`

	// Encryption protocol for SNMP messages. Valid values are `AES` and `DES`.
	PrivacyProtocol string `json:"privacy_protocol,omitempty"`

	// User ID recognized by the specified SNMP managers which Purity//FA is to use in communications with them.
	User string `json:"user,omitempty"`
}

// Validate validates this snmp v3
func (m *SNMPV3) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SNMPV3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SNMPV3) UnmarshalBinary(b []byte) error {
	var res SNMPV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
