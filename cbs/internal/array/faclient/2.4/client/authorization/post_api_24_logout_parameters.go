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

// NewPostAPI24LogoutParams creates a new PostAPI24LogoutParams object
// with the default values initialized.
func NewPostAPI24LogoutParams() *PostAPI24LogoutParams {
	var ()
	return &PostAPI24LogoutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPI24LogoutParamsWithTimeout creates a new PostAPI24LogoutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPI24LogoutParamsWithTimeout(timeout time.Duration) *PostAPI24LogoutParams {
	var ()
	return &PostAPI24LogoutParams{

		timeout: timeout,
	}
}

// NewPostAPI24LogoutParamsWithContext creates a new PostAPI24LogoutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPI24LogoutParamsWithContext(ctx context.Context) *PostAPI24LogoutParams {
	var ()
	return &PostAPI24LogoutParams{

		Context: ctx,
	}
}

// NewPostAPI24LogoutParamsWithHTTPClient creates a new PostAPI24LogoutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPI24LogoutParamsWithHTTPClient(client *http.Client) *PostAPI24LogoutParams {
	var ()
	return &PostAPI24LogoutParams{
		HTTPClient: client,
	}
}

/*PostAPI24LogoutParams contains all the parameters to send to the API endpoint
for the post API 24 logout operation typically these are written to a http.Request
*/
type PostAPI24LogoutParams struct {

	/*XAuthToken
	  Session token for a user.

	*/
	XAuthToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API 24 logout params
func (o *PostAPI24LogoutParams) WithTimeout(timeout time.Duration) *PostAPI24LogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API 24 logout params
func (o *PostAPI24LogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API 24 logout params
func (o *PostAPI24LogoutParams) WithContext(ctx context.Context) *PostAPI24LogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API 24 logout params
func (o *PostAPI24LogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API 24 logout params
func (o *PostAPI24LogoutParams) WithHTTPClient(client *http.Client) *PostAPI24LogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API 24 logout params
func (o *PostAPI24LogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXAuthToken adds the xAuthToken to the post API 24 logout params
func (o *PostAPI24LogoutParams) WithXAuthToken(xAuthToken *string) *PostAPI24LogoutParams {
	o.SetXAuthToken(xAuthToken)
	return o
}

// SetXAuthToken adds the xAuthToken to the post API 24 logout params
func (o *PostAPI24LogoutParams) SetXAuthToken(xAuthToken *string) {
	o.XAuthToken = xAuthToken
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPI24LogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XAuthToken != nil {

		// header param x-auth-token
		if err := r.SetHeaderParam("x-auth-token", *o.XAuthToken); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
