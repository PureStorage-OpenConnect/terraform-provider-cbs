// Code generated by go-swagger; DO NOT EDIT.

package protection_groups

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

// NewPatchAPI24ProtectionGroupsParams creates a new PatchAPI24ProtectionGroupsParams object
// with the default values initialized.
func NewPatchAPI24ProtectionGroupsParams() *PatchAPI24ProtectionGroupsParams {
	var ()
	return &PatchAPI24ProtectionGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24ProtectionGroupsParamsWithTimeout creates a new PatchAPI24ProtectionGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24ProtectionGroupsParamsWithTimeout(timeout time.Duration) *PatchAPI24ProtectionGroupsParams {
	var ()
	return &PatchAPI24ProtectionGroupsParams{

		timeout: timeout,
	}
}

// NewPatchAPI24ProtectionGroupsParamsWithContext creates a new PatchAPI24ProtectionGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24ProtectionGroupsParamsWithContext(ctx context.Context) *PatchAPI24ProtectionGroupsParams {
	var ()
	return &PatchAPI24ProtectionGroupsParams{

		Context: ctx,
	}
}

// NewPatchAPI24ProtectionGroupsParamsWithHTTPClient creates a new PatchAPI24ProtectionGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24ProtectionGroupsParamsWithHTTPClient(client *http.Client) *PatchAPI24ProtectionGroupsParams {
	var ()
	return &PatchAPI24ProtectionGroupsParams{
		HTTPClient: client,
	}
}

/*PatchAPI24ProtectionGroupsParams contains all the parameters to send to the API endpoint
for the patch API 24 protection groups operation typically these are written to a http.Request
*/
type PatchAPI24ProtectionGroupsParams struct {

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
	/*ProtectionGroup*/
	ProtectionGroup *models.ProtectionGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) WithTimeout(timeout time.Duration) *PatchAPI24ProtectionGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) WithContext(ctx context.Context) *PatchAPI24ProtectionGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) WithHTTPClient(client *http.Client) *PatchAPI24ProtectionGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) WithAuthorization(authorization *string) *PatchAPI24ProtectionGroupsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) WithXRequestID(xRequestID *string) *PatchAPI24ProtectionGroupsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNames adds the names to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) WithNames(names []string) *PatchAPI24ProtectionGroupsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) SetNames(names []string) {
	o.Names = names
}

// WithProtectionGroup adds the protectionGroup to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) WithProtectionGroup(protectionGroup *models.ProtectionGroup) *PatchAPI24ProtectionGroupsParams {
	o.SetProtectionGroup(protectionGroup)
	return o
}

// SetProtectionGroup adds the protectionGroup to the patch API 24 protection groups params
func (o *PatchAPI24ProtectionGroupsParams) SetProtectionGroup(protectionGroup *models.ProtectionGroup) {
	o.ProtectionGroup = protectionGroup
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24ProtectionGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ProtectionGroup != nil {
		if err := r.SetBodyParam(o.ProtectionGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
