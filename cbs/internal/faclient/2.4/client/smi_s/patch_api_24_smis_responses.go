// Code generated by go-swagger; DO NOT EDIT.

package smi_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PatchAPI24SMISReader is a Reader for the PatchAPI24SMIS structure.
type PatchAPI24SMISReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24SMISReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24SMISOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24SMISOK creates a PatchApi24SMISOK with default headers values
func NewPatchApi24SMISOK() *PatchApi24SMISOK {
	return &PatchApi24SMISOK{}
}

/*PatchApi24SMISOK handles this case with default header values.

OK
*/
type PatchApi24SMISOK struct {
	Payload *models.SMISResponse
}

func (o *PatchApi24SMISOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/smi-s][%d] patchApi24SmisOK  %+v", 200, o.Payload)
}

func (o *PatchApi24SMISOK) GetPayload() *models.SMISResponse {
	return o.Payload
}

func (o *PatchApi24SMISOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SMISResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}