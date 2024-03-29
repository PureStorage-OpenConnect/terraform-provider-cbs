// Code generated by go-swagger; DO NOT EDIT.

package protection_group_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PostAPI24ProtectionGroupSnapshotsReader is a Reader for the PostAPI24ProtectionGroupSnapshots structure.
type PostAPI24ProtectionGroupSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24ProtectionGroupSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24ProtectionGroupSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24ProtectionGroupSnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24ProtectionGroupSnapshotsOK creates a PostApi24ProtectionGroupSnapshotsOK with default headers values
func NewPostApi24ProtectionGroupSnapshotsOK() *PostApi24ProtectionGroupSnapshotsOK {
	return &PostApi24ProtectionGroupSnapshotsOK{}
}

/*PostApi24ProtectionGroupSnapshotsOK handles this case with default header values.

OK
*/
type PostApi24ProtectionGroupSnapshotsOK struct {
	Payload *models.ProtectionGroupSnapshotResponse
}

func (o *PostApi24ProtectionGroupSnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/protection-group-snapshots][%d] postApi24ProtectionGroupSnapshotsOK  %+v", 200, o.Payload)
}

func (o *PostApi24ProtectionGroupSnapshotsOK) GetPayload() *models.ProtectionGroupSnapshotResponse {
	return o.Payload
}

func (o *PostApi24ProtectionGroupSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionGroupSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24ProtectionGroupSnapshotsBadRequest creates a PostApi24ProtectionGroupSnapshotsBadRequest with default headers values
func NewPostApi24ProtectionGroupSnapshotsBadRequest() *PostApi24ProtectionGroupSnapshotsBadRequest {
	return &PostApi24ProtectionGroupSnapshotsBadRequest{}
}

/*PostApi24ProtectionGroupSnapshotsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24ProtectionGroupSnapshotsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24ProtectionGroupSnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/protection-group-snapshots][%d] postApi24ProtectionGroupSnapshotsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24ProtectionGroupSnapshotsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24ProtectionGroupSnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
