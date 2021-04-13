// Code generated by go-swagger; DO NOT EDIT.

package software

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24SoftwareInstallationsReader is a Reader for the PostAPI24SoftwareInstallations structure.
type PostAPI24SoftwareInstallationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24SoftwareInstallationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24SoftwareInstallationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24SoftwareInstallationsOK creates a PostApi24SoftwareInstallationsOK with default headers values
func NewPostApi24SoftwareInstallationsOK() *PostApi24SoftwareInstallationsOK {
	return &PostApi24SoftwareInstallationsOK{}
}

/*PostApi24SoftwareInstallationsOK handles this case with default header values.

The software upgrade record was created successfully.
*/
type PostApi24SoftwareInstallationsOK struct {
	Payload *models.SoftwareInstallationsResponse
}

func (o *PostApi24SoftwareInstallationsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/software-installations][%d] postApi24SoftwareInstallationsOK  %+v", 200, o.Payload)
}

func (o *PostApi24SoftwareInstallationsOK) GetPayload() *models.SoftwareInstallationsResponse {
	return o.Payload
}

func (o *PostApi24SoftwareInstallationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwareInstallationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}