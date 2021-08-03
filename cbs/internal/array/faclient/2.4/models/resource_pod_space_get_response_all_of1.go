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

// ResourcePodSpaceGetResponseAllOf1 resource pod space get response all of1
//
// swagger:model resourcePodSpaceGetResponseAllOf1
type ResourcePodSpaceGetResponseAllOf1 struct {

	// Displays a list of all items after filtering. The values are displayed for each name where meaningful. If `total_only=true`, the `items` list will be empty.
	Items []*ResourcePodSpace `json:"items"`

	// The aggregate value of all items after filtering. When applicable, the average value is displayed instead. The values are displayed for each field if meaningful.
	Total []*ResourcePodSpace `json:"total"`
}

// Validate validates this resource pod space get response all of1
func (m *ResourcePodSpaceGetResponseAllOf1) Validate(formats strfmt.Registry) error {
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

func (m *ResourcePodSpaceGetResponseAllOf1) validateItems(formats strfmt.Registry) error {

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

func (m *ResourcePodSpaceGetResponseAllOf1) validateTotal(formats strfmt.Registry) error {

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
func (m *ResourcePodSpaceGetResponseAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcePodSpaceGetResponseAllOf1) UnmarshalBinary(b []byte) error {
	var res ResourcePodSpaceGetResponseAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}