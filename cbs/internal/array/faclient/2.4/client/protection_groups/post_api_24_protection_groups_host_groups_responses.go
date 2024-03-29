// Code generated by go-swagger; DO NOT EDIT.

package protection_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PostAPI24ProtectionGroupsHostGroupsReader is a Reader for the PostAPI24ProtectionGroupsHostGroups structure.
type PostAPI24ProtectionGroupsHostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24ProtectionGroupsHostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24ProtectionGroupsHostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24ProtectionGroupsHostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24ProtectionGroupsHostGroupsOK creates a PostApi24ProtectionGroupsHostGroupsOK with default headers values
func NewPostApi24ProtectionGroupsHostGroupsOK() *PostApi24ProtectionGroupsHostGroupsOK {
	return &PostApi24ProtectionGroupsHostGroupsOK{}
}

/*PostApi24ProtectionGroupsHostGroupsOK handles this case with default header values.

OK
*/
type PostApi24ProtectionGroupsHostGroupsOK struct {
	Payload *models.MemberNoIDAllResponse
}

func (o *PostApi24ProtectionGroupsHostGroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/protection-groups/host-groups][%d] postApi24ProtectionGroupsHostGroupsOK  %+v", 200, o.Payload)
}

func (o *PostApi24ProtectionGroupsHostGroupsOK) GetPayload() *models.MemberNoIDAllResponse {
	return o.Payload
}

func (o *PostApi24ProtectionGroupsHostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24ProtectionGroupsHostGroupsBadRequest creates a PostApi24ProtectionGroupsHostGroupsBadRequest with default headers values
func NewPostApi24ProtectionGroupsHostGroupsBadRequest() *PostApi24ProtectionGroupsHostGroupsBadRequest {
	return &PostApi24ProtectionGroupsHostGroupsBadRequest{}
}

/*PostApi24ProtectionGroupsHostGroupsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24ProtectionGroupsHostGroupsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24ProtectionGroupsHostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/protection-groups/host-groups][%d] postApi24ProtectionGroupsHostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24ProtectionGroupsHostGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24ProtectionGroupsHostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
