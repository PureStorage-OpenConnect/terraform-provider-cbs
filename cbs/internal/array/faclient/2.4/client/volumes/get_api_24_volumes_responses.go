// Code generated by go-swagger; DO NOT EDIT.

package volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24VolumesReader is a Reader for the GetAPI24Volumes structure.
type GetAPI24VolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24VolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24VolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24VolumesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24VolumesOK creates a GetApi24VolumesOK with default headers values
func NewGetApi24VolumesOK() *GetApi24VolumesOK {
	return &GetApi24VolumesOK{}
}

/*GetApi24VolumesOK handles this case with default header values.

OK
*/
type GetApi24VolumesOK struct {
	Payload *models.VolumeGetResponse
}

func (o *GetApi24VolumesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volumes][%d] getApi24VolumesOK  %+v", 200, o.Payload)
}

func (o *GetApi24VolumesOK) GetPayload() *models.VolumeGetResponse {
	return o.Payload
}

func (o *GetApi24VolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VolumeGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24VolumesBadRequest creates a GetApi24VolumesBadRequest with default headers values
func NewGetApi24VolumesBadRequest() *GetApi24VolumesBadRequest {
	return &GetApi24VolumesBadRequest{}
}

/*GetApi24VolumesBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24VolumesBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24VolumesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volumes][%d] getApi24VolumesBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24VolumesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24VolumesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
