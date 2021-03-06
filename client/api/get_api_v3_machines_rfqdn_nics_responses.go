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

// GetAPIV3MachinesRfqdnNicsReader is a Reader for the GetAPIV3MachinesRfqdnNics structure.
type GetAPIV3MachinesRfqdnNicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3MachinesRfqdnNicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV3MachinesRfqdnNicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV3MachinesRfqdnNicsOK creates a GetAPIV3MachinesRfqdnNicsOK with default headers values
func NewGetAPIV3MachinesRfqdnNicsOK() *GetAPIV3MachinesRfqdnNicsOK {
	return &GetAPIV3MachinesRfqdnNicsOK{}
}

/*GetAPIV3MachinesRfqdnNicsOK handles this case with default header values.

Get all nics
*/
type GetAPIV3MachinesRfqdnNicsOK struct {
	Payload models.GetAPIV3MachinesRfqdnNicsOKBody
}

func (o *GetAPIV3MachinesRfqdnNicsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/machines/{rfqdn}/nics][%d] getApiV3MachinesRfqdnNicsOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3MachinesRfqdnNicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
