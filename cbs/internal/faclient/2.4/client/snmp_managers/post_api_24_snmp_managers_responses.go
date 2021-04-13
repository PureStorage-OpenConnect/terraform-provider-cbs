// Code generated by go-swagger; DO NOT EDIT.

package snmp_managers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24SNMPManagersReader is a Reader for the PostAPI24SNMPManagers structure.
type PostAPI24SNMPManagersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24SNMPManagersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24SNMPManagersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24SNMPManagersOK creates a PostApi24SNMPManagersOK with default headers values
func NewPostApi24SNMPManagersOK() *PostApi24SNMPManagersOK {
	return &PostApi24SNMPManagersOK{}
}

/*PostApi24SNMPManagersOK handles this case with default header values.

Returns the newly created snmp manager object.
*/
type PostApi24SNMPManagersOK struct {
	Payload *models.SNMPManagerResponse
}

func (o *PostApi24SNMPManagersOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/snmp-managers][%d] postApi24SnmpManagersOK  %+v", 200, o.Payload)
}

func (o *PostApi24SNMPManagersOK) GetPayload() *models.SNMPManagerResponse {
	return o.Payload
}

func (o *PostApi24SNMPManagersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SNMPManagerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
