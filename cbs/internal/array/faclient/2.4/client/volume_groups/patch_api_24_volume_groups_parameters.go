// Code generated by go-swagger; DO NOT EDIT.

package volume_groups

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

// NewPatchAPI24VolumeGroupsParams creates a new PatchAPI24VolumeGroupsParams object
// with the default values initialized.
func NewPatchAPI24VolumeGroupsParams() *PatchAPI24VolumeGroupsParams {
	var ()
	return &PatchAPI24VolumeGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24VolumeGroupsParamsWithTimeout creates a new PatchAPI24VolumeGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24VolumeGroupsParamsWithTimeout(timeout time.Duration) *PatchAPI24VolumeGroupsParams {
	var ()
	return &PatchAPI24VolumeGroupsParams{

		timeout: timeout,
	}
}

// NewPatchAPI24VolumeGroupsParamsWithContext creates a new PatchAPI24VolumeGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24VolumeGroupsParamsWithContext(ctx context.Context) *PatchAPI24VolumeGroupsParams {
	var ()
	return &PatchAPI24VolumeGroupsParams{

		Context: ctx,
	}
}

// NewPatchAPI24VolumeGroupsParamsWithHTTPClient creates a new PatchAPI24VolumeGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24VolumeGroupsParamsWithHTTPClient(client *http.Client) *PatchAPI24VolumeGroupsParams {
	var ()
	return &PatchAPI24VolumeGroupsParams{
		HTTPClient: client,
	}
}

/*PatchAPI24VolumeGroupsParams contains all the parameters to send to the API endpoint
for the patch API 24 volume groups operation typically these are written to a http.Request
*/
type PatchAPI24VolumeGroupsParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*DestroyContents
	  Set to `true` to destroy contents (e.g., volumes, protection groups, snapshots) and containers (e.g., pods, volume groups). This enables you to destroy containers with contents.

	*/
	DestroyContents *bool
	/*Ids
	  Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The `ids` and `names` parameters cannot be provided together.

	*/
	Ids []string
	/*Names
	  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, `name01,name02`.

	*/
	Names []string
	/*VolumeGroup*/
	VolumeGroup *models.VolumeGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithTimeout(timeout time.Duration) *PatchAPI24VolumeGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithContext(ctx context.Context) *PatchAPI24VolumeGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithHTTPClient(client *http.Client) *PatchAPI24VolumeGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithAuthorization(authorization *string) *PatchAPI24VolumeGroupsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithXRequestID(xRequestID *string) *PatchAPI24VolumeGroupsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithDestroyContents adds the destroyContents to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithDestroyContents(destroyContents *bool) *PatchAPI24VolumeGroupsParams {
	o.SetDestroyContents(destroyContents)
	return o
}

// SetDestroyContents adds the destroyContents to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetDestroyContents(destroyContents *bool) {
	o.DestroyContents = destroyContents
}

// WithIds adds the ids to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithIds(ids []string) *PatchAPI24VolumeGroupsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithNames adds the names to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithNames(names []string) *PatchAPI24VolumeGroupsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetNames(names []string) {
	o.Names = names
}

// WithVolumeGroup adds the volumeGroup to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) WithVolumeGroup(volumeGroup *models.VolumeGroup) *PatchAPI24VolumeGroupsParams {
	o.SetVolumeGroup(volumeGroup)
	return o
}

// SetVolumeGroup adds the volumeGroup to the patch API 24 volume groups params
func (o *PatchAPI24VolumeGroupsParams) SetVolumeGroup(volumeGroup *models.VolumeGroup) {
	o.VolumeGroup = volumeGroup
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24VolumeGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DestroyContents != nil {

		// query param destroy_contents
		var qrDestroyContents bool
		if o.DestroyContents != nil {
			qrDestroyContents = *o.DestroyContents
		}
		qDestroyContents := swag.FormatBool(qrDestroyContents)
		if qDestroyContents != "" {

			if err := r.SetQueryParam("destroy_contents", qDestroyContents); err != nil {
				return err
			}
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

	if o.VolumeGroup != nil {
		if err := r.SetBodyParam(o.VolumeGroup); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
