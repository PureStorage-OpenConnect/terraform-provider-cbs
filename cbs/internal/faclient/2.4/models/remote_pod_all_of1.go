// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RemotePodAllOf1 Remote pod that's not directly stretched to this array.
//
// swagger:model remotePodAllOf1
type RemotePodAllOf1 struct {

	// arrays
	// Min Items: 1
	Arrays []*Resource `json:"arrays"`
}

// Validate validates this remote pod all of1
func (m *RemotePodAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArrays(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RemotePodAllOf1) validateArrays(formats strfmt.Registry) error {

	if swag.IsZero(m.Arrays) { // not required
		return nil
	}

	iArraysSize := int64(len(m.Arrays))

	if err := validate.MinItems("arrays", "body", iArraysSize, 1); err != nil {
		return err
	}

	for i := 0; i < len(m.Arrays); i++ {
		if swag.IsZero(m.Arrays[i]) { // not required
			continue
		}

		if m.Arrays[i] != nil {
			if err := m.Arrays[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arrays" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *RemotePodAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RemotePodAllOf1) UnmarshalBinary(b []byte) error {
	var res RemotePodAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}