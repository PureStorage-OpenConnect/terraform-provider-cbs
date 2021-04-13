// Code generated by go-swagger; DO NOT EDIT.

package subnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PatchAPI24SubnetsReader is a Reader for the PatchAPI24Subnets structure.
type PatchAPI24SubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24SubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24SubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24SubnetsOK creates a PatchApi24SubnetsOK with default headers values
func NewPatchApi24SubnetsOK() *PatchApi24SubnetsOK {
	return &PatchApi24SubnetsOK{}
}

/*PatchApi24SubnetsOK handles this case with default header values.

Displays the updated subnet.
*/
type PatchApi24SubnetsOK struct {
	Payload *models.SubnetResponse
}

func (o *PatchApi24SubnetsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/subnets][%d] patchApi24SubnetsOK  %+v", 200, o.Payload)
}

func (o *PatchApi24SubnetsOK) GetPayload() *models.SubnetResponse {
	return o.Payload
}

func (o *PatchApi24SubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubnetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}