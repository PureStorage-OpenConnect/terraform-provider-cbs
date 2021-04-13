// Code generated by go-swagger; DO NOT EDIT.

package remote_volume_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24RemoteVolumeSnapshotsReader is a Reader for the GetAPI24RemoteVolumeSnapshots structure.
type GetAPI24RemoteVolumeSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24RemoteVolumeSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24RemoteVolumeSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24RemoteVolumeSnapshotsOK creates a GetApi24RemoteVolumeSnapshotsOK with default headers values
func NewGetApi24RemoteVolumeSnapshotsOK() *GetApi24RemoteVolumeSnapshotsOK {
	return &GetApi24RemoteVolumeSnapshotsOK{}
}

/*GetApi24RemoteVolumeSnapshotsOK handles this case with default header values.

OK
*/
type GetApi24RemoteVolumeSnapshotsOK struct {
	Payload *models.RemoteVolumeSnapshotGetResponse
}

func (o *GetApi24RemoteVolumeSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/remote-volume-snapshots][%d] getApi24RemoteVolumeSnapshotsOK  %+v", 200, o.Payload)
}

func (o *GetApi24RemoteVolumeSnapshotsOK) GetPayload() *models.RemoteVolumeSnapshotGetResponse {
	return o.Payload
}

func (o *GetApi24RemoteVolumeSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteVolumeSnapshotGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
