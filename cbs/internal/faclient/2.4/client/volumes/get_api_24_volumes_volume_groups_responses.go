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

// GetAPI24VolumesVolumeGroupsReader is a Reader for the GetAPI24VolumesVolumeGroups structure.
type GetAPI24VolumesVolumeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24VolumesVolumeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24VolumesVolumeGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24VolumesVolumeGroupsOK creates a GetApi24VolumesVolumeGroupsOK with default headers values
func NewGetApi24VolumesVolumeGroupsOK() *GetApi24VolumesVolumeGroupsOK {
	return &GetApi24VolumesVolumeGroupsOK{}
}

/*GetApi24VolumesVolumeGroupsOK handles this case with default header values.

OK
*/
type GetApi24VolumesVolumeGroupsOK struct {
	Payload *models.MemberGetResponse
}

func (o *GetApi24VolumesVolumeGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volumes/volume-groups][%d] getApi24VolumesVolumeGroupsOK  %+v", 200, o.Payload)
}

func (o *GetApi24VolumesVolumeGroupsOK) GetPayload() *models.MemberGetResponse {
	return o.Payload
}

func (o *GetApi24VolumesVolumeGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}