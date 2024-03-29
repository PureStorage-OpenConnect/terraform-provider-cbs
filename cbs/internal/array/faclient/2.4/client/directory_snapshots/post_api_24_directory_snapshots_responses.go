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

// PostAPI24DirectorySnapshotsReader is a Reader for the PostAPI24DirectorySnapshots structure.
type PostAPI24DirectorySnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24DirectorySnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24DirectorySnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24DirectorySnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24DirectorySnapshotsOK creates a PostApi24DirectorySnapshotsOK with default headers values
func NewPostApi24DirectorySnapshotsOK() *PostApi24DirectorySnapshotsOK {
	return &PostApi24DirectorySnapshotsOK{}
}

/*PostApi24DirectorySnapshotsOK handles this case with default header values.

OK
*/
type PostApi24DirectorySnapshotsOK struct {
	Payload *models.DirectorySnapshotResponse
}

func (o *PostApi24DirectorySnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/directory-snapshots][%d] postApi24DirectorySnapshotsOK  %+v", 200, o.Payload)
}

func (o *PostApi24DirectorySnapshotsOK) GetPayload() *models.DirectorySnapshotResponse {
	return o.Payload
}

func (o *PostApi24DirectorySnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectorySnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24DirectorySnapshotsBadRequest creates a PostApi24DirectorySnapshotsBadRequest with default headers values
func NewPostApi24DirectorySnapshotsBadRequest() *PostApi24DirectorySnapshotsBadRequest {
	return &PostApi24DirectorySnapshotsBadRequest{}
}

/*PostApi24DirectorySnapshotsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24DirectorySnapshotsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24DirectorySnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/directory-snapshots][%d] postApi24DirectorySnapshotsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24DirectorySnapshotsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24DirectorySnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
