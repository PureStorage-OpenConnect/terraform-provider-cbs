// Code generated by go-swagger; DO NOT EDIT.

package volume_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24VolumeGroupsVolumesReader is a Reader for the GetAPI24VolumeGroupsVolumes structure.
type GetAPI24VolumeGroupsVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24VolumeGroupsVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24VolumeGroupsVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24VolumeGroupsVolumesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24VolumeGroupsVolumesOK creates a GetApi24VolumeGroupsVolumesOK with default headers values
func NewGetApi24VolumeGroupsVolumesOK() *GetApi24VolumeGroupsVolumesOK {
	return &GetApi24VolumeGroupsVolumesOK{}
}

/*GetApi24VolumeGroupsVolumesOK handles this case with default header values.

OK
*/
type GetApi24VolumeGroupsVolumesOK struct {
	Payload *models.MemberGetResponse
}

func (o *GetApi24VolumeGroupsVolumesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volume-groups/volumes][%d] getApi24VolumeGroupsVolumesOK  %+v", 200, o.Payload)
}

func (o *GetApi24VolumeGroupsVolumesOK) GetPayload() *models.MemberGetResponse {
	return o.Payload
}

func (o *GetApi24VolumeGroupsVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MemberGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24VolumeGroupsVolumesBadRequest creates a GetApi24VolumeGroupsVolumesBadRequest with default headers values
func NewGetApi24VolumeGroupsVolumesBadRequest() *GetApi24VolumeGroupsVolumesBadRequest {
	return &GetApi24VolumeGroupsVolumesBadRequest{}
}

/*GetApi24VolumeGroupsVolumesBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24VolumeGroupsVolumesBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24VolumeGroupsVolumesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volume-groups/volumes][%d] getApi24VolumeGroupsVolumesBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24VolumeGroupsVolumesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24VolumeGroupsVolumesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
