// Code generated by go-swagger; DO NOT EDIT.

package administrators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PutAPI24AdminsCacheReader is a Reader for the PutAPI24AdminsCache structure.
type PutAPI24AdminsCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPI24AdminsCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutApi24AdminsCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutApi24AdminsCacheBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutApi24AdminsCacheOK creates a PutApi24AdminsCacheOK with default headers values
func NewPutApi24AdminsCacheOK() *PutApi24AdminsCacheOK {
	return &PutApi24AdminsCacheOK{}
}

/*PutApi24AdminsCacheOK handles this case with default header values.

OK
*/
type PutApi24AdminsCacheOK struct {
	Payload *models.AdminCacheResponse
}

func (o *PutApi24AdminsCacheOK) Error() string {
	return fmt.Sprintf("[PUT /api/2.4/admins/cache][%d] putApi24AdminsCacheOK  %+v", 200, o.Payload)
}

func (o *PutApi24AdminsCacheOK) GetPayload() *models.AdminCacheResponse {
	return o.Payload
}

func (o *PutApi24AdminsCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AdminCacheResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApi24AdminsCacheBadRequest creates a PutApi24AdminsCacheBadRequest with default headers values
func NewPutApi24AdminsCacheBadRequest() *PutApi24AdminsCacheBadRequest {
	return &PutApi24AdminsCacheBadRequest{}
}

/*PutApi24AdminsCacheBadRequest handles this case with default header values.

BadRequest
*/
type PutApi24AdminsCacheBadRequest struct {
	Payload *models.Error
}

func (o *PutApi24AdminsCacheBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/2.4/admins/cache][%d] putApi24AdminsCacheBadRequest  %+v", 400, o.Payload)
}

func (o *PutApi24AdminsCacheBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutApi24AdminsCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}