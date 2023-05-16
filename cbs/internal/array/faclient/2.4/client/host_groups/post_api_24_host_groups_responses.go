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

// PostAPI24HostGroupsReader is a Reader for the PostAPI24HostGroups structure.
type PostAPI24HostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24HostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24HostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24HostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24HostGroupsOK creates a PostApi24HostGroupsOK with default headers values
func NewPostApi24HostGroupsOK() *PostApi24HostGroupsOK {
	return &PostApi24HostGroupsOK{}
}

/*PostApi24HostGroupsOK handles this case with default header values.

OK
*/
type PostApi24HostGroupsOK struct {
	Payload *models.HostGroupResponse
}

func (o *PostApi24HostGroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/host-groups][%d] postApi24HostGroupsOK  %+v", 200, o.Payload)
}

func (o *PostApi24HostGroupsOK) GetPayload() *models.HostGroupResponse {
	return o.Payload
}

func (o *PostApi24HostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24HostGroupsBadRequest creates a PostApi24HostGroupsBadRequest with default headers values
func NewPostApi24HostGroupsBadRequest() *PostApi24HostGroupsBadRequest {
	return &PostApi24HostGroupsBadRequest{}
}

/*PostApi24HostGroupsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24HostGroupsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24HostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/host-groups][%d] postApi24HostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24HostGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24HostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
