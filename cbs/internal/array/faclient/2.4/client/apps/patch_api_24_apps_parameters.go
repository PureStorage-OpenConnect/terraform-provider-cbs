// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewPatchAPI24AppsParams creates a new PatchAPI24AppsParams object
// with the default values initialized.
func NewPatchAPI24AppsParams() *PatchAPI24AppsParams {
	var ()
	return &PatchAPI24AppsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24AppsParamsWithTimeout creates a new PatchAPI24AppsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24AppsParamsWithTimeout(timeout time.Duration) *PatchAPI24AppsParams {
	var ()
	return &PatchAPI24AppsParams{

		timeout: timeout,
	}
}

// NewPatchAPI24AppsParamsWithContext creates a new PatchAPI24AppsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24AppsParamsWithContext(ctx context.Context) *PatchAPI24AppsParams {
	var ()
	return &PatchAPI24AppsParams{

		Context: ctx,
	}
}

// NewPatchAPI24AppsParamsWithHTTPClient creates a new PatchAPI24AppsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24AppsParamsWithHTTPClient(client *http.Client) *PatchAPI24AppsParams {
	var ()
	return &PatchAPI24AppsParams{
		HTTPClient: client,
	}
}

/*PatchAPI24AppsParams contains all the parameters to send to the API endpoint
for the patch API 24 apps operation typically these are written to a http.Request
*/
type PatchAPI24AppsParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*App*/
	App *models.App
	/*Names
	  Performs the operation on the unique name specified. For example, `name01`. Enter multiple names in comma-separated format.

	*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 apps params
func (o *PatchAPI24AppsParams) WithTimeout(timeout time.Duration) *PatchAPI24AppsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 apps params
func (o *PatchAPI24AppsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 apps params
func (o *PatchAPI24AppsParams) WithContext(ctx context.Context) *PatchAPI24AppsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 apps params
func (o *PatchAPI24AppsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 apps params
func (o *PatchAPI24AppsParams) WithHTTPClient(client *http.Client) *PatchAPI24AppsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 apps params
func (o *PatchAPI24AppsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 apps params
func (o *PatchAPI24AppsParams) WithAuthorization(authorization *string) *PatchAPI24AppsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 apps params
func (o *PatchAPI24AppsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 apps params
func (o *PatchAPI24AppsParams) WithXRequestID(xRequestID *string) *PatchAPI24AppsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 apps params
func (o *PatchAPI24AppsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithApp adds the app to the patch API 24 apps params
func (o *PatchAPI24AppsParams) WithApp(app *models.App) *PatchAPI24AppsParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the patch API 24 apps params
func (o *PatchAPI24AppsParams) SetApp(app *models.App) {
	o.App = app
}

// WithNames adds the names to the patch API 24 apps params
func (o *PatchAPI24AppsParams) WithNames(names []string) *PatchAPI24AppsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the patch API 24 apps params
func (o *PatchAPI24AppsParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24AppsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.App != nil {
		if err := r.SetBodyParam(o.App); err != nil {
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
