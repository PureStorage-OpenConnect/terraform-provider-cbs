// Code generated by go-swagger; DO NOT EDIT.

package pods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24PodsArraysReader is a Reader for the DeleteAPI24PodsArrays structure.
type DeleteAPI24PodsArraysReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24PodsArraysReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24PodsArraysOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24PodsArraysOK creates a DeleteApi24PodsArraysOK with default headers values
func NewDeleteApi24PodsArraysOK() *DeleteApi24PodsArraysOK {
	return &DeleteApi24PodsArraysOK{}
}

/*DeleteApi24PodsArraysOK handles this case with default header values.

OK
*/
type DeleteApi24PodsArraysOK struct {
}

func (o *DeleteApi24PodsArraysOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/pods/arrays][%d] deleteApi24PodsArraysOK ", 200)
}

func (o *DeleteApi24PodsArraysOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
