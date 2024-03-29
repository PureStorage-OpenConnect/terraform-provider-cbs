// Code generated by go-swagger; DO NOT EDIT.

package software

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24SoftwareReader is a Reader for the GetAPI24Software structure.
type GetAPI24SoftwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24SoftwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24SoftwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24SoftwareBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24SoftwareOK creates a GetApi24SoftwareOK with default headers values
func NewGetApi24SoftwareOK() *GetApi24SoftwareOK {
	return &GetApi24SoftwareOK{}
}

/*GetApi24SoftwareOK handles this case with default header values.

OK
*/
type GetApi24SoftwareOK struct {
	Payload *models.SoftwareGetResponse
}

func (o *GetApi24SoftwareOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/software][%d] getApi24SoftwareOK  %+v", 200, o.Payload)
}

func (o *GetApi24SoftwareOK) GetPayload() *models.SoftwareGetResponse {
	return o.Payload
}

func (o *GetApi24SoftwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SoftwareGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24SoftwareBadRequest creates a GetApi24SoftwareBadRequest with default headers values
func NewGetApi24SoftwareBadRequest() *GetApi24SoftwareBadRequest {
	return &GetApi24SoftwareBadRequest{}
}

/*GetApi24SoftwareBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24SoftwareBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24SoftwareBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/software][%d] getApi24SoftwareBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24SoftwareBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24SoftwareBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
