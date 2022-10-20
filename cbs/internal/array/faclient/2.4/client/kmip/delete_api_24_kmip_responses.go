// Code generated by go-swagger; DO NOT EDIT.

package kmip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// DeleteAPI24KMIPReader is a Reader for the DeleteAPI24KMIP structure.
type DeleteAPI24KMIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24KMIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24KMIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24KMIPBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24KMIPOK creates a DeleteApi24KMIPOK with default headers values
func NewDeleteApi24KMIPOK() *DeleteApi24KMIPOK {
	return &DeleteApi24KMIPOK{}
}

/*DeleteApi24KMIPOK handles this case with default header values.

OK
*/
type DeleteApi24KMIPOK struct {
}

func (o *DeleteApi24KMIPOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/kmip][%d] deleteApi24KmipOK ", 200)
}

func (o *DeleteApi24KMIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24KMIPBadRequest creates a DeleteApi24KMIPBadRequest with default headers values
func NewDeleteApi24KMIPBadRequest() *DeleteApi24KMIPBadRequest {
	return &DeleteApi24KMIPBadRequest{}
}

/*DeleteApi24KMIPBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24KMIPBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24KMIPBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/kmip][%d] deleteApi24KmipBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24KMIPBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24KMIPBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
