// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PostOauth210TokenOKBody oauth_token_response
//
// swagger:model postOauth210TokenOKBody
type PostOauth210TokenOKBody struct {

	// The serialized OAuth 2.0 Bearer token used to perform authenticated requests. The access token must be added to the Authorization header of all API calls.
	AccessToken string `json:"access_token,omitempty"`

	// The duration after which the access token will expire. Measured in seconds. This differs from other duration fields that are expressed in milliseconds.
	ExpiresIn uint32 `json:"expires_in,omitempty"`

	// The type of token that is issued. The Pure Storage REST API supports OAuth 2.0 access tokens.
	IssuedTokenType string `json:"issued_token_type,omitempty"`

	// Indicates how the API client can use the access token issued. The Pure Storage REST API supports the `Bearer` token.
	TokenType string `json:"token_type,omitempty"`
}

// Validate validates this post oauth210 token o k body
func (m *PostOauth210TokenOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostOauth210TokenOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostOauth210TokenOKBody) UnmarshalBinary(b []byte) error {
	var res PostOauth210TokenOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
