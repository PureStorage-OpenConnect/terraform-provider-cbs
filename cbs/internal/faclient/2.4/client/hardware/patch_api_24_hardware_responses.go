// Code generated by go-swagger; DO NOT EDIT.

package hardware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PatchAPI24HardwareReader is a Reader for the PatchAPI24Hardware structure.
type PatchAPI24HardwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24HardwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24HardwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24HardwareOK creates a PatchApi24HardwareOK with default headers values
func NewPatchApi24HardwareOK() *PatchApi24HardwareOK {
	return &PatchApi24HardwareOK{}
}

/*PatchApi24HardwareOK handles this case with default header values.

OK
*/
type PatchApi24HardwareOK struct {
	Payload *models.HardwareResponse
}

func (o *PatchApi24HardwareOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/hardware][%d] patchApi24HardwareOK  %+v", 200, o.Payload)
}

func (o *PatchApi24HardwareOK) GetPayload() *models.HardwareResponse {
	return o.Payload
}

func (o *PatchApi24HardwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HardwareResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
