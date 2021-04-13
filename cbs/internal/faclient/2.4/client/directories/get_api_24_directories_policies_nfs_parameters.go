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

// NewGetAPI24DirectoriesPoliciesNfsParams creates a new GetAPI24DirectoriesPoliciesNfsParams object
// with the default values initialized.
func NewGetAPI24DirectoriesPoliciesNfsParams() *GetAPI24DirectoriesPoliciesNfsParams {
	var ()
	return &GetAPI24DirectoriesPoliciesNfsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPI24DirectoriesPoliciesNfsParamsWithTimeout creates a new GetAPI24DirectoriesPoliciesNfsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPI24DirectoriesPoliciesNfsParamsWithTimeout(timeout time.Duration) *GetAPI24DirectoriesPoliciesNfsParams {
	var ()
	return &GetAPI24DirectoriesPoliciesNfsParams{

		timeout: timeout,
	}
}

// NewGetAPI24DirectoriesPoliciesNfsParamsWithContext creates a new GetAPI24DirectoriesPoliciesNfsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPI24DirectoriesPoliciesNfsParamsWithContext(ctx context.Context) *GetAPI24DirectoriesPoliciesNfsParams {
	var ()
	return &GetAPI24DirectoriesPoliciesNfsParams{

		Context: ctx,
	}
}

// NewGetAPI24DirectoriesPoliciesNfsParamsWithHTTPClient creates a new GetAPI24DirectoriesPoliciesNfsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPI24DirectoriesPoliciesNfsParamsWithHTTPClient(client *http.Client) *GetAPI24DirectoriesPoliciesNfsParams {
	var ()
	return &GetAPI24DirectoriesPoliciesNfsParams{
		HTTPClient: client,
	}
}

/*GetAPI24DirectoriesPoliciesNfsParams contains all the parameters to send to the API endpoint
for the get API 24 directories policies nfs operation typically these are written to a http.Request
*/
type GetAPI24DirectoriesPoliciesNfsParams struct {

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
	/*Filter
	  Narrows down the results to only the response objects that satisfy the filter criteria.

	*/
	Filter *string
	/*Limit
	  Limits the size of the response to the specified number of objects on each page. To return the total number of resources, set `limit=0`. The total number of resources is returned as a `total_item_count` value. If the page size requested is larger than the system maximum limit, the server returns the maximum limit, disregarding the requested page size.

	*/
	Limit *int32
	/*MemberIds
	  Performs the operation on the unique member IDs specified. Enter multiple member IDs in comma-separated format. The `member_ids` and `member_names` parameters cannot be provided together.

	*/
	MemberIds []string
	/*MemberNames
	  Performs the operation on the unique member name specified. Examples of members include volumes, hosts, host groups, and directories. Enter multiple names in comma-separated format. For example, `vol01,vol02`.

	*/
	MemberNames []string
	/*Offset
	  The starting position based on the results of the query in relation to the full set of response objects returned.

	*/
	Offset *int32
	/*PolicyIds
	  Performs the operation on the unique policy IDs specified. Enter multiple policy IDs in comma-separated format. The `policy_ids` and `policy_names` parameters cannot be provided together.

	*/
	PolicyIds []string
	/*PolicyNames
	  Performs the operation on the policy names specified. Enter multiple policy names in comma-separated format. For example, `name01,name02`.

	*/
	PolicyNames []string
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

// WithTimeout adds the timeout to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithTimeout(timeout time.Duration) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithContext(ctx context.Context) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithHTTPClient(client *http.Client) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithAuthorization(authorization *string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetAuthorization(authorization *string) {
	o.Authorization = authorization
}

// WithXRequestID adds the xRequestID to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithXRequestID(xRequestID *string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithContinuationToken adds the continuationToken to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithContinuationToken(continuationToken *string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetContinuationToken(continuationToken)
	return o
}

// SetContinuationToken adds the continuationToken to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetContinuationToken(continuationToken *string) {
	o.ContinuationToken = continuationToken
}

// WithDestroyed adds the destroyed to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithDestroyed(destroyed *bool) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetDestroyed(destroyed)
	return o
}

// SetDestroyed adds the destroyed to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetDestroyed(destroyed *bool) {
	o.Destroyed = destroyed
}

// WithFilter adds the filter to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithFilter(filter *string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithLimit adds the limit to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithLimit(limit *int32) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithMemberIds adds the memberIds to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithMemberIds(memberIds []string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetMemberIds(memberIds)
	return o
}

// SetMemberIds adds the memberIds to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetMemberIds(memberIds []string) {
	o.MemberIds = memberIds
}

// WithMemberNames adds the memberNames to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithMemberNames(memberNames []string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetMemberNames(memberNames)
	return o
}

// SetMemberNames adds the memberNames to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetMemberNames(memberNames []string) {
	o.MemberNames = memberNames
}

// WithOffset adds the offset to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithOffset(offset *int32) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithPolicyIds adds the policyIds to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithPolicyIds(policyIds []string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetPolicyIds(policyIds)
	return o
}

// SetPolicyIds adds the policyIds to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetPolicyIds(policyIds []string) {
	o.PolicyIds = policyIds
}

// WithPolicyNames adds the policyNames to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithPolicyNames(policyNames []string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetPolicyNames(policyNames)
	return o
}

// SetPolicyNames adds the policyNames to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetPolicyNames(policyNames []string) {
	o.PolicyNames = policyNames
}

// WithSort adds the sort to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithSort(sort []string) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetSort(sort []string) {
	o.Sort = sort
}

// WithTotalItemCount adds the totalItemCount to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) WithTotalItemCount(totalItemCount *bool) *GetAPI24DirectoriesPoliciesNfsParams {
	o.SetTotalItemCount(totalItemCount)
	return o
}

// SetTotalItemCount adds the totalItemCount to the get API 24 directories policies nfs params
func (o *GetAPI24DirectoriesPoliciesNfsParams) SetTotalItemCount(totalItemCount *bool) {
	o.TotalItemCount = totalItemCount
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPI24DirectoriesPoliciesNfsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesMemberIds := o.MemberIds

	joinedMemberIds := swag.JoinByFormat(valuesMemberIds, "csv")
	// query array param member_ids
	if err := r.SetQueryParam("member_ids", joinedMemberIds...); err != nil {
		return err
	}

	valuesMemberNames := o.MemberNames

	joinedMemberNames := swag.JoinByFormat(valuesMemberNames, "csv")
	// query array param member_names
	if err := r.SetQueryParam("member_names", joinedMemberNames...); err != nil {
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

	valuesPolicyIds := o.PolicyIds

	joinedPolicyIds := swag.JoinByFormat(valuesPolicyIds, "csv")
	// query array param policy_ids
	if err := r.SetQueryParam("policy_ids", joinedPolicyIds...); err != nil {
		return err
	}

	valuesPolicyNames := o.PolicyNames

	joinedPolicyNames := swag.JoinByFormat(valuesPolicyNames, "csv")
	// query array param policy_names
	if err := r.SetQueryParam("policy_names", joinedPolicyNames...); err != nil {
		return err
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
