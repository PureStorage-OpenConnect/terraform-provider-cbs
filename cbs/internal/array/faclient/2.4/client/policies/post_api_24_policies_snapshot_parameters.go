// Code generated by go-swagger; DO NOT EDIT.

package policies

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

// NewPostAPI24PoliciesSnapshotParams creates a new PostAPI24PoliciesSnapshotParams object
// with the default values initialized.
func NewPostAPI24PoliciesSnapshotParams() *PostAPI24PoliciesSnapshotParams {
	var ()
	return &PostAPI24PoliciesSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24PoliciesSnapshotParamsWithTimeout creates a new PostAPI24PoliciesSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24PoliciesSnapshotParamsWithTimeout(timeout time.Duration) *PostAPI24PoliciesSnapshotParams {
	var ()
	return &PostAPI24PoliciesSnapshotParams{

		timeout: timeout,
	}
}

// NewPostAPI24PoliciesSnapshotParamsWithContext creates a new PostAPI24PoliciesSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24PoliciesSnapshotParamsWithContext(ctx context.Context) *PostAPI24PoliciesSnapshotParams {
	var ()
	return &PostAPI24PoliciesSnapshotParams{

		Context: ctx,
	}
}

// NewPostAPI24PoliciesSnapshotParamsWithHTTPClient creates a new PostAPI24PoliciesSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24PoliciesSnapshotParamsWithHTTPClient(client *http.Client) *PostAPI24PoliciesSnapshotParams {
	var ()
	return &PostAPI24PoliciesSnapshotParams{
		HTTPClient: client,
	}
}

/*PostAPI24PoliciesSnapshotParams contains all the parameters to send to the API endpoint
for the post API 24 policies snapshot operation typically these are written to a http.Request
*/
type PostAPI24PoliciesSnapshotParams struct {

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
	/*Policy*/
	Policy *models.PolicyPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) WithTimeout(timeout time.Duration) *PostAPI24PoliciesSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) WithContext(ctx context.Context) *PostAPI24PoliciesSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) WithHTTPClient(client *http.Client) *PostAPI24PoliciesSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) WithAuthorization(authorization *string) *PostAPI24PoliciesSnapshotParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) WithXRequestID(xRequestID *string) *PostAPI24PoliciesSnapshotParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNames adds the names to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) WithNames(names []string) *PostAPI24PoliciesSnapshotParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) SetNames(names []string) {
	o.Names = names
}

// WithPolicy adds the policy to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) WithPolicy(policy *models.PolicyPost) *PostAPI24PoliciesSnapshotParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the post API 24 policies snapshot params
func (o *PostAPI24PoliciesSnapshotParams) SetPolicy(policy *models.PolicyPost) {
	o.Policy = policy
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24PoliciesSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
