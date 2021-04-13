// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Qos qos
//
// swagger:model qos
type Qos struct {

	// The maximum QoS bandwidth limit for the volume. Whenever throughput exceeds the bandwidth limit, throttling occurs. Measured in bytes per second. Maximum limit is 512 GB/s.
	// Maximum: 5.49755813888e+11
	// Minimum: 1.048576e+06
	BandwidthLimit int64 `json:"bandwidth_limit,omitempty"`

	// The QoS IOPs limit for the volume.
	// Maximum: 1.048576e+08
	// Minimum: 100
	IopsLimit int64 `json:"iops_limit,omitempty"`
}

// Validate validates this qos
func (m *Qos) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBandwidthLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIopsLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Qos) validateBandwidthLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.BandwidthLimit) { // not required
		return nil
	}

	if err := validate.MinimumInt("bandwidth_limit", "body", int64(m.BandwidthLimit), 1.048576e+06, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("bandwidth_limit", "body", int64(m.BandwidthLimit), 5.49755813888e+11, false); err != nil {
		return err
	}

	return nil
}

func (m *Qos) validateIopsLimit(formats strfmt.Registry) error {

	if swag.IsZero(m.IopsLimit) { // not required
		return nil
	}

	if err := validate.MinimumInt("iops_limit", "body", int64(m.IopsLimit), 100, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("iops_limit", "body", int64(m.IopsLimit), 1.048576e+08, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Qos) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Qos) UnmarshalBinary(b []byte) error {
	var res Qos
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
