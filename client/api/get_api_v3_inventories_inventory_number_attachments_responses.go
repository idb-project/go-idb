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

// GetAPIV3InventoriesInventoryNumberAttachmentsReader is a Reader for the GetAPIV3InventoriesInventoryNumberAttachments structure.
type GetAPIV3InventoriesInventoryNumberAttachmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV3InventoriesInventoryNumberAttachmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV3InventoriesInventoryNumberAttachmentsOK creates a GetAPIV3InventoriesInventoryNumberAttachmentsOK with default headers values
func NewGetAPIV3InventoriesInventoryNumberAttachmentsOK() *GetAPIV3InventoriesInventoryNumberAttachmentsOK {
	return &GetAPIV3InventoriesInventoryNumberAttachmentsOK{}
}

/*GetAPIV3InventoriesInventoryNumberAttachmentsOK handles this case with default header values.

Get all attachments
*/
type GetAPIV3InventoriesInventoryNumberAttachmentsOK struct {
	Payload models.GetAPIV3InventoriesInventoryNumberAttachmentsOKBody
}

func (o *GetAPIV3InventoriesInventoryNumberAttachmentsOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/inventories/{inventory_number}/attachments][%d] getApiV3InventoriesInventoryNumberAttachmentsOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3InventoriesInventoryNumberAttachmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}