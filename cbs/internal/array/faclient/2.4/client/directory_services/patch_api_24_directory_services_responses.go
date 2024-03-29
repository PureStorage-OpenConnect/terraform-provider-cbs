// Code generated by go-swagger; DO NOT EDIT.

package directory_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24DirectoryServicesReader is a Reader for the PatchAPI24DirectoryServices structure.
type PatchAPI24DirectoryServicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24DirectoryServicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24DirectoryServicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24DirectoryServicesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24DirectoryServicesOK creates a PatchApi24DirectoryServicesOK with default headers values
func NewPatchApi24DirectoryServicesOK() *PatchApi24DirectoryServicesOK {
	return &PatchApi24DirectoryServicesOK{}
}

/*PatchApi24DirectoryServicesOK handles this case with default header values.

OK
*/
type PatchApi24DirectoryServicesOK struct {
	Payload *models.DirectoryServiceResponse
}

func (o *PatchApi24DirectoryServicesOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/directory-services][%d] patchApi24DirectoryServicesOK  %+v", 200, o.Payload)
}

func (o *PatchApi24DirectoryServicesOK) GetPayload() *models.DirectoryServiceResponse {
	return o.Payload
}

func (o *PatchApi24DirectoryServicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24DirectoryServicesBadRequest creates a PatchApi24DirectoryServicesBadRequest with default headers values
func NewPatchApi24DirectoryServicesBadRequest() *PatchApi24DirectoryServicesBadRequest {
	return &PatchApi24DirectoryServicesBadRequest{}
}

/*PatchApi24DirectoryServicesBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24DirectoryServicesBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24DirectoryServicesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/directory-services][%d] patchApi24DirectoryServicesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24DirectoryServicesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24DirectoryServicesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
