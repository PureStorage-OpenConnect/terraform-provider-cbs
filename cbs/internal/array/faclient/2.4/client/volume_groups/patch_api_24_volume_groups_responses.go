// Code generated by go-swagger; DO NOT EDIT.

package volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24VolumeGroupsReader is a Reader for the PatchAPI24VolumeGroups structure.
type PatchAPI24VolumeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24VolumeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24VolumeGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24VolumeGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24VolumeGroupsOK creates a PatchApi24VolumeGroupsOK with default headers values
func NewPatchApi24VolumeGroupsOK() *PatchApi24VolumeGroupsOK {
	return &PatchApi24VolumeGroupsOK{}
}

/*PatchApi24VolumeGroupsOK handles this case with default header values.

OK
*/
type PatchApi24VolumeGroupsOK struct {
	Payload *models.VolumeGroupResponse
}

func (o *PatchApi24VolumeGroupsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/volume-groups][%d] patchApi24VolumeGroupsOK  %+v", 200, o.Payload)
}

func (o *PatchApi24VolumeGroupsOK) GetPayload() *models.VolumeGroupResponse {
	return o.Payload
}

func (o *PatchApi24VolumeGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24VolumeGroupsBadRequest creates a PatchApi24VolumeGroupsBadRequest with default headers values
func NewPatchApi24VolumeGroupsBadRequest() *PatchApi24VolumeGroupsBadRequest {
	return &PatchApi24VolumeGroupsBadRequest{}
}

/*PatchApi24VolumeGroupsBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24VolumeGroupsBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24VolumeGroupsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/volume-groups][%d] patchApi24VolumeGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24VolumeGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24VolumeGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
