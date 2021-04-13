// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24PoliciesSnapshotRulesReader is a Reader for the GetAPI24PoliciesSnapshotRules structure.
type GetAPI24PoliciesSnapshotRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24PoliciesSnapshotRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24PoliciesSnapshotRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24PoliciesSnapshotRulesOK creates a GetApi24PoliciesSnapshotRulesOK with default headers values
func NewGetApi24PoliciesSnapshotRulesOK() *GetApi24PoliciesSnapshotRulesOK {
	return &GetApi24PoliciesSnapshotRulesOK{}
}

/*GetApi24PoliciesSnapshotRulesOK handles this case with default header values.

OK
*/
type GetApi24PoliciesSnapshotRulesOK struct {
	Payload *models.PolicyRuleSnapshotGetResponse
}

func (o *GetApi24PoliciesSnapshotRulesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/policies/snapshot/rules][%d] getApi24PoliciesSnapshotRulesOK  %+v", 200, o.Payload)
}

func (o *GetApi24PoliciesSnapshotRulesOK) GetPayload() *models.PolicyRuleSnapshotGetResponse {
	return o.Payload
}

func (o *GetApi24PoliciesSnapshotRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyRuleSnapshotGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
