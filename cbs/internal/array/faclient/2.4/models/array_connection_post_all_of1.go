// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ArrayConnectionPostAllOf1 array connection post all of1
//
// swagger:model arrayConnectionPostAllOf1
type ArrayConnectionPostAllOf1 struct {

	// The connection key of the target array.
	ConnectionKey string `json:"connection_key,omitempty"`

	// Management IP address of the target array.
	ManagementAddress string `json:"management_address,omitempty"`

	// IP addresses and FQDNs of the target arrays. Configurable only when `replication_transport` is set to `ip`. If not configured, will be set to all the replication addresses available on the target array at the time of the POST.
	ReplicationAddresses []string `json:"replication_addresses"`

	// The protocol used to transport data betwen the local array and the remote array. Valid values are `ip` and `fc`. The default is `ip`.
	ReplicationTransport string `json:"replication_transport,omitempty"`

	// The type of replication. Valid values are `async-replication` and `sync-replication`.
	Type string `json:"type,omitempty"`
}

// Validate validates this array connection post all of1
func (m *ArrayConnectionPostAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ArrayConnectionPostAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArrayConnectionPostAllOf1) UnmarshalBinary(b []byte) error {
	var res ArrayConnectionPostAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
