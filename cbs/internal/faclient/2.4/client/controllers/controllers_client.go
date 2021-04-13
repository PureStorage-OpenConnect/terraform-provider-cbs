// Code generated by go-swagger; DO NOT EDIT.

package controllers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new controllers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for controllers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPI24Controllers(params *GetAPI24ControllersParams) (*GetApi24ControllersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAPI24Controllers lists controller information and status

  Displays the name, mode, FlashArray model, Purity//FA software version, and status of each controller in the array.
*/
func (a *Client) GetAPI24Controllers(params *GetAPI24ControllersParams) (*GetApi24ControllersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPI24ControllersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPI24Controllers",
		Method:             "GET",
		PathPattern:        "/api/2.4/controllers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPI24ControllersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApi24ControllersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPI24Controllers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
