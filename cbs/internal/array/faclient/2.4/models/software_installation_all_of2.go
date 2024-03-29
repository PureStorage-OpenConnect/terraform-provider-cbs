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

// SoftwareInstallationAllOf2 Interactive software upgrade attempt that contains information about the current upgrade instance, including the current step, status of the upgrade attempt, etc.
//
// swagger:model softwareInstallationAllOf2
type SoftwareInstallationAllOf2 struct {

	// The `id` of the current step, or `null` if the upgrade is not active.
	CurrentStepID string `json:"current_step_id,omitempty"`

	// The detailed reason for the `status`.
	Details string `json:"details,omitempty"`

	// Mode that the upgrade is in. Valid values are `check-only`, `interactive`, `semi-interactive`, and `one-click`. In `check-only` mode, the upgrade only runs pre-upgrade checks and returns. In `interactive` mode, the upgrade process pauses at several points, at which users must enter certain commands to proceed. In `semi-interactive` mode, the upgrade pauses if there are any upgrade check failures, and functions like `one-click` mode otherwise. In `one-click` mode, the upgrade proceeds automatically without pausing.
	Mode string `json:"mode,omitempty"`

	// A list of upgrade checks whose failure is overridden during the upgrade. If any optional `args` are provided, they are validated later when the corresponding check script runs.
	OverrideChecks []*OverrideCheck `json:"override_checks"`

	// software
	Software *SoftwareInstallationAllOf2Software `json:"software,omitempty"`

	// Status of the upgrade. Valid values are `installing`, `paused`, `aborting`, `aborted`, and `finished`. A status of `installing` indicates that the upgrade is running. A status of `paused` indicates that the upgrade is paused and waiting for user input. A status of `aborting` indicates that the upgrade is being aborted. A status of `aborted` indicates that the upgrade has stopped due to an abort. A status of `finished` indicates that the upgrade has finished successfully.
	Status string `json:"status,omitempty"`
}

// Validate validates this software installation all of2
func (m *SoftwareInstallationAllOf2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOverrideChecks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSoftware(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SoftwareInstallationAllOf2) validateOverrideChecks(formats strfmt.Registry) error {

	if swag.IsZero(m.OverrideChecks) { // not required
		return nil
	}

	for i := 0; i < len(m.OverrideChecks); i++ {
		if swag.IsZero(m.OverrideChecks[i]) { // not required
			continue
		}

		if m.OverrideChecks[i] != nil {
			if err := m.OverrideChecks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("override_checks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SoftwareInstallationAllOf2) validateSoftware(formats strfmt.Registry) error {

	if swag.IsZero(m.Software) { // not required
		return nil
	}

	if m.Software != nil {
		if err := m.Software.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("software")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareInstallationAllOf2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareInstallationAllOf2) UnmarshalBinary(b []byte) error {
	var res SoftwareInstallationAllOf2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
