// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodReplicaLinkGetResponse pod replica link get response
//
// swagger:model podReplicaLinkGetResponse
type PodReplicaLinkGetResponse struct {
	PageInfo

	PodReplicaLinkResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *PodReplicaLinkGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PageInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PageInfo = aO0

	// AO1
	var aO1 PodReplicaLinkResponse
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.PodReplicaLinkResponse = aO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m PodReplicaLinkGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.PageInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.PodReplicaLinkResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this pod replica link get response
func (m *PodReplicaLinkGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PageInfo
	if err := m.PageInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with PodReplicaLinkResponse
	if err := m.PodReplicaLinkResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PodReplicaLinkGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodReplicaLinkGetResponse) UnmarshalBinary(b []byte) error {
	var res PodReplicaLinkGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
