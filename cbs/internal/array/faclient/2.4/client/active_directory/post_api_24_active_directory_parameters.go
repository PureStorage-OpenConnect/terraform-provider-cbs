// Code generated by go-swagger; DO NOT EDIT.

package active_directory

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

// NewPostAPI24ActiveDirectoryParams creates a new PostAPI24ActiveDirectoryParams object
// with the default values initialized.
func NewPostAPI24ActiveDirectoryParams() *PostAPI24ActiveDirectoryParams {
	var ()
	return &PostAPI24ActiveDirectoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24ActiveDirectoryParamsWithTimeout creates a new PostAPI24ActiveDirectoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24ActiveDirectoryParamsWithTimeout(timeout time.Duration) *PostAPI24ActiveDirectoryParams {
	var ()
	return &PostAPI24ActiveDirectoryParams{

		timeout: timeout,
	}
}

// NewPostAPI24ActiveDirectoryParamsWithContext creates a new PostAPI24ActiveDirectoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24ActiveDirectoryParamsWithContext(ctx context.Context) *PostAPI24ActiveDirectoryParams {
	var ()
	return &PostAPI24ActiveDirectoryParams{

		Context: ctx,
	}
}

// NewPostAPI24ActiveDirectoryParamsWithHTTPClient creates a new PostAPI24ActiveDirectoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24ActiveDirectoryParamsWithHTTPClient(client *http.Client) *PostAPI24ActiveDirectoryParams {
	var ()
	return &PostAPI24ActiveDirectoryParams{
		HTTPClient: client,
	}
}

/*PostAPI24ActiveDirectoryParams contains all the parameters to send to the API endpoint
for the post API 24 active directory operation typically these are written to a http.Request
*/
type PostAPI24ActiveDirectoryParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*ActiveDirectory*/
	ActiveDirectory *models.ActiveDirectoryPost
	/*Names
	  Performs the operation on the unique name specified. For example, `name01`. Enter multiple names in comma-separated format.

	*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) WithTimeout(timeout time.Duration) *PostAPI24ActiveDirectoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) WithContext(ctx context.Context) *PostAPI24ActiveDirectoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) WithHTTPClient(client *http.Client) *PostAPI24ActiveDirectoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) WithAuthorization(authorization *string) *PostAPI24ActiveDirectoryParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) WithXRequestID(xRequestID *string) *PostAPI24ActiveDirectoryParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithActiveDirectory adds the activeDirectory to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) WithActiveDirectory(activeDirectory *models.ActiveDirectoryPost) *PostAPI24ActiveDirectoryParams {
	o.SetActiveDirectory(activeDirectory)
	return o
}

// SetActiveDirectory adds the activeDirectory to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) SetActiveDirectory(activeDirectory *models.ActiveDirectoryPost) {
	o.ActiveDirectory = activeDirectory
}

// WithNames adds the names to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) WithNames(names []string) *PostAPI24ActiveDirectoryParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the post API 24 active directory params
func (o *PostAPI24ActiveDirectoryParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24ActiveDirectoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ActiveDirectory != nil {
		if err := r.SetBodyParam(o.ActiveDirectory); err != nil {
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
