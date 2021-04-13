// Code generated by go-swagger; DO NOT EDIT.

package volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24VolumeGroupsReader is a Reader for the PostAPI24VolumeGroups structure.
type PostAPI24VolumeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24VolumeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24VolumeGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24VolumeGroupsOK creates a PostApi24VolumeGroupsOK with default headers values
func NewPostApi24VolumeGroupsOK() *PostApi24VolumeGroupsOK {
	return &PostApi24VolumeGroupsOK{}
}

/*PostApi24VolumeGroupsOK handles this case with default header values.

OK
*/
type PostApi24VolumeGroupsOK struct {
	Payload *models.VolumeGroupResponse
}

func (o *PostApi24VolumeGroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/volume-groups][%d] postApi24VolumeGroupsOK  %+v", 200, o.Payload)
}

func (o *PostApi24VolumeGroupsOK) GetPayload() *models.VolumeGroupResponse {
	return o.Payload
}

func (o *PostApi24VolumeGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGroupResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}