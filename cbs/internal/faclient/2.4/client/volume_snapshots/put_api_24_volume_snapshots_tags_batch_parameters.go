// Code generated by go-swagger; DO NOT EDIT.

package volume_snapshots

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

// NewPutAPI24VolumeSnapshotsTagsBatchParams creates a new PutAPI24VolumeSnapshotsTagsBatchParams object
// with the default values initialized.
func NewPutAPI24VolumeSnapshotsTagsBatchParams() *PutAPI24VolumeSnapshotsTagsBatchParams {
	var ()
	return &PutAPI24VolumeSnapshotsTagsBatchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPI24VolumeSnapshotsTagsBatchParamsWithTimeout creates a new PutAPI24VolumeSnapshotsTagsBatchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPI24VolumeSnapshotsTagsBatchParamsWithTimeout(timeout time.Duration) *PutAPI24VolumeSnapshotsTagsBatchParams {
	var ()
	return &PutAPI24VolumeSnapshotsTagsBatchParams{

		timeout: timeout,
	}
}

// NewPutAPI24VolumeSnapshotsTagsBatchParamsWithContext creates a new PutAPI24VolumeSnapshotsTagsBatchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPI24VolumeSnapshotsTagsBatchParamsWithContext(ctx context.Context) *PutAPI24VolumeSnapshotsTagsBatchParams {
	var ()
	return &PutAPI24VolumeSnapshotsTagsBatchParams{

		Context: ctx,
	}
}

// NewPutAPI24VolumeSnapshotsTagsBatchParamsWithHTTPClient creates a new PutAPI24VolumeSnapshotsTagsBatchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPI24VolumeSnapshotsTagsBatchParamsWithHTTPClient(client *http.Client) *PutAPI24VolumeSnapshotsTagsBatchParams {
	var ()
	return &PutAPI24VolumeSnapshotsTagsBatchParams{
		HTTPClient: client,
	}
}

/*PutAPI24VolumeSnapshotsTagsBatchParams contains all the parameters to send to the API endpoint
for the put API 24 volume snapshots tags batch operation typically these are written to a http.Request
*/
type PutAPI24VolumeSnapshotsTagsBatchParams struct {

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
	  A list of tags to be created or modified.

	*/
	Tag []*models.Tag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WithTimeout(timeout time.Duration) *PutAPI24VolumeSnapshotsTagsBatchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WithContext(ctx context.Context) *PutAPI24VolumeSnapshotsTagsBatchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WithHTTPClient(client *http.Client) *PutAPI24VolumeSnapshotsTagsBatchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WithAuthorization(authorization *string) *PutAPI24VolumeSnapshotsTagsBatchParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WithXRequestID(xRequestID *string) *PutAPI24VolumeSnapshotsTagsBatchParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithResourceIds adds the resourceIds to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WithResourceIds(resourceIds []string) *PutAPI24VolumeSnapshotsTagsBatchParams {
	o.SetResourceIds(resourceIds)
	return o
}

// SetResourceIds adds the resourceIds to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) SetResourceIds(resourceIds []string) {
	o.ResourceIds = resourceIds
}

// WithResourceNames adds the resourceNames to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WithResourceNames(resourceNames []string) *PutAPI24VolumeSnapshotsTagsBatchParams {
	o.SetResourceNames(resourceNames)
	return o
}

// SetResourceNames adds the resourceNames to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) SetResourceNames(resourceNames []string) {
	o.ResourceNames = resourceNames
}

// WithTag adds the tag to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WithTag(tag []*models.Tag) *PutAPI24VolumeSnapshotsTagsBatchParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the put API 24 volume snapshots tags batch params
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) SetTag(tag []*models.Tag) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPI24VolumeSnapshotsTagsBatchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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