// Code generated by go-swagger; DO NOT EDIT.

package protection_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24ProtectionGroupsReader is a Reader for the DeleteAPI24ProtectionGroups structure.
type DeleteAPI24ProtectionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24ProtectionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24ProtectionGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24ProtectionGroupsOK creates a DeleteApi24ProtectionGroupsOK with default headers values
func NewDeleteApi24ProtectionGroupsOK() *DeleteApi24ProtectionGroupsOK {
	return &DeleteApi24ProtectionGroupsOK{}
}

/*DeleteApi24ProtectionGroupsOK handles this case with default header values.

OK
*/
type DeleteApi24ProtectionGroupsOK struct {
}

func (o *DeleteApi24ProtectionGroupsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/protection-groups][%d] deleteApi24ProtectionGroupsOK ", 200)
}

func (o *DeleteApi24ProtectionGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}