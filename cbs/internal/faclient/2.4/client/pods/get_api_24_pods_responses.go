// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24PodsReader is a Reader for the GetAPI24Pods structure.
type GetAPI24PodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24PodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24PodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24PodsOK creates a GetApi24PodsOK with default headers values
func NewGetApi24PodsOK() *GetApi24PodsOK {
	return &GetApi24PodsOK{}
}

/*GetApi24PodsOK handles this case with default header values.

OK
*/
type GetApi24PodsOK struct {
	Payload *models.PodGetResponse
}

func (o *GetApi24PodsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/pods][%d] getApi24PodsOK  %+v", 200, o.Payload)
}

func (o *GetApi24PodsOK) GetPayload() *models.PodGetResponse {
	return o.Payload
}

func (o *GetApi24PodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}