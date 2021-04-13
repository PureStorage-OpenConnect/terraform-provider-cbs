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

// GetAPI24ProtectionGroupsVolumesReader is a Reader for the GetAPI24ProtectionGroupsVolumes structure.
type GetAPI24ProtectionGroupsVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24ProtectionGroupsVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24ProtectionGroupsVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24ProtectionGroupsVolumesOK creates a GetApi24ProtectionGroupsVolumesOK with default headers values
func NewGetApi24ProtectionGroupsVolumesOK() *GetApi24ProtectionGroupsVolumesOK {
	return &GetApi24ProtectionGroupsVolumesOK{}
}

/*GetApi24ProtectionGroupsVolumesOK handles this case with default header values.

OK
*/
type GetApi24ProtectionGroupsVolumesOK struct {
	Payload *models.MemberNoIDAllGetResponse
}

func (o *GetApi24ProtectionGroupsVolumesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/protection-groups/volumes][%d] getApi24ProtectionGroupsVolumesOK  %+v", 200, o.Payload)
}

func (o *GetApi24ProtectionGroupsVolumesOK) GetPayload() *models.MemberNoIDAllGetResponse {
	return o.Payload
}

func (o *GetApi24ProtectionGroupsVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberNoIDAllGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
