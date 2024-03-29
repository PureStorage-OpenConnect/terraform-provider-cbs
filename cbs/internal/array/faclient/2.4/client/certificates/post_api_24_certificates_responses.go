// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PostAPI24CertificatesReader is a Reader for the PostAPI24Certificates structure.
type PostAPI24CertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24CertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24CertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24CertificatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24CertificatesOK creates a PostApi24CertificatesOK with default headers values
func NewPostApi24CertificatesOK() *PostApi24CertificatesOK {
	return &PostApi24CertificatesOK{}
}

/*PostApi24CertificatesOK handles this case with default header values.

Returns the newly created certificate object.
*/
type PostApi24CertificatesOK struct {
	Payload *models.CertificateResponse
}

func (o *PostApi24CertificatesOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/certificates][%d] postApi24CertificatesOK  %+v", 200, o.Payload)
}

func (o *PostApi24CertificatesOK) GetPayload() *models.CertificateResponse {
	return o.Payload
}

func (o *PostApi24CertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24CertificatesBadRequest creates a PostApi24CertificatesBadRequest with default headers values
func NewPostApi24CertificatesBadRequest() *PostApi24CertificatesBadRequest {
	return &PostApi24CertificatesBadRequest{}
}

/*PostApi24CertificatesBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24CertificatesBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24CertificatesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/certificates][%d] postApi24CertificatesBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24CertificatesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24CertificatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
