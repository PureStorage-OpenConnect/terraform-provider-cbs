// Code generated by go-swagger; DO NOT EDIT.

package remote_protection_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24RemoteProtectionGroupsReader is a Reader for the GetAPI24RemoteProtectionGroups structure.
type GetAPI24RemoteProtectionGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24RemoteProtectionGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24RemoteProtectionGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24RemoteProtectionGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24RemoteProtectionGroupsOK creates a GetApi24RemoteProtectionGroupsOK with default headers values
func NewGetApi24RemoteProtectionGroupsOK() *GetApi24RemoteProtectionGroupsOK {
	return &GetApi24RemoteProtectionGroupsOK{}
}

/*GetApi24RemoteProtectionGroupsOK handles this case with default header values.

OK
*/
type GetApi24RemoteProtectionGroupsOK struct {
	Payload *models.RemoteProtectionGroupGetResponse
}

func (o *GetApi24RemoteProtectionGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/remote-protection-groups][%d] getApi24RemoteProtectionGroupsOK  %+v", 200, o.Payload)
}

func (o *GetApi24RemoteProtectionGroupsOK) GetPayload() *models.RemoteProtectionGroupGetResponse {
	return o.Payload
}

func (o *GetApi24RemoteProtectionGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemoteProtectionGroupGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24RemoteProtectionGroupsBadRequest creates a GetApi24RemoteProtectionGroupsBadRequest with default headers values
func NewGetApi24RemoteProtectionGroupsBadRequest() *GetApi24RemoteProtectionGroupsBadRequest {
	return &GetApi24RemoteProtectionGroupsBadRequest{}
}

/*GetApi24RemoteProtectionGroupsBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24RemoteProtectionGroupsBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24RemoteProtectionGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/remote-protection-groups][%d] getApi24RemoteProtectionGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24RemoteProtectionGroupsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24RemoteProtectionGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
