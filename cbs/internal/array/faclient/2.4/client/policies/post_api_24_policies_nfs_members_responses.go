// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PostAPI24PoliciesNfsMembersReader is a Reader for the PostAPI24PoliciesNfsMembers structure.
type PostAPI24PoliciesNfsMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24PoliciesNfsMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24PoliciesNfsMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24PoliciesNfsMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24PoliciesNfsMembersOK creates a PostApi24PoliciesNfsMembersOK with default headers values
func NewPostApi24PoliciesNfsMembersOK() *PostApi24PoliciesNfsMembersOK {
	return &PostApi24PoliciesNfsMembersOK{}
}

/*PostApi24PoliciesNfsMembersOK handles this case with default header values.

OK
*/
type PostApi24PoliciesNfsMembersOK struct {
	Payload *models.PolicyMemberExportResponse
}

func (o *PostApi24PoliciesNfsMembersOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/policies/nfs/members][%d] postApi24PoliciesNfsMembersOK  %+v", 200, o.Payload)
}

func (o *PostApi24PoliciesNfsMembersOK) GetPayload() *models.PolicyMemberExportResponse {
	return o.Payload
}

func (o *PostApi24PoliciesNfsMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyMemberExportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24PoliciesNfsMembersBadRequest creates a PostApi24PoliciesNfsMembersBadRequest with default headers values
func NewPostApi24PoliciesNfsMembersBadRequest() *PostApi24PoliciesNfsMembersBadRequest {
	return &PostApi24PoliciesNfsMembersBadRequest{}
}

/*PostApi24PoliciesNfsMembersBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24PoliciesNfsMembersBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24PoliciesNfsMembersBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/policies/nfs/members][%d] postApi24PoliciesNfsMembersBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24PoliciesNfsMembersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24PoliciesNfsMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
