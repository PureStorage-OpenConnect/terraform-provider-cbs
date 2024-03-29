// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SNMPManagerGetResponse snmp manager get response
//
// swagger:model snmpManagerGetResponse
type SNMPManagerGetResponse struct {
	PageInfo

	SNMPManagerResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SNMPManagerGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PageInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PageInfo = aO0

	// AO1
	var aO1 SNMPManagerResponse
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SNMPManagerResponse = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SNMPManagerGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PageInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SNMPManagerResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this snmp manager get response
func (m *SNMPManagerGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PageInfo
	if err := m.PageInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SNMPManagerResponse
	if err := m.SNMPManagerResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SNMPManagerGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SNMPManagerGetResponse) UnmarshalBinary(b []byte) error {
	var res SNMPManagerGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
