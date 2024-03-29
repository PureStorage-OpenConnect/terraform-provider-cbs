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

// DeleteAPI24HostGroupsReader is a Reader for the DeleteAPI24HostGroups structure.
type DeleteAPI24HostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24HostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24HostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24HostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24HostGroupsOK creates a DeleteApi24HostGroupsOK with default headers values
func NewDeleteApi24HostGroupsOK() *DeleteApi24HostGroupsOK {
	return &DeleteApi24HostGroupsOK{}
}

/*DeleteApi24HostGroupsOK handles this case with default header values.

OK
*/
type DeleteApi24HostGroupsOK struct {
}

func (o *DeleteApi24HostGroupsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/host-groups][%d] deleteApi24HostGroupsOK ", 200)
}

func (o *DeleteApi24HostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24HostGroupsBadRequest creates a DeleteApi24HostGroupsBadRequest with default headers values
func NewDeleteApi24HostGroupsBadRequest() *DeleteApi24HostGroupsBadRequest {
	return &DeleteApi24HostGroupsBadRequest{}
}

/*DeleteApi24HostGroupsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24HostGroupsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24HostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/host-groups][%d] deleteApi24HostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24HostGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24HostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
