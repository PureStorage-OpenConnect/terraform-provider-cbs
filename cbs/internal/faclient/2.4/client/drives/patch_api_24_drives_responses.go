// Code generated by go-swagger; DO NOT EDIT.

package drives

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PatchAPI24DrivesReader is a Reader for the PatchAPI24Drives structure.
type PatchAPI24DrivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24DrivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24DrivesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24DrivesOK creates a PatchApi24DrivesOK with default headers values
func NewPatchApi24DrivesOK() *PatchApi24DrivesOK {
	return &PatchApi24DrivesOK{}
}

/*PatchApi24DrivesOK handles this case with default header values.

OK
*/
type PatchApi24DrivesOK struct {
	Payload *models.DriveResponse
}

func (o *PatchApi24DrivesOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/drives][%d] patchApi24DrivesOK  %+v", 200, o.Payload)
}

func (o *PatchApi24DrivesOK) GetPayload() *models.DriveResponse {
	return o.Payload
}

func (o *PatchApi24DrivesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DriveResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}