// Code generated by go-swagger; DO NOT EDIT.

package directory_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// DeleteAPI24DirectorySnapshotsReader is a Reader for the DeleteAPI24DirectorySnapshots structure.
type DeleteAPI24DirectorySnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24DirectorySnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24DirectorySnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24DirectorySnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24DirectorySnapshotsOK creates a DeleteApi24DirectorySnapshotsOK with default headers values
func NewDeleteApi24DirectorySnapshotsOK() *DeleteApi24DirectorySnapshotsOK {
	return &DeleteApi24DirectorySnapshotsOK{}
}

/*DeleteApi24DirectorySnapshotsOK handles this case with default header values.

OK
*/
type DeleteApi24DirectorySnapshotsOK struct {
}

func (o *DeleteApi24DirectorySnapshotsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/directory-snapshots][%d] deleteApi24DirectorySnapshotsOK ", 200)
}

func (o *DeleteApi24DirectorySnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24DirectorySnapshotsBadRequest creates a DeleteApi24DirectorySnapshotsBadRequest with default headers values
func NewDeleteApi24DirectorySnapshotsBadRequest() *DeleteApi24DirectorySnapshotsBadRequest {
	return &DeleteApi24DirectorySnapshotsBadRequest{}
}

/*DeleteApi24DirectorySnapshotsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24DirectorySnapshotsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24DirectorySnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/directory-snapshots][%d] deleteApi24DirectorySnapshotsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24DirectorySnapshotsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24DirectorySnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
