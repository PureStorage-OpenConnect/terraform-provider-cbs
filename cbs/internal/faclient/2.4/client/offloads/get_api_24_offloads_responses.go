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

// GetAPI24OffloadsReader is a Reader for the GetAPI24Offloads structure.
type GetAPI24OffloadsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24OffloadsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24OffloadsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24OffloadsOK creates a GetApi24OffloadsOK with default headers values
func NewGetApi24OffloadsOK() *GetApi24OffloadsOK {
	return &GetApi24OffloadsOK{}
}

/*GetApi24OffloadsOK handles this case with default header values.

OK
*/
type GetApi24OffloadsOK struct {
	Payload *models.OffloadGetResponse
}

func (o *GetApi24OffloadsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/offloads][%d] getApi24OffloadsOK  %+v", 200, o.Payload)
}

func (o *GetApi24OffloadsOK) GetPayload() *models.OffloadGetResponse {
	return o.Payload
}

func (o *GetApi24OffloadsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OffloadGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
