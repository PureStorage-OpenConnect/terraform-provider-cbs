// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24PoliciesNfsMembersReader is a Reader for the DeleteAPI24PoliciesNfsMembers structure.
type DeleteAPI24PoliciesNfsMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24PoliciesNfsMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24PoliciesNfsMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24PoliciesNfsMembersOK creates a DeleteApi24PoliciesNfsMembersOK with default headers values
func NewDeleteApi24PoliciesNfsMembersOK() *DeleteApi24PoliciesNfsMembersOK {
	return &DeleteApi24PoliciesNfsMembersOK{}
}

/*DeleteApi24PoliciesNfsMembersOK handles this case with default header values.

OK
*/
type DeleteApi24PoliciesNfsMembersOK struct {
}

func (o *DeleteApi24PoliciesNfsMembersOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/policies/nfs/members][%d] deleteApi24PoliciesNfsMembersOK ", 200)
}

func (o *DeleteApi24PoliciesNfsMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
