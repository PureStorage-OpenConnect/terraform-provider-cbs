// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DirectoryExportPost directory export post
//
// swagger:model directoryExportPost
type DirectoryExportPost struct {

	// The name of the export to create. Export names must be unique within the same protocol.
	ExportName string `json:"export_name,omitempty"`
}

// Validate validates this directory export post
func (m *DirectoryExportPost) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DirectoryExportPost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DirectoryExportPost) UnmarshalBinary(b []byte) error {
	var res DirectoryExportPost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
