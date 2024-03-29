// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new maintenance windows API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for maintenance windows API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPI24MaintenanceWindows(params *DeleteAPI24MaintenanceWindowsParams) (*DeleteApi24MaintenanceWindowsOK, error)

	GetAPI24MaintenanceWindows(params *GetAPI24MaintenanceWindowsParams) (*GetApi24MaintenanceWindowsOK, error)

	PostAPI24MaintenanceWindows(params *PostAPI24MaintenanceWindowsParams) (*PostApi24MaintenanceWindowsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPI24MaintenanceWindows deletes maintenance window

  Deletes an open maintenance window before its scheduled end (`expire`) time. The `names` parameter is required and must be set to `environment`.
*/
func (a *Client) DeleteAPI24MaintenanceWindows(params *DeleteAPI24MaintenanceWindowsParams) (*DeleteApi24MaintenanceWindowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPI24MaintenanceWindowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAPI24MaintenanceWindows",
		Method:             "DELETE",
		PathPattern:        "/api/2.4/maintenance-windows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAPI24MaintenanceWindowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteApi24MaintenanceWindowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPI24MaintenanceWindows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPI24MaintenanceWindows lists maintenance window details

  Displays maintenance window details, including start time, end time, and maintenance type.
*/
func (a *Client) GetAPI24MaintenanceWindows(params *GetAPI24MaintenanceWindowsParams) (*GetApi24MaintenanceWindowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPI24MaintenanceWindowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPI24MaintenanceWindows",
		Method:             "GET",
		PathPattern:        "/api/2.4/maintenance-windows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPI24MaintenanceWindowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApi24MaintenanceWindowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPI24MaintenanceWindows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPI24MaintenanceWindows creates a maintenance window

  Creates a maintenance window that suppresses alerts for a specified period of time. A maintenance window can be manually closed at any time. The `names` and `timeout` parameters are required. Set the `names` parameter to `environment`.
*/
func (a *Client) PostAPI24MaintenanceWindows(params *PostAPI24MaintenanceWindowsParams) (*PostApi24MaintenanceWindowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPI24MaintenanceWindowsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPI24MaintenanceWindows",
		Method:             "POST",
		PathPattern:        "/api/2.4/maintenance-windows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPI24MaintenanceWindowsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostApi24MaintenanceWindowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPI24MaintenanceWindows: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
