// Code generated by go-swagger; DO NOT EDIT.

package host_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24HostGroupsHostsReader is a Reader for the GetAPI24HostGroupsHosts structure.
type GetAPI24HostGroupsHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24HostGroupsHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24HostGroupsHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24HostGroupsHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24HostGroupsHostsOK creates a GetApi24HostGroupsHostsOK with default headers values
func NewGetApi24HostGroupsHostsOK() *GetApi24HostGroupsHostsOK {
	return &GetApi24HostGroupsHostsOK{}
}

/*GetApi24HostGroupsHostsOK handles this case with default header values.

OK
*/
type GetApi24HostGroupsHostsOK struct {
	Payload *models.MemberNoIDAllGetResponse
}

func (o *GetApi24HostGroupsHostsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/host-groups/hosts][%d] getApi24HostGroupsHostsOK  %+v", 200, o.Payload)
}

func (o *GetApi24HostGroupsHostsOK) GetPayload() *models.MemberNoIDAllGetResponse {
	return o.Payload
}

func (o *GetApi24HostGroupsHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24HostGroupsHostsBadRequest creates a GetApi24HostGroupsHostsBadRequest with default headers values
func NewGetApi24HostGroupsHostsBadRequest() *GetApi24HostGroupsHostsBadRequest {
	return &GetApi24HostGroupsHostsBadRequest{}
}

/*GetApi24HostGroupsHostsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24HostGroupsHostsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24HostGroupsHostsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/host-groups/hosts][%d] getApi24HostGroupsHostsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24HostGroupsHostsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24HostGroupsHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}