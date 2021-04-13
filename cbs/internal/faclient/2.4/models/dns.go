// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DNS Configuration information for the domain name servers.
//
// swagger:model DNS
type DNS struct {

	// Domain suffix to be appended by the appliance when performing DNS lookups.
	Domain string `json:"domain,omitempty"`

	// List of DNS server IP addresses.
	// Max Length: 3
	Nameservers []string `json:"nameservers"`
}

// Validate validates this DNS
func (m *DNS) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNameservers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DNS) validateNameservers(formats strfmt.Registry) error {

	if swag.IsZero(m.Nameservers) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DNS) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNS) UnmarshalBinary(b []byte) error {
	var res DNS
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
