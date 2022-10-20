// Code generated by go-swagger; DO NOT EDIT.

package directories

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

// NewPostAPI24DirectoriesPoliciesSnapshotParams creates a new PostAPI24DirectoriesPoliciesSnapshotParams object
// with the default values initialized.
func NewPostAPI24DirectoriesPoliciesSnapshotParams() *PostAPI24DirectoriesPoliciesSnapshotParams {
	var ()
	return &PostAPI24DirectoriesPoliciesSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24DirectoriesPoliciesSnapshotParamsWithTimeout creates a new PostAPI24DirectoriesPoliciesSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24DirectoriesPoliciesSnapshotParamsWithTimeout(timeout time.Duration) *PostAPI24DirectoriesPoliciesSnapshotParams {
	var ()
	return &PostAPI24DirectoriesPoliciesSnapshotParams{

		timeout: timeout,
	}
}

// NewPostAPI24DirectoriesPoliciesSnapshotParamsWithContext creates a new PostAPI24DirectoriesPoliciesSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24DirectoriesPoliciesSnapshotParamsWithContext(ctx context.Context) *PostAPI24DirectoriesPoliciesSnapshotParams {
	var ()
	return &PostAPI24DirectoriesPoliciesSnapshotParams{

		Context: ctx,
	}
}

// NewPostAPI24DirectoriesPoliciesSnapshotParamsWithHTTPClient creates a new PostAPI24DirectoriesPoliciesSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24DirectoriesPoliciesSnapshotParamsWithHTTPClient(client *http.Client) *PostAPI24DirectoriesPoliciesSnapshotParams {
	var ()
	return &PostAPI24DirectoriesPoliciesSnapshotParams{
		HTTPClient: client,
	}
}

/*PostAPI24DirectoriesPoliciesSnapshotParams contains all the parameters to send to the API endpoint
for the post API 24 directories policies snapshot operation typically these are written to a http.Request
*/
type PostAPI24DirectoriesPoliciesSnapshotParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*MemberIds
	  Performs the operation on the unique member IDs specified. Enter multiple member IDs in comma-separated format. The `member_ids` and `member_names` parameters cannot be provided together.

	*/
	MemberIds []string
	/*MemberNames
	  Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, `vol01,vol02`.

	*/
	MemberNames []string
	/*Policies*/
	Policies *models.DirectoryPolicyPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WithTimeout(timeout time.Duration) *PostAPI24DirectoriesPoliciesSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WithContext(ctx context.Context) *PostAPI24DirectoriesPoliciesSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WithHTTPClient(client *http.Client) *PostAPI24DirectoriesPoliciesSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WithAuthorization(authorization *string) *PostAPI24DirectoriesPoliciesSnapshotParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WithXRequestID(xRequestID *string) *PostAPI24DirectoriesPoliciesSnapshotParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithMemberIds adds the memberIds to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WithMemberIds(memberIds []string) *PostAPI24DirectoriesPoliciesSnapshotParams {
	o.SetMemberIds(memberIds)
	return o
}

// SetMemberIds adds the memberIds to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) SetMemberIds(memberIds []string) {
	o.MemberIds = memberIds
}

// WithMemberNames adds the memberNames to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WithMemberNames(memberNames []string) *PostAPI24DirectoriesPoliciesSnapshotParams {
	o.SetMemberNames(memberNames)
	return o
}

// SetMemberNames adds the memberNames to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) SetMemberNames(memberNames []string) {
	o.MemberNames = memberNames
}

// WithPolicies adds the policies to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WithPolicies(policies *models.DirectoryPolicyPost) *PostAPI24DirectoriesPoliciesSnapshotParams {
	o.SetPolicies(policies)
	return o
}

// SetPolicies adds the policies to the post API 24 directories policies snapshot params
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) SetPolicies(policies *models.DirectoryPolicyPost) {
	o.Policies = policies
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24DirectoriesPoliciesSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesMemberIds := o.MemberIds

	joinedMemberIds := swag.JoinByFormat(valuesMemberIds, "csv")
	// query array param member_ids
	if err := r.SetQueryParam("member_ids", joinedMemberIds...); err != nil {
		return err
	}

	valuesMemberNames := o.MemberNames

	joinedMemberNames := swag.JoinByFormat(valuesMemberNames, "csv")
	// query array param member_names
	if err := r.SetQueryParam("member_names", joinedMemberNames...); err != nil {
		return err
	}

	if o.Policies != nil {
		if err := r.SetBodyParam(o.Policies); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
