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

// DeleteAPI24VolumeGroupsReader is a Reader for the DeleteAPI24VolumeGroups structure.
type DeleteAPI24VolumeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24VolumeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24VolumeGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24VolumeGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24VolumeGroupsOK creates a DeleteApi24VolumeGroupsOK with default headers values
func NewDeleteApi24VolumeGroupsOK() *DeleteApi24VolumeGroupsOK {
	return &DeleteApi24VolumeGroupsOK{}
}

/*DeleteApi24VolumeGroupsOK handles this case with default header values.

OK
*/
type DeleteApi24VolumeGroupsOK struct {
}

func (o *DeleteApi24VolumeGroupsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/volume-groups][%d] deleteApi24VolumeGroupsOK ", 200)
}

func (o *DeleteApi24VolumeGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24VolumeGroupsBadRequest creates a DeleteApi24VolumeGroupsBadRequest with default headers values
func NewDeleteApi24VolumeGroupsBadRequest() *DeleteApi24VolumeGroupsBadRequest {
	return &DeleteApi24VolumeGroupsBadRequest{}
}

/*DeleteApi24VolumeGroupsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24VolumeGroupsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24VolumeGroupsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/volume-groups][%d] deleteApi24VolumeGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24VolumeGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24VolumeGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
