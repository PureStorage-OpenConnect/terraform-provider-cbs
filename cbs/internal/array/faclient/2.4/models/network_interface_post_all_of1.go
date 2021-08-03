// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkInterfacePostAllOf1 network interface post all of1
//
// swagger:model networkInterfacePostAllOf1
type NetworkInterfacePostAllOf1 struct {

	// eth
	Eth *NetworkInterfacePostAllOf1Eth `json:"eth,omitempty"`
}

// Validate validates this network interface post all of1
func (m *NetworkInterfacePostAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInterfacePostAllOf1) validateEth(formats strfmt.Registry) error {

	if swag.IsZero(m.Eth) { // not required
		return nil
	}

	if m.Eth != nil {
		if err := m.Eth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterfacePostAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfacePostAllOf1) UnmarshalBinary(b []byte) error {
	var res NetworkInterfacePostAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}