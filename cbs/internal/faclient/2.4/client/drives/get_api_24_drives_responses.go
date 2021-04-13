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

// GetAPI24DrivesReader is a Reader for the GetAPI24Drives structure.
type GetAPI24DrivesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24DrivesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24DrivesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24DrivesOK creates a GetApi24DrivesOK with default headers values
func NewGetApi24DrivesOK() *GetApi24DrivesOK {
	return &GetApi24DrivesOK{}
}

/*GetApi24DrivesOK handles this case with default header values.

OK
*/
type GetApi24DrivesOK struct {
	Payload *models.DriveGetResponse
}

func (o *GetApi24DrivesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/drives][%d] getApi24DrivesOK  %+v", 200, o.Payload)
}

func (o *GetApi24DrivesOK) GetPayload() *models.DriveGetResponse {
	return o.Payload
}

func (o *GetApi24DrivesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DriveGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
