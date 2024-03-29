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

// DeleteAPI24AdminsCacheReader is a Reader for the DeleteAPI24AdminsCache structure.
type DeleteAPI24AdminsCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24AdminsCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24AdminsCacheOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24AdminsCacheBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24AdminsCacheOK creates a DeleteApi24AdminsCacheOK with default headers values
func NewDeleteApi24AdminsCacheOK() *DeleteApi24AdminsCacheOK {
	return &DeleteApi24AdminsCacheOK{}
}

/*DeleteApi24AdminsCacheOK handles this case with default header values.

OK
*/
type DeleteApi24AdminsCacheOK struct {
}

func (o *DeleteApi24AdminsCacheOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/admins/cache][%d] deleteApi24AdminsCacheOK ", 200)
}

func (o *DeleteApi24AdminsCacheOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24AdminsCacheBadRequest creates a DeleteApi24AdminsCacheBadRequest with default headers values
func NewDeleteApi24AdminsCacheBadRequest() *DeleteApi24AdminsCacheBadRequest {
	return &DeleteApi24AdminsCacheBadRequest{}
}

/*DeleteApi24AdminsCacheBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24AdminsCacheBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24AdminsCacheBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/admins/cache][%d] deleteApi24AdminsCacheBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24AdminsCacheBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24AdminsCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
