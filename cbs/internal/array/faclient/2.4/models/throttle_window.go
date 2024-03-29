// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ThrottleWindow The time during which the `window_limit` threshold is in effect.
//
// swagger:model throttleWindow
type ThrottleWindow struct {
	TimeWindow
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ThrottleWindow) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 TimeWindow
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.TimeWindow = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ThrottleWindow) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.TimeWindow)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this throttle window
func (m *ThrottleWindow) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with TimeWindow
	if err := m.TimeWindow.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *ThrottleWindow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ThrottleWindow) UnmarshalBinary(b []byte) error {
	var res ThrottleWindow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
