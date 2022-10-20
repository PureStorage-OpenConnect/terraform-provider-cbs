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

// PostAPI24ProtectionGroupsHostsReader is a Reader for the PostAPI24ProtectionGroupsHosts structure.
type PostAPI24ProtectionGroupsHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24ProtectionGroupsHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24ProtectionGroupsHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24ProtectionGroupsHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24ProtectionGroupsHostsOK creates a PostApi24ProtectionGroupsHostsOK with default headers values
func NewPostApi24ProtectionGroupsHostsOK() *PostApi24ProtectionGroupsHostsOK {
	return &PostApi24ProtectionGroupsHostsOK{}
}

/*PostApi24ProtectionGroupsHostsOK handles this case with default header values.

OK
*/
type PostApi24ProtectionGroupsHostsOK struct {
	Payload *models.MemberNoIDAllResponse
}

func (o *PostApi24ProtectionGroupsHostsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/protection-groups/hosts][%d] postApi24ProtectionGroupsHostsOK  %+v", 200, o.Payload)
}

func (o *PostApi24ProtectionGroupsHostsOK) GetPayload() *models.MemberNoIDAllResponse {
	return o.Payload
}

func (o *PostApi24ProtectionGroupsHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24ProtectionGroupsHostsBadRequest creates a PostApi24ProtectionGroupsHostsBadRequest with default headers values
func NewPostApi24ProtectionGroupsHostsBadRequest() *PostApi24ProtectionGroupsHostsBadRequest {
	return &PostApi24ProtectionGroupsHostsBadRequest{}
}

/*PostApi24ProtectionGroupsHostsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24ProtectionGroupsHostsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24ProtectionGroupsHostsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/protection-groups/hosts][%d] postApi24ProtectionGroupsHostsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24ProtectionGroupsHostsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24ProtectionGroupsHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
