// Code generated by go-swagger; DO NOT EDIT.

package syslog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new syslog API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for syslog API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteAPI24SyslogServers(params *DeleteAPI24SyslogServersParams) (*DeleteApi24SyslogServersOK, error)

	GetAPI24SyslogServers(params *GetAPI24SyslogServersParams) (*GetApi24SyslogServersOK, error)

	GetAPI24SyslogServersSettings(params *GetAPI24SyslogServersSettingsParams) (*GetApi24SyslogServersSettingsOK, error)

	GetAPI24SyslogServersTest(params *GetAPI24SyslogServersTestParams) (*GetApi24SyslogServersTestOK, error)

	PatchAPI24SyslogServers(params *PatchAPI24SyslogServersParams) (*PatchApi24SyslogServersOK, error)

	PatchAPI24SyslogServersSettings(params *PatchAPI24SyslogServersSettingsParams) (*PatchApi24SyslogServersSettingsOK, error)

	PostAPI24SyslogServers(params *PostAPI24SyslogServersParams) (*PostApi24SyslogServersOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteAPI24SyslogServers deletes syslog server

  Deletes a configured syslog server and stop forwarding syslog messages.
*/
func (a *Client) DeleteAPI24SyslogServers(params *DeleteAPI24SyslogServersParams) (*DeleteApi24SyslogServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPI24SyslogServersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAPI24SyslogServers",
		Method:             "DELETE",
		PathPattern:        "/api/2.4/syslog-servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteAPI24SyslogServersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteApi24SyslogServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for DeleteAPI24SyslogServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPI24SyslogServers lists syslog servers

  Displays a list of configured syslog servers.
*/
func (a *Client) GetAPI24SyslogServers(params *GetAPI24SyslogServersParams) (*GetApi24SyslogServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPI24SyslogServersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPI24SyslogServers",
		Method:             "GET",
		PathPattern:        "/api/2.4/syslog-servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPI24SyslogServersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApi24SyslogServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPI24SyslogServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPI24SyslogServersSettings lists syslog settings

  Displays syslog settings. Values include `continuation_token`, `items`, `more_items_remaining`, and `total_item_count`.
*/
func (a *Client) GetAPI24SyslogServersSettings(params *GetAPI24SyslogServersSettingsParams) (*GetApi24SyslogServersSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPI24SyslogServersSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPI24SyslogServersSettings",
		Method:             "GET",
		PathPattern:        "/api/2.4/syslog-servers/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPI24SyslogServersSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApi24SyslogServersSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPI24SyslogServersSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetAPI24SyslogServersTest lists syslog server test results

  Displays syslog server test results, which indicate whether the syslog is working and configured correctly.
*/
func (a *Client) GetAPI24SyslogServersTest(params *GetAPI24SyslogServersTestParams) (*GetApi24SyslogServersTestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPI24SyslogServersTestParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPI24SyslogServersTest",
		Method:             "GET",
		PathPattern:        "/api/2.4/syslog-servers/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPI24SyslogServersTestReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApi24SyslogServersTestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPI24SyslogServersTest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPI24SyslogServers modifies syslog server

  Modifies the URI of a configured syslog server.
*/
func (a *Client) PatchAPI24SyslogServers(params *PatchAPI24SyslogServersParams) (*PatchApi24SyslogServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPI24SyslogServersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAPI24SyslogServers",
		Method:             "PATCH",
		PathPattern:        "/api/2.4/syslog-servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPI24SyslogServersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchApi24SyslogServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPI24SyslogServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PatchAPI24SyslogServersSettings modifies syslog settings

  Modifies syslog settings. Values include `continuation_token`, `items`, `more_items_remaining`, and `total_item_count`.
*/
func (a *Client) PatchAPI24SyslogServersSettings(params *PatchAPI24SyslogServersSettingsParams) (*PatchApi24SyslogServersSettingsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchAPI24SyslogServersSettingsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PatchAPI24SyslogServersSettings",
		Method:             "PATCH",
		PathPattern:        "/api/2.4/syslog-servers/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PatchAPI24SyslogServersSettingsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchApi24SyslogServersSettingsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PatchAPI24SyslogServersSettings: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPI24SyslogServers creates syslog server

  Creates a new syslog server. Transmission of syslog messages is enabled immediately.
*/
func (a *Client) PostAPI24SyslogServers(params *PostAPI24SyslogServersParams) (*PostApi24SyslogServersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPI24SyslogServersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPI24SyslogServers",
		Method:             "POST",
		PathPattern:        "/api/2.4/syslog-servers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPI24SyslogServersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostApi24SyslogServersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPI24SyslogServers: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
