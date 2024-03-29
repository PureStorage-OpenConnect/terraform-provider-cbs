// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodReplicaLinkPatch pod replica link patch
//
// swagger:model podReplicaLinkPatch
type PodReplicaLinkPatch struct {

	// Returns a value of `true` if the replica link is to be created in a `paused` state. Returns a value of `false` if the replica link is to be created not in a `paused` state. If not specified, defaults to `false`.
	Paused bool `json:"paused,omitempty"`
}

// Validate validates this pod replica link patch
func (m *PodReplicaLinkPatch) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PodReplicaLinkPatch) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodReplicaLinkPatch) UnmarshalBinary(b []byte) error {
	var res PodReplicaLinkPatch
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
