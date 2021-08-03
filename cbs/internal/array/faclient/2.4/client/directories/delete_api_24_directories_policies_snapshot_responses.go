// Code generated by go-swagger; DO NOT EDIT.

package directories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// DeleteAPI24DirectoriesPoliciesSnapshotReader is a Reader for the DeleteAPI24DirectoriesPoliciesSnapshot structure.
type DeleteAPI24DirectoriesPoliciesSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24DirectoriesPoliciesSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24DirectoriesPoliciesSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24DirectoriesPoliciesSnapshotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24DirectoriesPoliciesSnapshotOK creates a DeleteApi24DirectoriesPoliciesSnapshotOK with default headers values
func NewDeleteApi24DirectoriesPoliciesSnapshotOK() *DeleteApi24DirectoriesPoliciesSnapshotOK {
	return &DeleteApi24DirectoriesPoliciesSnapshotOK{}
}

/*DeleteApi24DirectoriesPoliciesSnapshotOK handles this case with default header values.

OK
*/
type DeleteApi24DirectoriesPoliciesSnapshotOK struct {
}

func (o *DeleteApi24DirectoriesPoliciesSnapshotOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/directories/policies/snapshot][%d] deleteApi24DirectoriesPoliciesSnapshotOK ", 200)
}

func (o *DeleteApi24DirectoriesPoliciesSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24DirectoriesPoliciesSnapshotBadRequest creates a DeleteApi24DirectoriesPoliciesSnapshotBadRequest with default headers values
func NewDeleteApi24DirectoriesPoliciesSnapshotBadRequest() *DeleteApi24DirectoriesPoliciesSnapshotBadRequest {
	return &DeleteApi24DirectoriesPoliciesSnapshotBadRequest{}
}

/*DeleteApi24DirectoriesPoliciesSnapshotBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24DirectoriesPoliciesSnapshotBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24DirectoriesPoliciesSnapshotBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/directories/policies/snapshot][%d] deleteApi24DirectoriesPoliciesSnapshotBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24DirectoriesPoliciesSnapshotBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24DirectoriesPoliciesSnapshotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}