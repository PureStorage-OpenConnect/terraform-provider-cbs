// Code generated by go-swagger; DO NOT EDIT.

package network_interfaces

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24NetworkInterfacesReader is a Reader for the DeleteAPI24NetworkInterfaces structure.
type DeleteAPI24NetworkInterfacesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24NetworkInterfacesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24NetworkInterfacesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24NetworkInterfacesOK creates a DeleteApi24NetworkInterfacesOK with default headers values
func NewDeleteApi24NetworkInterfacesOK() *DeleteApi24NetworkInterfacesOK {
	return &DeleteApi24NetworkInterfacesOK{}
}

/*DeleteApi24NetworkInterfacesOK handles this case with default header values.

OK
*/
type DeleteApi24NetworkInterfacesOK struct {
}

func (o *DeleteApi24NetworkInterfacesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/network-interfaces][%d] deleteApi24NetworkInterfacesOK ", 200)
}

func (o *DeleteApi24NetworkInterfacesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
