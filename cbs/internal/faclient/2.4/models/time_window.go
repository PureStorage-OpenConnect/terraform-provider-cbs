// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TimeWindow time window
//
// swagger:model timeWindow
type TimeWindow struct {

	// The window end time. Measured in milliseconds since midnight. The time must be set on the hour. (e.g., `28800000`, which is equal to 8:00 AM).
	End int64 `json:"end,omitempty"`

	// The window start time. Measured in milliseconds since midnight. The time must be set on the hour. (e.g., `18000000`, which is equal to 5:00 AM).
	Start int64 `json:"start,omitempty"`
}

// Validate validates this time window
func (m *TimeWindow) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TimeWindow) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeWindow) UnmarshalBinary(b []byte) error {
	var res TimeWindow
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
