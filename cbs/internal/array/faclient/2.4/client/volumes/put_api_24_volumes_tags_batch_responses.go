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

// PutAPI24VolumesTagsBatchReader is a Reader for the PutAPI24VolumesTagsBatch structure.
type PutAPI24VolumesTagsBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPI24VolumesTagsBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutApi24VolumesTagsBatchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutApi24VolumesTagsBatchBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutApi24VolumesTagsBatchOK creates a PutApi24VolumesTagsBatchOK with default headers values
func NewPutApi24VolumesTagsBatchOK() *PutApi24VolumesTagsBatchOK {
	return &PutApi24VolumesTagsBatchOK{}
}

/*PutApi24VolumesTagsBatchOK handles this case with default header values.

OK
*/
type PutApi24VolumesTagsBatchOK struct {
	Payload *models.TagResponse
}

func (o *PutApi24VolumesTagsBatchOK) Error() string {
	return fmt.Sprintf("[PUT /api/2.4/volumes/tags/batch][%d] putApi24VolumesTagsBatchOK  %+v", 200, o.Payload)
}

func (o *PutApi24VolumesTagsBatchOK) GetPayload() *models.TagResponse {
	return o.Payload
}

func (o *PutApi24VolumesTagsBatchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApi24VolumesTagsBatchBadRequest creates a PutApi24VolumesTagsBatchBadRequest with default headers values
func NewPutApi24VolumesTagsBatchBadRequest() *PutApi24VolumesTagsBatchBadRequest {
	return &PutApi24VolumesTagsBatchBadRequest{}
}

/*PutApi24VolumesTagsBatchBadRequest handles this case with default header values.

BadRequest
*/
type PutApi24VolumesTagsBatchBadRequest struct {
	Payload *models.Error
}

func (o *PutApi24VolumesTagsBatchBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/2.4/volumes/tags/batch][%d] putApi24VolumesTagsBatchBadRequest  %+v", 400, o.Payload)
}

func (o *PutApi24VolumesTagsBatchBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutApi24VolumesTagsBatchBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
