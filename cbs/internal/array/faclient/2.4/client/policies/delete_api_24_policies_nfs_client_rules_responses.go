// Code generated by go-swagger; DO NOT EDIT.

package policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// DeleteAPI24PoliciesNfsClientRulesReader is a Reader for the DeleteAPI24PoliciesNfsClientRules structure.
type DeleteAPI24PoliciesNfsClientRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24PoliciesNfsClientRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24PoliciesNfsClientRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24PoliciesNfsClientRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24PoliciesNfsClientRulesOK creates a DeleteApi24PoliciesNfsClientRulesOK with default headers values
func NewDeleteApi24PoliciesNfsClientRulesOK() *DeleteApi24PoliciesNfsClientRulesOK {
	return &DeleteApi24PoliciesNfsClientRulesOK{}
}

/*DeleteApi24PoliciesNfsClientRulesOK handles this case with default header values.

OK
*/
type DeleteApi24PoliciesNfsClientRulesOK struct {
}

func (o *DeleteApi24PoliciesNfsClientRulesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/policies/nfs/client-rules][%d] deleteApi24PoliciesNfsClientRulesOK ", 200)
}

func (o *DeleteApi24PoliciesNfsClientRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24PoliciesNfsClientRulesBadRequest creates a DeleteApi24PoliciesNfsClientRulesBadRequest with default headers values
func NewDeleteApi24PoliciesNfsClientRulesBadRequest() *DeleteApi24PoliciesNfsClientRulesBadRequest {
	return &DeleteApi24PoliciesNfsClientRulesBadRequest{}
}

/*DeleteApi24PoliciesNfsClientRulesBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24PoliciesNfsClientRulesBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24PoliciesNfsClientRulesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/policies/nfs/client-rules][%d] deleteApi24PoliciesNfsClientRulesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24PoliciesNfsClientRulesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24PoliciesNfsClientRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
