// Code generated by go-swagger; DO NOT EDIT.

package subnets

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

// NewPatchAPI24SubnetsParams creates a new PatchAPI24SubnetsParams object
// with the default values initialized.
func NewPatchAPI24SubnetsParams() *PatchAPI24SubnetsParams {
	var ()
	return &PatchAPI24SubnetsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24SubnetsParamsWithTimeout creates a new PatchAPI24SubnetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24SubnetsParamsWithTimeout(timeout time.Duration) *PatchAPI24SubnetsParams {
	var ()
	return &PatchAPI24SubnetsParams{

		timeout: timeout,
	}
}

// NewPatchAPI24SubnetsParamsWithContext creates a new PatchAPI24SubnetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24SubnetsParamsWithContext(ctx context.Context) *PatchAPI24SubnetsParams {
	var ()
	return &PatchAPI24SubnetsParams{

		Context: ctx,
	}
}

// NewPatchAPI24SubnetsParamsWithHTTPClient creates a new PatchAPI24SubnetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24SubnetsParamsWithHTTPClient(client *http.Client) *PatchAPI24SubnetsParams {
	var ()
	return &PatchAPI24SubnetsParams{
		HTTPClient: client,
	}
}

/*PatchAPI24SubnetsParams contains all the parameters to send to the API endpoint
for the patch API 24 subnets operation typically these are written to a http.Request
*/
type PatchAPI24SubnetsParams struct {

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
	/*Subnet*/
	Subnet *models.SubnetPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) WithTimeout(timeout time.Duration) *PatchAPI24SubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) WithContext(ctx context.Context) *PatchAPI24SubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) WithHTTPClient(client *http.Client) *PatchAPI24SubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) WithAuthorization(authorization *string) *PatchAPI24SubnetsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) WithXRequestID(xRequestID *string) *PatchAPI24SubnetsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNames adds the names to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) WithNames(names []string) *PatchAPI24SubnetsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) SetNames(names []string) {
	o.Names = names
}

// WithSubnet adds the subnet to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) WithSubnet(subnet *models.SubnetPatch) *PatchAPI24SubnetsParams {
	o.SetSubnet(subnet)
	return o
}

// SetSubnet adds the subnet to the patch API 24 subnets params
func (o *PatchAPI24SubnetsParams) SetSubnet(subnet *models.SubnetPatch) {
	o.Subnet = subnet
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24SubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Subnet != nil {
		if err := r.SetBodyParam(o.Subnet); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
