// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// GetAPI24CertificatesReader is a Reader for the GetAPI24Certificates structure.
type GetAPI24CertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24CertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24CertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetApi24CertificatesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24CertificatesOK creates a GetApi24CertificatesOK with default headers values
func NewGetApi24CertificatesOK() *GetApi24CertificatesOK {
	return &GetApi24CertificatesOK{}
}

/*GetApi24CertificatesOK handles this case with default header values.

OK
*/
type GetApi24CertificatesOK struct {
	Payload *models.CertificateGetResponse
}

func (o *GetApi24CertificatesOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/certificates][%d] getApi24CertificatesOK  %+v", 200, o.Payload)
}

func (o *GetApi24CertificatesOK) GetPayload() *models.CertificateGetResponse {
	return o.Payload
}

func (o *GetApi24CertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetApi24CertificatesBadRequest creates a GetApi24CertificatesBadRequest with default headers values
func NewGetApi24CertificatesBadRequest() *GetApi24CertificatesBadRequest {
	return &GetApi24CertificatesBadRequest{}
}

/*GetApi24CertificatesBadRequest handles this case with default header values.

BadRequest
*/
type GetApi24CertificatesBadRequest struct {
	Payload *models.Error
}

func (o *GetApi24CertificatesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/2.4/certificates][%d] getApi24CertificatesBadRequest  %+v", 400, o.Payload)
}

func (o *GetApi24CertificatesBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetApi24CertificatesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}