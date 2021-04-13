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
	"github.com/go-openapi/swag"
)

// NewDeleteAPI24ArraysParams creates a new DeleteAPI24ArraysParams object
// with the default values initialized.
func NewDeleteAPI24ArraysParams() *DeleteAPI24ArraysParams {
	var ()
	return &DeleteAPI24ArraysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPI24ArraysParamsWithTimeout creates a new DeleteAPI24ArraysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPI24ArraysParamsWithTimeout(timeout time.Duration) *DeleteAPI24ArraysParams {
	var ()
	return &DeleteAPI24ArraysParams{

		timeout: timeout,
	}
}

// NewDeleteAPI24ArraysParamsWithContext creates a new DeleteAPI24ArraysParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPI24ArraysParamsWithContext(ctx context.Context) *DeleteAPI24ArraysParams {
	var ()
	return &DeleteAPI24ArraysParams{

		Context: ctx,
	}
}

// NewDeleteAPI24ArraysParamsWithHTTPClient creates a new DeleteAPI24ArraysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPI24ArraysParamsWithHTTPClient(client *http.Client) *DeleteAPI24ArraysParams {
	var ()
	return &DeleteAPI24ArraysParams{
		HTTPClient: client,
	}
}

/*DeleteAPI24ArraysParams contains all the parameters to send to the API endpoint
for the delete API 24 arrays operation typically these are written to a http.Request
*/
type DeleteAPI24ArraysParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*EradicateAllData
	  Set to `true` to perform a factory reset.

	*/
	EradicateAllData *bool
	/*FactoryResetToken
	  A token required to perform a factory reset.

	*/
	FactoryResetToken *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) WithTimeout(timeout time.Duration) *DeleteAPI24ArraysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) WithContext(ctx context.Context) *DeleteAPI24ArraysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) WithHTTPClient(client *http.Client) *DeleteAPI24ArraysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) WithAuthorization(authorization *string) *DeleteAPI24ArraysParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) WithXRequestID(xRequestID *string) *DeleteAPI24ArraysParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEradicateAllData adds the eradicateAllData to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) WithEradicateAllData(eradicateAllData *bool) *DeleteAPI24ArraysParams {
	o.SetEradicateAllData(eradicateAllData)
	return o
}

// SetEradicateAllData adds the eradicateAllData to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) SetEradicateAllData(eradicateAllData *bool) {
	o.EradicateAllData = eradicateAllData
}

// WithFactoryResetToken adds the factoryResetToken to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) WithFactoryResetToken(factoryResetToken *int64) *DeleteAPI24ArraysParams {
	o.SetFactoryResetToken(factoryResetToken)
	return o
}

// SetFactoryResetToken adds the factoryResetToken to the delete API 24 arrays params
func (o *DeleteAPI24ArraysParams) SetFactoryResetToken(factoryResetToken *int64) {
	o.FactoryResetToken = factoryResetToken
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPI24ArraysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EradicateAllData != nil {

		// query param eradicate_all_data
		var qrEradicateAllData bool
		if o.EradicateAllData != nil {
			qrEradicateAllData = *o.EradicateAllData
		}
		qEradicateAllData := swag.FormatBool(qrEradicateAllData)
		if qEradicateAllData != "" {

			if err := r.SetQueryParam("eradicate_all_data", qEradicateAllData); err != nil {
				return err
			}
		}

	}

	if o.FactoryResetToken != nil {

		// query param factory_reset_token
		var qrFactoryResetToken int64
		if o.FactoryResetToken != nil {
			qrFactoryResetToken = *o.FactoryResetToken
		}
		qFactoryResetToken := swag.FormatInt64(qrFactoryResetToken)
		if qFactoryResetToken != "" {

			if err := r.SetQueryParam("factory_reset_token", qFactoryResetToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
