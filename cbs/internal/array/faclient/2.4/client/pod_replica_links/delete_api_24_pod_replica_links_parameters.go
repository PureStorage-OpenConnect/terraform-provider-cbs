// Code generated by go-swagger; DO NOT EDIT.

package pod_replica_links

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

// NewDeleteAPI24PodReplicaLinksParams creates a new DeleteAPI24PodReplicaLinksParams object
// with the default values initialized.
func NewDeleteAPI24PodReplicaLinksParams() *DeleteAPI24PodReplicaLinksParams {
	var ()
	return &DeleteAPI24PodReplicaLinksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24PodReplicaLinksParamsWithTimeout creates a new DeleteAPI24PodReplicaLinksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24PodReplicaLinksParamsWithTimeout(timeout time.Duration) *DeleteAPI24PodReplicaLinksParams {
	var ()
	return &DeleteAPI24PodReplicaLinksParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24PodReplicaLinksParamsWithContext creates a new DeleteAPI24PodReplicaLinksParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24PodReplicaLinksParamsWithContext(ctx context.Context) *DeleteAPI24PodReplicaLinksParams {
	var ()
	return &DeleteAPI24PodReplicaLinksParams{

		Context: ctx,
	}
}

// NewDeleteAPI24PodReplicaLinksParamsWithHTTPClient creates a new DeleteAPI24PodReplicaLinksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24PodReplicaLinksParamsWithHTTPClient(client *http.Client) *DeleteAPI24PodReplicaLinksParams {
	var ()
	return &DeleteAPI24PodReplicaLinksParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24PodReplicaLinksParams contains all the parameters to send to the API endpoint
for the delete API 24 pod replica links operation typically these are written to a http.Request
*/
type DeleteAPI24PodReplicaLinksParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*Ids
	  Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The `ids` and `names` parameters cannot be provided together.

	*/
	Ids []string
	/*LocalPodIds
	  A comma-separated list of local pod IDs. If, after filtering, there is not at least one resource that matches each of the elements, then an error is returned. This cannot be provided together with the `local_pod_names` query parameter.

	*/
	LocalPodIds []string
	/*LocalPodNames
	  A comma-separated list of local pod names. If, after filtering, there is not at least one resource that matches each of the elements, then an error is returned. This cannot be provided together with the `local_pod_ids` query parameter.

	*/
	LocalPodNames []string
	/*RemotePodIds
	  A comma-separated list of remote pod IDs. If, after filtering, there is not at least one resource that matches each of the elements, then an error is returned. This cannot be provided together with the `remote_pod_names` query parameter.

	*/
	RemotePodIds []string
	/*RemotePodNames
	  A comma-separated list of remote pod names. If, after filtering, there is not at least one resource that matches each of the elements, then an error is returned. This cannot be provided together with the `remote_pod_ids` query parameter.

	*/
	RemotePodNames []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithTimeout(timeout time.Duration) *DeleteAPI24PodReplicaLinksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithContext(ctx context.Context) *DeleteAPI24PodReplicaLinksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithHTTPClient(client *http.Client) *DeleteAPI24PodReplicaLinksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithAuthorization(authorization *string) *DeleteAPI24PodReplicaLinksParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithXRequestID(xRequestID *string) *DeleteAPI24PodReplicaLinksParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithIds adds the ids to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithIds(ids []string) *DeleteAPI24PodReplicaLinksParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithLocalPodIds adds the localPodIds to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithLocalPodIds(localPodIds []string) *DeleteAPI24PodReplicaLinksParams {
	o.SetLocalPodIds(localPodIds)
	return o
}

// SetLocalPodIds adds the localPodIds to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetLocalPodIds(localPodIds []string) {
	o.LocalPodIds = localPodIds
}

// WithLocalPodNames adds the localPodNames to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithLocalPodNames(localPodNames []string) *DeleteAPI24PodReplicaLinksParams {
	o.SetLocalPodNames(localPodNames)
	return o
}

// SetLocalPodNames adds the localPodNames to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetLocalPodNames(localPodNames []string) {
	o.LocalPodNames = localPodNames
}

// WithRemotePodIds adds the remotePodIds to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithRemotePodIds(remotePodIds []string) *DeleteAPI24PodReplicaLinksParams {
	o.SetRemotePodIds(remotePodIds)
	return o
}

// SetRemotePodIds adds the remotePodIds to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetRemotePodIds(remotePodIds []string) {
	o.RemotePodIds = remotePodIds
}

// WithRemotePodNames adds the remotePodNames to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) WithRemotePodNames(remotePodNames []string) *DeleteAPI24PodReplicaLinksParams {
	o.SetRemotePodNames(remotePodNames)
	return o
}

// SetRemotePodNames adds the remotePodNames to the delete API 24 pod replica links params
func (o *DeleteAPI24PodReplicaLinksParams) SetRemotePodNames(remotePodNames []string) {
	o.RemotePodNames = remotePodNames
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24PodReplicaLinksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "csv")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
	}

	valuesLocalPodIds := o.LocalPodIds

	joinedLocalPodIds := swag.JoinByFormat(valuesLocalPodIds, "csv")
	// query array param local_pod_ids
	if err := r.SetQueryParam("local_pod_ids", joinedLocalPodIds...); err != nil {
		return err
	}

	valuesLocalPodNames := o.LocalPodNames

	joinedLocalPodNames := swag.JoinByFormat(valuesLocalPodNames, "csv")
	// query array param local_pod_names
	if err := r.SetQueryParam("local_pod_names", joinedLocalPodNames...); err != nil {
		return err
	}

	valuesRemotePodIds := o.RemotePodIds

	joinedRemotePodIds := swag.JoinByFormat(valuesRemotePodIds, "csv")
	// query array param remote_pod_ids
	if err := r.SetQueryParam("remote_pod_ids", joinedRemotePodIds...); err != nil {
		return err
	}

	valuesRemotePodNames := o.RemotePodNames

	joinedRemotePodNames := swag.JoinByFormat(valuesRemotePodNames, "csv")
	// query array param remote_pod_names
	if err := r.SetQueryParam("remote_pod_names", joinedRemotePodNames...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
