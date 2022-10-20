// Code generated by go-swagger; DO NOT EDIT.

package api_clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PostAPI24APIClientsReader is a Reader for the PostAPI24APIClients structure.
type PostAPI24APIClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24APIClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24APIClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24APIClientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24APIClientsOK creates a PostApi24APIClientsOK with default headers values
func NewPostApi24APIClientsOK() *PostApi24APIClientsOK {
	return &PostApi24APIClientsOK{}
}

/*PostApi24APIClientsOK handles this case with default header values.

OK
*/
type PostApi24APIClientsOK struct {
	Payload *models.APIClientResponse
}

func (o *PostApi24APIClientsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/api-clients][%d] postApi24ApiClientsOK  %+v", 200, o.Payload)
}

func (o *PostApi24APIClientsOK) GetPayload() *models.APIClientResponse {
	return o.Payload
}

func (o *PostApi24APIClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIClientResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24APIClientsBadRequest creates a PostApi24APIClientsBadRequest with default headers values
func NewPostApi24APIClientsBadRequest() *PostApi24APIClientsBadRequest {
	return &PostApi24APIClientsBadRequest{}
}

/*PostApi24APIClientsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24APIClientsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24APIClientsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/api-clients][%d] postApi24ApiClientsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24APIClientsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24APIClientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
