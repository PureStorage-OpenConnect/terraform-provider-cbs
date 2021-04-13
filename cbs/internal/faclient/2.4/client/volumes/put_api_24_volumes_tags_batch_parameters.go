// Code generated by go-swagger; DO NOT EDIT.

package volumes

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

// NewPutAPI24VolumesTagsBatchParams creates a new PutAPI24VolumesTagsBatchParams object
// with the default values initialized.
func NewPutAPI24VolumesTagsBatchParams() *PutAPI24VolumesTagsBatchParams {
	var ()
	return &PutAPI24VolumesTagsBatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPI24VolumesTagsBatchParamsWithTimeout creates a new PutAPI24VolumesTagsBatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPI24VolumesTagsBatchParamsWithTimeout(timeout time.Duration) *PutAPI24VolumesTagsBatchParams {
	var ()
	return &PutAPI24VolumesTagsBatchParams{

		timeout: timeout,
	}
}

// NewPutAPI24VolumesTagsBatchParamsWithContext creates a new PutAPI24VolumesTagsBatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPI24VolumesTagsBatchParamsWithContext(ctx context.Context) *PutAPI24VolumesTagsBatchParams {
	var ()
	return &PutAPI24VolumesTagsBatchParams{

		Context: ctx,
	}
}

// NewPutAPI24VolumesTagsBatchParamsWithHTTPClient creates a new PutAPI24VolumesTagsBatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPI24VolumesTagsBatchParamsWithHTTPClient(client *http.Client) *PutAPI24VolumesTagsBatchParams {
	var ()
	return &PutAPI24VolumesTagsBatchParams{
		HTTPClient: client,
	}
}

/*PutAPI24VolumesTagsBatchParams contains all the parameters to send to the API endpoint
for the put API 24 volumes tags batch operation typically these are written to a http.Request
*/
type PutAPI24VolumesTagsBatchParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*ResourceIds
	  A comma-separated list of resource IDs. The `resource_ids` and `resource_names` parameters cannot be provided together.

	*/
	ResourceIds []string
	/*ResourceNames
	  A comma-separated list of resource names. The `resource_ids` and `resource_names` parameters cannot be provided together.

	*/
	ResourceNames []string
	/*Tag
	  A list of tags to be created or, if already existing, updated.

	*/
	Tag []*models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) WithTimeout(timeout time.Duration) *PutAPI24VolumesTagsBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) WithContext(ctx context.Context) *PutAPI24VolumesTagsBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) WithHTTPClient(client *http.Client) *PutAPI24VolumesTagsBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) WithAuthorization(authorization *string) *PutAPI24VolumesTagsBatchParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) WithXRequestID(xRequestID *string) *PutAPI24VolumesTagsBatchParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithResourceIds adds the resourceIds to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) WithResourceIds(resourceIds []string) *PutAPI24VolumesTagsBatchParams {
	o.SetResourceIds(resourceIds)
	return o
}

// SetResourceIds adds the resourceIds to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) SetResourceIds(resourceIds []string) {
	o.ResourceIds = resourceIds
}

// WithResourceNames adds the resourceNames to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) WithResourceNames(resourceNames []string) *PutAPI24VolumesTagsBatchParams {
	o.SetResourceNames(resourceNames)
	return o
}

// SetResourceNames adds the resourceNames to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) SetResourceNames(resourceNames []string) {
	o.ResourceNames = resourceNames
}

// WithTag adds the tag to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) WithTag(tag []*models.Tag) *PutAPI24VolumesTagsBatchParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the put API 24 volumes tags batch params
func (o *PutAPI24VolumesTagsBatchParams) SetTag(tag []*models.Tag) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPI24VolumesTagsBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesResourceIds := o.ResourceIds

	joinedResourceIds := swag.JoinByFormat(valuesResourceIds, "csv")
	// query array param resource_ids
	if err := r.SetQueryParam("resource_ids", joinedResourceIds...); err != nil {
		return err
	}

	valuesResourceNames := o.ResourceNames

	joinedResourceNames := swag.JoinByFormat(valuesResourceNames, "csv")
	// query array param resource_names
	if err := r.SetQueryParam("resource_names", joinedResourceNames...); err != nil {
		return err
	}

	if o.Tag != nil {
		if err := r.SetBodyParam(o.Tag); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
