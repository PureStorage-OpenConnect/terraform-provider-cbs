// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// FileSystemOAIGenAllOf1 file system Oai gen all of1
//
// swagger:model fileSystemOaiGenAllOf1
type FileSystemOAIGenAllOf1 struct {

	// The file system creation time, measured in milliseconds since the UNIX epoch.
	// Read Only: true
	Created int64 `json:"created,omitempty"`

	// Returns a value of `true` if the file system has been destroyed and is pending eradication. The `time_remaining` value displays the amount of time left until the destroyed file system is permanently eradicated. Before the `time_remaining` period has elapsed, the destroyed file system can be recovered by setting `destroyed=false`. Once the `time_remaining` period has elapsed, the file system is permanently eradicated and can no longer be recovered.
	Destroyed bool `json:"destroyed,omitempty"`

	// The amount of time left, measured in milliseconds until the destroyed file system is permanently eradicated. Before the `time_remaining` period has elapsed, the destroyed file system can be recovered by setting `destroyed=false`.
	// Read Only: true
	TimeRemaining int64 `json:"time_remaining,omitempty"`
}

// Validate validates this file system Oai gen all of1
func (m *FileSystemOAIGenAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *FileSystemOAIGenAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FileSystemOAIGenAllOf1) UnmarshalBinary(b []byte) error {
	var res FileSystemOAIGenAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
