// Code generated by go-swagger; DO NOT EDIT.

package remote_protection_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24RemoteProtectionGroupsReader is a Reader for the PatchAPI24RemoteProtectionGroups structure.
type PatchAPI24RemoteProtectionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24RemoteProtectionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24RemoteProtectionGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24RemoteProtectionGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24RemoteProtectionGroupsOK creates a PatchApi24RemoteProtectionGroupsOK with default headers values
func NewPatchApi24RemoteProtectionGroupsOK() *PatchApi24RemoteProtectionGroupsOK {
	return &PatchApi24RemoteProtectionGroupsOK{}
}

/*PatchApi24RemoteProtectionGroupsOK handles this case with default header values.

OK
*/
type PatchApi24RemoteProtectionGroupsOK struct {
	Payload *models.RemoteProtectionGroupResponse
}

func (o *PatchApi24RemoteProtectionGroupsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/remote-protection-groups][%d] patchApi24RemoteProtectionGroupsOK  %+v", 200, o.Payload)
}

func (o *PatchApi24RemoteProtectionGroupsOK) GetPayload() *models.RemoteProtectionGroupResponse {
	return o.Payload
}

func (o *PatchApi24RemoteProtectionGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProtectionGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24RemoteProtectionGroupsBadRequest creates a PatchApi24RemoteProtectionGroupsBadRequest with default headers values
func NewPatchApi24RemoteProtectionGroupsBadRequest() *PatchApi24RemoteProtectionGroupsBadRequest {
	return &PatchApi24RemoteProtectionGroupsBadRequest{}
}

/*PatchApi24RemoteProtectionGroupsBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24RemoteProtectionGroupsBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24RemoteProtectionGroupsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/remote-protection-groups][%d] patchApi24RemoteProtectionGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24RemoteProtectionGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24RemoteProtectionGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
