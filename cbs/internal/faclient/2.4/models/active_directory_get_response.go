// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActiveDirectoryGetResponse active directory get response
//
// swagger:model activeDirectoryGetResponse
type ActiveDirectoryGetResponse struct {
	PageInfo

	ActiveDirectoryResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ActiveDirectoryGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PageInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PageInfo = aO0

	// AO1
	var aO1 ActiveDirectoryResponse
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ActiveDirectoryResponse = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ActiveDirectoryGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PageInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ActiveDirectoryResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this active directory get response
func (m *ActiveDirectoryGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PageInfo
	if err := m.PageInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ActiveDirectoryResponse
	if err := m.ActiveDirectoryResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ActiveDirectoryGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveDirectoryGetResponse) UnmarshalBinary(b []byte) error {
	var res ActiveDirectoryGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
