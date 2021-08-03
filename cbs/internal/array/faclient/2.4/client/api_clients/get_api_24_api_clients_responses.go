// Code generated by go-swagger; DO NOT EDIT.

package api_clients

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24APIClientsReader is a Reader for the GetAPI24APIClients structure.
type GetAPI24APIClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24APIClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24APIClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24APIClientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24APIClientsOK creates a GetApi24APIClientsOK with default headers values
func NewGetApi24APIClientsOK() *GetApi24APIClientsOK {
	return &GetApi24APIClientsOK{}
}

/*GetApi24APIClientsOK handles this case with default header values.

OK
*/
type GetApi24APIClientsOK struct {
	Payload *models.APIClientGetResponse
}

func (o *GetApi24APIClientsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/api-clients][%d] getApi24ApiClientsOK  %+v", 200, o.Payload)
}

func (o *GetApi24APIClientsOK) GetPayload() *models.APIClientGetResponse {
	return o.Payload
}

func (o *GetApi24APIClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIClientGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24APIClientsBadRequest creates a GetApi24APIClientsBadRequest with default headers values
func NewGetApi24APIClientsBadRequest() *GetApi24APIClientsBadRequest {
	return &GetApi24APIClientsBadRequest{}
}

/*GetApi24APIClientsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24APIClientsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24APIClientsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/api-clients][%d] getApi24ApiClientsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24APIClientsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24APIClientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}