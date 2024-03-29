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
)

// NewDeleteAPI24VolumeSnapshotsTagsParams creates a new DeleteAPI24VolumeSnapshotsTagsParams object
// with the default values initialized.
func NewDeleteAPI24VolumeSnapshotsTagsParams() *DeleteAPI24VolumeSnapshotsTagsParams {
	var ()
	return &DeleteAPI24VolumeSnapshotsTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24VolumeSnapshotsTagsParamsWithTimeout creates a new DeleteAPI24VolumeSnapshotsTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24VolumeSnapshotsTagsParamsWithTimeout(timeout time.Duration) *DeleteAPI24VolumeSnapshotsTagsParams {
	var ()
	return &DeleteAPI24VolumeSnapshotsTagsParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24VolumeSnapshotsTagsParamsWithContext creates a new DeleteAPI24VolumeSnapshotsTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24VolumeSnapshotsTagsParamsWithContext(ctx context.Context) *DeleteAPI24VolumeSnapshotsTagsParams {
	var ()
	return &DeleteAPI24VolumeSnapshotsTagsParams{

		Context: ctx,
	}
}

// NewDeleteAPI24VolumeSnapshotsTagsParamsWithHTTPClient creates a new DeleteAPI24VolumeSnapshotsTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24VolumeSnapshotsTagsParamsWithHTTPClient(client *http.Client) *DeleteAPI24VolumeSnapshotsTagsParams {
	var ()
	return &DeleteAPI24VolumeSnapshotsTagsParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24VolumeSnapshotsTagsParams contains all the parameters to send to the API endpoint
for the delete API 24 volume snapshots tags operation typically these are written to a http.Request
*/
type DeleteAPI24VolumeSnapshotsTagsParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*Keys
	  A comma-separated list of tag keys.

	*/
	Keys []string
	/*Namespaces
	  A comma-separated list of namespaces.

	*/
	Namespaces []string
	/*ResourceIds
	  A comma-separated list of resource IDs. The `resource_ids` and `resource_names` parameters cannot be provided together.

	*/
	ResourceIds []string
	/*ResourceNames
	  A comma-separated list of resource names. The `resource_ids` and `resource_names` parameters cannot be provided together.

	*/
	ResourceNames []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithTimeout(timeout time.Duration) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithContext(ctx context.Context) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithHTTPClient(client *http.Client) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithAuthorization(authorization *string) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithXRequestID(xRequestID *string) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithKeys adds the keys to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithKeys(keys []string) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithNamespaces adds the namespaces to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithNamespaces(namespaces []string) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetNamespaces(namespaces)
	return o
}

// SetNamespaces adds the namespaces to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetNamespaces(namespaces []string) {
	o.Namespaces = namespaces
}

// WithResourceIds adds the resourceIds to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithResourceIds(resourceIds []string) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetResourceIds(resourceIds)
	return o
}

// SetResourceIds adds the resourceIds to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetResourceIds(resourceIds []string) {
	o.ResourceIds = resourceIds
}

// WithResourceNames adds the resourceNames to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WithResourceNames(resourceNames []string) *DeleteAPI24VolumeSnapshotsTagsParams {
	o.SetResourceNames(resourceNames)
	return o
}

// SetResourceNames adds the resourceNames to the delete API 24 volume snapshots tags params
func (o *DeleteAPI24VolumeSnapshotsTagsParams) SetResourceNames(resourceNames []string) {
	o.ResourceNames = resourceNames
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24VolumeSnapshotsTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "csv")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	valuesNamespaces := o.Namespaces

	joinedNamespaces := swag.JoinByFormat(valuesNamespaces, "csv")
	// query array param namespaces
	if err := r.SetQueryParam("namespaces", joinedNamespaces...); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
