// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new authorization API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for authorization API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetAPIAPIVersion(params *GetAPIAPIVersionParams) (*GetAPIAPIVersionOK, error)

	PostAPI24Login(params *PostAPI24LoginParams) (*PostApi24LoginOK, error)

	PostAPI24Logout(params *PostAPI24LogoutParams) (*PostApi24LogoutOK, error)

	PostOauth210Token(params *PostOauth210TokenParams) (*PostOauth210TokenOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetAPIAPIVersion lists available API versions

  Returns a list of available API versions. No authentication is required to access this endpoint.

*/
func (a *Client) GetAPIAPIVersion(params *GetAPIAPIVersionParams) (*GetAPIAPIVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIAPIVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPIAPIVersion",
		Method:             "GET",
		PathPattern:        "/api/api_version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetAPIAPIVersionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetAPIAPIVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetAPIAPIVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPI24Login ps o s t login placeholder

  Exchange an API token for a session token.

*/
func (a *Client) PostAPI24Login(params *PostAPI24LoginParams) (*PostApi24LoginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPI24LoginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPI24Login",
		Method:             "POST",
		PathPattern:        "/api/2.4/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPI24LoginReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostApi24LoginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPI24Login: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostAPI24Logout ps o s t logout placeholder

  Invalidate a session token.

*/
func (a *Client) PostAPI24Logout(params *PostAPI24LogoutParams) (*PostApi24LogoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPI24LogoutParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPI24Logout",
		Method:             "POST",
		PathPattern:        "/api/2.4/logout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostAPI24LogoutReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostApi24LogoutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostAPI24Logout: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostOauth210Token gets access token

  Exchanges an ID Token for an OAuth 2.0 access token.

*/
func (a *Client) PostOauth210Token(params *PostOauth210TokenParams) (*PostOauth210TokenOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOauth210TokenParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostOauth210Token",
		Method:             "POST",
		PathPattern:        "/oauth2/1.0/token",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/x-www-form-urlencoded"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostOauth210TokenReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOauth210TokenOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostOauth210Token: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
