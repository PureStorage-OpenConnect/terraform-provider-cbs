// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyRuleSmbClientGetResponse policy rule smb client get response
//
// swagger:model policyRuleSmbClientGetResponse
type PolicyRuleSmbClientGetResponse struct {
	PageInfo

	PolicyRuleSmbClientResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PolicyRuleSmbClientGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PageInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PageInfo = aO0

	// AO1
	var aO1 PolicyRuleSmbClientResponse
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PolicyRuleSmbClientResponse = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PolicyRuleSmbClientGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PageInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PolicyRuleSmbClientResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this policy rule smb client get response
func (m *PolicyRuleSmbClientGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PageInfo
	if err := m.PageInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PolicyRuleSmbClientResponse
	if err := m.PolicyRuleSmbClientResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRuleSmbClientGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRuleSmbClientGetResponse) UnmarshalBinary(b []byte) error {
	var res PolicyRuleSmbClientGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}