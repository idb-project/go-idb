// Code generated by go-swagger; DO NOT EDIT.

package nics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/idb-project/go-idb/models"
)

// GetAPIV3NicsIDReader is a Reader for the GetAPIV3NicsID structure.
type GetAPIV3NicsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3NicsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV3NicsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV3NicsIDOK creates a GetAPIV3NicsIDOK with default headers values
func NewGetAPIV3NicsIDOK() *GetAPIV3NicsIDOK {
	return &GetAPIV3NicsIDOK{}
}

/*GetAPIV3NicsIDOK handles this case with default header values.

Get a nic by id
*/
type GetAPIV3NicsIDOK struct {
	Payload *models.Nic
}

func (o *GetAPIV3NicsIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/nics/{id}][%d] getApiV3NicsIdOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3NicsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Nic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
