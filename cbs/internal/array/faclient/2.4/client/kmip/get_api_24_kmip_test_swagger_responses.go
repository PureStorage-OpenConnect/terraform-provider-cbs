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

// GetAPI24KMIPTestReader is a Reader for the GetAPI24KMIPTest structure.
type GetAPI24KMIPTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24KMIPTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24KMIPTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24KMIPTestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24KMIPTestOK creates a GetApi24KMIPTestOK with default headers values
func NewGetApi24KMIPTestOK() *GetApi24KMIPTestOK {
	return &GetApi24KMIPTestOK{}
}

/*GetApi24KMIPTestOK handles this case with default header values.

OK
*/
type GetApi24KMIPTestOK struct {
	Payload *models.KMIPTestResultGetResponse
}

func (o *GetApi24KMIPTestOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/kmip/test][%d] getApi24KmipTestOK  %+v", 200, o.Payload)
}

func (o *GetApi24KMIPTestOK) GetPayload() *models.KMIPTestResultGetResponse {
	return o.Payload
}

func (o *GetApi24KMIPTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KMIPTestResultGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24KMIPTestBadRequest creates a GetApi24KMIPTestBadRequest with default headers values
func NewGetApi24KMIPTestBadRequest() *GetApi24KMIPTestBadRequest {
	return &GetApi24KMIPTestBadRequest{}
}

/*GetApi24KMIPTestBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24KMIPTestBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24KMIPTestBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/kmip/test][%d] getApi24KmipTestBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24KMIPTestBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24KMIPTestBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
