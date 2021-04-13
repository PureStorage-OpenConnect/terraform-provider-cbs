// Code generated by go-swagger; DO NOT EDIT.

package network_interfaces

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

// NewGetAPI24NetworkInterfacesPerformanceParams creates a new GetAPI24NetworkInterfacesPerformanceParams object
// with the default values initialized.
func NewGetAPI24NetworkInterfacesPerformanceParams() *GetAPI24NetworkInterfacesPerformanceParams {
	var ()
	return &GetAPI24NetworkInterfacesPerformanceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPI24NetworkInterfacesPerformanceParamsWithTimeout creates a new GetAPI24NetworkInterfacesPerformanceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPI24NetworkInterfacesPerformanceParamsWithTimeout(timeout time.Duration) *GetAPI24NetworkInterfacesPerformanceParams {
	var ()
	return &GetAPI24NetworkInterfacesPerformanceParams{

		timeout: timeout,
	}
}

// NewGetAPI24NetworkInterfacesPerformanceParamsWithContext creates a new GetAPI24NetworkInterfacesPerformanceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPI24NetworkInterfacesPerformanceParamsWithContext(ctx context.Context) *GetAPI24NetworkInterfacesPerformanceParams {
	var ()
	return &GetAPI24NetworkInterfacesPerformanceParams{

		Context: ctx,
	}
}

// NewGetAPI24NetworkInterfacesPerformanceParamsWithHTTPClient creates a new GetAPI24NetworkInterfacesPerformanceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPI24NetworkInterfacesPerformanceParamsWithHTTPClient(client *http.Client) *GetAPI24NetworkInterfacesPerformanceParams {
	var ()
	return &GetAPI24NetworkInterfacesPerformanceParams{
		HTTPClient: client,
	}
}

/*GetAPI24NetworkInterfacesPerformanceParams contains all the parameters to send to the API endpoint
for the get API 24 network interfaces performance operation typically these are written to a http.Request
*/
type GetAPI24NetworkInterfacesPerformanceParams struct {

	/*Authorization
	  Access token (in JWT format) required to use any API endpoint (except `/oauth2`, `/login`, and `/logout`)

	*/
	Authorization *string
	/*XRequestID
	  Supplied by client during request or generated by server.

	*/
	XRequestID *string
	/*EndTime
	  Displays historical performance data for the specified time window, where `start_time` is the beginning of the time window, and `end_time` is the end of the time window. The `start_time` and `end_time` parameters are specified in milliseconds since the UNIX epoch. If `start_time` is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. If `end_time`is not specified, the end time will default to the current time. Include the `resolution` parameter to display the performance data at the specified resolution. If not specified, `resolution` defaults to the lowest valid resolution.

	*/
	EndTime *int64
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
	/*Resolution
	  The number of milliseconds between samples of historical data. For array-wide performance metrics (`/arrays/performance` endpoint), valid values are `1000` (1 second), `30000` (30 seconds), `300000` (5 minutes), `1800000` (30 minutes), `7200000` (2 hours), `28800000` (8 hours), and `86400000` (24 hours). For performance metrics on storage objects (`<object name>/performance` endpoint), such as volumes, valid values are `30000` (30 seconds), `300000` (5 minutes), `1800000` (30 minutes), `7200000` (2 hours), `28800000` (8 hours), and `86400000` (24 hours). For space metrics, (`<object name>/space` endpoint), valid values are `300000` (5 minutes), `1800000` (30 minutes), `7200000` (2 hours), `28800000` (8 hours), and `86400000` (24 hours). Include the `start_time` parameter to display the performance data starting at the specified start time. If `start_time` is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. Include the `end_time` parameter to display the performance data until the specified end time. If `end_time`is not specified, the end time will default to the current time. If the `resolution` parameter is not specified but either the `start_time` or `end_time` parameter is, then `resolution` will default to the lowest valid resolution.

	*/
	Resolution *int64
	/*Sort
	  Returns the response objects in the order specified. Set `sort` to the name in the response by which to sort. Sorting can be performed on any of the names in the response, and the objects can be sorted in ascending or descending order. By default, the response objects are sorted in ascending order. To sort in descending order, append the minus sign (`-`) to the name. A single request can be sorted on multiple objects. For example, you can sort all volumes from largest to smallest volume size, and then sort volumes of the same size in ascending order by volume name. To sort on multiple names, list the names as comma-separated values.

	*/
	Sort []string
	/*StartTime
	  Displays historical performance data for the specified time window, where `start_time` is the beginning of the time window, and `end_time` is the end of the time window. The `start_time` and `end_time` parameters are specified in milliseconds since the UNIX epoch. If `start_time` is not specified, the start time will default to one resolution before the end time, meaning that the most recent sample of performance data will be displayed. If `end_time`is not specified, the end time will default to the current time. Include the `resolution` parameter to display the performance data at the specified resolution. If not specified, `resolution` defaults to the lowest valid resolution.

	*/
	StartTime *int64
	/*TotalItemCount
	  If set to `true`, the `total_item_count` matching the specified query parameters is calculated and returned in the response. If set to `false`, the `total_item_count` is `null` in the response. This may speed up queries where the `total_item_count` is large. If not specified, defaults to `false`.

	*/
	TotalItemCount *bool
	/*TotalOnly
	  If set to `true`, returns the aggregate value of all items after filtering. Where it makes more sense, the average value is displayed instead. The values are displayed for each name where meaningful. If `total_only=true`, the `items` list will be empty.

	*/
	TotalOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithTimeout(timeout time.Duration) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithContext(ctx context.Context) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithHTTPClient(client *http.Client) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithAuthorization(authorization *string) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithXRequestID(xRequestID *string) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithEndTime adds the endTime to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithEndTime(endTime *int64) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetEndTime(endTime *int64) {
	o.EndTime = endTime
}

// WithFilter adds the filter to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithFilter(filter *string) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithLimit(limit *int32) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNames adds the names to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithNames(names []string) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetNames(names []string) {
	o.Names = names
}

// WithOffset adds the offset to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithOffset(offset *int32) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithResolution adds the resolution to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithResolution(resolution *int64) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetResolution(resolution)
	return o
}

// SetResolution adds the resolution to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetResolution(resolution *int64) {
	o.Resolution = resolution
}

// WithSort adds the sort to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithSort(sort []string) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithStartTime adds the startTime to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithStartTime(startTime *int64) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetStartTime(startTime *int64) {
	o.StartTime = startTime
}

// WithTotalItemCount adds the totalItemCount to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithTotalItemCount(totalItemCount *bool) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetTotalItemCount(totalItemCount)
	return o
}

// SetTotalItemCount adds the totalItemCount to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetTotalItemCount(totalItemCount *bool) {
	o.TotalItemCount = totalItemCount
}

// WithTotalOnly adds the totalOnly to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) WithTotalOnly(totalOnly *bool) *GetAPI24NetworkInterfacesPerformanceParams {
	o.SetTotalOnly(totalOnly)
	return o
}

// SetTotalOnly adds the totalOnly to the get API 24 network interfaces performance params
func (o *GetAPI24NetworkInterfacesPerformanceParams) SetTotalOnly(totalOnly *bool) {
	o.TotalOnly = totalOnly
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPI24NetworkInterfacesPerformanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EndTime != nil {

		// query param end_time
		var qrEndTime int64
		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := swag.FormatInt64(qrEndTime)
		if qEndTime != "" {

			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
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

	if o.Resolution != nil {

		// query param resolution
		var qrResolution int64
		if o.Resolution != nil {
			qrResolution = *o.Resolution
		}
		qResolution := swag.FormatInt64(qrResolution)
		if qResolution != "" {

			if err := r.SetQueryParam("resolution", qResolution); err != nil {
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

	if o.StartTime != nil {

		// query param start_time
		var qrStartTime int64
		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := swag.FormatInt64(qrStartTime)
		if qStartTime != "" {

			if err := r.SetQueryParam("start_time", qStartTime); err != nil {
				return err
			}
		}

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

	if o.TotalOnly != nil {

		// query param total_only
		var qrTotalOnly bool
		if o.TotalOnly != nil {
			qrTotalOnly = *o.TotalOnly
		}
		qTotalOnly := swag.FormatBool(qrTotalOnly)
		if qTotalOnly != "" {

			if err := r.SetQueryParam("total_only", qTotalOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
