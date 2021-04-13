// Code generated by go-swagger; DO NOT EDIT.

package kmip

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PostAPI24KMIPReader is a Reader for the PostAPI24KMIP structure.
type PostAPI24KMIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24KMIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24KMIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24KMIPOK creates a PostApi24KMIPOK with default headers values
func NewPostApi24KMIPOK() *PostApi24KMIPOK {
	return &PostApi24KMIPOK{}
}

/*PostApi24KMIPOK handles this case with default header values.

Returns the newly created KMIP server object.
*/
type PostApi24KMIPOK struct {
	Payload *models.KMIPResponse
}

func (o *PostApi24KMIPOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/kmip][%d] postApi24KmipOK  %+v", 200, o.Payload)
}

func (o *PostApi24KMIPOK) GetPayload() *models.KMIPResponse {
	return o.Payload
}

func (o *PostApi24KMIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.KMIPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}