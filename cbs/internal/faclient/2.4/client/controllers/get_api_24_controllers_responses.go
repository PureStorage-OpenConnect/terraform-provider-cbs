// Code generated by go-swagger; DO NOT EDIT.

package controllers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24ControllersReader is a Reader for the GetAPI24Controllers structure.
type GetAPI24ControllersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24ControllersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24ControllersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24ControllersOK creates a GetApi24ControllersOK with default headers values
func NewGetApi24ControllersOK() *GetApi24ControllersOK {
	return &GetApi24ControllersOK{}
}

/*GetApi24ControllersOK handles this case with default header values.

OK
*/
type GetApi24ControllersOK struct {
	Payload *models.ControllerGetResponse
}

func (o *GetApi24ControllersOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/controllers][%d] getApi24ControllersOK  %+v", 200, o.Payload)
}

func (o *GetApi24ControllersOK) GetPayload() *models.ControllerGetResponse {
	return o.Payload
}

func (o *GetApi24ControllersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControllerGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}