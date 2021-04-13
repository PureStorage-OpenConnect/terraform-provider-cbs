// Code generated by go-swagger; DO NOT EDIT.

package hosts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24HostsReader is a Reader for the GetAPI24Hosts structure.
type GetAPI24HostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24HostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24HostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24HostsOK creates a GetApi24HostsOK with default headers values
func NewGetApi24HostsOK() *GetApi24HostsOK {
	return &GetApi24HostsOK{}
}

/*GetApi24HostsOK handles this case with default header values.

OK
*/
type GetApi24HostsOK struct {
	Payload *models.HostGetResponse
}

func (o *GetApi24HostsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/hosts][%d] getApi24HostsOK  %+v", 200, o.Payload)
}

func (o *GetApi24HostsOK) GetPayload() *models.HostGetResponse {
	return o.Payload
}

func (o *GetApi24HostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HostGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}