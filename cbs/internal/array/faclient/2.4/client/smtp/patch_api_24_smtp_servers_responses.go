// Code generated by go-swagger; DO NOT EDIT.

package smtp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24SMTPServersReader is a Reader for the PatchAPI24SMTPServers structure.
type PatchAPI24SMTPServersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24SMTPServersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24SMTPServersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24SMTPServersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24SMTPServersOK creates a PatchApi24SMTPServersOK with default headers values
func NewPatchApi24SMTPServersOK() *PatchApi24SMTPServersOK {
	return &PatchApi24SMTPServersOK{}
}

/*PatchApi24SMTPServersOK handles this case with default header values.

OK
*/
type PatchApi24SMTPServersOK struct {
	Payload *models.SMTPServerResponse
}

func (o *PatchApi24SMTPServersOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/smtp-servers][%d] patchApi24SmtpServersOK  %+v", 200, o.Payload)
}

func (o *PatchApi24SMTPServersOK) GetPayload() *models.SMTPServerResponse {
	return o.Payload
}

func (o *PatchApi24SMTPServersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SMTPServerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24SMTPServersBadRequest creates a PatchApi24SMTPServersBadRequest with default headers values
func NewPatchApi24SMTPServersBadRequest() *PatchApi24SMTPServersBadRequest {
	return &PatchApi24SMTPServersBadRequest{}
}

/*PatchApi24SMTPServersBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24SMTPServersBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24SMTPServersBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/smtp-servers][%d] patchApi24SmtpServersBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24SMTPServersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24SMTPServersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}