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

// GetAPI24VolumeGroupsPerformanceReader is a Reader for the GetAPI24VolumeGroupsPerformance structure.
type GetAPI24VolumeGroupsPerformanceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24VolumeGroupsPerformanceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24VolumeGroupsPerformanceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24VolumeGroupsPerformanceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24VolumeGroupsPerformanceOK creates a GetApi24VolumeGroupsPerformanceOK with default headers values
func NewGetApi24VolumeGroupsPerformanceOK() *GetApi24VolumeGroupsPerformanceOK {
	return &GetApi24VolumeGroupsPerformanceOK{}
}

/*GetApi24VolumeGroupsPerformanceOK handles this case with default header values.

OK
*/
type GetApi24VolumeGroupsPerformanceOK struct {
	Payload *models.ResourcePerformanceGetResponse
}

func (o *GetApi24VolumeGroupsPerformanceOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volume-groups/performance][%d] getApi24VolumeGroupsPerformanceOK  %+v", 200, o.Payload)
}

func (o *GetApi24VolumeGroupsPerformanceOK) GetPayload() *models.ResourcePerformanceGetResponse {
	return o.Payload
}

func (o *GetApi24VolumeGroupsPerformanceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourcePerformanceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24VolumeGroupsPerformanceBadRequest creates a GetApi24VolumeGroupsPerformanceBadRequest with default headers values
func NewGetApi24VolumeGroupsPerformanceBadRequest() *GetApi24VolumeGroupsPerformanceBadRequest {
	return &GetApi24VolumeGroupsPerformanceBadRequest{}
}

/*GetApi24VolumeGroupsPerformanceBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24VolumeGroupsPerformanceBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24VolumeGroupsPerformanceBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/volume-groups/performance][%d] getApi24VolumeGroupsPerformanceBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24VolumeGroupsPerformanceBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24VolumeGroupsPerformanceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
