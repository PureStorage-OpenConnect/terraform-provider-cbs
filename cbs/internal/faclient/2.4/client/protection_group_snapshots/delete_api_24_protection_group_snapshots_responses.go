// Code generated by go-swagger; DO NOT EDIT.

package protection_group_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24ProtectionGroupSnapshotsReader is a Reader for the DeleteAPI24ProtectionGroupSnapshots structure.
type DeleteAPI24ProtectionGroupSnapshotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24ProtectionGroupSnapshotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24ProtectionGroupSnapshotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24ProtectionGroupSnapshotsOK creates a DeleteApi24ProtectionGroupSnapshotsOK with default headers values
func NewDeleteApi24ProtectionGroupSnapshotsOK() *DeleteApi24ProtectionGroupSnapshotsOK {
	return &DeleteApi24ProtectionGroupSnapshotsOK{}
}

/*DeleteApi24ProtectionGroupSnapshotsOK handles this case with default header values.

OK
*/
type DeleteApi24ProtectionGroupSnapshotsOK struct {
}

func (o *DeleteApi24ProtectionGroupSnapshotsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/protection-group-snapshots][%d] deleteApi24ProtectionGroupSnapshotsOK ", 200)
}

func (o *DeleteApi24ProtectionGroupSnapshotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}