// Code generated by go-swagger; DO NOT EDIT.

package apps

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

// NewGetAPI24AppsNodesParams creates a new GetAPI24AppsNodesParams object
// with the default values initialized.
func NewGetAPI24AppsNodesParams() *GetAPI24AppsNodesParams {
	var ()
	return &GetAPI24AppsNodesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPI24AppsNodesParamsWithTimeout creates a new GetAPI24AppsNodesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPI24AppsNodesParamsWithTimeout(timeout time.Duration) *GetAPI24AppsNodesParams {
	var ()
	return &GetAPI24AppsNodesParams{

		timeout: timeout,
	}
}

// NewGetAPI24AppsNodesParamsWithContext creates a new GetAPI24AppsNodesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPI24AppsNodesParamsWithContext(ctx context.Context) *GetAPI24AppsNodesParams {
	var ()
	return &GetAPI24AppsNodesParams{

		Context: ctx,
	}
}

// NewGetAPI24AppsNodesParamsWithHTTPClient creates a new GetAPI24AppsNodesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPI24AppsNodesParamsWithHTTPClient(client *http.Client) *GetAPI24AppsNodesParams {
	var ()
	return &GetAPI24AppsNodesParams{
		HTTPClient: client,
	}
}

/*GetAPI24AppsNodesParams contains all the parameters to send to the API endpoint
for the get API 24 apps nodes operation typically these are written to a http.Request
*/
type GetAPI24AppsNodesParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*AppNames
	  The name of the installed app. Enter multiple names in comma-separated format.

	*/
	AppNames []string
	/*ContinuationToken
	  A token used to retrieve the next page of data with some consistency guaranteed. The token is a Base64 encoded value. Set `continuation_token` to the system-generated token taken from the `x-next-token` header field of the response. A query has reached its last page when the response does not include a token. Pagination requires the `limit` and `continuation_token` query parameters.

	*/
	ContinuationToken *string
	/*Filter
	  Narrows down the results to only the response objects that satisfy the filter criteria.

	*/
	Filter *string
	/*Limit
	  Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set `limit=0`. The total number of resources is returned as a `total_item_count` value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size.

	*/
	Limit *int32
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

// WithTimeout adds the timeout to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithTimeout(timeout time.Duration) *GetAPI24AppsNodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithContext(ctx context.Context) *GetAPI24AppsNodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithHTTPClient(client *http.Client) *GetAPI24AppsNodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithAuthorization(authorization *string) *GetAPI24AppsNodesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithXRequestID(xRequestID *string) *GetAPI24AppsNodesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithAppNames adds the appNames to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithAppNames(appNames []string) *GetAPI24AppsNodesParams {
	o.SetAppNames(appNames)
	return o
}

// SetAppNames adds the appNames to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetAppNames(appNames []string) {
	o.AppNames = appNames
}

// WithContinuationToken adds the continuationToken to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithContinuationToken(continuationToken *string) *GetAPI24AppsNodesParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetContinuationToken(continuationToken *string) {
	o.ContinuationToken = continuationToken
}

// WithFilter adds the filter to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithFilter(filter *string) *GetAPI24AppsNodesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithLimit(limit *int32) *GetAPI24AppsNodesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOffset adds the offset to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithOffset(offset *int32) *GetAPI24AppsNodesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithSort(sort []string) *GetAPI24AppsNodesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithTotalItemCount adds the totalItemCount to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) WithTotalItemCount(totalItemCount *bool) *GetAPI24AppsNodesParams {
	o.SetTotalItemCount(totalItemCount)
	return o
}

// SetTotalItemCount adds the totalItemCount to the get API 24 apps nodes params
func (o *GetAPI24AppsNodesParams) SetTotalItemCount(totalItemCount *bool) {
	o.TotalItemCount = totalItemCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPI24AppsNodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesAppNames := o.AppNames

	joinedAppNames := swag.JoinByFormat(valuesAppNames, "csv")
	// query array param app_names
	if err := r.SetQueryParam("app_names", joinedAppNames...); err != nil {
		return err
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