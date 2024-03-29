// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24PoliciesNfsReader is a Reader for the PatchAPI24PoliciesNfs structure.
type PatchAPI24PoliciesNfsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24PoliciesNfsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24PoliciesNfsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24PoliciesNfsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24PoliciesNfsOK creates a PatchApi24PoliciesNfsOK with default headers values
func NewPatchApi24PoliciesNfsOK() *PatchApi24PoliciesNfsOK {
	return &PatchApi24PoliciesNfsOK{}
}

/*PatchApi24PoliciesNfsOK handles this case with default header values.

OK
*/
type PatchApi24PoliciesNfsOK struct {
	Payload *models.PolicyResponse
}

func (o *PatchApi24PoliciesNfsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/policies/nfs][%d] patchApi24PoliciesNfsOK  %+v", 200, o.Payload)
}

func (o *PatchApi24PoliciesNfsOK) GetPayload() *models.PolicyResponse {
	return o.Payload
}

func (o *PatchApi24PoliciesNfsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24PoliciesNfsBadRequest creates a PatchApi24PoliciesNfsBadRequest with default headers values
func NewPatchApi24PoliciesNfsBadRequest() *PatchApi24PoliciesNfsBadRequest {
	return &PatchApi24PoliciesNfsBadRequest{}
}

/*PatchApi24PoliciesNfsBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24PoliciesNfsBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24PoliciesNfsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/policies/nfs][%d] patchApi24PoliciesNfsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24PoliciesNfsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24PoliciesNfsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
