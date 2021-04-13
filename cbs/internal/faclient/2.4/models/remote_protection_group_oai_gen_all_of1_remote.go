// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RemoteProtectionGroupOAIGenAllOf1Remote The offload target that the remote protection group is on.
//
// swagger:model remoteProtectionGroupOaiGenAllOf1Remote
type RemoteProtectionGroupOAIGenAllOf1Remote struct {
	FixedReference
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *RemoteProtectionGroupOAIGenAllOf1Remote) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 FixedReference
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.FixedReference = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m RemoteProtectionGroupOAIGenAllOf1Remote) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.FixedReference)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this remote protection group Oai gen all of1 remote
func (m *RemoteProtectionGroupOAIGenAllOf1Remote) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with FixedReference
	if err := m.FixedReference.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *RemoteProtectionGroupOAIGenAllOf1Remote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemoteProtectionGroupOAIGenAllOf1Remote) UnmarshalBinary(b []byte) error {
	var res RemoteProtectionGroupOAIGenAllOf1Remote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
