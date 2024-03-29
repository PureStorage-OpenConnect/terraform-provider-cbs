// Code generated by go-swagger; DO NOT EDIT.

package syslog

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

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// NewPostAPI24SyslogServersParams creates a new PostAPI24SyslogServersParams object
// with the default values initialized.
func NewPostAPI24SyslogServersParams() *PostAPI24SyslogServersParams {
	var ()
	return &PostAPI24SyslogServersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24SyslogServersParamsWithTimeout creates a new PostAPI24SyslogServersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24SyslogServersParamsWithTimeout(timeout time.Duration) *PostAPI24SyslogServersParams {
	var ()
	return &PostAPI24SyslogServersParams{

		timeout: timeout,
	}
}

// NewPostAPI24SyslogServersParamsWithContext creates a new PostAPI24SyslogServersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24SyslogServersParamsWithContext(ctx context.Context) *PostAPI24SyslogServersParams {
	var ()
	return &PostAPI24SyslogServersParams{

		Context: ctx,
	}
}

// NewPostAPI24SyslogServersParamsWithHTTPClient creates a new PostAPI24SyslogServersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24SyslogServersParamsWithHTTPClient(client *http.Client) *PostAPI24SyslogServersParams {
	var ()
	return &PostAPI24SyslogServersParams{
		HTTPClient: client,
	}
}

/*PostAPI24SyslogServersParams contains all the parameters to send to the API endpoint
for the post API 24 syslog servers operation typically these are written to a http.Request
*/
type PostAPI24SyslogServersParams struct {

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
	/*SyslogServer*/
	SyslogServer *models.SyslogServer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) WithTimeout(timeout time.Duration) *PostAPI24SyslogServersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) WithContext(ctx context.Context) *PostAPI24SyslogServersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) WithHTTPClient(client *http.Client) *PostAPI24SyslogServersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) WithAuthorization(authorization *string) *PostAPI24SyslogServersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) WithXRequestID(xRequestID *string) *PostAPI24SyslogServersParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNames adds the names to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) WithNames(names []string) *PostAPI24SyslogServersParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) SetNames(names []string) {
	o.Names = names
}

// WithSyslogServer adds the syslogServer to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) WithSyslogServer(syslogServer *models.SyslogServer) *PostAPI24SyslogServersParams {
	o.SetSyslogServer(syslogServer)
	return o
}

// SetSyslogServer adds the syslogServer to the post API 24 syslog servers params
func (o *PostAPI24SyslogServersParams) SetSyslogServer(syslogServer *models.SyslogServer) {
	o.SyslogServer = syslogServer
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24SyslogServersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SyslogServer != nil {
		if err := r.SetBodyParam(o.SyslogServer); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
