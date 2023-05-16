// Code generated by go-swagger; DO NOT EDIT.

package directory_services

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

// NewPatchAPI24DirectoryServicesRolesParams creates a new PatchAPI24DirectoryServicesRolesParams object
// with the default values initialized.
func NewPatchAPI24DirectoryServicesRolesParams() *PatchAPI24DirectoryServicesRolesParams {
	var ()
	return &PatchAPI24DirectoryServicesRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPI24DirectoryServicesRolesParamsWithTimeout creates a new PatchAPI24DirectoryServicesRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPI24DirectoryServicesRolesParamsWithTimeout(timeout time.Duration) *PatchAPI24DirectoryServicesRolesParams {
	var ()
	return &PatchAPI24DirectoryServicesRolesParams{

		timeout: timeout,
	}
}

// NewPatchAPI24DirectoryServicesRolesParamsWithContext creates a new PatchAPI24DirectoryServicesRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPI24DirectoryServicesRolesParamsWithContext(ctx context.Context) *PatchAPI24DirectoryServicesRolesParams {
	var ()
	return &PatchAPI24DirectoryServicesRolesParams{

		Context: ctx,
	}
}

// NewPatchAPI24DirectoryServicesRolesParamsWithHTTPClient creates a new PatchAPI24DirectoryServicesRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPI24DirectoryServicesRolesParamsWithHTTPClient(client *http.Client) *PatchAPI24DirectoryServicesRolesParams {
	var ()
	return &PatchAPI24DirectoryServicesRolesParams{
		HTTPClient: client,
	}
}

/*PatchAPI24DirectoryServicesRolesParams contains all the parameters to send to the API endpoint
for the patch API 24 directory services roles operation typically these are written to a http.Request
*/
type PatchAPI24DirectoryServicesRolesParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*DirectoryServiceRoles*/
	DirectoryServiceRoles *models.DirectoryServiceRole
	/*RoleNames
	  Performs the operation on the unique roles specified. For example, `array_admin`. Enter multiple roles in comma-separated format.

	*/
	RoleNames []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) WithTimeout(timeout time.Duration) *PatchAPI24DirectoryServicesRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) WithContext(ctx context.Context) *PatchAPI24DirectoryServicesRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) WithHTTPClient(client *http.Client) *PatchAPI24DirectoryServicesRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) WithAuthorization(authorization *string) *PatchAPI24DirectoryServicesRolesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) WithXRequestID(xRequestID *string) *PatchAPI24DirectoryServicesRolesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithDirectoryServiceRoles adds the directoryServiceRoles to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) WithDirectoryServiceRoles(directoryServiceRoles *models.DirectoryServiceRole) *PatchAPI24DirectoryServicesRolesParams {
	o.SetDirectoryServiceRoles(directoryServiceRoles)
	return o
}

// SetDirectoryServiceRoles adds the directoryServiceRoles to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) SetDirectoryServiceRoles(directoryServiceRoles *models.DirectoryServiceRole) {
	o.DirectoryServiceRoles = directoryServiceRoles
}

// WithRoleNames adds the roleNames to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) WithRoleNames(roleNames []string) *PatchAPI24DirectoryServicesRolesParams {
	o.SetRoleNames(roleNames)
	return o
}

// SetRoleNames adds the roleNames to the patch API 24 directory services roles params
func (o *PatchAPI24DirectoryServicesRolesParams) SetRoleNames(roleNames []string) {
	o.RoleNames = roleNames
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPI24DirectoryServicesRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DirectoryServiceRoles != nil {
		if err := r.SetBodyParam(o.DirectoryServiceRoles); err != nil {
			return err
		}
	}

	valuesRoleNames := o.RoleNames

	joinedRoleNames := swag.JoinByFormat(valuesRoleNames, "csv")
	// query array param role_names
	if err := r.SetQueryParam("role_names", joinedRoleNames...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
