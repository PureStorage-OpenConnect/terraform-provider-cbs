// Code generated by go-swagger; DO NOT EDIT.

package snmp_managers

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

// NewDeleteAPI24SNMPManagersParams creates a new DeleteAPI24SNMPManagersParams object
// with the default values initialized.
func NewDeleteAPI24SNMPManagersParams() *DeleteAPI24SNMPManagersParams {
	var ()
	return &DeleteAPI24SNMPManagersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24SNMPManagersParamsWithTimeout creates a new DeleteAPI24SNMPManagersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24SNMPManagersParamsWithTimeout(timeout time.Duration) *DeleteAPI24SNMPManagersParams {
	var ()
	return &DeleteAPI24SNMPManagersParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24SNMPManagersParamsWithContext creates a new DeleteAPI24SNMPManagersParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24SNMPManagersParamsWithContext(ctx context.Context) *DeleteAPI24SNMPManagersParams {
	var ()
	return &DeleteAPI24SNMPManagersParams{

		Context: ctx,
	}
}

// NewDeleteAPI24SNMPManagersParamsWithHTTPClient creates a new DeleteAPI24SNMPManagersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24SNMPManagersParamsWithHTTPClient(client *http.Client) *DeleteAPI24SNMPManagersParams {
	var ()
	return &DeleteAPI24SNMPManagersParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24SNMPManagersParams contains all the parameters to send to the API endpoint
for the delete API 24 SNMP managers operation typically these are written to a http.Request
*/
type DeleteAPI24SNMPManagersParams struct {

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

// WithTimeout adds the timeout to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) WithTimeout(timeout time.Duration) *DeleteAPI24SNMPManagersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) WithContext(ctx context.Context) *DeleteAPI24SNMPManagersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) WithHTTPClient(client *http.Client) *DeleteAPI24SNMPManagersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) WithAuthorization(authorization *string) *DeleteAPI24SNMPManagersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) WithXRequestID(xRequestID *string) *DeleteAPI24SNMPManagersParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNames adds the names to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) WithNames(names []string) *DeleteAPI24SNMPManagersParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the delete API 24 SNMP managers params
func (o *DeleteAPI24SNMPManagersParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24SNMPManagersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
