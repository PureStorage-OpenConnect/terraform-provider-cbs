// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PodOAIGen1AllOf1 pod Oai gen1 all of1
//
// swagger:model podOaiGen1AllOf1
type PodOAIGen1AllOf1 struct {

	// Number of source pods that link to the pod.
	LinkSourceCount int64 `json:"link_source_count,omitempty"`

	// Number of target pods that link to the pod.
	LinkTargetCount int64 `json:"link_target_count,omitempty"`

	// Current promotion status of a pod. Valid values are `promoted`, `demoted`, and `promoting`. The `promoted` status indicates that the pod has been promoted. The pod takes writes from hosts instead of incorporating replicated data. This is the default mode for a pod when it is created. The `demoted` status indicates that the pod has been demoted. The pod does not accept write requests and is ready to be used as a link target. The `promoting` status indicates that the pod is in an intermediate status between `demoted` and `promoted` while the promotion process is taking place.
	PromotionStatus string `json:"promotion_status,omitempty"`

	// Valid values are `promoted` and `demoted`. Patch `requested_promotion_state` to `demoted` to demote the pod so that it can be used as a link target for continuous replication between pods. Demoted pods do not accept write requests, and a destroyed version of the pod with `undo-demote` appended to the pod name is created on the array with the state of the pod when it was in the promoted state. Patch `requested_promotion_state` to `promoted` to start the process of promoting the pod. The `promotion_status` indicates when the pod has been successfully promoted. Promoted pods stop incorporating replicated data from the source pod and start accepting write requests. The replication process does not stop as the source pod continues replicating data to the pod. The space consumed by the unique replicated data is tracked by the `space.journal` field of the pod.
	RequestedPromotionState string `json:"requested_promotion_state,omitempty"`

	// space
	Space *PodSpace `json:"space,omitempty"`
}

// Validate validates this pod Oai gen1 all of1
func (m *PodOAIGen1AllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSpace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PodOAIGen1AllOf1) validateSpace(formats strfmt.Registry) error {

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
func (m *PodOAIGen1AllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PodOAIGen1AllOf1) UnmarshalBinary(b []byte) error {
	var res PodOAIGen1AllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}