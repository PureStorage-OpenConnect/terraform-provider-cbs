// Code generated by go-swagger; DO NOT EDIT.

package kmip

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

// NewPostAPI24KMIPParams creates a new PostAPI24KMIPParams object
// with the default values initialized.
func NewPostAPI24KMIPParams() *PostAPI24KMIPParams {
	var ()
	return &PostAPI24KMIPParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24KMIPParamsWithTimeout creates a new PostAPI24KMIPParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24KMIPParamsWithTimeout(timeout time.Duration) *PostAPI24KMIPParams {
	var ()
	return &PostAPI24KMIPParams{

		timeout: timeout,
	}
}

// NewPostAPI24KMIPParamsWithContext creates a new PostAPI24KMIPParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24KMIPParamsWithContext(ctx context.Context) *PostAPI24KMIPParams {
	var ()
	return &PostAPI24KMIPParams{

		Context: ctx,
	}
}

// NewPostAPI24KMIPParamsWithHTTPClient creates a new PostAPI24KMIPParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24KMIPParamsWithHTTPClient(client *http.Client) *PostAPI24KMIPParams {
	var ()
	return &PostAPI24KMIPParams{
		HTTPClient: client,
	}
}

/*PostAPI24KMIPParams contains all the parameters to send to the API endpoint
for the post API 24 KMIP operation typically these are written to a http.Request
*/
type PostAPI24KMIPParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*KMIP*/
	KMIP *models.KMIPPost
	/*Names
	  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, `name01,name02`.

	*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) WithTimeout(timeout time.Duration) *PostAPI24KMIPParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) WithContext(ctx context.Context) *PostAPI24KMIPParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) WithHTTPClient(client *http.Client) *PostAPI24KMIPParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) WithAuthorization(authorization *string) *PostAPI24KMIPParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) WithXRequestID(xRequestID *string) *PostAPI24KMIPParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithKMIP adds the kmip to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) WithKMIP(kmip *models.KMIPPost) *PostAPI24KMIPParams {
	o.SetKMIP(kmip)
	return o
}

// SetKMIP adds the kmip to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) SetKMIP(kmip *models.KMIPPost) {
	o.KMIP = kmip
}

// WithNames adds the names to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) WithNames(names []string) *PostAPI24KMIPParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the post API 24 KMIP params
func (o *PostAPI24KMIPParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24KMIPParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.KMIP != nil {
		if err := r.SetBodyParam(o.KMIP); err != nil {
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
