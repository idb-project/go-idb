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

// PostAPIV3MachinesRfqdnAliasesReader is a Reader for the PostAPIV3MachinesRfqdnAliases structure.
type PostAPIV3MachinesRfqdnAliasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIV3MachinesRfqdnAliasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostAPIV3MachinesRfqdnAliasesCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIV3MachinesRfqdnAliasesCreated creates a PostAPIV3MachinesRfqdnAliasesCreated with default headers values
func NewPostAPIV3MachinesRfqdnAliasesCreated() *PostAPIV3MachinesRfqdnAliasesCreated {
	return &PostAPIV3MachinesRfqdnAliasesCreated{}
}

/*PostAPIV3MachinesRfqdnAliasesCreated handles this case with default header values.

Create an alias
*/
type PostAPIV3MachinesRfqdnAliasesCreated struct {
	Payload *models.MachineAlias
}

func (o *PostAPIV3MachinesRfqdnAliasesCreated) Error() string {
	return fmt.Sprintf("[POST /api/v3/machines/{rfqdn}/aliases][%d] postApiV3MachinesRfqdnAliasesCreated  %+v", 201, o.Payload)
}

func (o *PostAPIV3MachinesRfqdnAliasesCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MachineAlias)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
