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

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// NewPatchAPI24SNMPManagersParams creates a new PatchAPI24SNMPManagersParams object
// with the default values initialized.
func NewPatchAPI24SNMPManagersParams() *PatchAPI24SNMPManagersParams {
	var ()
	return &PatchAPI24SNMPManagersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24SNMPManagersParamsWithTimeout creates a new PatchAPI24SNMPManagersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24SNMPManagersParamsWithTimeout(timeout time.Duration) *PatchAPI24SNMPManagersParams {
	var ()
	return &PatchAPI24SNMPManagersParams{

		timeout: timeout,
	}
}

// NewPatchAPI24SNMPManagersParamsWithContext creates a new PatchAPI24SNMPManagersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24SNMPManagersParamsWithContext(ctx context.Context) *PatchAPI24SNMPManagersParams {
	var ()
	return &PatchAPI24SNMPManagersParams{

		Context: ctx,
	}
}

// NewPatchAPI24SNMPManagersParamsWithHTTPClient creates a new PatchAPI24SNMPManagersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24SNMPManagersParamsWithHTTPClient(client *http.Client) *PatchAPI24SNMPManagersParams {
	var ()
	return &PatchAPI24SNMPManagersParams{
		HTTPClient: client,
	}
}

/*PatchAPI24SNMPManagersParams contains all the parameters to send to the API endpoint
for the patch API 24 SNMP managers operation typically these are written to a http.Request
*/
type PatchAPI24SNMPManagersParams struct {

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
	/*SNMPManager*/
	SNMPManager *models.SNMPManagerPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) WithTimeout(timeout time.Duration) *PatchAPI24SNMPManagersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) WithContext(ctx context.Context) *PatchAPI24SNMPManagersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) WithHTTPClient(client *http.Client) *PatchAPI24SNMPManagersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) WithAuthorization(authorization *string) *PatchAPI24SNMPManagersParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) WithXRequestID(xRequestID *string) *PatchAPI24SNMPManagersParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNames adds the names to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) WithNames(names []string) *PatchAPI24SNMPManagersParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) SetNames(names []string) {
	o.Names = names
}

// WithSNMPManager adds the sNMPManager to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) WithSNMPManager(sNMPManager *models.SNMPManagerPatch) *PatchAPI24SNMPManagersParams {
	o.SetSNMPManager(sNMPManager)
	return o
}

// SetSNMPManager adds the snmpManager to the patch API 24 SNMP managers params
func (o *PatchAPI24SNMPManagersParams) SetSNMPManager(sNMPManager *models.SNMPManagerPatch) {
	o.SNMPManager = sNMPManager
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24SNMPManagersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SNMPManager != nil {
		if err := r.SetBodyParam(o.SNMPManager); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
