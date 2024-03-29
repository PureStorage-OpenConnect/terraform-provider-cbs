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

// DeleteAPI24ProtectionGroupsHostsReader is a Reader for the DeleteAPI24ProtectionGroupsHosts structure.
type DeleteAPI24ProtectionGroupsHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24ProtectionGroupsHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24ProtectionGroupsHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24ProtectionGroupsHostsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24ProtectionGroupsHostsOK creates a DeleteApi24ProtectionGroupsHostsOK with default headers values
func NewDeleteApi24ProtectionGroupsHostsOK() *DeleteApi24ProtectionGroupsHostsOK {
	return &DeleteApi24ProtectionGroupsHostsOK{}
}

/*DeleteApi24ProtectionGroupsHostsOK handles this case with default header values.

OK
*/
type DeleteApi24ProtectionGroupsHostsOK struct {
}

func (o *DeleteApi24ProtectionGroupsHostsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/protection-groups/hosts][%d] deleteApi24ProtectionGroupsHostsOK ", 200)
}

func (o *DeleteApi24ProtectionGroupsHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24ProtectionGroupsHostsBadRequest creates a DeleteApi24ProtectionGroupsHostsBadRequest with default headers values
func NewDeleteApi24ProtectionGroupsHostsBadRequest() *DeleteApi24ProtectionGroupsHostsBadRequest {
	return &DeleteApi24ProtectionGroupsHostsBadRequest{}
}

/*DeleteApi24ProtectionGroupsHostsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24ProtectionGroupsHostsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24ProtectionGroupsHostsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/protection-groups/hosts][%d] deleteApi24ProtectionGroupsHostsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24ProtectionGroupsHostsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24ProtectionGroupsHostsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
