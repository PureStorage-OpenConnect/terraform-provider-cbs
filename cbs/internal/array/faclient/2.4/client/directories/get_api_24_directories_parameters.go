// Code generated by go-swagger; DO NOT EDIT.

package directories

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

// NewGetAPI24DirectoriesParams creates a new GetAPI24DirectoriesParams object
// with the default values initialized.
func NewGetAPI24DirectoriesParams() *GetAPI24DirectoriesParams {
	var ()
	return &GetAPI24DirectoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPI24DirectoriesParamsWithTimeout creates a new GetAPI24DirectoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPI24DirectoriesParamsWithTimeout(timeout time.Duration) *GetAPI24DirectoriesParams {
	var ()
	return &GetAPI24DirectoriesParams{

		timeout: timeout,
	}
}

// NewGetAPI24DirectoriesParamsWithContext creates a new GetAPI24DirectoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPI24DirectoriesParamsWithContext(ctx context.Context) *GetAPI24DirectoriesParams {
	var ()
	return &GetAPI24DirectoriesParams{

		Context: ctx,
	}
}

// NewGetAPI24DirectoriesParamsWithHTTPClient creates a new GetAPI24DirectoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPI24DirectoriesParamsWithHTTPClient(client *http.Client) *GetAPI24DirectoriesParams {
	var ()
	return &GetAPI24DirectoriesParams{
		HTTPClient: client,
	}
}

/*GetAPI24DirectoriesParams contains all the parameters to send to the API endpoint
for the get API 24 directories operation typically these are written to a http.Request
*/
type GetAPI24DirectoriesParams struct {

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
	/*Destroyed
	  If set to `true`, lists only destroyed objects that are in the eradication pending state. If set to `false`, lists only objects that are not destroyed. For destroyed objects, the time remaining is displayed in milliseconds.

	*/
	Destroyed *bool
	/*FileSystemIds
	  Performs the operation on the file system ID specified. Enter multiple file system IDs in comma-separated format. The `file_system_ids` and `file_system_names` parameters cannot be provided together.

	*/
	FileSystemIds []string
	/*FileSystemNames
	  Performs the operation on the file system name specified. Enter multiple file system names in comma-separated format. For example, `filesystem01,filesystem02`.

	*/
	FileSystemNames []string
	/*Filter
	  Narrows down the results to only the response objects that satisfy the filter criteria.

	*/
	Filter *string
	/*Ids
	  Performs the operation on the unique resource IDs specified. Enter multiple resource IDs in comma-separated format. The `ids` and `names` parameters cannot be provided together.

	*/
	Ids []string
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
	/*TotalOnly
	  If set to `true`, returns the aggregate value of all items after filtering. Where it makes more sense, the average value is displayed instead. The values are displayed for each name where meaningful. If `total_only=true`, the `items` list will be empty.

	*/
	TotalOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithTimeout(timeout time.Duration) *GetAPI24DirectoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithContext(ctx context.Context) *GetAPI24DirectoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithHTTPClient(client *http.Client) *GetAPI24DirectoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithAuthorization(authorization *string) *GetAPI24DirectoriesParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithXRequestID(xRequestID *string) *GetAPI24DirectoriesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithContinuationToken adds the continuationToken to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithContinuationToken(continuationToken *string) *GetAPI24DirectoriesParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetContinuationToken(continuationToken *string) {
	o.ContinuationToken = continuationToken
}

// WithDestroyed adds the destroyed to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithDestroyed(destroyed *bool) *GetAPI24DirectoriesParams {
	o.SetDestroyed(destroyed)
	return o
}

// SetDestroyed adds the destroyed to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetDestroyed(destroyed *bool) {
	o.Destroyed = destroyed
}

// WithFileSystemIds adds the fileSystemIds to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithFileSystemIds(fileSystemIds []string) *GetAPI24DirectoriesParams {
	o.SetFileSystemIds(fileSystemIds)
	return o
}

// SetFileSystemIds adds the fileSystemIds to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetFileSystemIds(fileSystemIds []string) {
	o.FileSystemIds = fileSystemIds
}

// WithFileSystemNames adds the fileSystemNames to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithFileSystemNames(fileSystemNames []string) *GetAPI24DirectoriesParams {
	o.SetFileSystemNames(fileSystemNames)
	return o
}

// SetFileSystemNames adds the fileSystemNames to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetFileSystemNames(fileSystemNames []string) {
	o.FileSystemNames = fileSystemNames
}

// WithFilter adds the filter to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithFilter(filter *string) *GetAPI24DirectoriesParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithIds adds the ids to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithIds(ids []string) *GetAPI24DirectoriesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetIds(ids []string) {
	o.Ids = ids
}

// WithLimit adds the limit to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithLimit(limit *int32) *GetAPI24DirectoriesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithNames adds the names to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithNames(names []string) *GetAPI24DirectoriesParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetNames(names []string) {
	o.Names = names
}

// WithOffset adds the offset to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithOffset(offset *int32) *GetAPI24DirectoriesParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSort adds the sort to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithSort(sort []string) *GetAPI24DirectoriesParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithTotalItemCount adds the totalItemCount to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithTotalItemCount(totalItemCount *bool) *GetAPI24DirectoriesParams {
	o.SetTotalItemCount(totalItemCount)
	return o
}

// SetTotalItemCount adds the totalItemCount to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetTotalItemCount(totalItemCount *bool) {
	o.TotalItemCount = totalItemCount
}

// WithTotalOnly adds the totalOnly to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) WithTotalOnly(totalOnly *bool) *GetAPI24DirectoriesParams {
	o.SetTotalOnly(totalOnly)
	return o
}

// SetTotalOnly adds the totalOnly to the get API 24 directories params
func (o *GetAPI24DirectoriesParams) SetTotalOnly(totalOnly *bool) {
	o.TotalOnly = totalOnly
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPI24DirectoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Destroyed != nil {

		// query param destroyed
		var qrDestroyed bool
		if o.Destroyed != nil {
			qrDestroyed = *o.Destroyed
		}
		qDestroyed := swag.FormatBool(qrDestroyed)
		if qDestroyed != "" {

			if err := r.SetQueryParam("destroyed", qDestroyed); err != nil {
				return err
			}
		}

	}

	valuesFileSystemIds := o.FileSystemIds

	joinedFileSystemIds := swag.JoinByFormat(valuesFileSystemIds, "csv")
	// query array param file_system_ids
	if err := r.SetQueryParam("file_system_ids", joinedFileSystemIds...); err != nil {
		return err
	}

	valuesFileSystemNames := o.FileSystemNames

	joinedFileSystemNames := swag.JoinByFormat(valuesFileSystemNames, "csv")
	// query array param file_system_names
	if err := r.SetQueryParam("file_system_names", joinedFileSystemNames...); err != nil {
		return err
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

	valuesIds := o.Ids

	joinedIds := swag.JoinByFormat(valuesIds, "csv")
	// query array param ids
	if err := r.SetQueryParam("ids", joinedIds...); err != nil {
		return err
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
