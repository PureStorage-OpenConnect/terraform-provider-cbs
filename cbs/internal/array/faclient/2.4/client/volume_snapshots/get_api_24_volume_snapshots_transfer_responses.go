// Code generated by go-swagger; DO NOT EDIT.

package volume_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24VolumeSnapshotsTransferReader is a Reader for the GetAPI24VolumeSnapshotsTransfer structure.
type GetAPI24VolumeSnapshotsTransferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24VolumeSnapshotsTransferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24VolumeSnapshotsTransferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24VolumeSnapshotsTransferBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24VolumeSnapshotsTransferOK creates a GetApi24VolumeSnapshotsTransferOK with default headers values
func NewGetApi24VolumeSnapshotsTransferOK() *GetApi24VolumeSnapshotsTransferOK {
	return &GetApi24VolumeSnapshotsTransferOK{}
}

/*GetApi24VolumeSnapshotsTransferOK handles this case with default header values.

OK
*/
type GetApi24VolumeSnapshotsTransferOK struct {
	Payload *models.VolumeSnapshotTransferGetResponse
}

func (o *GetApi24VolumeSnapshotsTransferOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volume-snapshots/transfer][%d] getApi24VolumeSnapshotsTransferOK  %+v", 200, o.Payload)
}

func (o *GetApi24VolumeSnapshotsTransferOK) GetPayload() *models.VolumeSnapshotTransferGetResponse {
	return o.Payload
}

func (o *GetApi24VolumeSnapshotsTransferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeSnapshotTransferGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24VolumeSnapshotsTransferBadRequest creates a GetApi24VolumeSnapshotsTransferBadRequest with default headers values
func NewGetApi24VolumeSnapshotsTransferBadRequest() *GetApi24VolumeSnapshotsTransferBadRequest {
	return &GetApi24VolumeSnapshotsTransferBadRequest{}
}

/*GetApi24VolumeSnapshotsTransferBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24VolumeSnapshotsTransferBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24VolumeSnapshotsTransferBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volume-snapshots/transfer][%d] getApi24VolumeSnapshotsTransferBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24VolumeSnapshotsTransferBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24VolumeSnapshotsTransferBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}