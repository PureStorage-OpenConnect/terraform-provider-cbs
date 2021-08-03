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

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// NewPostAPI24PoliciesSmbClientRulesParams creates a new PostAPI24PoliciesSmbClientRulesParams object
// with the default values initialized.
func NewPostAPI24PoliciesSmbClientRulesParams() *PostAPI24PoliciesSmbClientRulesParams {
	var ()
	return &PostAPI24PoliciesSmbClientRulesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24PoliciesSmbClientRulesParamsWithTimeout creates a new PostAPI24PoliciesSmbClientRulesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24PoliciesSmbClientRulesParamsWithTimeout(timeout time.Duration) *PostAPI24PoliciesSmbClientRulesParams {
	var ()
	return &PostAPI24PoliciesSmbClientRulesParams{

		timeout: timeout,
	}
}

// NewPostAPI24PoliciesSmbClientRulesParamsWithContext creates a new PostAPI24PoliciesSmbClientRulesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24PoliciesSmbClientRulesParamsWithContext(ctx context.Context) *PostAPI24PoliciesSmbClientRulesParams {
	var ()
	return &PostAPI24PoliciesSmbClientRulesParams{

		Context: ctx,
	}
}

// NewPostAPI24PoliciesSmbClientRulesParamsWithHTTPClient creates a new PostAPI24PoliciesSmbClientRulesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24PoliciesSmbClientRulesParamsWithHTTPClient(client *http.Client) *PostAPI24PoliciesSmbClientRulesParams {
	var ()
	return &PostAPI24PoliciesSmbClientRulesParams{
		HTTPClient: client,
	}
}

/*PostAPI24PoliciesSmbClientRulesParams contains all the parameters to send to the API endpoint
for the post API 24 policies smb client rules operation typically these are written to a http.Request
*/
type PostAPI24PoliciesSmbClientRulesParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*PolicyIds
	  Performs the operation on the unique policy IDs specified. Enter multiple policy IDs in comma-separated format. The `policy_ids` and `policy_names` parameters cannot be provided together.

	*/
	PolicyIds []string
	/*PolicyNames
	  Performs the operation on the policy names specified. Enter multiple policy names in comma-separated format. For example, `name01,name02`.

	*/
	PolicyNames []string
	/*Rules*/
	Rules *models.PolicyRuleSmbClientPost

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) WithTimeout(timeout time.Duration) *PostAPI24PoliciesSmbClientRulesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) WithContext(ctx context.Context) *PostAPI24PoliciesSmbClientRulesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) WithHTTPClient(client *http.Client) *PostAPI24PoliciesSmbClientRulesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) WithAuthorization(authorization *string) *PostAPI24PoliciesSmbClientRulesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) WithXRequestID(xRequestID *string) *PostAPI24PoliciesSmbClientRulesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithPolicyIds adds the policyIds to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) WithPolicyIds(policyIds []string) *PostAPI24PoliciesSmbClientRulesParams {
	o.SetPolicyIds(policyIds)
	return o
}

// SetPolicyIds adds the policyIds to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) SetPolicyIds(policyIds []string) {
	o.PolicyIds = policyIds
}

// WithPolicyNames adds the policyNames to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) WithPolicyNames(policyNames []string) *PostAPI24PoliciesSmbClientRulesParams {
	o.SetPolicyNames(policyNames)
	return o
}

// SetPolicyNames adds the policyNames to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) SetPolicyNames(policyNames []string) {
	o.PolicyNames = policyNames
}

// WithRules adds the rules to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) WithRules(rules *models.PolicyRuleSmbClientPost) *PostAPI24PoliciesSmbClientRulesParams {
	o.SetRules(rules)
	return o
}

// SetRules adds the rules to the post API 24 policies smb client rules params
func (o *PostAPI24PoliciesSmbClientRulesParams) SetRules(rules *models.PolicyRuleSmbClientPost) {
	o.Rules = rules
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24PoliciesSmbClientRulesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesPolicyIds := o.PolicyIds

	joinedPolicyIds := swag.JoinByFormat(valuesPolicyIds, "csv")
	// query array param policy_ids
	if err := r.SetQueryParam("policy_ids", joinedPolicyIds...); err != nil {
		return err
	}

	valuesPolicyNames := o.PolicyNames

	joinedPolicyNames := swag.JoinByFormat(valuesPolicyNames, "csv")
	// query array param policy_names
	if err := r.SetQueryParam("policy_names", joinedPolicyNames...); err != nil {
		return err
	}

	if o.Rules != nil {
		if err := r.SetBodyParam(o.Rules); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}