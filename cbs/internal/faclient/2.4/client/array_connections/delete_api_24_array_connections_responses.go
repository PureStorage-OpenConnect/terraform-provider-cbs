// Code generated by go-swagger; DO NOT EDIT.

package array_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24ArrayConnectionsReader is a Reader for the DeleteAPI24ArrayConnections structure.
type DeleteAPI24ArrayConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24ArrayConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24ArrayConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24ArrayConnectionsOK creates a DeleteApi24ArrayConnectionsOK with default headers values
func NewDeleteApi24ArrayConnectionsOK() *DeleteApi24ArrayConnectionsOK {
	return &DeleteApi24ArrayConnectionsOK{}
}

/*DeleteApi24ArrayConnectionsOK handles this case with default header values.

OK
*/
type DeleteApi24ArrayConnectionsOK struct {
}

func (o *DeleteApi24ArrayConnectionsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/array-connections][%d] deleteApi24ArrayConnectionsOK ", 200)
}

func (o *DeleteApi24ArrayConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
