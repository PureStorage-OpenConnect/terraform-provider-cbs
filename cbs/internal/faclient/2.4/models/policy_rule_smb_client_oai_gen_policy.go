// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyRuleSmbClientOAIGenPolicy The policy to which this rule belongs.
//
// swagger:model policyRuleSmbClientOaiGenPolicy
type PolicyRuleSmbClientOAIGenPolicy struct {
	FixedReferenceWithType
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PolicyRuleSmbClientOAIGenPolicy) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 FixedReferenceWithType
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.FixedReferenceWithType = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PolicyRuleSmbClientOAIGenPolicy) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.FixedReferenceWithType)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this policy rule smb client Oai gen policy
func (m *PolicyRuleSmbClientOAIGenPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FixedReferenceWithType
	if err := m.FixedReferenceWithType.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyRuleSmbClientOAIGenPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyRuleSmbClientOAIGenPolicy) UnmarshalBinary(b []byte) error {
	var res PolicyRuleSmbClientOAIGenPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
