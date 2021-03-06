// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAPIV3MachinesRfqdnAttachmentsFingerprintReader is a Reader for the DeleteAPIV3MachinesRfqdnAttachmentsFingerprint structure.
type DeleteAPIV3MachinesRfqdnAttachmentsFingerprintReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV3MachinesRfqdnAttachmentsFingerprintReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent creates a DeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent with default headers values
func NewDeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent() *DeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent {
	return &DeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent{}
}

/*DeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent handles this case with default header values.

Delete an attachment
*/
type DeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent struct {
}

func (o *DeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/machines/{rfqdn}/attachments/{fingerprint}][%d] deleteApiV3MachinesRfqdnAttachmentsFingerprintNoContent ", 204)
}

func (o *DeleteAPIV3MachinesRfqdnAttachmentsFingerprintNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
