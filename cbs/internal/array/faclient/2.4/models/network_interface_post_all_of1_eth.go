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

// NetworkInterfacePostAllOf1Eth Ethernet network interface properties.
//
// swagger:model networkInterfacePostAllOf1Eth
type NetworkInterfacePostAllOf1Eth struct {

	// The IPv4 or IPv6 address to be associated with the specified network interface.
	Address string `json:"address,omitempty"`

	// List of network interfaces configured to be a subinterface of the specified network interface.
	Subinterfaces []*FixedReferenceNoID `json:"subinterfaces"`

	// subnet
	Subnet *NetworkInterfacePostAllOf1EthSubnet `json:"subnet,omitempty"`

	// The subtype of the specified network interface. Only interfaces of subtype `virtual` and `lacp_bond` can be created. Configurable on POST only. Valid values are `failover_bond`, `lacp_bond`, `physical`, and `virtual`.
	Subtype string `json:"subtype,omitempty"`
}

// Validate validates this network interface post all of1 eth
func (m *NetworkInterfacePostAllOf1Eth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubinterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkInterfacePostAllOf1Eth) validateSubinterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.Subinterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Subinterfaces); i++ {
		if swag.IsZero(m.Subinterfaces[i]) { // not required
			continue
		}

		if m.Subinterfaces[i] != nil {
			if err := m.Subinterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subinterfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInterfacePostAllOf1Eth) validateSubnet(formats strfmt.Registry) error {

	if swag.IsZero(m.Subnet) { // not required
		return nil
	}

	if m.Subnet != nil {
		if err := m.Subnet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subnet")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterfacePostAllOf1Eth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfacePostAllOf1Eth) UnmarshalBinary(b []byte) error {
	var res NetworkInterfacePostAllOf1Eth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
