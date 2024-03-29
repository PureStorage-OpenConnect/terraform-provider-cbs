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

// PatchAPI24ProtectionGroupSnapshotsReader is a Reader for the PatchAPI24ProtectionGroupSnapshots structure.
type PatchAPI24ProtectionGroupSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24ProtectionGroupSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24ProtectionGroupSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24ProtectionGroupSnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24ProtectionGroupSnapshotsOK creates a PatchApi24ProtectionGroupSnapshotsOK with default headers values
func NewPatchApi24ProtectionGroupSnapshotsOK() *PatchApi24ProtectionGroupSnapshotsOK {
	return &PatchApi24ProtectionGroupSnapshotsOK{}
}

/*PatchApi24ProtectionGroupSnapshotsOK handles this case with default header values.

OK
*/
type PatchApi24ProtectionGroupSnapshotsOK struct {
	Payload *models.ProtectionGroupSnapshotResponse
}

func (o *PatchApi24ProtectionGroupSnapshotsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/protection-group-snapshots][%d] patchApi24ProtectionGroupSnapshotsOK  %+v", 200, o.Payload)
}

func (o *PatchApi24ProtectionGroupSnapshotsOK) GetPayload() *models.ProtectionGroupSnapshotResponse {
	return o.Payload
}

func (o *PatchApi24ProtectionGroupSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionGroupSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24ProtectionGroupSnapshotsBadRequest creates a PatchApi24ProtectionGroupSnapshotsBadRequest with default headers values
func NewPatchApi24ProtectionGroupSnapshotsBadRequest() *PatchApi24ProtectionGroupSnapshotsBadRequest {
	return &PatchApi24ProtectionGroupSnapshotsBadRequest{}
}

/*PatchApi24ProtectionGroupSnapshotsBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24ProtectionGroupSnapshotsBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24ProtectionGroupSnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/protection-group-snapshots][%d] patchApi24ProtectionGroupSnapshotsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24ProtectionGroupSnapshotsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24ProtectionGroupSnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
