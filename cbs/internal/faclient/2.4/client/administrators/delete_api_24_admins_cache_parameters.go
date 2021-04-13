// Code generated by go-swagger; DO NOT EDIT.

package administrators

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

// NewDeleteAPI24AdminsCacheParams creates a new DeleteAPI24AdminsCacheParams object
// with the default values initialized.
func NewDeleteAPI24AdminsCacheParams() *DeleteAPI24AdminsCacheParams {
	var ()
	return &DeleteAPI24AdminsCacheParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24AdminsCacheParamsWithTimeout creates a new DeleteAPI24AdminsCacheParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24AdminsCacheParamsWithTimeout(timeout time.Duration) *DeleteAPI24AdminsCacheParams {
	var ()
	return &DeleteAPI24AdminsCacheParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24AdminsCacheParamsWithContext creates a new DeleteAPI24AdminsCacheParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24AdminsCacheParamsWithContext(ctx context.Context) *DeleteAPI24AdminsCacheParams {
	var ()
	return &DeleteAPI24AdminsCacheParams{

		Context: ctx,
	}
}

// NewDeleteAPI24AdminsCacheParamsWithHTTPClient creates a new DeleteAPI24AdminsCacheParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24AdminsCacheParamsWithHTTPClient(client *http.Client) *DeleteAPI24AdminsCacheParams {
	var ()
	return &DeleteAPI24AdminsCacheParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24AdminsCacheParams contains all the parameters to send to the API endpoint
for the delete API 24 admins cache operation typically these are written to a http.Request
*/
type DeleteAPI24AdminsCacheParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*RemoveAllEntries
	  If set to `true`, removes all entries from the administrator cache.

	*/
	RemoveAllEntries bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) WithTimeout(timeout time.Duration) *DeleteAPI24AdminsCacheParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) WithContext(ctx context.Context) *DeleteAPI24AdminsCacheParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) WithHTTPClient(client *http.Client) *DeleteAPI24AdminsCacheParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) WithAuthorization(authorization *string) *DeleteAPI24AdminsCacheParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) WithXRequestID(xRequestID *string) *DeleteAPI24AdminsCacheParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithRemoveAllEntries adds the removeAllEntries to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) WithRemoveAllEntries(removeAllEntries bool) *DeleteAPI24AdminsCacheParams {
	o.SetRemoveAllEntries(removeAllEntries)
	return o
}

// SetRemoveAllEntries adds the removeAllEntries to the delete API 24 admins cache params
func (o *DeleteAPI24AdminsCacheParams) SetRemoveAllEntries(removeAllEntries bool) {
	o.RemoveAllEntries = removeAllEntries
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24AdminsCacheParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param remove_all_entries
	qrRemoveAllEntries := o.RemoveAllEntries
	qRemoveAllEntries := swag.FormatBool(qrRemoveAllEntries)
	if qRemoveAllEntries != "" {

		if err := r.SetQueryParam("remove_all_entries", qRemoveAllEntries); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
