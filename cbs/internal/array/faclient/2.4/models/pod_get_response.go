// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodGetResponse pod get response
//
// swagger:model podGetResponse
type PodGetResponse struct {
	PageInfo

	PodResponse

	PodGetResponseAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PodGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PageInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PageInfo = aO0

	// AO1
	var aO1 PodResponse
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PodResponse = aO1

	// AO2
	var aO2 PodGetResponseAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.PodGetResponseAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PodGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.PageInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PodResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.PodGetResponseAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pod get response
func (m *PodGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PageInfo
	if err := m.PageInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PodResponse
	if err := m.PodResponse.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PodGetResponseAllOf2
	if err := m.PodGetResponseAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PodGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodGetResponse) UnmarshalBinary(b []byte) error {
	var res PodGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
