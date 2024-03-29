// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OffloadOAIGenAllOf2 offload Oai gen all of2
//
// swagger:model offloadOaiGenAllOf2
type OffloadOAIGenAllOf2 struct {

	// Type of offload. Valid values include `azure`, `google-cloud`, `nfs`, and `s3`.
	// Read Only: true
	Protocol string `json:"protocol,omitempty"`

	// space
	Space *ArraySpaceOAIGenAllOf1SpaceAllOf0 `json:"space,omitempty"`

	// Offload status. Valid values are `connecting`, `connected`, `disconnecting`, `not connected`, and `scanning`.
	// Read Only: true
	Status string `json:"status,omitempty"`

	// Unique ID for the offload target. When multiple connections to one offload target are created, they each have distinct IDs but share the same `target_id`.
	// Read Only: true
	TargetID string `json:"target_id,omitempty"`
}

// Validate validates this offload Oai gen all of2
func (m *OffloadOAIGenAllOf2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OffloadOAIGenAllOf2) validateSpace(formats strfmt.Registry) error {

	if swag.IsZero(m.Space) { // not required
		return nil
	}

	if m.Space != nil {
		if err := m.Space.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("space")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OffloadOAIGenAllOf2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffloadOAIGenAllOf2) UnmarshalBinary(b []byte) error {
	var res OffloadOAIGenAllOf2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
