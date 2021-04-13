// Code generated by go-swagger; DO NOT EDIT.

package directory_exports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24DirectoryExportsReader is a Reader for the PostAPI24DirectoryExports structure.
type PostAPI24DirectoryExportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24DirectoryExportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24DirectoryExportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24DirectoryExportsOK creates a PostApi24DirectoryExportsOK with default headers values
func NewPostApi24DirectoryExportsOK() *PostApi24DirectoryExportsOK {
	return &PostApi24DirectoryExportsOK{}
}

/*PostApi24DirectoryExportsOK handles this case with default header values.

OK
*/
type PostApi24DirectoryExportsOK struct {
	Payload *models.DirectoryExportResponse
}

func (o *PostApi24DirectoryExportsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/directory-exports][%d] postApi24DirectoryExportsOK  %+v", 200, o.Payload)
}

func (o *PostApi24DirectoryExportsOK) GetPayload() *models.DirectoryExportResponse {
	return o.Payload
}

func (o *PostApi24DirectoryExportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryExportResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
