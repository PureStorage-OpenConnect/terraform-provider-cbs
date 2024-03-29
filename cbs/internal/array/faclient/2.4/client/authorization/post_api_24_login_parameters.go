// Code generated by go-swagger; DO NOT EDIT.

package authorization

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

// NewPostAPI24LoginParams creates a new PostAPI24LoginParams object
// with the default values initialized.
func NewPostAPI24LoginParams() *PostAPI24LoginParams {
	var ()
	return &PostAPI24LoginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24LoginParamsWithTimeout creates a new PostAPI24LoginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24LoginParamsWithTimeout(timeout time.Duration) *PostAPI24LoginParams {
	var ()
	return &PostAPI24LoginParams{

		timeout: timeout,
	}
}

// NewPostAPI24LoginParamsWithContext creates a new PostAPI24LoginParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24LoginParamsWithContext(ctx context.Context) *PostAPI24LoginParams {
	var ()
	return &PostAPI24LoginParams{

		Context: ctx,
	}
}

// NewPostAPI24LoginParamsWithHTTPClient creates a new PostAPI24LoginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24LoginParamsWithHTTPClient(client *http.Client) *PostAPI24LoginParams {
	var ()
	return &PostAPI24LoginParams{
		HTTPClient: client,
	}
}

/*PostAPI24LoginParams contains all the parameters to send to the API endpoint
for the post API 24 login operation typically these are written to a http.Request
*/
type PostAPI24LoginParams struct {

	/*APIToken
	  API token for a user.

	*/
	APIToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 login params
func (o *PostAPI24LoginParams) WithTimeout(timeout time.Duration) *PostAPI24LoginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 login params
func (o *PostAPI24LoginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 login params
func (o *PostAPI24LoginParams) WithContext(ctx context.Context) *PostAPI24LoginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 login params
func (o *PostAPI24LoginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 login params
func (o *PostAPI24LoginParams) WithHTTPClient(client *http.Client) *PostAPI24LoginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 login params
func (o *PostAPI24LoginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIToken adds the aPIToken to the post API 24 login params
func (o *PostAPI24LoginParams) WithAPIToken(aPIToken *string) *PostAPI24LoginParams {
	o.SetAPIToken(aPIToken)
	return o
}

// SetAPIToken adds the apiToken to the post API 24 login params
func (o *PostAPI24LoginParams) SetAPIToken(aPIToken *string) {
	o.APIToken = aPIToken
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24LoginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIToken != nil {

		// header param api-token
		if err := r.SetHeaderParam("api-token", *o.APIToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
