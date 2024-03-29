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

// GetAPI24HostsHostGroupsReader is a Reader for the GetAPI24HostsHostGroups structure.
type GetAPI24HostsHostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24HostsHostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24HostsHostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24HostsHostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24HostsHostGroupsOK creates a GetApi24HostsHostGroupsOK with default headers values
func NewGetApi24HostsHostGroupsOK() *GetApi24HostsHostGroupsOK {
	return &GetApi24HostsHostGroupsOK{}
}

/*GetApi24HostsHostGroupsOK handles this case with default header values.

OK
*/
type GetApi24HostsHostGroupsOK struct {
	Payload *models.MemberNoIDAllGetResponse
}

func (o *GetApi24HostsHostGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/hosts/host-groups][%d] getApi24HostsHostGroupsOK  %+v", 200, o.Payload)
}

func (o *GetApi24HostsHostGroupsOK) GetPayload() *models.MemberNoIDAllGetResponse {
	return o.Payload
}

func (o *GetApi24HostsHostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24HostsHostGroupsBadRequest creates a GetApi24HostsHostGroupsBadRequest with default headers values
func NewGetApi24HostsHostGroupsBadRequest() *GetApi24HostsHostGroupsBadRequest {
	return &GetApi24HostsHostGroupsBadRequest{}
}

/*GetApi24HostsHostGroupsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24HostsHostGroupsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24HostsHostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/hosts/host-groups][%d] getApi24HostsHostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24HostsHostGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24HostsHostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
