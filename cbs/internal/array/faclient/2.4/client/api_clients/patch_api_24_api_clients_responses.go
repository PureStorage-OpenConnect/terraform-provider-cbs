// Code generated by go-swagger; DO NOT EDIT.

package api_clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24APIClientsReader is a Reader for the PatchAPI24APIClients structure.
type PatchAPI24APIClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24APIClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24APIClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24APIClientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24APIClientsOK creates a PatchApi24APIClientsOK with default headers values
func NewPatchApi24APIClientsOK() *PatchApi24APIClientsOK {
	return &PatchApi24APIClientsOK{}
}

/*PatchApi24APIClientsOK handles this case with default header values.

OK
*/
type PatchApi24APIClientsOK struct {
	Payload *models.APIClientResponse
}

func (o *PatchApi24APIClientsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/api-clients][%d] patchApi24ApiClientsOK  %+v", 200, o.Payload)
}

func (o *PatchApi24APIClientsOK) GetPayload() *models.APIClientResponse {
	return o.Payload
}

func (o *PatchApi24APIClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIClientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24APIClientsBadRequest creates a PatchApi24APIClientsBadRequest with default headers values
func NewPatchApi24APIClientsBadRequest() *PatchApi24APIClientsBadRequest {
	return &PatchApi24APIClientsBadRequest{}
}

/*PatchApi24APIClientsBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24APIClientsBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24APIClientsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/api-clients][%d] patchApi24ApiClientsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24APIClientsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24APIClientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
