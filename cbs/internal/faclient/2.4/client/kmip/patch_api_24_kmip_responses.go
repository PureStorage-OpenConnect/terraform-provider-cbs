// Code generated by go-swagger; DO NOT EDIT.

package kmip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PatchAPI24KMIPReader is a Reader for the PatchAPI24KMIP structure.
type PatchAPI24KMIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24KMIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24KMIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24KMIPOK creates a PatchApi24KMIPOK with default headers values
func NewPatchApi24KMIPOK() *PatchApi24KMIPOK {
	return &PatchApi24KMIPOK{}
}

/*PatchApi24KMIPOK handles this case with default header values.

OK
*/
type PatchApi24KMIPOK struct {
	Payload *models.KMIPResponse
}

func (o *PatchApi24KMIPOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/kmip][%d] patchApi24KmipOK  %+v", 200, o.Payload)
}

func (o *PatchApi24KMIPOK) GetPayload() *models.KMIPResponse {
	return o.Payload
}

func (o *PatchApi24KMIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KMIPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}