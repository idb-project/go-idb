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

// PostAPIV3NicsReader is a Reader for the PostAPIV3Nics structure.
type PostAPIV3NicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3NicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAPIV3NicsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIV3NicsCreated creates a PostAPIV3NicsCreated with default headers values
func NewPostAPIV3NicsCreated() *PostAPIV3NicsCreated {
	return &PostAPIV3NicsCreated{}
}

/*PostAPIV3NicsCreated handles this case with default header values.

Create a new nic
*/
type PostAPIV3NicsCreated struct {
	Payload *models.Nic
}

func (o *PostAPIV3NicsCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/nics][%d] postApiV3NicsCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3NicsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Nic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}