// Code generated by go-swagger; DO NOT EDIT.

package remote_protection_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24RemoteProtectionGroupsReader is a Reader for the DeleteAPI24RemoteProtectionGroups structure.
type DeleteAPI24RemoteProtectionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24RemoteProtectionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24RemoteProtectionGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24RemoteProtectionGroupsOK creates a DeleteApi24RemoteProtectionGroupsOK with default headers values
func NewDeleteApi24RemoteProtectionGroupsOK() *DeleteApi24RemoteProtectionGroupsOK {
	return &DeleteApi24RemoteProtectionGroupsOK{}
}

/*DeleteApi24RemoteProtectionGroupsOK handles this case with default header values.

OK
*/
type DeleteApi24RemoteProtectionGroupsOK struct {
}

func (o *DeleteApi24RemoteProtectionGroupsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/remote-protection-groups][%d] deleteApi24RemoteProtectionGroupsOK ", 200)
}

func (o *DeleteApi24RemoteProtectionGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
