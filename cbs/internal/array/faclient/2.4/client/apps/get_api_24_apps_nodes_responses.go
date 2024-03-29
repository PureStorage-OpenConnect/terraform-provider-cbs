// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24AppsNodesReader is a Reader for the GetAPI24AppsNodes structure.
type GetAPI24AppsNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24AppsNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24AppsNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24AppsNodesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24AppsNodesOK creates a GetApi24AppsNodesOK with default headers values
func NewGetApi24AppsNodesOK() *GetApi24AppsNodesOK {
	return &GetApi24AppsNodesOK{}
}

/*GetApi24AppsNodesOK handles this case with default header values.

OK
*/
type GetApi24AppsNodesOK struct {
	Payload *models.AppNodeGetResponse
}

func (o *GetApi24AppsNodesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/apps/nodes][%d] getApi24AppsNodesOK  %+v", 200, o.Payload)
}

func (o *GetApi24AppsNodesOK) GetPayload() *models.AppNodeGetResponse {
	return o.Payload
}

func (o *GetApi24AppsNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppNodeGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24AppsNodesBadRequest creates a GetApi24AppsNodesBadRequest with default headers values
func NewGetApi24AppsNodesBadRequest() *GetApi24AppsNodesBadRequest {
	return &GetApi24AppsNodesBadRequest{}
}

/*GetApi24AppsNodesBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24AppsNodesBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24AppsNodesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/apps/nodes][%d] getApi24AppsNodesBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24AppsNodesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24AppsNodesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
