// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TestResultWithResourceOAIGenAllOf1 test result with resource Oai gen all of1
//
// swagger:model testResultWithResourceOaiGenAllOf1
type TestResultWithResourceOAIGenAllOf1 struct {

	// resource
	Resource *TestResultWithResourceOAIGenAllOf1Resource `json:"resource,omitempty"`
}

// Validate validates this test result with resource Oai gen all of1
func (m *TestResultWithResourceOAIGenAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TestResultWithResourceOAIGenAllOf1) validateResource(formats strfmt.Registry) error {

	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TestResultWithResourceOAIGenAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TestResultWithResourceOAIGenAllOf1) UnmarshalBinary(b []byte) error {
	var res TestResultWithResourceOAIGenAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
