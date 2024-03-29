// Code generated by go-swagger; DO NOT EDIT.

package arrays

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
)

// NewDeleteAPI24ArraysFactoryResetTokenParams creates a new DeleteAPI24ArraysFactoryResetTokenParams object
// with the default values initialized.
func NewDeleteAPI24ArraysFactoryResetTokenParams() *DeleteAPI24ArraysFactoryResetTokenParams {

	return &DeleteAPI24ArraysFactoryResetTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24ArraysFactoryResetTokenParamsWithTimeout creates a new DeleteAPI24ArraysFactoryResetTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24ArraysFactoryResetTokenParamsWithTimeout(timeout time.Duration) *DeleteAPI24ArraysFactoryResetTokenParams {

	return &DeleteAPI24ArraysFactoryResetTokenParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24ArraysFactoryResetTokenParamsWithContext creates a new DeleteAPI24ArraysFactoryResetTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24ArraysFactoryResetTokenParamsWithContext(ctx context.Context) *DeleteAPI24ArraysFactoryResetTokenParams {

	return &DeleteAPI24ArraysFactoryResetTokenParams{

		Context: ctx,
	}
}

// NewDeleteAPI24ArraysFactoryResetTokenParamsWithHTTPClient creates a new DeleteAPI24ArraysFactoryResetTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24ArraysFactoryResetTokenParamsWithHTTPClient(client *http.Client) *DeleteAPI24ArraysFactoryResetTokenParams {

	return &DeleteAPI24ArraysFactoryResetTokenParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24ArraysFactoryResetTokenParams contains all the parameters to send to the API endpoint
for the delete API 24 arrays factory reset token operation typically these are written to a http.Request
*/
type DeleteAPI24ArraysFactoryResetTokenParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 arrays factory reset token params
func (o *DeleteAPI24ArraysFactoryResetTokenParams) WithTimeout(timeout time.Duration) *DeleteAPI24ArraysFactoryResetTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 arrays factory reset token params
func (o *DeleteAPI24ArraysFactoryResetTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 arrays factory reset token params
func (o *DeleteAPI24ArraysFactoryResetTokenParams) WithContext(ctx context.Context) *DeleteAPI24ArraysFactoryResetTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 arrays factory reset token params
func (o *DeleteAPI24ArraysFactoryResetTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 arrays factory reset token params
func (o *DeleteAPI24ArraysFactoryResetTokenParams) WithHTTPClient(client *http.Client) *DeleteAPI24ArraysFactoryResetTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 arrays factory reset token params
func (o *DeleteAPI24ArraysFactoryResetTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24ArraysFactoryResetTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
