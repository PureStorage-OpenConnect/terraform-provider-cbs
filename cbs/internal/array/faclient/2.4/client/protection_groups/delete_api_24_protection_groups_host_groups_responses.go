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

// DeleteAPI24ProtectionGroupsHostGroupsReader is a Reader for the DeleteAPI24ProtectionGroupsHostGroups structure.
type DeleteAPI24ProtectionGroupsHostGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24ProtectionGroupsHostGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24ProtectionGroupsHostGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24ProtectionGroupsHostGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24ProtectionGroupsHostGroupsOK creates a DeleteApi24ProtectionGroupsHostGroupsOK with default headers values
func NewDeleteApi24ProtectionGroupsHostGroupsOK() *DeleteApi24ProtectionGroupsHostGroupsOK {
	return &DeleteApi24ProtectionGroupsHostGroupsOK{}
}

/*DeleteApi24ProtectionGroupsHostGroupsOK handles this case with default header values.

OK
*/
type DeleteApi24ProtectionGroupsHostGroupsOK struct {
}

func (o *DeleteApi24ProtectionGroupsHostGroupsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/protection-groups/host-groups][%d] deleteApi24ProtectionGroupsHostGroupsOK ", 200)
}

func (o *DeleteApi24ProtectionGroupsHostGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24ProtectionGroupsHostGroupsBadRequest creates a DeleteApi24ProtectionGroupsHostGroupsBadRequest with default headers values
func NewDeleteApi24ProtectionGroupsHostGroupsBadRequest() *DeleteApi24ProtectionGroupsHostGroupsBadRequest {
	return &DeleteApi24ProtectionGroupsHostGroupsBadRequest{}
}

/*DeleteApi24ProtectionGroupsHostGroupsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24ProtectionGroupsHostGroupsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24ProtectionGroupsHostGroupsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/protection-groups/host-groups][%d] deleteApi24ProtectionGroupsHostGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24ProtectionGroupsHostGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24ProtectionGroupsHostGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
