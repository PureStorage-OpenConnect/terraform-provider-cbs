// Code generated by go-swagger; DO NOT EDIT.

package syslog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.dev.purestorage.com/FlashArray/terraform-provider-cbs/cbs/internal/faclient/2.4/models"
)

// GetAPI24SyslogServersTestReader is a Reader for the GetAPI24SyslogServersTest structure.
type GetAPI24SyslogServersTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24SyslogServersTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24SyslogServersTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24SyslogServersTestOK creates a GetApi24SyslogServersTestOK with default headers values
func NewGetApi24SyslogServersTestOK() *GetApi24SyslogServersTestOK {
	return &GetApi24SyslogServersTestOK{}
}

/*GetApi24SyslogServersTestOK handles this case with default header values.

OK
*/
type GetApi24SyslogServersTestOK struct {
	Payload *models.TestResultWithResourceGetResponse
}

func (o *GetApi24SyslogServersTestOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/syslog-servers/test][%d] getApi24SyslogServersTestOK  %+v", 200, o.Payload)
}

func (o *GetApi24SyslogServersTestOK) GetPayload() *models.TestResultWithResourceGetResponse {
	return o.Payload
}

func (o *GetApi24SyslogServersTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TestResultWithResourceGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
