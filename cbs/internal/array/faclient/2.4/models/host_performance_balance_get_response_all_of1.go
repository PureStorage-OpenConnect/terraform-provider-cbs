// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// HostPerformanceBalanceGetResponseAllOf1 host performance balance get response all of1
//
// swagger:model hostPerformanceBalanceGetResponseAllOf1
type HostPerformanceBalanceGetResponseAllOf1 struct {

	// A list of entries indicating count and percentage of I/O operations across various data paths between the host and the array.
	Items []*HostPerformanceBalance `json:"items"`
}

// Validate validates this host performance balance get response all of1
func (m *HostPerformanceBalanceGetResponseAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostPerformanceBalanceGetResponseAllOf1) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostPerformanceBalanceGetResponseAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostPerformanceBalanceGetResponseAllOf1) UnmarshalBinary(b []byte) error {
	var res HostPerformanceBalanceGetResponseAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
