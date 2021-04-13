// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SNMPAgentMibGetResponse snmp agent mib get response
//
// swagger:model snmpAgentMibGetResponse
type SNMPAgentMibGetResponse struct {
	PageInfo

	SNMPAgentMibResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *SNMPAgentMibGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PageInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PageInfo = aO0

	// AO1
	var aO1 SNMPAgentMibResponse
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.SNMPAgentMibResponse = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m SNMPAgentMibGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PageInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.SNMPAgentMibResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this snmp agent mib get response
func (m *SNMPAgentMibGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PageInfo
	if err := m.PageInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SNMPAgentMibResponse
	if err := m.SNMPAgentMibResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SNMPAgentMibGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SNMPAgentMibGetResponse) UnmarshalBinary(b []byte) error {
	var res SNMPAgentMibGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
