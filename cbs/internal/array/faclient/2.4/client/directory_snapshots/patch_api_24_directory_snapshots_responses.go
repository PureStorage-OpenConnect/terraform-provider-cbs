// Code generated by go-swagger; DO NOT EDIT.

package directory_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24DirectorySnapshotsReader is a Reader for the PatchAPI24DirectorySnapshots structure.
type PatchAPI24DirectorySnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24DirectorySnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24DirectorySnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24DirectorySnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24DirectorySnapshotsOK creates a PatchApi24DirectorySnapshotsOK with default headers values
func NewPatchApi24DirectorySnapshotsOK() *PatchApi24DirectorySnapshotsOK {
	return &PatchApi24DirectorySnapshotsOK{}
}

/*PatchApi24DirectorySnapshotsOK handles this case with default header values.

OK
*/
type PatchApi24DirectorySnapshotsOK struct {
	Payload *models.DirectorySnapshotResponse
}

func (o *PatchApi24DirectorySnapshotsOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/directory-snapshots][%d] patchApi24DirectorySnapshotsOK  %+v", 200, o.Payload)
}

func (o *PatchApi24DirectorySnapshotsOK) GetPayload() *models.DirectorySnapshotResponse {
	return o.Payload
}

func (o *PatchApi24DirectorySnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectorySnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24DirectorySnapshotsBadRequest creates a PatchApi24DirectorySnapshotsBadRequest with default headers values
func NewPatchApi24DirectorySnapshotsBadRequest() *PatchApi24DirectorySnapshotsBadRequest {
	return &PatchApi24DirectorySnapshotsBadRequest{}
}

/*PatchApi24DirectorySnapshotsBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24DirectorySnapshotsBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24DirectorySnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/directory-snapshots][%d] patchApi24DirectorySnapshotsBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24DirectorySnapshotsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24DirectorySnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}