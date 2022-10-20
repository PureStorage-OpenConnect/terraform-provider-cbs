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

// PatchAPI24ProtectionGroupsTargetsReader is a Reader for the PatchAPI24ProtectionGroupsTargets structure.
type PatchAPI24ProtectionGroupsTargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24ProtectionGroupsTargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24ProtectionGroupsTargetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24ProtectionGroupsTargetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24ProtectionGroupsTargetsOK creates a PatchApi24ProtectionGroupsTargetsOK with default headers values
func NewPatchApi24ProtectionGroupsTargetsOK() *PatchApi24ProtectionGroupsTargetsOK {
	return &PatchApi24ProtectionGroupsTargetsOK{}
}

/*PatchApi24ProtectionGroupsTargetsOK handles this case with default header values.

OK
*/
type PatchApi24ProtectionGroupsTargetsOK struct {
	Payload *models.ProtectionGroupTargetResponse
}

func (o *PatchApi24ProtectionGroupsTargetsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/protection-groups/targets][%d] patchApi24ProtectionGroupsTargetsOK  %+v", 200, o.Payload)
}

func (o *PatchApi24ProtectionGroupsTargetsOK) GetPayload() *models.ProtectionGroupTargetResponse {
	return o.Payload
}

func (o *PatchApi24ProtectionGroupsTargetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionGroupTargetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24ProtectionGroupsTargetsBadRequest creates a PatchApi24ProtectionGroupsTargetsBadRequest with default headers values
func NewPatchApi24ProtectionGroupsTargetsBadRequest() *PatchApi24ProtectionGroupsTargetsBadRequest {
	return &PatchApi24ProtectionGroupsTargetsBadRequest{}
}

/*PatchApi24ProtectionGroupsTargetsBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24ProtectionGroupsTargetsBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24ProtectionGroupsTargetsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/protection-groups/targets][%d] patchApi24ProtectionGroupsTargetsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24ProtectionGroupsTargetsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24ProtectionGroupsTargetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
