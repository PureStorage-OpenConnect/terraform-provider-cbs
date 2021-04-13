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

// GetAPI24SNMPAgentsReader is a Reader for the GetAPI24SNMPAgents structure.
type GetAPI24SNMPAgentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24SNMPAgentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24SNMPAgentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24SNMPAgentsOK creates a GetApi24SNMPAgentsOK with default headers values
func NewGetApi24SNMPAgentsOK() *GetApi24SNMPAgentsOK {
	return &GetApi24SNMPAgentsOK{}
}

/*GetApi24SNMPAgentsOK handles this case with default header values.

OK
*/
type GetApi24SNMPAgentsOK struct {
	Payload *models.SNMPAgentGetResponse
}

func (o *GetApi24SNMPAgentsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/snmp-agents][%d] getApi24SnmpAgentsOK  %+v", 200, o.Payload)
}

func (o *GetApi24SNMPAgentsOK) GetPayload() *models.SNMPAgentGetResponse {
	return o.Payload
}

func (o *GetApi24SNMPAgentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SNMPAgentGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
