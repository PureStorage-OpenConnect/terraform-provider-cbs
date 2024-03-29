// Code generated by go-swagger; DO NOT EDIT.

package host_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24HostGroupsPerformanceByArrayReader is a Reader for the GetAPI24HostGroupsPerformanceByArray structure.
type GetAPI24HostGroupsPerformanceByArrayReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24HostGroupsPerformanceByArrayReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24HostGroupsPerformanceByArrayOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24HostGroupsPerformanceByArrayBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24HostGroupsPerformanceByArrayOK creates a GetApi24HostGroupsPerformanceByArrayOK with default headers values
func NewGetApi24HostGroupsPerformanceByArrayOK() *GetApi24HostGroupsPerformanceByArrayOK {
	return &GetApi24HostGroupsPerformanceByArrayOK{}
}

/*GetApi24HostGroupsPerformanceByArrayOK handles this case with default header values.

OK
*/
type GetApi24HostGroupsPerformanceByArrayOK struct {
	Payload *models.ResourcePerformanceNoIDByArrayGetResponse
}

func (o *GetApi24HostGroupsPerformanceByArrayOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/host-groups/performance/by-array][%d] getApi24HostGroupsPerformanceByArrayOK  %+v", 200, o.Payload)
}

func (o *GetApi24HostGroupsPerformanceByArrayOK) GetPayload() *models.ResourcePerformanceNoIDByArrayGetResponse {
	return o.Payload
}

func (o *GetApi24HostGroupsPerformanceByArrayOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourcePerformanceNoIDByArrayGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24HostGroupsPerformanceByArrayBadRequest creates a GetApi24HostGroupsPerformanceByArrayBadRequest with default headers values
func NewGetApi24HostGroupsPerformanceByArrayBadRequest() *GetApi24HostGroupsPerformanceByArrayBadRequest {
	return &GetApi24HostGroupsPerformanceByArrayBadRequest{}
}

/*GetApi24HostGroupsPerformanceByArrayBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24HostGroupsPerformanceByArrayBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24HostGroupsPerformanceByArrayBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/host-groups/performance/by-array][%d] getApi24HostGroupsPerformanceByArrayBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24HostGroupsPerformanceByArrayBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24HostGroupsPerformanceByArrayBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
