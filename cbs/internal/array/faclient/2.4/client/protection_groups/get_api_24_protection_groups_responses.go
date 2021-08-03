// Code generated by go-swagger; DO NOT EDIT.

package protection_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24ProtectionGroupsReader is a Reader for the GetAPI24ProtectionGroups structure.
type GetAPI24ProtectionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24ProtectionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24ProtectionGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24ProtectionGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24ProtectionGroupsOK creates a GetApi24ProtectionGroupsOK with default headers values
func NewGetApi24ProtectionGroupsOK() *GetApi24ProtectionGroupsOK {
	return &GetApi24ProtectionGroupsOK{}
}

/*GetApi24ProtectionGroupsOK handles this case with default header values.

OK
*/
type GetApi24ProtectionGroupsOK struct {
	Payload *models.ProtectionGroupGetResponse
}

func (o *GetApi24ProtectionGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/protection-groups][%d] getApi24ProtectionGroupsOK  %+v", 200, o.Payload)
}

func (o *GetApi24ProtectionGroupsOK) GetPayload() *models.ProtectionGroupGetResponse {
	return o.Payload
}

func (o *GetApi24ProtectionGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProtectionGroupGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24ProtectionGroupsBadRequest creates a GetApi24ProtectionGroupsBadRequest with default headers values
func NewGetApi24ProtectionGroupsBadRequest() *GetApi24ProtectionGroupsBadRequest {
	return &GetApi24ProtectionGroupsBadRequest{}
}

/*GetApi24ProtectionGroupsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24ProtectionGroupsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24ProtectionGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/protection-groups][%d] getApi24ProtectionGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24ProtectionGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24ProtectionGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}