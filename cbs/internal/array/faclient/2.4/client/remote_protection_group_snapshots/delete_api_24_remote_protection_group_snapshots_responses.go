// Code generated by go-swagger; DO NOT EDIT.

package remote_protection_group_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// DeleteAPI24RemoteProtectionGroupSnapshotsReader is a Reader for the DeleteAPI24RemoteProtectionGroupSnapshots structure.
type DeleteAPI24RemoteProtectionGroupSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24RemoteProtectionGroupSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24RemoteProtectionGroupSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24RemoteProtectionGroupSnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24RemoteProtectionGroupSnapshotsOK creates a DeleteApi24RemoteProtectionGroupSnapshotsOK with default headers values
func NewDeleteApi24RemoteProtectionGroupSnapshotsOK() *DeleteApi24RemoteProtectionGroupSnapshotsOK {
	return &DeleteApi24RemoteProtectionGroupSnapshotsOK{}
}

/*DeleteApi24RemoteProtectionGroupSnapshotsOK handles this case with default header values.

OK
*/
type DeleteApi24RemoteProtectionGroupSnapshotsOK struct {
}

func (o *DeleteApi24RemoteProtectionGroupSnapshotsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/remote-protection-group-snapshots][%d] deleteApi24RemoteProtectionGroupSnapshotsOK ", 200)
}

func (o *DeleteApi24RemoteProtectionGroupSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24RemoteProtectionGroupSnapshotsBadRequest creates a DeleteApi24RemoteProtectionGroupSnapshotsBadRequest with default headers values
func NewDeleteApi24RemoteProtectionGroupSnapshotsBadRequest() *DeleteApi24RemoteProtectionGroupSnapshotsBadRequest {
	return &DeleteApi24RemoteProtectionGroupSnapshotsBadRequest{}
}

/*DeleteApi24RemoteProtectionGroupSnapshotsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24RemoteProtectionGroupSnapshotsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24RemoteProtectionGroupSnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/remote-protection-group-snapshots][%d] deleteApi24RemoteProtectionGroupSnapshotsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24RemoteProtectionGroupSnapshotsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24RemoteProtectionGroupSnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}