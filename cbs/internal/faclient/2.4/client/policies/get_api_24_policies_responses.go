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

// GetAPI24PoliciesReader is a Reader for the GetAPI24Policies structure.
type GetAPI24PoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24PoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24PoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24PoliciesOK creates a GetApi24PoliciesOK with default headers values
func NewGetApi24PoliciesOK() *GetApi24PoliciesOK {
	return &GetApi24PoliciesOK{}
}

/*GetApi24PoliciesOK handles this case with default header values.

OK
*/
type GetApi24PoliciesOK struct {
	Payload *models.PolicyGetResponse
}

func (o *GetApi24PoliciesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/policies][%d] getApi24PoliciesOK  %+v", 200, o.Payload)
}

func (o *GetApi24PoliciesOK) GetPayload() *models.PolicyGetResponse {
	return o.Payload
}

func (o *GetApi24PoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
