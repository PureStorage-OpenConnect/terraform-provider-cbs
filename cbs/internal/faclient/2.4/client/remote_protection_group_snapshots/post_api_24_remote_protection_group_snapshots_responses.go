// Code generated by go-swagger; DO NOT EDIT.

package remote_protection_group_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24RemoteProtectionGroupSnapshotsReader is a Reader for the PostAPI24RemoteProtectionGroupSnapshots structure.
type PostAPI24RemoteProtectionGroupSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24RemoteProtectionGroupSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24RemoteProtectionGroupSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24RemoteProtectionGroupSnapshotsOK creates a PostApi24RemoteProtectionGroupSnapshotsOK with default headers values
func NewPostApi24RemoteProtectionGroupSnapshotsOK() *PostApi24RemoteProtectionGroupSnapshotsOK {
	return &PostApi24RemoteProtectionGroupSnapshotsOK{}
}

/*PostApi24RemoteProtectionGroupSnapshotsOK handles this case with default header values.

OK
*/
type PostApi24RemoteProtectionGroupSnapshotsOK struct {
	Payload *models.RemoteProtectionGroupSnapshotResponse
}

func (o *PostApi24RemoteProtectionGroupSnapshotsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/remote-protection-group-snapshots][%d] postApi24RemoteProtectionGroupSnapshotsOK  %+v", 200, o.Payload)
}

func (o *PostApi24RemoteProtectionGroupSnapshotsOK) GetPayload() *models.RemoteProtectionGroupSnapshotResponse {
	return o.Payload
}

func (o *PostApi24RemoteProtectionGroupSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProtectionGroupSnapshotResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
