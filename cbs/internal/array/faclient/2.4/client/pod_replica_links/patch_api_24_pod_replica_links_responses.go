// Code generated by go-swagger; DO NOT EDIT.

package pod_replica_links

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/PureStorage-OpenConnect/terraform-provider-cbs/cbs/internal/array/faclient/2.4/models"
)

// PatchAPI24PodReplicaLinksReader is a Reader for the PatchAPI24PodReplicaLinks structure.
type PatchAPI24PodReplicaLinksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchAPI24PodReplicaLinksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchApi24PodReplicaLinksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchApi24PodReplicaLinksBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchApi24PodReplicaLinksOK creates a PatchApi24PodReplicaLinksOK with default headers values
func NewPatchApi24PodReplicaLinksOK() *PatchApi24PodReplicaLinksOK {
	return &PatchApi24PodReplicaLinksOK{}
}

/*PatchApi24PodReplicaLinksOK handles this case with default header values.

OK
*/
type PatchApi24PodReplicaLinksOK struct {
	Payload *models.PodReplicaLinkResponse
}

func (o *PatchApi24PodReplicaLinksOK) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/pod-replica-links][%d] patchApi24PodReplicaLinksOK  %+v", 200, o.Payload)
}

func (o *PatchApi24PodReplicaLinksOK) GetPayload() *models.PodReplicaLinkResponse {
	return o.Payload
}

func (o *PatchApi24PodReplicaLinksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PodReplicaLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchApi24PodReplicaLinksBadRequest creates a PatchApi24PodReplicaLinksBadRequest with default headers values
func NewPatchApi24PodReplicaLinksBadRequest() *PatchApi24PodReplicaLinksBadRequest {
	return &PatchApi24PodReplicaLinksBadRequest{}
}

/*PatchApi24PodReplicaLinksBadRequest handles this case with default header values.

BadRequest
*/
type PatchApi24PodReplicaLinksBadRequest struct {
	Payload *models.Error
}

func (o *PatchApi24PodReplicaLinksBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/2.4/pod-replica-links][%d] patchApi24PodReplicaLinksBadRequest  %+v", 400, o.Payload)
}

func (o *PatchApi24PodReplicaLinksBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PatchApi24PodReplicaLinksBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
