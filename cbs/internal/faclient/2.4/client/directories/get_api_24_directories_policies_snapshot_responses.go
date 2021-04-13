// Code generated by go-swagger; DO NOT EDIT.

package directories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24DirectoriesPoliciesSnapshotReader is a Reader for the GetAPI24DirectoriesPoliciesSnapshot structure.
type GetAPI24DirectoriesPoliciesSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24DirectoriesPoliciesSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24DirectoriesPoliciesSnapshotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24DirectoriesPoliciesSnapshotOK creates a GetApi24DirectoriesPoliciesSnapshotOK with default headers values
func NewGetApi24DirectoriesPoliciesSnapshotOK() *GetApi24DirectoriesPoliciesSnapshotOK {
	return &GetApi24DirectoriesPoliciesSnapshotOK{}
}

/*GetApi24DirectoriesPoliciesSnapshotOK handles this case with default header values.

OK
*/
type GetApi24DirectoriesPoliciesSnapshotOK struct {
	Payload *models.PolicyMemberGetResponse
}

func (o *GetApi24DirectoriesPoliciesSnapshotOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/directories/policies/snapshot][%d] getApi24DirectoriesPoliciesSnapshotOK  %+v", 200, o.Payload)
}

func (o *GetApi24DirectoriesPoliciesSnapshotOK) GetPayload() *models.PolicyMemberGetResponse {
	return o.Payload
}

func (o *GetApi24DirectoriesPoliciesSnapshotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyMemberGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
