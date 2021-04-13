// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyrulesnapshotpostRules policyrulesnapshotpost rules
//
// swagger:model policyrulesnapshotpostRules
type PolicyrulesnapshotpostRules struct {

	// Specifies the number of milliseconds since midnight at which to take a snapshot. The `at` value can only be set to an hour and must be between 0 and 82800000, inclusive. The `at` value can only be set on the rule with the smallest `every` value. The `at` value cannot be set if the `every` value is not measured in days. The `at` value can only be set for at most one rule in the same policy.
	At int64 `json:"at,omitempty"`

	// The snapshot client name. A full snapshot name is constructed in the form of `DIR.CLIENT_NAME.SUFFIX` where `DIR` is the managed directory name, `CLIENT_NAME` is the snapshot client name, and `SUFFIX` is a monotonically increasing number generated by the system. The client visible snapshot name is `CLIENT_NAME.SUFFIX`.
	ClientName string `json:"client_name,omitempty"`

	// Specifies the interval between snapshots, in milliseconds. The `every` value for all rules must be multiples of one another. The `every` value must be unique for each rule in the same policy. The `every` value must be between 5 minutes and 1 year.
	Every int64 `json:"every,omitempty"`

	// Specifies the period that snapshots are retained before they are eradicated, in milliseconds. The `keep_for` value cannot be less than the `every` value of the rule. The `keep_for` value must be unique for each rule in the same policy. The `keep_for` value must be between 5 minutes and 1 year. The `keep_for` value cannot be less than the `keep_for` value of any rule in the same policy with a smaller `every` value.
	KeepFor int64 `json:"keep_for,omitempty"`
}

// Validate validates this policyrulesnapshotpost rules
func (m *PolicyrulesnapshotpostRules) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyrulesnapshotpostRules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyrulesnapshotpostRules) UnmarshalBinary(b []byte) error {
	var res PolicyrulesnapshotpostRules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
