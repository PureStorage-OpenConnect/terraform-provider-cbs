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

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// NewPatchAPI24VolumeSnapshotsParams creates a new PatchAPI24VolumeSnapshotsParams object
// with the default values initialized.
func NewPatchAPI24VolumeSnapshotsParams() *PatchAPI24VolumeSnapshotsParams {
	var ()
	return &PatchAPI24VolumeSnapshotsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24VolumeSnapshotsParamsWithTimeout creates a new PatchAPI24VolumeSnapshotsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24VolumeSnapshotsParamsWithTimeout(timeout time.Duration) *PatchAPI24VolumeSnapshotsParams {
	var ()
	return &PatchAPI24VolumeSnapshotsParams{

		timeout: timeout,
	}
}

// NewPatchAPI24VolumeSnapshotsParamsWithContext creates a new PatchAPI24VolumeSnapshotsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24VolumeSnapshotsParamsWithContext(ctx context.Context) *PatchAPI24VolumeSnapshotsParams {
	var ()
	return &PatchAPI24VolumeSnapshotsParams{

		Context: ctx,
	}
}

// NewPatchAPI24VolumeSnapshotsParamsWithHTTPClient creates a new PatchAPI24VolumeSnapshotsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24VolumeSnapshotsParamsWithHTTPClient(client *http.Client) *PatchAPI24VolumeSnapshotsParams {
	var ()
	return &PatchAPI24VolumeSnapshotsParams{
		HTTPClient: client,
	}
}

/*PatchAPI24VolumeSnapshotsParams contains all the parameters to send to the API endpoint
for the patch API 24 volume snapshots operation typically these are written to a http.Request
*/
type PatchAPI24VolumeSnapshotsParams struct {

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
	/*Names
	  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, `name01,name02`.

	*/
	Names []string
	/*ReplicationSnapshot
	  If set to `true`, allow destruction/eradication of snapshots in use by replication. If set to `false`, allow destruction/eradication of snapshots not in use by replication. If not specified, defaults to `false`.

	*/
	ReplicationSnapshot *bool
	/*VolumeSnapshot*/
	VolumeSnapshot *models.VolumeSnapshotPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithTimeout(timeout time.Duration) *PatchAPI24VolumeSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithContext(ctx context.Context) *PatchAPI24VolumeSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithHTTPClient(client *http.Client) *PatchAPI24VolumeSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithAuthorization(authorization *string) *PatchAPI24VolumeSnapshotsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithXRequestID(xRequestID *string) *PatchAPI24VolumeSnapshotsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithIds adds the ids to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithIds(ids []string) *PatchAPI24VolumeSnapshotsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithNames adds the names to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithNames(names []string) *PatchAPI24VolumeSnapshotsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetNames(names []string) {
	o.Names = names
}

// WithReplicationSnapshot adds the replicationSnapshot to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithReplicationSnapshot(replicationSnapshot *bool) *PatchAPI24VolumeSnapshotsParams {
	o.SetReplicationSnapshot(replicationSnapshot)
	return o
}

// SetReplicationSnapshot adds the replicationSnapshot to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetReplicationSnapshot(replicationSnapshot *bool) {
	o.ReplicationSnapshot = replicationSnapshot
}

// WithVolumeSnapshot adds the volumeSnapshot to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) WithVolumeSnapshot(volumeSnapshot *models.VolumeSnapshotPatch) *PatchAPI24VolumeSnapshotsParams {
	o.SetVolumeSnapshot(volumeSnapshot)
	return o
}

// SetVolumeSnapshot adds the volumeSnapshot to the patch API 24 volume snapshots params
func (o *PatchAPI24VolumeSnapshotsParams) SetVolumeSnapshot(volumeSnapshot *models.VolumeSnapshotPatch) {
	o.VolumeSnapshot = volumeSnapshot
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24VolumeSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesNames := o.Names

	joinedNames := swag.JoinByFormat(valuesNames, "csv")
	// query array param names
	if err := r.SetQueryParam("names", joinedNames...); err != nil {
		return err
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

	if o.VolumeSnapshot != nil {
		if err := r.SetBodyParam(o.VolumeSnapshot); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
