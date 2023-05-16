// Code generated by go-swagger; DO NOT EDIT.

package array_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24ArrayConnectionsReader is a Reader for the GetAPI24ArrayConnections structure.
type GetAPI24ArrayConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24ArrayConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24ArrayConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24ArrayConnectionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24ArrayConnectionsOK creates a GetApi24ArrayConnectionsOK with default headers values
func NewGetApi24ArrayConnectionsOK() *GetApi24ArrayConnectionsOK {
	return &GetApi24ArrayConnectionsOK{}
}

/*GetApi24ArrayConnectionsOK handles this case with default header values.

OK
*/
type GetApi24ArrayConnectionsOK struct {
	Payload *models.ArrayConnectionGetResponse
}

func (o *GetApi24ArrayConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/array-connections][%d] getApi24ArrayConnectionsOK  %+v", 200, o.Payload)
}

func (o *GetApi24ArrayConnectionsOK) GetPayload() *models.ArrayConnectionGetResponse {
	return o.Payload
}

func (o *GetApi24ArrayConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ArrayConnectionGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24ArrayConnectionsBadRequest creates a GetApi24ArrayConnectionsBadRequest with default headers values
func NewGetApi24ArrayConnectionsBadRequest() *GetApi24ArrayConnectionsBadRequest {
	return &GetApi24ArrayConnectionsBadRequest{}
}

/*GetApi24ArrayConnectionsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24ArrayConnectionsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24ArrayConnectionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/array-connections][%d] getApi24ArrayConnectionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24ArrayConnectionsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24ArrayConnectionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
