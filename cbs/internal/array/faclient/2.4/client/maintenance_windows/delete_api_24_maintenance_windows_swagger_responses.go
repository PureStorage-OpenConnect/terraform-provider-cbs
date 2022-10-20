// Code generated by go-swagger; DO NOT EDIT.

package maintenance_windows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// DeleteAPI24MaintenanceWindowsReader is a Reader for the DeleteAPI24MaintenanceWindows structure.
type DeleteAPI24MaintenanceWindowsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24MaintenanceWindowsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24MaintenanceWindowsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24MaintenanceWindowsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24MaintenanceWindowsOK creates a DeleteApi24MaintenanceWindowsOK with default headers values
func NewDeleteApi24MaintenanceWindowsOK() *DeleteApi24MaintenanceWindowsOK {
	return &DeleteApi24MaintenanceWindowsOK{}
}

/*DeleteApi24MaintenanceWindowsOK handles this case with default header values.

OK
*/
type DeleteApi24MaintenanceWindowsOK struct {
}

func (o *DeleteApi24MaintenanceWindowsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/maintenance-windows][%d] deleteApi24MaintenanceWindowsOK ", 200)
}

func (o *DeleteApi24MaintenanceWindowsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24MaintenanceWindowsBadRequest creates a DeleteApi24MaintenanceWindowsBadRequest with default headers values
func NewDeleteApi24MaintenanceWindowsBadRequest() *DeleteApi24MaintenanceWindowsBadRequest {
	return &DeleteApi24MaintenanceWindowsBadRequest{}
}

/*DeleteApi24MaintenanceWindowsBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24MaintenanceWindowsBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24MaintenanceWindowsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/maintenance-windows][%d] deleteApi24MaintenanceWindowsBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24MaintenanceWindowsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24MaintenanceWindowsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
