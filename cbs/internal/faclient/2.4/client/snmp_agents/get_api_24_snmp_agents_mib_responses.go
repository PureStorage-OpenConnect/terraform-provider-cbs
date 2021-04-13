// Code generated by go-swagger; DO NOT EDIT.

package snmp_agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24SNMPAgentsMibReader is a Reader for the GetAPI24SNMPAgentsMib structure.
type GetAPI24SNMPAgentsMibReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24SNMPAgentsMibReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24SNMPAgentsMibOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24SNMPAgentsMibOK creates a GetApi24SNMPAgentsMibOK with default headers values
func NewGetApi24SNMPAgentsMibOK() *GetApi24SNMPAgentsMibOK {
	return &GetApi24SNMPAgentsMibOK{}
}

/*GetApi24SNMPAgentsMibOK handles this case with default header values.

OK
*/
type GetApi24SNMPAgentsMibOK struct {
	Payload *models.SNMPAgentMibGetResponse
}

func (o *GetApi24SNMPAgentsMibOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/snmp-agents/mib][%d] getApi24SnmpAgentsMibOK  %+v", 200, o.Payload)
}

func (o *GetApi24SNMPAgentsMibOK) GetPayload() *models.SNMPAgentMibGetResponse {
	return o.Payload
}

func (o *GetApi24SNMPAgentsMibOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SNMPAgentMibGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}