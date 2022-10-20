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

// GetAPI24PodsArraysReader is a Reader for the GetAPI24PodsArrays structure.
type GetAPI24PodsArraysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24PodsArraysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24PodsArraysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24PodsArraysBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24PodsArraysOK creates a GetApi24PodsArraysOK with default headers values
func NewGetApi24PodsArraysOK() *GetApi24PodsArraysOK {
	return &GetApi24PodsArraysOK{}
}

/*GetApi24PodsArraysOK handles this case with default header values.

OK
*/
type GetApi24PodsArraysOK struct {
	Payload *models.MemberGetResponse
}

func (o *GetApi24PodsArraysOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/pods/arrays][%d] getApi24PodsArraysOK  %+v", 200, o.Payload)
}

func (o *GetApi24PodsArraysOK) GetPayload() *models.MemberGetResponse {
	return o.Payload
}

func (o *GetApi24PodsArraysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24PodsArraysBadRequest creates a GetApi24PodsArraysBadRequest with default headers values
func NewGetApi24PodsArraysBadRequest() *GetApi24PodsArraysBadRequest {
	return &GetApi24PodsArraysBadRequest{}
}

/*GetApi24PodsArraysBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24PodsArraysBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24PodsArraysBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/pods/arrays][%d] getApi24PodsArraysBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24PodsArraysBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24PodsArraysBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
