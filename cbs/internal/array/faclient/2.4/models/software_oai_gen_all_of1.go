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

// SoftwareOAIGenAllOf1 software Oai gen all of1
//
// swagger:model softwareOaiGenAllOf1
type SoftwareOAIGenAllOf1 struct {

	// A list of steps that are planned to run during the upgrade in an optimal scenario (i.e., all upgrade checks pass, no step is retried, and the upgrade is not aborted). Steps are listed in the order that they should occur.
	UpgradePlan []*SoftwareOAIGenAllOf1UpgradePlanItems `json:"upgrade_plan"`
}

// Validate validates this software Oai gen all of1
func (m *SoftwareOAIGenAllOf1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUpgradePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareOAIGenAllOf1) validateUpgradePlan(formats strfmt.Registry) error {

	if swag.IsZero(m.UpgradePlan) { // not required
		return nil
	}

	for i := 0; i < len(m.UpgradePlan); i++ {
		if swag.IsZero(m.UpgradePlan[i]) { // not required
			continue
		}

		if m.UpgradePlan[i] != nil {
			if err := m.UpgradePlan[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("upgrade_plan" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareOAIGenAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareOAIGenAllOf1) UnmarshalBinary(b []byte) error {
	var res SoftwareOAIGenAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
