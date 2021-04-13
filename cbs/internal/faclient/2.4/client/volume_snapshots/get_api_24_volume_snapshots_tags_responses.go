// Code generated by go-swagger; DO NOT EDIT.

package volume_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24VolumeSnapshotsTagsReader is a Reader for the GetAPI24VolumeSnapshotsTags structure.
type GetAPI24VolumeSnapshotsTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24VolumeSnapshotsTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24VolumeSnapshotsTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24VolumeSnapshotsTagsOK creates a GetApi24VolumeSnapshotsTagsOK with default headers values
func NewGetApi24VolumeSnapshotsTagsOK() *GetApi24VolumeSnapshotsTagsOK {
	return &GetApi24VolumeSnapshotsTagsOK{}
}

/*GetApi24VolumeSnapshotsTagsOK handles this case with default header values.

OK
*/
type GetApi24VolumeSnapshotsTagsOK struct {
	Payload *models.TagGetResponse
}

func (o *GetApi24VolumeSnapshotsTagsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volume-snapshots/tags][%d] getApi24VolumeSnapshotsTagsOK  %+v", 200, o.Payload)
}

func (o *GetApi24VolumeSnapshotsTagsOK) GetPayload() *models.TagGetResponse {
	return o.Payload
}

func (o *GetApi24VolumeSnapshotsTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TagGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
