// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/idb-project/go-idb/models"
)

// PutAPIV3SwitchesFqdnPortsNumberReader is a Reader for the PutAPIV3SwitchesFqdnPortsNumber structure.
type PutAPIV3SwitchesFqdnPortsNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3SwitchesFqdnPortsNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAPIV3SwitchesFqdnPortsNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAPIV3SwitchesFqdnPortsNumberOK creates a PutAPIV3SwitchesFqdnPortsNumberOK with default headers values
func NewPutAPIV3SwitchesFqdnPortsNumberOK() *PutAPIV3SwitchesFqdnPortsNumberOK {
	return &PutAPIV3SwitchesFqdnPortsNumberOK{}
}

/*PutAPIV3SwitchesFqdnPortsNumberOK handles this case with default header values.

Update a switch port
*/
type PutAPIV3SwitchesFqdnPortsNumberOK struct {
	Payload *models.SwitchPort
}

func (o *PutAPIV3SwitchesFqdnPortsNumberOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/switches/{fqdn}/ports/{number}][%d] putApiV3SwitchesFqdnPortsNumberOK  %+v", 200, o.Payload)
}

func (o *PutAPIV3SwitchesFqdnPortsNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwitchPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
