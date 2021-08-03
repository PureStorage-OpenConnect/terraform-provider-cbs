// Code generated by go-swagger; DO NOT EDIT.

package software

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24SoftwareInstallationsReader is a Reader for the GetAPI24SoftwareInstallations structure.
type GetAPI24SoftwareInstallationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24SoftwareInstallationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24SoftwareInstallationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24SoftwareInstallationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24SoftwareInstallationsOK creates a GetApi24SoftwareInstallationsOK with default headers values
func NewGetApi24SoftwareInstallationsOK() *GetApi24SoftwareInstallationsOK {
	return &GetApi24SoftwareInstallationsOK{}
}

/*GetApi24SoftwareInstallationsOK handles this case with default header values.

OK
*/
type GetApi24SoftwareInstallationsOK struct {
	Payload *models.SoftwareInstallationsGetResponse
}

func (o *GetApi24SoftwareInstallationsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/software-installations][%d] getApi24SoftwareInstallationsOK  %+v", 200, o.Payload)
}

func (o *GetApi24SoftwareInstallationsOK) GetPayload() *models.SoftwareInstallationsGetResponse {
	return o.Payload
}

func (o *GetApi24SoftwareInstallationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwareInstallationsGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24SoftwareInstallationsBadRequest creates a GetApi24SoftwareInstallationsBadRequest with default headers values
func NewGetApi24SoftwareInstallationsBadRequest() *GetApi24SoftwareInstallationsBadRequest {
	return &GetApi24SoftwareInstallationsBadRequest{}
}

/*GetApi24SoftwareInstallationsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24SoftwareInstallationsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24SoftwareInstallationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/software-installations][%d] getApi24SoftwareInstallationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24SoftwareInstallationsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24SoftwareInstallationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}