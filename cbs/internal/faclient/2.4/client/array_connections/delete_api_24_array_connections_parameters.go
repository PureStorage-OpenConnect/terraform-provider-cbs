// Code generated by go-swagger; DO NOT EDIT.

package array_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDeleteAPI24ArrayConnectionsParams creates a new DeleteAPI24ArrayConnectionsParams object
// with the default values initialized.
func NewDeleteAPI24ArrayConnectionsParams() *DeleteAPI24ArrayConnectionsParams {
	var ()
	return &DeleteAPI24ArrayConnectionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24ArrayConnectionsParamsWithTimeout creates a new DeleteAPI24ArrayConnectionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24ArrayConnectionsParamsWithTimeout(timeout time.Duration) *DeleteAPI24ArrayConnectionsParams {
	var ()
	return &DeleteAPI24ArrayConnectionsParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24ArrayConnectionsParamsWithContext creates a new DeleteAPI24ArrayConnectionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24ArrayConnectionsParamsWithContext(ctx context.Context) *DeleteAPI24ArrayConnectionsParams {
	var ()
	return &DeleteAPI24ArrayConnectionsParams{

		Context: ctx,
	}
}

// NewDeleteAPI24ArrayConnectionsParamsWithHTTPClient creates a new DeleteAPI24ArrayConnectionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24ArrayConnectionsParamsWithHTTPClient(client *http.Client) *DeleteAPI24ArrayConnectionsParams {
	var ()
	return &DeleteAPI24ArrayConnectionsParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24ArrayConnectionsParams contains all the parameters to send to the API endpoint
for the delete API 24 array connections operation typically these are written to a http.Request
*/
type DeleteAPI24ArrayConnectionsParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*Names
	  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, `name01,name02`.

	*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) WithTimeout(timeout time.Duration) *DeleteAPI24ArrayConnectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) WithContext(ctx context.Context) *DeleteAPI24ArrayConnectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) WithHTTPClient(client *http.Client) *DeleteAPI24ArrayConnectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) WithAuthorization(authorization *string) *DeleteAPI24ArrayConnectionsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) WithXRequestID(xRequestID *string) *DeleteAPI24ArrayConnectionsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNames adds the names to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) WithNames(names []string) *DeleteAPI24ArrayConnectionsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the delete API 24 array connections params
func (o *DeleteAPI24ArrayConnectionsParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24ArrayConnectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Authorization != nil {

		// header param Authorization
		if err := r.SetHeaderParam("Authorization", *o.Authorization); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-ID
		if err := r.SetHeaderParam("X-Request-ID", *o.XRequestID); err != nil {
			return err
		}

	}

	valuesNames := o.Names

	joinedNames := swag.JoinByFormat(valuesNames, "csv")
	// query array param names
	if err := r.SetQueryParam("names", joinedNames...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
