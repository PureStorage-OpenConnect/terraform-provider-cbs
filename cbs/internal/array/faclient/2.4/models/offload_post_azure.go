// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OffloadPostAzure Microsoft Azure Blob storage settings.
//
// swagger:model offloadPostAzure
type OffloadPostAzure struct {
	OffloadAzure
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OffloadPostAzure) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 OffloadAzure
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.OffloadAzure = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OffloadPostAzure) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.OffloadAzure)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this offload post azure
func (m *OffloadPostAzure) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OffloadAzure
	if err := m.OffloadAzure.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OffloadPostAzure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffloadPostAzure) UnmarshalBinary(b []byte) error {
	var res OffloadPostAzure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}