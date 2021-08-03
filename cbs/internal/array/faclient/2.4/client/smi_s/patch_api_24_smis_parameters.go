// Code generated by go-swagger; DO NOT EDIT.

package smi_s

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

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// NewPatchAPI24SMISParams creates a new PatchAPI24SMISParams object
// with the default values initialized.
func NewPatchAPI24SMISParams() *PatchAPI24SMISParams {
	var ()
	return &PatchAPI24SMISParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24SMISParamsWithTimeout creates a new PatchAPI24SMISParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24SMISParamsWithTimeout(timeout time.Duration) *PatchAPI24SMISParams {
	var ()
	return &PatchAPI24SMISParams{

		timeout: timeout,
	}
}

// NewPatchAPI24SMISParamsWithContext creates a new PatchAPI24SMISParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24SMISParamsWithContext(ctx context.Context) *PatchAPI24SMISParams {
	var ()
	return &PatchAPI24SMISParams{

		Context: ctx,
	}
}

// NewPatchAPI24SMISParamsWithHTTPClient creates a new PatchAPI24SMISParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24SMISParamsWithHTTPClient(client *http.Client) *PatchAPI24SMISParams {
	var ()
	return &PatchAPI24SMISParams{
		HTTPClient: client,
	}
}

/*PatchAPI24SMISParams contains all the parameters to send to the API endpoint
for the patch API 24 SMIS operation typically these are written to a http.Request
*/
type PatchAPI24SMISParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*SMIs*/
	SMIs *models.SMIS

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) WithTimeout(timeout time.Duration) *PatchAPI24SMISParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) WithContext(ctx context.Context) *PatchAPI24SMISParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) WithHTTPClient(client *http.Client) *PatchAPI24SMISParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) WithAuthorization(authorization *string) *PatchAPI24SMISParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) WithXRequestID(xRequestID *string) *PatchAPI24SMISParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSMIs adds the sMIs to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) WithSMIs(sMIs *models.SMIS) *PatchAPI24SMISParams {
	o.SetSMIs(sMIs)
	return o
}

// SetSMIs adds the smiS to the patch API 24 SMIS params
func (o *PatchAPI24SMISParams) SetSMIs(sMIs *models.SMIS) {
	o.SMIs = sMIs
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24SMISParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SMIs != nil {
		if err := r.SetBodyParam(o.SMIs); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}