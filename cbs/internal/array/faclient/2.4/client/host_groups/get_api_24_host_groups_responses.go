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

// GetAPI24HostGroupsReader is a Reader for the GetAPI24HostGroups structure.
type GetAPI24HostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24HostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24HostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24HostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24HostGroupsOK creates a GetApi24HostGroupsOK with default headers values
func NewGetApi24HostGroupsOK() *GetApi24HostGroupsOK {
	return &GetApi24HostGroupsOK{}
}

/*GetApi24HostGroupsOK handles this case with default header values.

OK
*/
type GetApi24HostGroupsOK struct {
	Payload *models.HostGroupGetResponse
}

func (o *GetApi24HostGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/host-groups][%d] getApi24HostGroupsOK  %+v", 200, o.Payload)
}

func (o *GetApi24HostGroupsOK) GetPayload() *models.HostGroupGetResponse {
	return o.Payload
}

func (o *GetApi24HostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostGroupGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24HostGroupsBadRequest creates a GetApi24HostGroupsBadRequest with default headers values
func NewGetApi24HostGroupsBadRequest() *GetApi24HostGroupsBadRequest {
	return &GetApi24HostGroupsBadRequest{}
}

/*GetApi24HostGroupsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24HostGroupsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24HostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/host-groups][%d] getApi24HostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24HostGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24HostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
