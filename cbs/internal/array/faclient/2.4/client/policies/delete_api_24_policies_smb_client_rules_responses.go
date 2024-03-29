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

// DeleteAPI24PoliciesSmbClientRulesReader is a Reader for the DeleteAPI24PoliciesSmbClientRules structure.
type DeleteAPI24PoliciesSmbClientRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24PoliciesSmbClientRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24PoliciesSmbClientRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteApi24PoliciesSmbClientRulesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24PoliciesSmbClientRulesOK creates a DeleteApi24PoliciesSmbClientRulesOK with default headers values
func NewDeleteApi24PoliciesSmbClientRulesOK() *DeleteApi24PoliciesSmbClientRulesOK {
	return &DeleteApi24PoliciesSmbClientRulesOK{}
}

/*DeleteApi24PoliciesSmbClientRulesOK handles this case with default header values.

OK
*/
type DeleteApi24PoliciesSmbClientRulesOK struct {
}

func (o *DeleteApi24PoliciesSmbClientRulesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/policies/smb/client-rules][%d] deleteApi24PoliciesSmbClientRulesOK ", 200)
}

func (o *DeleteApi24PoliciesSmbClientRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteApi24PoliciesSmbClientRulesBadRequest creates a DeleteApi24PoliciesSmbClientRulesBadRequest with default headers values
func NewDeleteApi24PoliciesSmbClientRulesBadRequest() *DeleteApi24PoliciesSmbClientRulesBadRequest {
	return &DeleteApi24PoliciesSmbClientRulesBadRequest{}
}

/*DeleteApi24PoliciesSmbClientRulesBadRequest handles this case with default header values.

BadRequest
*/
type DeleteApi24PoliciesSmbClientRulesBadRequest struct {
	Payload *models.Error
}

func (o *DeleteApi24PoliciesSmbClientRulesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/policies/smb/client-rules][%d] deleteApi24PoliciesSmbClientRulesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteApi24PoliciesSmbClientRulesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteApi24PoliciesSmbClientRulesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
