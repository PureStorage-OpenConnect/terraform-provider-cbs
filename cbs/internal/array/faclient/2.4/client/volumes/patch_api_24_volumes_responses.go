// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24VolumesReader is a Reader for the PatchAPI24Volumes structure.
type PatchAPI24VolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24VolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24VolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24VolumesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24VolumesOK creates a PatchApi24VolumesOK with default headers values
func NewPatchApi24VolumesOK() *PatchApi24VolumesOK {
	return &PatchApi24VolumesOK{}
}

/*PatchApi24VolumesOK handles this case with default header values.

OK
*/
type PatchApi24VolumesOK struct {
	Payload *models.VolumeResponse
}

func (o *PatchApi24VolumesOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/volumes][%d] patchApi24VolumesOK  %+v", 200, o.Payload)
}

func (o *PatchApi24VolumesOK) GetPayload() *models.VolumeResponse {
	return o.Payload
}

func (o *PatchApi24VolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24VolumesBadRequest creates a PatchApi24VolumesBadRequest with default headers values
func NewPatchApi24VolumesBadRequest() *PatchApi24VolumesBadRequest {
	return &PatchApi24VolumesBadRequest{}
}

/*PatchApi24VolumesBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24VolumesBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24VolumesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/volumes][%d] patchApi24VolumesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24VolumesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24VolumesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
