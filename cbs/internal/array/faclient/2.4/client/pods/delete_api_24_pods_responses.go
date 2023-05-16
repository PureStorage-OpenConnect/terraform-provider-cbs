// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// DeleteAPI24PodsReader is a Reader for the DeleteAPI24Pods structure.
type DeleteAPI24PodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24PodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24PodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24PodsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24PodsOK creates a DeleteApi24PodsOK with default headers values
func NewDeleteApi24PodsOK() *DeleteApi24PodsOK {
	return &DeleteApi24PodsOK{}
}

/*DeleteApi24PodsOK handles this case with default header values.

OK
*/
type DeleteApi24PodsOK struct {
}

func (o *DeleteApi24PodsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/pods][%d] deleteApi24PodsOK ", 200)
}

func (o *DeleteApi24PodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24PodsBadRequest creates a DeleteApi24PodsBadRequest with default headers values
func NewDeleteApi24PodsBadRequest() *DeleteApi24PodsBadRequest {
	return &DeleteApi24PodsBadRequest{}
}

/*DeleteApi24PodsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24PodsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24PodsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/pods][%d] deleteApi24PodsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24PodsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24PodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
