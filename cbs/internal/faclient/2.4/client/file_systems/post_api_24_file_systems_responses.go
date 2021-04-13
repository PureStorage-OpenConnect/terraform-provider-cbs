// Code generated by go-swagger; DO NOT EDIT.

package file_systems

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24FileSystemsReader is a Reader for the PostAPI24FileSystems structure.
type PostAPI24FileSystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24FileSystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24FileSystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24FileSystemsOK creates a PostApi24FileSystemsOK with default headers values
func NewPostApi24FileSystemsOK() *PostApi24FileSystemsOK {
	return &PostApi24FileSystemsOK{}
}

/*PostApi24FileSystemsOK handles this case with default header values.

OK
*/
type PostApi24FileSystemsOK struct {
	Payload *models.FileSystemResponse
}

func (o *PostApi24FileSystemsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/file-systems][%d] postApi24FileSystemsOK  %+v", 200, o.Payload)
}

func (o *PostApi24FileSystemsOK) GetPayload() *models.FileSystemResponse {
	return o.Payload
}

func (o *PostApi24FileSystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FileSystemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}