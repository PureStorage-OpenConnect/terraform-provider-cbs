// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SNMPV3Patch The v3 configurations of SNMP.
//
// swagger:model snmpV3Patch
type SNMPV3Patch struct {

	// Passphrase used by Purity//FA to authenticate the array with the specified managers.
	// Max Length: 32
	AuthPassphrase string `json:"auth_passphrase,omitempty"`

	// Hash algorithm used to validate the authentication passphrase. Valid values are `MD5` and `SHA`.
	AuthProtocol string `json:"auth_protocol,omitempty"`

	// Passphrase used to encrypt SNMP messages.
	// Max Length: 63
	// Min Length: 8
	PrivacyPassphrase string `json:"privacy_passphrase,omitempty"`

	// Encryption protocol for SNMP messages. Valid values are `AES` and `DES`.
	PrivacyProtocol string `json:"privacy_protocol,omitempty"`

	// User ID recognized by the specified SNMP managers which Purity//FA is to use in communications with them.
	User string `json:"user,omitempty"`
}

// Validate validates this snmp v3 patch
func (m *SNMPV3Patch) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthPassphrase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivacyPassphrase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SNMPV3Patch) validateAuthPassphrase(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthPassphrase) { // not required
		return nil
	}

	if err := validate.MaxLength("auth_passphrase", "body", string(m.AuthPassphrase), 32); err != nil {
		return err
	}

	return nil
}

func (m *SNMPV3Patch) validatePrivacyPassphrase(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivacyPassphrase) { // not required
		return nil
	}

	if err := validate.MinLength("privacy_passphrase", "body", string(m.PrivacyPassphrase), 8); err != nil {
		return err
	}

	if err := validate.MaxLength("privacy_passphrase", "body", string(m.PrivacyPassphrase), 63); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SNMPV3Patch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SNMPV3Patch) UnmarshalBinary(b []byte) error {
	var res SNMPV3Patch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
