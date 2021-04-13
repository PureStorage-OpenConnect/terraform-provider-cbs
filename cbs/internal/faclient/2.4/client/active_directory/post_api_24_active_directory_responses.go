// Code generated by go-swagger; DO NOT EDIT.

package active_directory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24ActiveDirectoryReader is a Reader for the PostAPI24ActiveDirectory structure.
type PostAPI24ActiveDirectoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24ActiveDirectoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24ActiveDirectoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24ActiveDirectoryOK creates a PostApi24ActiveDirectoryOK with default headers values
func NewPostApi24ActiveDirectoryOK() *PostApi24ActiveDirectoryOK {
	return &PostApi24ActiveDirectoryOK{}
}

/*PostApi24ActiveDirectoryOK handles this case with default header values.

OK
*/
type PostApi24ActiveDirectoryOK struct {
	Payload *models.ActiveDirectoryResponse
}

func (o *PostApi24ActiveDirectoryOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/active-directory][%d] postApi24ActiveDirectoryOK  %+v", 200, o.Payload)
}

func (o *PostApi24ActiveDirectoryOK) GetPayload() *models.ActiveDirectoryResponse {
	return o.Payload
}

func (o *PostApi24ActiveDirectoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActiveDirectoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
