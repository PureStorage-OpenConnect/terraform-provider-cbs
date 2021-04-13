// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OffloadPostS3 S3 settings.
//
// swagger:model offloadPostS3
type OffloadPostS3 struct {
	OffloadS3
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OffloadPostS3) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 OffloadS3
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.OffloadS3 = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OffloadPostS3) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.OffloadS3)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this offload post s3
func (m *OffloadPostS3) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with OffloadS3
	if err := m.OffloadS3.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OffloadPostS3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffloadPostS3) UnmarshalBinary(b []byte) error {
	var res OffloadPostS3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}