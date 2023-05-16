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

// GetAPI24PoliciesSmbMembersReader is a Reader for the GetAPI24PoliciesSmbMembers structure.
type GetAPI24PoliciesSmbMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24PoliciesSmbMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24PoliciesSmbMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24PoliciesSmbMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24PoliciesSmbMembersOK creates a GetApi24PoliciesSmbMembersOK with default headers values
func NewGetApi24PoliciesSmbMembersOK() *GetApi24PoliciesSmbMembersOK {
	return &GetApi24PoliciesSmbMembersOK{}
}

/*GetApi24PoliciesSmbMembersOK handles this case with default header values.

OK
*/
type GetApi24PoliciesSmbMembersOK struct {
	Payload *models.PolicyMemberExportGetResponse
}

func (o *GetApi24PoliciesSmbMembersOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/policies/smb/members][%d] getApi24PoliciesSmbMembersOK  %+v", 200, o.Payload)
}

func (o *GetApi24PoliciesSmbMembersOK) GetPayload() *models.PolicyMemberExportGetResponse {
	return o.Payload
}

func (o *GetApi24PoliciesSmbMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyMemberExportGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24PoliciesSmbMembersBadRequest creates a GetApi24PoliciesSmbMembersBadRequest with default headers values
func NewGetApi24PoliciesSmbMembersBadRequest() *GetApi24PoliciesSmbMembersBadRequest {
	return &GetApi24PoliciesSmbMembersBadRequest{}
}

/*GetApi24PoliciesSmbMembersBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24PoliciesSmbMembersBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24PoliciesSmbMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/policies/smb/members][%d] getApi24PoliciesSmbMembersBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24PoliciesSmbMembersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24PoliciesSmbMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
