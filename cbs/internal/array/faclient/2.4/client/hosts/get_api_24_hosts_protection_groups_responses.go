// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24HostsProtectionGroupsReader is a Reader for the GetAPI24HostsProtectionGroups structure.
type GetAPI24HostsProtectionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24HostsProtectionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24HostsProtectionGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24HostsProtectionGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24HostsProtectionGroupsOK creates a GetApi24HostsProtectionGroupsOK with default headers values
func NewGetApi24HostsProtectionGroupsOK() *GetApi24HostsProtectionGroupsOK {
	return &GetApi24HostsProtectionGroupsOK{}
}

/*GetApi24HostsProtectionGroupsOK handles this case with default header values.

OK
*/
type GetApi24HostsProtectionGroupsOK struct {
	Payload *models.MemberNoIDAllGetResponse
}

func (o *GetApi24HostsProtectionGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/hosts/protection-groups][%d] getApi24HostsProtectionGroupsOK  %+v", 200, o.Payload)
}

func (o *GetApi24HostsProtectionGroupsOK) GetPayload() *models.MemberNoIDAllGetResponse {
	return o.Payload
}

func (o *GetApi24HostsProtectionGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24HostsProtectionGroupsBadRequest creates a GetApi24HostsProtectionGroupsBadRequest with default headers values
func NewGetApi24HostsProtectionGroupsBadRequest() *GetApi24HostsProtectionGroupsBadRequest {
	return &GetApi24HostsProtectionGroupsBadRequest{}
}

/*GetApi24HostsProtectionGroupsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24HostsProtectionGroupsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24HostsProtectionGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/hosts/protection-groups][%d] getApi24HostsProtectionGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24HostsProtectionGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24HostsProtectionGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
