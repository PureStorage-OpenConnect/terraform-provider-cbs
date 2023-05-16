// Code generated by go-swagger; DO NOT EDIT.

package remote_volume_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PostAPI24RemoteVolumeSnapshotsReader is a Reader for the PostAPI24RemoteVolumeSnapshots structure.
type PostAPI24RemoteVolumeSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24RemoteVolumeSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24RemoteVolumeSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24RemoteVolumeSnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24RemoteVolumeSnapshotsOK creates a PostApi24RemoteVolumeSnapshotsOK with default headers values
func NewPostApi24RemoteVolumeSnapshotsOK() *PostApi24RemoteVolumeSnapshotsOK {
	return &PostApi24RemoteVolumeSnapshotsOK{}
}

/*PostApi24RemoteVolumeSnapshotsOK handles this case with default header values.

OK
*/
type PostApi24RemoteVolumeSnapshotsOK struct {
	Payload *models.RemoteVolumeSnapshotResponse
}

func (o *PostApi24RemoteVolumeSnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/remote-volume-snapshots][%d] postApi24RemoteVolumeSnapshotsOK  %+v", 200, o.Payload)
}

func (o *PostApi24RemoteVolumeSnapshotsOK) GetPayload() *models.RemoteVolumeSnapshotResponse {
	return o.Payload
}

func (o *PostApi24RemoteVolumeSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteVolumeSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24RemoteVolumeSnapshotsBadRequest creates a PostApi24RemoteVolumeSnapshotsBadRequest with default headers values
func NewPostApi24RemoteVolumeSnapshotsBadRequest() *PostApi24RemoteVolumeSnapshotsBadRequest {
	return &PostApi24RemoteVolumeSnapshotsBadRequest{}
}

/*PostApi24RemoteVolumeSnapshotsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24RemoteVolumeSnapshotsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24RemoteVolumeSnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/remote-volume-snapshots][%d] postApi24RemoteVolumeSnapshotsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24RemoteVolumeSnapshotsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24RemoteVolumeSnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
