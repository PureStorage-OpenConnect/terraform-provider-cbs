// Code generated by go-swagger; DO NOT EDIT.

package file_systems

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

// NewDeleteAPI24FileSystemsParams creates a new DeleteAPI24FileSystemsParams object
// with the default values initialized.
func NewDeleteAPI24FileSystemsParams() *DeleteAPI24FileSystemsParams {
	var ()
	return &DeleteAPI24FileSystemsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24FileSystemsParamsWithTimeout creates a new DeleteAPI24FileSystemsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24FileSystemsParamsWithTimeout(timeout time.Duration) *DeleteAPI24FileSystemsParams {
	var ()
	return &DeleteAPI24FileSystemsParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24FileSystemsParamsWithContext creates a new DeleteAPI24FileSystemsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24FileSystemsParamsWithContext(ctx context.Context) *DeleteAPI24FileSystemsParams {
	var ()
	return &DeleteAPI24FileSystemsParams{

		Context: ctx,
	}
}

// NewDeleteAPI24FileSystemsParamsWithHTTPClient creates a new DeleteAPI24FileSystemsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24FileSystemsParamsWithHTTPClient(client *http.Client) *DeleteAPI24FileSystemsParams {
	var ()
	return &DeleteAPI24FileSystemsParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24FileSystemsParams contains all the parameters to send to the API endpoint
for the delete API 24 file systems operation typically these are written to a http.Request
*/
type DeleteAPI24FileSystemsParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*Ids
	  Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The `ids` and `names` parameters cannot be provided together.

	*/
	Ids []string
	/*Names
	  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, `name01,name02`.

	*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) WithTimeout(timeout time.Duration) *DeleteAPI24FileSystemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) WithContext(ctx context.Context) *DeleteAPI24FileSystemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) WithHTTPClient(client *http.Client) *DeleteAPI24FileSystemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) WithAuthorization(authorization *string) *DeleteAPI24FileSystemsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) WithXRequestID(xRequestID *string) *DeleteAPI24FileSystemsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithIds adds the ids to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) WithIds(ids []string) *DeleteAPI24FileSystemsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithNames adds the names to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) WithNames(names []string) *DeleteAPI24FileSystemsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the delete API 24 file systems params
func (o *DeleteAPI24FileSystemsParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24FileSystemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "csv")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
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