// Code generated by go-swagger; DO NOT EDIT.

package host_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24HostGroupsHostsReader is a Reader for the PostAPI24HostGroupsHosts structure.
type PostAPI24HostGroupsHostsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24HostGroupsHostsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24HostGroupsHostsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24HostGroupsHostsOK creates a PostApi24HostGroupsHostsOK with default headers values
func NewPostApi24HostGroupsHostsOK() *PostApi24HostGroupsHostsOK {
	return &PostApi24HostGroupsHostsOK{}
}

/*PostApi24HostGroupsHostsOK handles this case with default header values.

OK
*/
type PostApi24HostGroupsHostsOK struct {
	Payload *models.MemberNoIDAllResponse
}

func (o *PostApi24HostGroupsHostsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/host-groups/hosts][%d] postApi24HostGroupsHostsOK  %+v", 200, o.Payload)
}

func (o *PostApi24HostGroupsHostsOK) GetPayload() *models.MemberNoIDAllResponse {
	return o.Payload
}

func (o *PostApi24HostGroupsHostsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}