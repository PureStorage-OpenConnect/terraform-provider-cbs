// Code generated by go-swagger; DO NOT EDIT.

package directory_exports

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

// NewDeleteAPI24DirectoryExportsParams creates a new DeleteAPI24DirectoryExportsParams object
// with the default values initialized.
func NewDeleteAPI24DirectoryExportsParams() *DeleteAPI24DirectoryExportsParams {
	var ()
	return &DeleteAPI24DirectoryExportsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24DirectoryExportsParamsWithTimeout creates a new DeleteAPI24DirectoryExportsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24DirectoryExportsParamsWithTimeout(timeout time.Duration) *DeleteAPI24DirectoryExportsParams {
	var ()
	return &DeleteAPI24DirectoryExportsParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24DirectoryExportsParamsWithContext creates a new DeleteAPI24DirectoryExportsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24DirectoryExportsParamsWithContext(ctx context.Context) *DeleteAPI24DirectoryExportsParams {
	var ()
	return &DeleteAPI24DirectoryExportsParams{

		Context: ctx,
	}
}

// NewDeleteAPI24DirectoryExportsParamsWithHTTPClient creates a new DeleteAPI24DirectoryExportsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24DirectoryExportsParamsWithHTTPClient(client *http.Client) *DeleteAPI24DirectoryExportsParams {
	var ()
	return &DeleteAPI24DirectoryExportsParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24DirectoryExportsParams contains all the parameters to send to the API endpoint
for the delete API 24 directory exports operation typically these are written to a http.Request
*/
type DeleteAPI24DirectoryExportsParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*ExportNames
	  Performs the operation on the export names specified. Enter multiple names in comma-separated format. For example, `name01,name02`.

	*/
	ExportNames []string
	/*PolicyIds
	  Performs the operation on the unique policy IDs specified. Enter multiple policy IDs in comma-separated format. The `policy_ids` and `policy_names` parameters cannot be provided together.

	*/
	PolicyIds []string
	/*PolicyNames
	  Performs the operation on the policy names specified. Enter multiple policy names in comma-separated format. For example, `name01,name02`.

	*/
	PolicyNames []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) WithTimeout(timeout time.Duration) *DeleteAPI24DirectoryExportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) WithContext(ctx context.Context) *DeleteAPI24DirectoryExportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) WithHTTPClient(client *http.Client) *DeleteAPI24DirectoryExportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) WithAuthorization(authorization *string) *DeleteAPI24DirectoryExportsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) WithXRequestID(xRequestID *string) *DeleteAPI24DirectoryExportsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithExportNames adds the exportNames to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) WithExportNames(exportNames []string) *DeleteAPI24DirectoryExportsParams {
	o.SetExportNames(exportNames)
	return o
}

// SetExportNames adds the exportNames to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) SetExportNames(exportNames []string) {
	o.ExportNames = exportNames
}

// WithPolicyIds adds the policyIds to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) WithPolicyIds(policyIds []string) *DeleteAPI24DirectoryExportsParams {
	o.SetPolicyIds(policyIds)
	return o
}

// SetPolicyIds adds the policyIds to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) SetPolicyIds(policyIds []string) {
	o.PolicyIds = policyIds
}

// WithPolicyNames adds the policyNames to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) WithPolicyNames(policyNames []string) *DeleteAPI24DirectoryExportsParams {
	o.SetPolicyNames(policyNames)
	return o
}

// SetPolicyNames adds the policyNames to the delete API 24 directory exports params
func (o *DeleteAPI24DirectoryExportsParams) SetPolicyNames(policyNames []string) {
	o.PolicyNames = policyNames
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24DirectoryExportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesExportNames := o.ExportNames

	joinedExportNames := swag.JoinByFormat(valuesExportNames, "csv")
	// query array param export_names
	if err := r.SetQueryParam("export_names", joinedExportNames...); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
