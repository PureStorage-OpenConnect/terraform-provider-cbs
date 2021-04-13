// Code generated by go-swagger; DO NOT EDIT.

package offloads

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24OffloadsReader is a Reader for the PostAPI24Offloads structure.
type PostAPI24OffloadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24OffloadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24OffloadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24OffloadsOK creates a PostApi24OffloadsOK with default headers values
func NewPostApi24OffloadsOK() *PostApi24OffloadsOK {
	return &PostApi24OffloadsOK{}
}

/*PostApi24OffloadsOK handles this case with default header values.

OK
*/
type PostApi24OffloadsOK struct {
	Payload *models.OffloadResponse
}

func (o *PostApi24OffloadsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/offloads][%d] postApi24OffloadsOK  %+v", 200, o.Payload)
}

func (o *PostApi24OffloadsOK) GetPayload() *models.OffloadResponse {
	return o.Payload
}

func (o *PostApi24OffloadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OffloadResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}