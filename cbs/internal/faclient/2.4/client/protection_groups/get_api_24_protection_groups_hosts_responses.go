// Code generated by go-swagger; DO NOT EDIT.

package protection_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24ProtectionGroupsHostsReader is a Reader for the GetAPI24ProtectionGroupsHosts structure.
type GetAPI24ProtectionGroupsHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24ProtectionGroupsHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24ProtectionGroupsHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24ProtectionGroupsHostsOK creates a GetApi24ProtectionGroupsHostsOK with default headers values
func NewGetApi24ProtectionGroupsHostsOK() *GetApi24ProtectionGroupsHostsOK {
	return &GetApi24ProtectionGroupsHostsOK{}
}

/*GetApi24ProtectionGroupsHostsOK handles this case with default header values.

OK
*/
type GetApi24ProtectionGroupsHostsOK struct {
	Payload *models.MemberNoIDAllGetResponse
}

func (o *GetApi24ProtectionGroupsHostsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/protection-groups/hosts][%d] getApi24ProtectionGroupsHostsOK  %+v", 200, o.Payload)
}

func (o *GetApi24ProtectionGroupsHostsOK) GetPayload() *models.MemberNoIDAllGetResponse {
	return o.Payload
}

func (o *GetApi24ProtectionGroupsHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
