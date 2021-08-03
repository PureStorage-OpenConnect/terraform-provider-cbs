// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PostAPI24PoliciesSnapshotReader is a Reader for the PostAPI24PoliciesSnapshot structure.
type PostAPI24PoliciesSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24PoliciesSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24PoliciesSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24PoliciesSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24PoliciesSnapshotOK creates a PostApi24PoliciesSnapshotOK with default headers values
func NewPostApi24PoliciesSnapshotOK() *PostApi24PoliciesSnapshotOK {
	return &PostApi24PoliciesSnapshotOK{}
}

/*PostApi24PoliciesSnapshotOK handles this case with default header values.

OK
*/
type PostApi24PoliciesSnapshotOK struct {
	Payload *models.PolicyResponse
}

func (o *PostApi24PoliciesSnapshotOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/policies/snapshot][%d] postApi24PoliciesSnapshotOK  %+v", 200, o.Payload)
}

func (o *PostApi24PoliciesSnapshotOK) GetPayload() *models.PolicyResponse {
	return o.Payload
}

func (o *PostApi24PoliciesSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24PoliciesSnapshotBadRequest creates a PostApi24PoliciesSnapshotBadRequest with default headers values
func NewPostApi24PoliciesSnapshotBadRequest() *PostApi24PoliciesSnapshotBadRequest {
	return &PostApi24PoliciesSnapshotBadRequest{}
}

/*PostApi24PoliciesSnapshotBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24PoliciesSnapshotBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24PoliciesSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/policies/snapshot][%d] postApi24PoliciesSnapshotBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24PoliciesSnapshotBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24PoliciesSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}