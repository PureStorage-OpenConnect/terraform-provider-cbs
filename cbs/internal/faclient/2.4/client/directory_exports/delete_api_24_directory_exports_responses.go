// Code generated by go-swagger; DO NOT EDIT.

package directory_exports

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteAPI24DirectoryExportsReader is a Reader for the DeleteAPI24DirectoryExports structure.
type DeleteAPI24DirectoryExportsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPI24DirectoryExportsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteApi24DirectoryExportsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteApi24DirectoryExportsOK creates a DeleteApi24DirectoryExportsOK with default headers values
func NewDeleteApi24DirectoryExportsOK() *DeleteApi24DirectoryExportsOK {
	return &DeleteApi24DirectoryExportsOK{}
}

/*DeleteApi24DirectoryExportsOK handles this case with default header values.

OK
*/
type DeleteApi24DirectoryExportsOK struct {
}

func (o *DeleteApi24DirectoryExportsOK) Error() string {
	return fmt.Sprintf("[DELETE /api/2.4/directory-exports][%d] deleteApi24DirectoryExportsOK ", 200)
}

func (o *DeleteApi24DirectoryExportsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
