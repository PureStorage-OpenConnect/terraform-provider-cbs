// Code generated by go-swagger; DO NOT EDIT.

package pods

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

// NewDeleteAPI24PodsParams creates a new DeleteAPI24PodsParams object
// with the default values initialized.
func NewDeleteAPI24PodsParams() *DeleteAPI24PodsParams {
	var ()
	return &DeleteAPI24PodsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24PodsParamsWithTimeout creates a new DeleteAPI24PodsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24PodsParamsWithTimeout(timeout time.Duration) *DeleteAPI24PodsParams {
	var ()
	return &DeleteAPI24PodsParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24PodsParamsWithContext creates a new DeleteAPI24PodsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24PodsParamsWithContext(ctx context.Context) *DeleteAPI24PodsParams {
	var ()
	return &DeleteAPI24PodsParams{

		Context: ctx,
	}
}

// NewDeleteAPI24PodsParamsWithHTTPClient creates a new DeleteAPI24PodsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24PodsParamsWithHTTPClient(client *http.Client) *DeleteAPI24PodsParams {
	var ()
	return &DeleteAPI24PodsParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24PodsParams contains all the parameters to send to the API endpoint
for the delete API 24 pods operation typically these are written to a http.Request
*/
type DeleteAPI24PodsParams struct {

	/*EradicateContents
	  Set to `true` to eradicate contents (e.g., volumes, protection groups, snapshots) and containers (e.g., pods, volume groups). This enables you to eradicate containers with contents.

	*/
	EradicateContents *bool
	/*Ids
	  Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The `ids` and `names` parameters cannot be provided together.

	*/
	Ids []string
	/*Names
	  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, `name01,name02`.

	*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) WithTimeout(timeout time.Duration) *DeleteAPI24PodsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) WithContext(ctx context.Context) *DeleteAPI24PodsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) WithHTTPClient(client *http.Client) *DeleteAPI24PodsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEradicateContents adds the eradicateContents to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) WithEradicateContents(eradicateContents *bool) *DeleteAPI24PodsParams {
	o.SetEradicateContents(eradicateContents)
	return o
}

// SetEradicateContents adds the eradicateContents to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) SetEradicateContents(eradicateContents *bool) {
	o.EradicateContents = eradicateContents
}

// WithIds adds the ids to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) WithIds(ids []string) *DeleteAPI24PodsParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithNames adds the names to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) WithNames(names []string) *DeleteAPI24PodsParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the delete API 24 pods params
func (o *DeleteAPI24PodsParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24PodsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EradicateContents != nil {

		// query param eradicate_contents
		var qrEradicateContents bool
		if o.EradicateContents != nil {
			qrEradicateContents = *o.EradicateContents
		}
		qEradicateContents := swag.FormatBool(qrEradicateContents)
		if qEradicateContents != "" {

			if err := r.SetQueryParam("eradicate_contents", qEradicateContents); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
