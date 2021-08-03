// Code generated by go-swagger; DO NOT EDIT.

package subnets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// DeleteAPI24SubnetsReader is a Reader for the DeleteAPI24Subnets structure.
type DeleteAPI24SubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24SubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24SubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24SubnetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24SubnetsOK creates a DeleteApi24SubnetsOK with default headers values
func NewDeleteApi24SubnetsOK() *DeleteApi24SubnetsOK {
	return &DeleteApi24SubnetsOK{}
}

/*DeleteApi24SubnetsOK handles this case with default header values.

OK
*/
type DeleteApi24SubnetsOK struct {
}

func (o *DeleteApi24SubnetsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/subnets][%d] deleteApi24SubnetsOK ", 200)
}

func (o *DeleteApi24SubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24SubnetsBadRequest creates a DeleteApi24SubnetsBadRequest with default headers values
func NewDeleteApi24SubnetsBadRequest() *DeleteApi24SubnetsBadRequest {
	return &DeleteApi24SubnetsBadRequest{}
}

/*DeleteApi24SubnetsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24SubnetsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24SubnetsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/subnets][%d] deleteApi24SubnetsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24SubnetsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24SubnetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}