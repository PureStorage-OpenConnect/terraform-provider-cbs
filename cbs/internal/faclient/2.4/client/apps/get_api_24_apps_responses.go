// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24AppsReader is a Reader for the GetAPI24Apps structure.
type GetAPI24AppsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24AppsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24AppsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24AppsOK creates a GetApi24AppsOK with default headers values
func NewGetApi24AppsOK() *GetApi24AppsOK {
	return &GetApi24AppsOK{}
}

/*GetApi24AppsOK handles this case with default header values.

OK
*/
type GetApi24AppsOK struct {
	Payload *models.AppGetResponse
}

func (o *GetApi24AppsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/apps][%d] getApi24AppsOK  %+v", 200, o.Payload)
}

func (o *GetApi24AppsOK) GetPayload() *models.AppGetResponse {
	return o.Payload
}

func (o *GetApi24AppsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}