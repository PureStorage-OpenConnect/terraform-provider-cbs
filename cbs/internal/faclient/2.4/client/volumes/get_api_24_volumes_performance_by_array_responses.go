// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24VolumesPerformanceByArrayReader is a Reader for the GetAPI24VolumesPerformanceByArray structure.
type GetAPI24VolumesPerformanceByArrayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24VolumesPerformanceByArrayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24VolumesPerformanceByArrayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24VolumesPerformanceByArrayOK creates a GetApi24VolumesPerformanceByArrayOK with default headers values
func NewGetApi24VolumesPerformanceByArrayOK() *GetApi24VolumesPerformanceByArrayOK {
	return &GetApi24VolumesPerformanceByArrayOK{}
}

/*GetApi24VolumesPerformanceByArrayOK handles this case with default header values.

OK
*/
type GetApi24VolumesPerformanceByArrayOK struct {
	Payload *models.ResourcePerformanceByArrayGetResponse
}

func (o *GetApi24VolumesPerformanceByArrayOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volumes/performance/by-array][%d] getApi24VolumesPerformanceByArrayOK  %+v", 200, o.Payload)
}

func (o *GetApi24VolumesPerformanceByArrayOK) GetPayload() *models.ResourcePerformanceByArrayGetResponse {
	return o.Payload
}

func (o *GetApi24VolumesPerformanceByArrayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourcePerformanceByArrayGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}