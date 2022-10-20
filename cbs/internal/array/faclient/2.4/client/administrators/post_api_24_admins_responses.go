// Code generated by go-swagger; DO NOT EDIT.

package administrators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PostAPI24AdminsReader is a Reader for the PostAPI24Admins structure.
type PostAPI24AdminsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24AdminsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24AdminsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24AdminsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24AdminsOK creates a PostApi24AdminsOK with default headers values
func NewPostApi24AdminsOK() *PostApi24AdminsOK {
	return &PostApi24AdminsOK{}
}

/*PostApi24AdminsOK handles this case with default header values.

OK
*/
type PostApi24AdminsOK struct {
	Payload *models.AdminResponse
}

func (o *PostApi24AdminsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/admins][%d] postApi24AdminsOK  %+v", 200, o.Payload)
}

func (o *PostApi24AdminsOK) GetPayload() *models.AdminResponse {
	return o.Payload
}

func (o *PostApi24AdminsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24AdminsBadRequest creates a PostApi24AdminsBadRequest with default headers values
func NewPostApi24AdminsBadRequest() *PostApi24AdminsBadRequest {
	return &PostApi24AdminsBadRequest{}
}

/*PostApi24AdminsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24AdminsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24AdminsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/admins][%d] postApi24AdminsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24AdminsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24AdminsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
