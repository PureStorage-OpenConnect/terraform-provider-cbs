// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourcePerformanceNoIDByArrayAllOf1 resource performance no Id by array all of1
//
// swagger:model resourcePerformanceNoIdByArrayAllOf1
type ResourcePerformanceNoIDByArrayAllOf1 struct {

	// array
	Array *ResourcePerformanceNoIDByArrayAllOf1Array `json:"array,omitempty"`
}

// Validate validates this resource performance no Id by array all of1
func (m *ResourcePerformanceNoIDByArrayAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArray(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcePerformanceNoIDByArrayAllOf1) validateArray(formats strfmt.Registry) error {

	if swag.IsZero(m.Array) { // not required
		return nil
	}

	if m.Array != nil {
		if err := m.Array.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("array")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcePerformanceNoIDByArrayAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcePerformanceNoIDByArrayAllOf1) UnmarshalBinary(b []byte) error {
	var res ResourcePerformanceNoIDByArrayAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}