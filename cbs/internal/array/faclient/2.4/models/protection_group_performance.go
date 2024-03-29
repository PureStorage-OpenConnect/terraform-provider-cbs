// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProtectionGroupPerformance protection group performance
//
// swagger:model ProtectionGroupPerformance
type ProtectionGroupPerformance struct {
	BuiltInResourceNoID

	AggregateReplicationPerformance

	ProtectionGroupPerformanceOAIGenAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ProtectionGroupPerformance) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 BuiltInResourceNoID
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.BuiltInResourceNoID = aO0

	// AO1
	var aO1 AggregateReplicationPerformance
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.AggregateReplicationPerformance = aO1

	// AO2
	var aO2 ProtectionGroupPerformanceOAIGenAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.ProtectionGroupPerformanceOAIGenAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ProtectionGroupPerformance) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.BuiltInResourceNoID)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.AggregateReplicationPerformance)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.ProtectionGroupPerformanceOAIGenAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this protection group performance
func (m *ProtectionGroupPerformance) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with BuiltInResourceNoID
	if err := m.BuiltInResourceNoID.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with AggregateReplicationPerformance
	if err := m.AggregateReplicationPerformance.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ProtectionGroupPerformanceOAIGenAllOf2
	if err := m.ProtectionGroupPerformanceOAIGenAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ProtectionGroupPerformance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProtectionGroupPerformance) UnmarshalBinary(b []byte) error {
	var res ProtectionGroupPerformance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
