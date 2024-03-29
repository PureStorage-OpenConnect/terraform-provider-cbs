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

// NetworkInterfacePatchAllOf0Eth Ethernet network interface properties.
//
// swagger:model networkInterfacePatchAllOf0Eth
type NetworkInterfacePatchAllOf0Eth struct {

	// Slave devices to be added to the specified bond interface.
	AddSubinterfaces []*FixedReferenceNoID `json:"add_subinterfaces"`

	// The IPv4 or IPv6 address to be associated with the specified network interface.
	Address string `json:"address,omitempty"`

	// The IPv4 or IPv6 address of the gateway through which the specified network interface is to communicate with the network.
	Gateway string `json:"gateway,omitempty"`

	// Maximum message transfer unit (packet) size for the network interface in bytes. MTU setting cannot exceed the MTU of the corresponding physical interface.
	Mtu int32 `json:"mtu,omitempty"`

	// Netmask of the specified network interface that, when combined with the address of the interface, determines the network address of the interface.
	Netmask string `json:"netmask,omitempty"`

	// Slave devices to be removed from the specified bond interface.
	RemoveSubinterfaces []*FixedReferenceNoID `json:"remove_subinterfaces"`

	// Slave devices to be added to the specified bond interface.
	Subinterfaces []*FixedReferenceNoID `json:"subinterfaces"`

	// subnet
	Subnet *NetworkInterfacePatchAllOf0EthSubnet `json:"subnet,omitempty"`
}

// Validate validates this network interface patch all of0 eth
func (m *NetworkInterfacePatchAllOf0Eth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddSubinterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoveSubinterfaces(formats); err != nil {
		res = append(res, err)
	}

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

func (m *NetworkInterfacePatchAllOf0Eth) validateAddSubinterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.AddSubinterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.AddSubinterfaces); i++ {
		if swag.IsZero(m.AddSubinterfaces[i]) { // not required
			continue
		}

		if m.AddSubinterfaces[i] != nil {
			if err := m.AddSubinterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("add_subinterfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInterfacePatchAllOf0Eth) validateRemoveSubinterfaces(formats strfmt.Registry) error {

	if swag.IsZero(m.RemoveSubinterfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.RemoveSubinterfaces); i++ {
		if swag.IsZero(m.RemoveSubinterfaces[i]) { // not required
			continue
		}

		if m.RemoveSubinterfaces[i] != nil {
			if err := m.RemoveSubinterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("remove_subinterfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *NetworkInterfacePatchAllOf0Eth) validateSubinterfaces(formats strfmt.Registry) error {

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

func (m *NetworkInterfacePatchAllOf0Eth) validateSubnet(formats strfmt.Registry) error {

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
func (m *NetworkInterfacePatchAllOf0Eth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfacePatchAllOf0Eth) UnmarshalBinary(b []byte) error {
	var res NetworkInterfacePatchAllOf0Eth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
