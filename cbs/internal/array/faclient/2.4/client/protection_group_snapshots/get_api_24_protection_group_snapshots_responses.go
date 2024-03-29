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

// GetAPI24ProtectionGroupSnapshotsReader is a Reader for the GetAPI24ProtectionGroupSnapshots structure.
type GetAPI24ProtectionGroupSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24ProtectionGroupSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24ProtectionGroupSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24ProtectionGroupSnapshotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24ProtectionGroupSnapshotsOK creates a GetApi24ProtectionGroupSnapshotsOK with default headers values
func NewGetApi24ProtectionGroupSnapshotsOK() *GetApi24ProtectionGroupSnapshotsOK {
	return &GetApi24ProtectionGroupSnapshotsOK{}
}

/*GetApi24ProtectionGroupSnapshotsOK handles this case with default header values.

OK
*/
type GetApi24ProtectionGroupSnapshotsOK struct {
	Payload *models.ProtectionGroupSnapshotGetResponse
}

func (o *GetApi24ProtectionGroupSnapshotsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/protection-group-snapshots][%d] getApi24ProtectionGroupSnapshotsOK  %+v", 200, o.Payload)
}

func (o *GetApi24ProtectionGroupSnapshotsOK) GetPayload() *models.ProtectionGroupSnapshotGetResponse {
	return o.Payload
}

func (o *GetApi24ProtectionGroupSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionGroupSnapshotGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24ProtectionGroupSnapshotsBadRequest creates a GetApi24ProtectionGroupSnapshotsBadRequest with default headers values
func NewGetApi24ProtectionGroupSnapshotsBadRequest() *GetApi24ProtectionGroupSnapshotsBadRequest {
	return &GetApi24ProtectionGroupSnapshotsBadRequest{}
}

/*GetApi24ProtectionGroupSnapshotsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24ProtectionGroupSnapshotsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24ProtectionGroupSnapshotsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/protection-group-snapshots][%d] getApi24ProtectionGroupSnapshotsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24ProtectionGroupSnapshotsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24ProtectionGroupSnapshotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
