// Code generated by go-swagger; DO NOT EDIT.

package directory_exports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24DirectoryExportsReader is a Reader for the GetAPI24DirectoryExports structure.
type GetAPI24DirectoryExportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24DirectoryExportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24DirectoryExportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24DirectoryExportsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24DirectoryExportsOK creates a GetApi24DirectoryExportsOK with default headers values
func NewGetApi24DirectoryExportsOK() *GetApi24DirectoryExportsOK {
	return &GetApi24DirectoryExportsOK{}
}

/*GetApi24DirectoryExportsOK handles this case with default header values.

OK
*/
type GetApi24DirectoryExportsOK struct {
	Payload *models.DirectoryExportGetResponse
}

func (o *GetApi24DirectoryExportsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/directory-exports][%d] getApi24DirectoryExportsOK  %+v", 200, o.Payload)
}

func (o *GetApi24DirectoryExportsOK) GetPayload() *models.DirectoryExportGetResponse {
	return o.Payload
}

func (o *GetApi24DirectoryExportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DirectoryExportGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24DirectoryExportsBadRequest creates a GetApi24DirectoryExportsBadRequest with default headers values
func NewGetApi24DirectoryExportsBadRequest() *GetApi24DirectoryExportsBadRequest {
	return &GetApi24DirectoryExportsBadRequest{}
}

/*GetApi24DirectoryExportsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24DirectoryExportsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24DirectoryExportsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/directory-exports][%d] getApi24DirectoryExportsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24DirectoryExportsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24DirectoryExportsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}