// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodReplicaLinkPerformanceReplicationResponse pod replica link performance replication response
//
// swagger:model podReplicaLinkPerformanceReplicationResponse
type PodReplicaLinkPerformanceReplicationResponse struct {

	// A list of pod replica link performance objects.
	Items []*PodReplicaLinkPerformanceReplication `json:"items"`

	// The aggregate value of all items after filtering. For real-time performance, the values are aggregated for the latest timestamp. For historical performance, the values are aggregated for each timestamp from `start_time` to `end_time`. Where it makes more sense, the average value is displayed instead. The values are displayed for each field where meaningful.
	Total []*PodReplicaLinkPerformanceReplication `json:"total"`
}

// Validate validates this pod replica link performance replication response
func (m *PodReplicaLinkPerformanceReplicationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodReplicaLinkPerformanceReplicationResponse) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PodReplicaLinkPerformanceReplicationResponse) validateTotal(formats strfmt.Registry) error {

	if swag.IsZero(m.Total) { // not required
		return nil
	}

	for i := 0; i < len(m.Total); i++ {
		if swag.IsZero(m.Total[i]) { // not required
			continue
		}

		if m.Total[i] != nil {
			if err := m.Total[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("total" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PodReplicaLinkPerformanceReplicationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodReplicaLinkPerformanceReplicationResponse) UnmarshalBinary(b []byte) error {
	var res PodReplicaLinkPerformanceReplicationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
