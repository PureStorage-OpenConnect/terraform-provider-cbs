// Code generated by go-swagger; DO NOT EDIT.

package protection_groups

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

// NewDeleteAPI24ProtectionGroupsVolumesParams creates a new DeleteAPI24ProtectionGroupsVolumesParams object
// with the default values initialized.
func NewDeleteAPI24ProtectionGroupsVolumesParams() *DeleteAPI24ProtectionGroupsVolumesParams {
	var ()
	return &DeleteAPI24ProtectionGroupsVolumesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24ProtectionGroupsVolumesParamsWithTimeout creates a new DeleteAPI24ProtectionGroupsVolumesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24ProtectionGroupsVolumesParamsWithTimeout(timeout time.Duration) *DeleteAPI24ProtectionGroupsVolumesParams {
	var ()
	return &DeleteAPI24ProtectionGroupsVolumesParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24ProtectionGroupsVolumesParamsWithContext creates a new DeleteAPI24ProtectionGroupsVolumesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24ProtectionGroupsVolumesParamsWithContext(ctx context.Context) *DeleteAPI24ProtectionGroupsVolumesParams {
	var ()
	return &DeleteAPI24ProtectionGroupsVolumesParams{

		Context: ctx,
	}
}

// NewDeleteAPI24ProtectionGroupsVolumesParamsWithHTTPClient creates a new DeleteAPI24ProtectionGroupsVolumesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24ProtectionGroupsVolumesParamsWithHTTPClient(client *http.Client) *DeleteAPI24ProtectionGroupsVolumesParams {
	var ()
	return &DeleteAPI24ProtectionGroupsVolumesParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24ProtectionGroupsVolumesParams contains all the parameters to send to the API endpoint
for the delete API 24 protection groups volumes operation typically these are written to a http.Request
*/
type DeleteAPI24ProtectionGroupsVolumesParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*GroupNames
	  Performs the operation on the unique group name specified. Examples of groups include host groups, pods, protection groups, and volume groups. Enter multiple names in comma-separated format. For example, `hgroup01,hgroup02`.

	*/
	GroupNames []string
	/*MemberNames
	  Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, `vol01,vol02`.

	*/
	MemberNames []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) WithTimeout(timeout time.Duration) *DeleteAPI24ProtectionGroupsVolumesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) WithContext(ctx context.Context) *DeleteAPI24ProtectionGroupsVolumesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) WithHTTPClient(client *http.Client) *DeleteAPI24ProtectionGroupsVolumesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) WithAuthorization(authorization *string) *DeleteAPI24ProtectionGroupsVolumesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) WithXRequestID(xRequestID *string) *DeleteAPI24ProtectionGroupsVolumesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithGroupNames adds the groupNames to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) WithGroupNames(groupNames []string) *DeleteAPI24ProtectionGroupsVolumesParams {
	o.SetGroupNames(groupNames)
	return o
}

// SetGroupNames adds the groupNames to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) SetGroupNames(groupNames []string) {
	o.GroupNames = groupNames
}

// WithMemberNames adds the memberNames to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) WithMemberNames(memberNames []string) *DeleteAPI24ProtectionGroupsVolumesParams {
	o.SetMemberNames(memberNames)
	return o
}

// SetMemberNames adds the memberNames to the delete API 24 protection groups volumes params
func (o *DeleteAPI24ProtectionGroupsVolumesParams) SetMemberNames(memberNames []string) {
	o.MemberNames = memberNames
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24ProtectionGroupsVolumesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesGroupNames := o.GroupNames

	joinedGroupNames := swag.JoinByFormat(valuesGroupNames, "csv")
	// query array param group_names
	if err := r.SetQueryParam("group_names", joinedGroupNames...); err != nil {
		return err
	}

	valuesMemberNames := o.MemberNames

	joinedMemberNames := swag.JoinByFormat(valuesMemberNames, "csv")
	// query array param member_names
	if err := r.SetQueryParam("member_names", joinedMemberNames...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
