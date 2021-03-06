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

// PutAPIV3NicsIDReader is a Reader for the PutAPIV3NicsID structure.
type PutAPIV3NicsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutAPIV3NicsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutAPIV3NicsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutAPIV3NicsIDOK creates a PutAPIV3NicsIDOK with default headers values
func NewPutAPIV3NicsIDOK() *PutAPIV3NicsIDOK {
	return &PutAPIV3NicsIDOK{}
}

/*PutAPIV3NicsIDOK handles this case with default header values.

Update a single nic
*/
type PutAPIV3NicsIDOK struct {
	Payload *models.Nic
}

func (o *PutAPIV3NicsIDOK) Error() string {
	return fmt.Sprintf("[PUT /api/v3/nics/{id}][%d] putApiV3NicsIdOK  %+v", 200, o.Payload)
}

func (o *PutAPIV3NicsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Nic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
