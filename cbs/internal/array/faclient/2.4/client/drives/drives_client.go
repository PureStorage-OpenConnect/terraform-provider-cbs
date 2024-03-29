// Code generated by go-swagger; DO NOT EDIT.

package drives

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new drives API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for drives API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPI24Drives(params *GetAPI24DrivesParams) (*GetApi24DrivesOK, error)

	PatchAPI24Drives(params *PatchAPI24DrivesParams) (*PatchApi24DrivesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAPI24Drives lists flash n v RAM and cache modules

  Displays a list of flash, NVRAM, and cache modules that are installed in the array along with their attributes and status.
*/
func (a *Client) GetAPI24Drives(params *GetAPI24DrivesParams) (*GetApi24DrivesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPI24DrivesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPI24Drives",
		Method:             "GET",
		PathPattern:        "/api/2.4/drives",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPI24DrivesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApi24DrivesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPI24Drives: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPI24Drives modifies flash and n v RAM modules

  Modifies flash and NVRAM modules that have been added or connected but not yet admitted to the array.
*/
func (a *Client) PatchAPI24Drives(params *PatchAPI24DrivesParams) (*PatchApi24DrivesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPI24DrivesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAPI24Drives",
		Method:             "PATCH",
		PathPattern:        "/api/2.4/drives",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPI24DrivesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchApi24DrivesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPI24Drives: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
