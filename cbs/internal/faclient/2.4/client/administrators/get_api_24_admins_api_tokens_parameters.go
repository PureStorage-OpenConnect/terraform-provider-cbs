// Code generated by go-swagger; DO NOT EDIT.

package administrators

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

// NewGetAPI24AdminsAPITokensParams creates a new GetAPI24AdminsAPITokensParams object
// with the default values initialized.
func NewGetAPI24AdminsAPITokensParams() *GetAPI24AdminsAPITokensParams {
	var ()
	return &GetAPI24AdminsAPITokensParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPI24AdminsAPITokensParamsWithTimeout creates a new GetAPI24AdminsAPITokensParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPI24AdminsAPITokensParamsWithTimeout(timeout time.Duration) *GetAPI24AdminsAPITokensParams {
	var ()
	return &GetAPI24AdminsAPITokensParams{

		timeout: timeout,
	}
}

// NewGetAPI24AdminsAPITokensParamsWithContext creates a new GetAPI24AdminsAPITokensParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPI24AdminsAPITokensParamsWithContext(ctx context.Context) *GetAPI24AdminsAPITokensParams {
	var ()
	return &GetAPI24AdminsAPITokensParams{

		Context: ctx,
	}
}

// NewGetAPI24AdminsAPITokensParamsWithHTTPClient creates a new GetAPI24AdminsAPITokensParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPI24AdminsAPITokensParamsWithHTTPClient(client *http.Client) *GetAPI24AdminsAPITokensParams {
	var ()
	return &GetAPI24AdminsAPITokensParams{
		HTTPClient: client,
	}
}

/*GetAPI24AdminsAPITokensParams contains all the parameters to send to the API endpoint
for the get API 24 admins API tokens operation typically these are written to a http.Request
*/
type GetAPI24AdminsAPITokensParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*ContinuationToken
	  A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set `continuation_token` to the system-generated token taken from the `x-next-token` header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the `limit` and `continuation_token` query parameters.

	*/
	ContinuationToken *string
	/*ExposeAPIToken
	  If `true`, exposes the API token of the current user.

	*/
	ExposeAPIToken *bool
	/*Filter
	  Narrows down the results to only the response objects that satisfy the filter criteria.

	*/
	Filter *string
	/*Limit
	  Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set `limit=0`. The total number of resources is returned as a `total_item_count` value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size.

	*/
	Limit *int32
	/*Names
	  Performs the operation on the unique name specified. Enter multiple names in comma-separated format. For example, `name01,name02`.

	*/
	Names []string
	/*Offset
	  The starting position based on the results of the query in relation to the full set of response objects returned.

	*/
	Offset *int32
	/*Sort
	  Returns the response objects in the order specified. Set `sort` to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (`-`) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values.

	*/
	Sort []string
	/*TotalItemCount
	  If set to `true`, the `total_item_count` matching the specified query parameters is calculated and returned in the response. If set to `false`, the `total_item_count` is `null` in the response. This may speed up queries where the `total_item_count` is large. If not specified, defaults to `false`.

	*/
	TotalItemCount *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithTimeout(timeout time.Duration) *GetAPI24AdminsAPITokensParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithContext(ctx context.Context) *GetAPI24AdminsAPITokensParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithHTTPClient(client *http.Client) *GetAPI24AdminsAPITokensParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithAuthorization(authorization *string) *GetAPI24AdminsAPITokensParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithXRequestID(xRequestID *string) *GetAPI24AdminsAPITokensParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithContinuationToken adds the continuationToken to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithContinuationToken(continuationToken *string) *GetAPI24AdminsAPITokensParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetContinuationToken(continuationToken *string) {
	o.ContinuationToken = continuationToken
}

// WithExposeAPIToken adds the exposeAPIToken to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithExposeAPIToken(exposeAPIToken *bool) *GetAPI24AdminsAPITokensParams {
	o.SetExposeAPIToken(exposeAPIToken)
	return o
}

// SetExposeAPIToken adds the exposeApiToken to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetExposeAPIToken(exposeAPIToken *bool) {
	o.ExposeAPIToken = exposeAPIToken
}

// WithFilter adds the filter to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithFilter(filter *string) *GetAPI24AdminsAPITokensParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithLimit(limit *int32) *GetAPI24AdminsAPITokensParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNames adds the names to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithNames(names []string) *GetAPI24AdminsAPITokensParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetNames(names []string) {
	o.Names = names
}

// WithOffset adds the offset to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithOffset(offset *int32) *GetAPI24AdminsAPITokensParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithSort(sort []string) *GetAPI24AdminsAPITokensParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithTotalItemCount adds the totalItemCount to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) WithTotalItemCount(totalItemCount *bool) *GetAPI24AdminsAPITokensParams {
	o.SetTotalItemCount(totalItemCount)
	return o
}

// SetTotalItemCount adds the totalItemCount to the get API 24 admins API tokens params
func (o *GetAPI24AdminsAPITokensParams) SetTotalItemCount(totalItemCount *bool) {
	o.TotalItemCount = totalItemCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPI24AdminsAPITokensParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.ContinuationToken != nil {

		// query param continuation_token
		var qrContinuationToken string
		if o.ContinuationToken != nil {
			qrContinuationToken = *o.ContinuationToken
		}
		qContinuationToken := qrContinuationToken
		if qContinuationToken != "" {

			if err := r.SetQueryParam("continuation_token", qContinuationToken); err != nil {
				return err
			}
		}

	}

	if o.ExposeAPIToken != nil {

		// query param expose_api_token
		var qrExposeAPIToken bool
		if o.ExposeAPIToken != nil {
			qrExposeAPIToken = *o.ExposeAPIToken
		}
		qExposeAPIToken := swag.FormatBool(qrExposeAPIToken)
		if qExposeAPIToken != "" {

			if err := r.SetQueryParam("expose_api_token", qExposeAPIToken); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	valuesNames := o.Names

	joinedNames := swag.JoinByFormat(valuesNames, "csv")
	// query array param names
	if err := r.SetQueryParam("names", joinedNames...); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	valuesSort := o.Sort

	joinedSort := swag.JoinByFormat(valuesSort, "csv")
	// query array param sort
	if err := r.SetQueryParam("sort", joinedSort...); err != nil {
		return err
	}

	if o.TotalItemCount != nil {

		// query param total_item_count
		var qrTotalItemCount bool
		if o.TotalItemCount != nil {
			qrTotalItemCount = *o.TotalItemCount
		}
		qTotalItemCount := swag.FormatBool(qrTotalItemCount)
		if qTotalItemCount != "" {

			if err := r.SetQueryParam("total_item_count", qTotalItemCount); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
