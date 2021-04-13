// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResourcePerformanceNoID resource performance no Id
//
// swagger:model resourcePerformanceNoId
type ResourcePerformanceNoID struct {
	Performance

	ResourceNoID
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ResourcePerformanceNoID) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 Performance
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.Performance = aO0

	// AO1
	var aO1 ResourceNoID
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.ResourceNoID = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ResourcePerformanceNoID) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.Performance)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.ResourceNoID)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this resource performance no Id
func (m *ResourcePerformanceNoID) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with Performance
	if err := m.Performance.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ResourceNoID
	if err := m.ResourceNoID.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ResourcePerformanceNoID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcePerformanceNoID) UnmarshalBinary(b []byte) error {
	var res ResourcePerformanceNoID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
