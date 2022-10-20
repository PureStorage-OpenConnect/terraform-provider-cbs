// Code generated by go-swagger; DO NOT EDIT.

package host_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24HostGroupsProtectionGroupsReader is a Reader for the GetAPI24HostGroupsProtectionGroups structure.
type GetAPI24HostGroupsProtectionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24HostGroupsProtectionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24HostGroupsProtectionGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24HostGroupsProtectionGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24HostGroupsProtectionGroupsOK creates a GetApi24HostGroupsProtectionGroupsOK with default headers values
func NewGetApi24HostGroupsProtectionGroupsOK() *GetApi24HostGroupsProtectionGroupsOK {
	return &GetApi24HostGroupsProtectionGroupsOK{}
}

/*GetApi24HostGroupsProtectionGroupsOK handles this case with default header values.

OK
*/
type GetApi24HostGroupsProtectionGroupsOK struct {
	Payload *models.MemberNoIDAllGetResponse
}

func (o *GetApi24HostGroupsProtectionGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/host-groups/protection-groups][%d] getApi24HostGroupsProtectionGroupsOK  %+v", 200, o.Payload)
}

func (o *GetApi24HostGroupsProtectionGroupsOK) GetPayload() *models.MemberNoIDAllGetResponse {
	return o.Payload
}

func (o *GetApi24HostGroupsProtectionGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24HostGroupsProtectionGroupsBadRequest creates a GetApi24HostGroupsProtectionGroupsBadRequest with default headers values
func NewGetApi24HostGroupsProtectionGroupsBadRequest() *GetApi24HostGroupsProtectionGroupsBadRequest {
	return &GetApi24HostGroupsProtectionGroupsBadRequest{}
}

/*GetApi24HostGroupsProtectionGroupsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24HostGroupsProtectionGroupsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24HostGroupsProtectionGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/host-groups/protection-groups][%d] getApi24HostGroupsProtectionGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24HostGroupsProtectionGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24HostGroupsProtectionGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
