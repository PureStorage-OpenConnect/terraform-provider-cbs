// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OffloadAzure offload azure
//
// swagger:model offloadAzure
type OffloadAzure struct {

	// The name of the existing Azure Blob storage account.
	AccountName string `json:"account_name,omitempty"`

	// The name of the container in the Azure Blob storage account to where the data will be offloaded. The name must be a valid DNS name. If not specified, defaults to `offload`.
	ContainerName string `json:"container_name,omitempty"`

	// The secret access key that goes with the account name (`account_name`) of the Azure Blob storage account. The secret access key is only accepted when creating the connection between the array and the Azure Blob storage account. The `account_name` and `container_name`, and `secret_access_key` parameters must be set together.
	SecretAccessKey string `json:"secret_access_key,omitempty"`
}

// Validate validates this offload azure
func (m *OffloadAzure) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OffloadAzure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OffloadAzure) UnmarshalBinary(b []byte) error {
	var res OffloadAzure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
