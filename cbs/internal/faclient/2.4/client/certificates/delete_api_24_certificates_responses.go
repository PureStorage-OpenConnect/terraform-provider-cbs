// Code generated by go-swagger; DO NOT EDIT.

package certificates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24CertificatesReader is a Reader for the DeleteAPI24Certificates structure.
type DeleteAPI24CertificatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24CertificatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24CertificatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24CertificatesOK creates a DeleteApi24CertificatesOK with default headers values
func NewDeleteApi24CertificatesOK() *DeleteApi24CertificatesOK {
	return &DeleteApi24CertificatesOK{}
}

/*DeleteApi24CertificatesOK handles this case with default header values.

OK
*/
type DeleteApi24CertificatesOK struct {
}

func (o *DeleteApi24CertificatesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/certificates][%d] deleteApi24CertificatesOK ", 200)
}

func (o *DeleteApi24CertificatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}