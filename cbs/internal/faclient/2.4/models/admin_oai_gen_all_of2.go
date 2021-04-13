// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AdminOAIGenAllOf2 admin Oai gen all of2
//
// swagger:model adminOaiGenAllOf2
type AdminOAIGenAllOf2 struct {

	// api token
	APIToken *APIToken `json:"api_token,omitempty"`

	// Returns a value of `true` if the user is local to the machine, otherwise `false`.
	// Read Only: true
	IsLocal *bool `json:"is_local,omitempty"`

	// Returns a value of `true` if the user is currently locked out, otherwise `false`. Can be patched to `false` to unlock a user. This field is only visible to `array_admin` roles. For all other users, the value is always `null`.
	Locked bool `json:"locked,omitempty"`

	// The remaining lockout period, in milliseconds, if the user is locked out. This field is only visible to `array_admin` roles. For all other users, the value is always `null`.
	// Read Only: true
	LockoutRemaining int64 `json:"lockout_remaining,omitempty"`

	// Password associated with the account.
	Password string `json:"password,omitempty"`

	// Public key for SSH access.
	PublicKey string `json:"public_key,omitempty"`

	// role
	Role *AdminRole `json:"role,omitempty"`
}

// Validate validates this admin Oai gen all of2
func (m *AdminOAIGenAllOf2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAPIToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdminOAIGenAllOf2) validateAPIToken(formats strfmt.Registry) error {

	if swag.IsZero(m.APIToken) { // not required
		return nil
	}

	if m.APIToken != nil {
		if err := m.APIToken.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("api_token")
			}
			return err
		}
	}

	return nil
}

func (m *AdminOAIGenAllOf2) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdminOAIGenAllOf2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdminOAIGenAllOf2) UnmarshalBinary(b []byte) error {
	var res AdminOAIGenAllOf2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}