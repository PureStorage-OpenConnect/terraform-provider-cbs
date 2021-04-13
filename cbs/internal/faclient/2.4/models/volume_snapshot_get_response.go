// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// VolumeSnapshotGetResponse volume snapshot get response
//
// swagger:model volumeSnapshotGetResponse
type VolumeSnapshotGetResponse struct {
	PageInfo

	VolumeSnapshotResponse

	VolumeSnapshotGetResponseAllOf2
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *VolumeSnapshotGetResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 PageInfo
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.PageInfo = aO0

	// AO1
	var aO1 VolumeSnapshotResponse
	if err := swag.ReadJSON(raw, &aO1); err != nil {
		return err
	}
	m.VolumeSnapshotResponse = aO1

	// AO2
	var aO2 VolumeSnapshotGetResponseAllOf2
	if err := swag.ReadJSON(raw, &aO2); err != nil {
		return err
	}
	m.VolumeSnapshotGetResponseAllOf2 = aO2

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m VolumeSnapshotGetResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 3)

	aO0, err := swag.WriteJSON(m.PageInfo)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	aO1, err := swag.WriteJSON(m.VolumeSnapshotResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO1)

	aO2, err := swag.WriteJSON(m.VolumeSnapshotGetResponseAllOf2)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO2)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this volume snapshot get response
func (m *VolumeSnapshotGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with PageInfo
	if err := m.PageInfo.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VolumeSnapshotResponse
	if err := m.VolumeSnapshotResponse.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with VolumeSnapshotGetResponseAllOf2
	if err := m.VolumeSnapshotGetResponseAllOf2.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *VolumeSnapshotGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolumeSnapshotGetResponse) UnmarshalBinary(b []byte) error {
	var res VolumeSnapshotGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
