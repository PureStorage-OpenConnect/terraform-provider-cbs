// Code generated by go-swagger; DO NOT EDIT.

package directories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24DirectoriesReader is a Reader for the PatchAPI24Directories structure.
type PatchAPI24DirectoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24DirectoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24DirectoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24DirectoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24DirectoriesOK creates a PatchApi24DirectoriesOK with default headers values
func NewPatchApi24DirectoriesOK() *PatchApi24DirectoriesOK {
	return &PatchApi24DirectoriesOK{}
}

/*PatchApi24DirectoriesOK handles this case with default header values.

OK
*/
type PatchApi24DirectoriesOK struct {
	Payload *models.DirectoryResponse
}

func (o *PatchApi24DirectoriesOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/directories][%d] patchApi24DirectoriesOK  %+v", 200, o.Payload)
}

func (o *PatchApi24DirectoriesOK) GetPayload() *models.DirectoryResponse {
	return o.Payload
}

func (o *PatchApi24DirectoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24DirectoriesBadRequest creates a PatchApi24DirectoriesBadRequest with default headers values
func NewPatchApi24DirectoriesBadRequest() *PatchApi24DirectoriesBadRequest {
	return &PatchApi24DirectoriesBadRequest{}
}

/*PatchApi24DirectoriesBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24DirectoriesBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24DirectoriesBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/directories][%d] patchApi24DirectoriesBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24DirectoriesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24DirectoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
