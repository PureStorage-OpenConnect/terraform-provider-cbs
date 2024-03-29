// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodPerformanceReplicationOAIGenContinuousBytesPerSec Total bytes transmitted or received per second for continuous replication. The continuous replication feature is used for disaster recovery on FlashArray and provides a recovery point objective (RPO) of significantly less than 30s.
//
// swagger:model podPerformanceReplicationOaiGenContinuousBytesPerSec
type PodPerformanceReplicationOAIGenContinuousBytesPerSec struct {
	ReplicationPerformanceWithTotal
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PodPerformanceReplicationOAIGenContinuousBytesPerSec) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ReplicationPerformanceWithTotal
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ReplicationPerformanceWithTotal = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PodPerformanceReplicationOAIGenContinuousBytesPerSec) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ReplicationPerformanceWithTotal)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pod performance replication Oai gen continuous bytes per sec
func (m *PodPerformanceReplicationOAIGenContinuousBytesPerSec) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ReplicationPerformanceWithTotal
	if err := m.ReplicationPerformanceWithTotal.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PodPerformanceReplicationOAIGenContinuousBytesPerSec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodPerformanceReplicationOAIGenContinuousBytesPerSec) UnmarshalBinary(b []byte) error {
	var res PodPerformanceReplicationOAIGenContinuousBytesPerSec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
