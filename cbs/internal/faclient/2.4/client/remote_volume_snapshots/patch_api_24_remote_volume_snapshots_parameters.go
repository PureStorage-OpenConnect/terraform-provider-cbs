// Code generated by go-swagger; DO NOT EDIT.

package remote_volume_snapshots

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

// NewPatchAPI24RemoteVolumeSnapshotsParams creates a new PatchAPI24RemoteVolumeSnapshotsParams object
// with the default values initialized.
func NewPatchAPI24RemoteVolumeSnapshotsParams() *PatchAPI24RemoteVolumeSnapshotsParams {
	var ()
	return &PatchAPI24RemoteVolumeSnapshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24RemoteVolumeSnapshotsParamsWithTimeout creates a new PatchAPI24RemoteVolumeSnapshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24RemoteVolumeSnapshotsParamsWithTimeout(timeout time.Duration) *PatchAPI24RemoteVolumeSnapshotsParams {
	var ()
	return &PatchAPI24RemoteVolumeSnapshotsParams{

		timeout: timeout,
	}
}

// NewPatchAPI24RemoteVolumeSnapshotsParamsWithContext creates a new PatchAPI24RemoteVolumeSnapshotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24RemoteVolumeSnapshotsParamsWithContext(ctx context.Context) *PatchAPI24RemoteVolumeSnapshotsParams {
	var ()
	return &PatchAPI24RemoteVolumeSnapshotsParams{

		Context: ctx,
	}
}

// NewPatchAPI24RemoteVolumeSnapshotsParamsWithHTTPClient creates a new PatchAPI24RemoteVolumeSnapshotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24RemoteVolumeSnapshotsParamsWithHTTPClient(client *http.Client) *PatchAPI24RemoteVolumeSnapshotsParams {
	var ()
	return &PatchAPI24RemoteVolumeSnapshotsParams{
		HTTPClient: client,
	}
}

/*PatchAPI24RemoteVolumeSnapshotsParams contains all the parameters to send to the API endpoint
for the patch API 24 remote volume snapshots operation typically these are written to a http.Request
*/
type PatchAPI24RemoteVolumeSnapshotsParams struct {

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
	/*On
	  Performs the operation on the target name specified. For example, `targetName01`.

	*/
	On *string
	/*RemoteVolumeSnapshot*/
	RemoteVolumeSnapshot *models.DestroyedPatchPost
	/*ReplicationSnapshot
	  If set to `true`, allow destruction/eradication of snapshots in use by replication. If set to `false`, allow destruction/eradication of snapshots not in use by replication. If not specified, defaults to `false`.

	*/
	ReplicationSnapshot *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithTimeout(timeout time.Duration) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithContext(ctx context.Context) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithHTTPClient(client *http.Client) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithAuthorization(authorization *string) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithXRequestID(xRequestID *string) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNames adds the names to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithNames(names []string) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetNames(names []string) {
	o.Names = names
}

// WithOn adds the on to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithOn(on *string) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetOn(on)
	return o
}

// SetOn adds the on to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetOn(on *string) {
	o.On = on
}

// WithRemoteVolumeSnapshot adds the remoteVolumeSnapshot to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithRemoteVolumeSnapshot(remoteVolumeSnapshot *models.DestroyedPatchPost) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetRemoteVolumeSnapshot(remoteVolumeSnapshot)
	return o
}

// SetRemoteVolumeSnapshot adds the remoteVolumeSnapshot to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetRemoteVolumeSnapshot(remoteVolumeSnapshot *models.DestroyedPatchPost) {
	o.RemoteVolumeSnapshot = remoteVolumeSnapshot
}

// WithReplicationSnapshot adds the replicationSnapshot to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WithReplicationSnapshot(replicationSnapshot *bool) *PatchAPI24RemoteVolumeSnapshotsParams {
	o.SetReplicationSnapshot(replicationSnapshot)
	return o
}

// SetReplicationSnapshot adds the replicationSnapshot to the patch API 24 remote volume snapshots params
func (o *PatchAPI24RemoteVolumeSnapshotsParams) SetReplicationSnapshot(replicationSnapshot *bool) {
	o.ReplicationSnapshot = replicationSnapshot
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24RemoteVolumeSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.On != nil {

		// query param on
		var qrOn string
		if o.On != nil {
			qrOn = *o.On
		}
		qOn := qrOn
		if qOn != "" {

			if err := r.SetQueryParam("on", qOn); err != nil {
				return err
			}
		}

	}

	if o.RemoteVolumeSnapshot != nil {
		if err := r.SetBodyParam(o.RemoteVolumeSnapshot); err != nil {
			return err
		}
	}

	if o.ReplicationSnapshot != nil {

		// query param replication_snapshot
		var qrReplicationSnapshot bool
		if o.ReplicationSnapshot != nil {
			qrReplicationSnapshot = *o.ReplicationSnapshot
		}
		qReplicationSnapshot := swag.FormatBool(qrReplicationSnapshot)
		if qReplicationSnapshot != "" {

			if err := r.SetQueryParam("replication_snapshot", qReplicationSnapshot); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}