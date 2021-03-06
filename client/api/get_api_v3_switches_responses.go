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

// GetAPIV3SwitchesReader is a Reader for the GetAPIV3Switches structure.
type GetAPIV3SwitchesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3SwitchesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV3SwitchesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV3SwitchesOK creates a GetAPIV3SwitchesOK with default headers values
func NewGetAPIV3SwitchesOK() *GetAPIV3SwitchesOK {
	return &GetAPIV3SwitchesOK{}
}

/*GetAPIV3SwitchesOK handles this case with default header values.

Return a list of switches, possibly filtered
*/
type GetAPIV3SwitchesOK struct {
	Payload models.GetAPIV3SwitchesOKBody
}

func (o *GetAPIV3SwitchesOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/switches][%d] getApiV3SwitchesOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3SwitchesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
