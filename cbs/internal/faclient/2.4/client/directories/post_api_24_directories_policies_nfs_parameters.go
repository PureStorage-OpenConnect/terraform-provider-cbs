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

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// NewPostAPI24DirectoriesPoliciesNfsParams creates a new PostAPI24DirectoriesPoliciesNfsParams object
// with the default values initialized.
func NewPostAPI24DirectoriesPoliciesNfsParams() *PostAPI24DirectoriesPoliciesNfsParams {
	var ()
	return &PostAPI24DirectoriesPoliciesNfsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24DirectoriesPoliciesNfsParamsWithTimeout creates a new PostAPI24DirectoriesPoliciesNfsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24DirectoriesPoliciesNfsParamsWithTimeout(timeout time.Duration) *PostAPI24DirectoriesPoliciesNfsParams {
	var ()
	return &PostAPI24DirectoriesPoliciesNfsParams{

		timeout: timeout,
	}
}

// NewPostAPI24DirectoriesPoliciesNfsParamsWithContext creates a new PostAPI24DirectoriesPoliciesNfsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24DirectoriesPoliciesNfsParamsWithContext(ctx context.Context) *PostAPI24DirectoriesPoliciesNfsParams {
	var ()
	return &PostAPI24DirectoriesPoliciesNfsParams{

		Context: ctx,
	}
}

// NewPostAPI24DirectoriesPoliciesNfsParamsWithHTTPClient creates a new PostAPI24DirectoriesPoliciesNfsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24DirectoriesPoliciesNfsParamsWithHTTPClient(client *http.Client) *PostAPI24DirectoriesPoliciesNfsParams {
	var ()
	return &PostAPI24DirectoriesPoliciesNfsParams{
		HTTPClient: client,
	}
}

/*PostAPI24DirectoriesPoliciesNfsParams contains all the parameters to send to the API endpoint
for the post API 24 directories policies nfs operation typically these are written to a http.Request
*/
type PostAPI24DirectoriesPoliciesNfsParams struct {

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
	Policies *models.DirectoryPolicyExportPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) WithTimeout(timeout time.Duration) *PostAPI24DirectoriesPoliciesNfsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) WithContext(ctx context.Context) *PostAPI24DirectoriesPoliciesNfsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) WithHTTPClient(client *http.Client) *PostAPI24DirectoriesPoliciesNfsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) WithAuthorization(authorization *string) *PostAPI24DirectoriesPoliciesNfsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) WithXRequestID(xRequestID *string) *PostAPI24DirectoriesPoliciesNfsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithMemberIds adds the memberIds to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) WithMemberIds(memberIds []string) *PostAPI24DirectoriesPoliciesNfsParams {
	o.SetMemberIds(memberIds)
	return o
}

// SetMemberIds adds the memberIds to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) SetMemberIds(memberIds []string) {
	o.MemberIds = memberIds
}

// WithMemberNames adds the memberNames to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) WithMemberNames(memberNames []string) *PostAPI24DirectoriesPoliciesNfsParams {
	o.SetMemberNames(memberNames)
	return o
}

// SetMemberNames adds the memberNames to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) SetMemberNames(memberNames []string) {
	o.MemberNames = memberNames
}

// WithPolicies adds the policies to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) WithPolicies(policies *models.DirectoryPolicyExportPost) *PostAPI24DirectoriesPoliciesNfsParams {
	o.SetPolicies(policies)
	return o
}

// SetPolicies adds the policies to the post API 24 directories policies nfs params
func (o *PostAPI24DirectoriesPoliciesNfsParams) SetPolicies(policies *models.DirectoryPolicyExportPost) {
	o.Policies = policies
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24DirectoriesPoliciesNfsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
