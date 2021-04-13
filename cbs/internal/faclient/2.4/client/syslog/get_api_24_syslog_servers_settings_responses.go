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

// GetAPI24SyslogServersSettingsReader is a Reader for the GetAPI24SyslogServersSettings structure.
type GetAPI24SyslogServersSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPI24SyslogServersSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetApi24SyslogServersSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetApi24SyslogServersSettingsOK creates a GetApi24SyslogServersSettingsOK with default headers values
func NewGetApi24SyslogServersSettingsOK() *GetApi24SyslogServersSettingsOK {
	return &GetApi24SyslogServersSettingsOK{}
}

/*GetApi24SyslogServersSettingsOK handles this case with default header values.

OK
*/
type GetApi24SyslogServersSettingsOK struct {
	Payload *models.SyslogServerSettingsGetResponse
}

func (o *GetApi24SyslogServersSettingsOK) Error() string {
	return fmt.Sprintf("[GET /api/2.4/syslog-servers/settings][%d] getApi24SyslogServersSettingsOK  %+v", 200, o.Payload)
}

func (o *GetApi24SyslogServersSettingsOK) GetPayload() *models.SyslogServerSettingsGetResponse {
	return o.Payload
}

func (o *GetApi24SyslogServersSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SyslogServerSettingsGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}