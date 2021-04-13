// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24VolumesReader is a Reader for the PostAPI24Volumes structure.
type PostAPI24VolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24VolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24VolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24VolumesOK creates a PostApi24VolumesOK with default headers values
func NewPostApi24VolumesOK() *PostApi24VolumesOK {
	return &PostApi24VolumesOK{}
}

/*PostApi24VolumesOK handles this case with default header values.

OK
*/
type PostApi24VolumesOK struct {
	Payload *models.VolumeResponse
}

func (o *PostApi24VolumesOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/volumes][%d] postApi24VolumesOK  %+v", 200, o.Payload)
}

func (o *PostApi24VolumesOK) GetPayload() *models.VolumeResponse {
	return o.Payload
}

func (o *PostApi24VolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
