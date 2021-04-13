// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SNMPAgentPatchAllOf1 snmp agent patch all of1
//
// swagger:model snmpAgentPatchAllOf1
type SNMPAgentPatchAllOf1 struct {

	// The administration domain unique name of an SNMP agent.
	// Read Only: true
	// Max Length: 32
	EngineID string `json:"engine_id,omitempty"`

	// v2c
	V2c *SNMPV2c `json:"v2c,omitempty"`

	// v3
	V3 *SNMPV3Patch `json:"v3,omitempty"`

	// Version of the SNMP protocol to be used by Purity//FA in communications with the specified manager. Valid values are `v2c` and `v3`.
	Version string `json:"version,omitempty"`
}

// Validate validates this snmp agent patch all of1
func (m *SNMPAgentPatchAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEngineID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateV2c(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateV3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SNMPAgentPatchAllOf1) validateEngineID(formats strfmt.Registry) error {

	if swag.IsZero(m.EngineID) { // not required
		return nil
	}

	if err := validate.MaxLength("engine_id", "body", string(m.EngineID), 32); err != nil {
		return err
	}

	return nil
}

func (m *SNMPAgentPatchAllOf1) validateV2c(formats strfmt.Registry) error {

	if swag.IsZero(m.V2c) { // not required
		return nil
	}

	if m.V2c != nil {
		if err := m.V2c.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v2c")
			}
			return err
		}
	}

	return nil
}

func (m *SNMPAgentPatchAllOf1) validateV3(formats strfmt.Registry) error {

	if swag.IsZero(m.V3) { // not required
		return nil
	}

	if m.V3 != nil {
		if err := m.V3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("v3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SNMPAgentPatchAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SNMPAgentPatchAllOf1) UnmarshalBinary(b []byte) error {
	var res SNMPAgentPatchAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
