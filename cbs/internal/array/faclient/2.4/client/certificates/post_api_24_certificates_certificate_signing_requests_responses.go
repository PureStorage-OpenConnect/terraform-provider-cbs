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

// PostAPI24CertificatesCertificateSigningRequestsReader is a Reader for the PostAPI24CertificatesCertificateSigningRequests structure.
type PostAPI24CertificatesCertificateSigningRequestsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPI24CertificatesCertificateSigningRequestsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostApi24CertificatesCertificateSigningRequestsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostApi24CertificatesCertificateSigningRequestsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostApi24CertificatesCertificateSigningRequestsOK creates a PostApi24CertificatesCertificateSigningRequestsOK with default headers values
func NewPostApi24CertificatesCertificateSigningRequestsOK() *PostApi24CertificatesCertificateSigningRequestsOK {
	return &PostApi24CertificatesCertificateSigningRequestsOK{}
}

/*PostApi24CertificatesCertificateSigningRequestsOK handles this case with default header values.

Returns the newly created certificate object.
*/
type PostApi24CertificatesCertificateSigningRequestsOK struct {
	Payload *models.CertificateSigningRequestResponse
}

func (o *PostApi24CertificatesCertificateSigningRequestsOK) Error() string {
	return fmt.Sprintf("[POST /api/2.4/certificates/certificate-signing-requests][%d] postApi24CertificatesCertificateSigningRequestsOK  %+v", 200, o.Payload)
}

func (o *PostApi24CertificatesCertificateSigningRequestsOK) GetPayload() *models.CertificateSigningRequestResponse {
	return o.Payload
}

func (o *PostApi24CertificatesCertificateSigningRequestsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateSigningRequestResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostApi24CertificatesCertificateSigningRequestsBadRequest creates a PostApi24CertificatesCertificateSigningRequestsBadRequest with default headers values
func NewPostApi24CertificatesCertificateSigningRequestsBadRequest() *PostApi24CertificatesCertificateSigningRequestsBadRequest {
	return &PostApi24CertificatesCertificateSigningRequestsBadRequest{}
}

/*PostApi24CertificatesCertificateSigningRequestsBadRequest handles this case with default header values.

BadRequest
*/
type PostApi24CertificatesCertificateSigningRequestsBadRequest struct {
	Payload *models.Error
}

func (o *PostApi24CertificatesCertificateSigningRequestsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/2.4/certificates/certificate-signing-requests][%d] postApi24CertificatesCertificateSigningRequestsBadRequest  %+v", 400, o.Payload)
}

func (o *PostApi24CertificatesCertificateSigningRequestsBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PostApi24CertificatesCertificateSigningRequestsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
