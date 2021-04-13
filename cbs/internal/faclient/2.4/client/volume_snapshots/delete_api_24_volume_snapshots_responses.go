// Code generated by go-swagger; DO NOT EDIT.

package volume_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24VolumeSnapshotsReader is a Reader for the DeleteAPI24VolumeSnapshots structure.
type DeleteAPI24VolumeSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24VolumeSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24VolumeSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24VolumeSnapshotsOK creates a DeleteApi24VolumeSnapshotsOK with default headers values
func NewDeleteApi24VolumeSnapshotsOK() *DeleteApi24VolumeSnapshotsOK {
	return &DeleteApi24VolumeSnapshotsOK{}
}

/*DeleteApi24VolumeSnapshotsOK handles this case with default header values.

OK
*/
type DeleteApi24VolumeSnapshotsOK struct {
}

func (o *DeleteApi24VolumeSnapshotsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/volume-snapshots][%d] deleteApi24VolumeSnapshotsOK ", 200)
}

func (o *DeleteApi24VolumeSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}