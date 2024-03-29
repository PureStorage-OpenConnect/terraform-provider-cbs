// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
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
	case 400:
		result := NewDeleteApi24PoliciesNfsMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

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

// NewDeleteApi24PoliciesNfsMembersBadRequest creates a DeleteApi24PoliciesNfsMembersBadRequest with default headers values
func NewDeleteApi24PoliciesNfsMembersBadRequest() *DeleteApi24PoliciesNfsMembersBadRequest {
	return &DeleteApi24PoliciesNfsMembersBadRequest{}
}

/*DeleteApi24PoliciesNfsMembersBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24PoliciesNfsMembersBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24PoliciesNfsMembersBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/policies/nfs/members][%d] deleteApi24PoliciesNfsMembersBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24PoliciesNfsMembersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24PoliciesNfsMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
