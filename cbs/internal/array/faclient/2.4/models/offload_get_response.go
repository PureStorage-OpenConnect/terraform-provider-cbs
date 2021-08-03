// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OffloadGetResponse offload get response
//
// swagger:model offloadGetResponse
type OffloadGetResponse struct {
	PageInfo

	OffloadResponse

	OffloadGetResponseAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *OffloadGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PageInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PageInfo = aO0

	// AO1
	var aO1 OffloadResponse
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.OffloadResponse = aO1

	// AO2
	var aO2 OffloadGetResponseAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.OffloadGetResponseAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m OffloadGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.PageInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.OffloadResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.OffloadGetResponseAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this offload get response
func (m *OffloadGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PageInfo
	if err := m.PageInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OffloadResponse
	if err := m.OffloadResponse.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with OffloadGetResponseAllOf2
	if err := m.OffloadGetResponseAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OffloadGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffloadGetResponse) UnmarshalBinary(b []byte) error {
	var res OffloadGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}