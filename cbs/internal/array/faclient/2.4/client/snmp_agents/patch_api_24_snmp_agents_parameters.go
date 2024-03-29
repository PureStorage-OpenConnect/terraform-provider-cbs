// Code generated by go-swagger; DO NOT EDIT.

package snmp_agents

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

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// NewPatchAPI24SNMPAgentsParams creates a new PatchAPI24SNMPAgentsParams object
// with the default values initialized.
func NewPatchAPI24SNMPAgentsParams() *PatchAPI24SNMPAgentsParams {
	var ()
	return &PatchAPI24SNMPAgentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24SNMPAgentsParamsWithTimeout creates a new PatchAPI24SNMPAgentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24SNMPAgentsParamsWithTimeout(timeout time.Duration) *PatchAPI24SNMPAgentsParams {
	var ()
	return &PatchAPI24SNMPAgentsParams{

		timeout: timeout,
	}
}

// NewPatchAPI24SNMPAgentsParamsWithContext creates a new PatchAPI24SNMPAgentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24SNMPAgentsParamsWithContext(ctx context.Context) *PatchAPI24SNMPAgentsParams {
	var ()
	return &PatchAPI24SNMPAgentsParams{

		Context: ctx,
	}
}

// NewPatchAPI24SNMPAgentsParamsWithHTTPClient creates a new PatchAPI24SNMPAgentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24SNMPAgentsParamsWithHTTPClient(client *http.Client) *PatchAPI24SNMPAgentsParams {
	var ()
	return &PatchAPI24SNMPAgentsParams{
		HTTPClient: client,
	}
}

/*PatchAPI24SNMPAgentsParams contains all the parameters to send to the API endpoint
for the patch API 24 SNMP agents operation typically these are written to a http.Request
*/
type PatchAPI24SNMPAgentsParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*SNMPAgent*/
	SNMPAgent *models.SNMPAgentPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) WithTimeout(timeout time.Duration) *PatchAPI24SNMPAgentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) WithContext(ctx context.Context) *PatchAPI24SNMPAgentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) WithHTTPClient(client *http.Client) *PatchAPI24SNMPAgentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) WithAuthorization(authorization *string) *PatchAPI24SNMPAgentsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) WithXRequestID(xRequestID *string) *PatchAPI24SNMPAgentsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithSNMPAgent adds the sNMPAgent to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) WithSNMPAgent(sNMPAgent *models.SNMPAgentPatch) *PatchAPI24SNMPAgentsParams {
	o.SetSNMPAgent(sNMPAgent)
	return o
}

// SetSNMPAgent adds the snmpAgent to the patch API 24 SNMP agents params
func (o *PatchAPI24SNMPAgentsParams) SetSNMPAgent(sNMPAgent *models.SNMPAgentPatch) {
	o.SNMPAgent = sNMPAgent
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24SNMPAgentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SNMPAgent != nil {
		if err := r.SetBodyParam(o.SNMPAgent); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
