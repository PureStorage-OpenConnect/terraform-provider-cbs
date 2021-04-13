// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// PatchAPI24CertificatesReader is a Reader for the PatchAPI24Certificates structure.
type PatchAPI24CertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24CertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24CertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24CertificatesOK creates a PatchApi24CertificatesOK with default headers values
func NewPatchApi24CertificatesOK() *PatchApi24CertificatesOK {
	return &PatchApi24CertificatesOK{}
}

/*PatchApi24CertificatesOK handles this case with default header values.

Returns the newly updated certificate object.
*/
type PatchApi24CertificatesOK struct {
	Payload *models.CertificateResponse
}

func (o *PatchApi24CertificatesOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/certificates][%d] patchApi24CertificatesOK  %+v", 200, o.Payload)
}

func (o *PatchApi24CertificatesOK) GetPayload() *models.CertificateResponse {
	return o.Payload
}

func (o *PatchApi24CertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CertificateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
