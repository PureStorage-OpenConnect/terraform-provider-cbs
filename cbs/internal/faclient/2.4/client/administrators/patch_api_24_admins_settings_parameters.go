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

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// NewPatchAPI24AdminsSettingsParams creates a new PatchAPI24AdminsSettingsParams object
// with the default values initialized.
func NewPatchAPI24AdminsSettingsParams() *PatchAPI24AdminsSettingsParams {
	var ()
	return &PatchAPI24AdminsSettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24AdminsSettingsParamsWithTimeout creates a new PatchAPI24AdminsSettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24AdminsSettingsParamsWithTimeout(timeout time.Duration) *PatchAPI24AdminsSettingsParams {
	var ()
	return &PatchAPI24AdminsSettingsParams{

		timeout: timeout,
	}
}

// NewPatchAPI24AdminsSettingsParamsWithContext creates a new PatchAPI24AdminsSettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24AdminsSettingsParamsWithContext(ctx context.Context) *PatchAPI24AdminsSettingsParams {
	var ()
	return &PatchAPI24AdminsSettingsParams{

		Context: ctx,
	}
}

// NewPatchAPI24AdminsSettingsParamsWithHTTPClient creates a new PatchAPI24AdminsSettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24AdminsSettingsParamsWithHTTPClient(client *http.Client) *PatchAPI24AdminsSettingsParams {
	var ()
	return &PatchAPI24AdminsSettingsParams{
		HTTPClient: client,
	}
}

/*PatchAPI24AdminsSettingsParams contains all the parameters to send to the API endpoint
for the patch API 24 admins settings operation typically these are written to a http.Request
*/
type PatchAPI24AdminsSettingsParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*AdminSettings*/
	AdminSettings *models.AdminSettings

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) WithTimeout(timeout time.Duration) *PatchAPI24AdminsSettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) WithContext(ctx context.Context) *PatchAPI24AdminsSettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) WithHTTPClient(client *http.Client) *PatchAPI24AdminsSettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) WithAuthorization(authorization *string) *PatchAPI24AdminsSettingsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) WithXRequestID(xRequestID *string) *PatchAPI24AdminsSettingsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAdminSettings adds the adminSettings to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) WithAdminSettings(adminSettings *models.AdminSettings) *PatchAPI24AdminsSettingsParams {
	o.SetAdminSettings(adminSettings)
	return o
}

// SetAdminSettings adds the adminSettings to the patch API 24 admins settings params
func (o *PatchAPI24AdminsSettingsParams) SetAdminSettings(adminSettings *models.AdminSettings) {
	o.AdminSettings = adminSettings
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24AdminsSettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.AdminSettings != nil {
		if err := r.SetBodyParam(o.AdminSettings); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
