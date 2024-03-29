// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PortCommon port common
//
// swagger:model portCommon
type PortCommon struct {

	// The iSCSI Qualified Name (or `null` if target is not iSCSI).
	Iqn string `json:"iqn,omitempty"`

	// NVMe Qualified Name (or `null` if target is not NVMeoF).
	Nqn string `json:"nqn,omitempty"`

	// IP and port number (or `null` if target is not iSCSI).
	Portal string `json:"portal,omitempty"`

	// Fibre Channel World Wide Name (or `null` if target is not Fibre Channel).
	Wwn string `json:"wwn,omitempty"`
}

// Validate validates this port common
func (m *PortCommon) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PortCommon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PortCommon) UnmarshalBinary(b []byte) error {
	var res PortCommon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
