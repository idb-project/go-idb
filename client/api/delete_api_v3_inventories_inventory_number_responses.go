// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAPIV3InventoriesInventoryNumberReader is a Reader for the DeleteAPIV3InventoriesInventoryNumber structure.
type DeleteAPIV3InventoriesInventoryNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAPIV3InventoriesInventoryNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAPIV3InventoriesInventoryNumberNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAPIV3InventoriesInventoryNumberNoContent creates a DeleteAPIV3InventoriesInventoryNumberNoContent with default headers values
func NewDeleteAPIV3InventoriesInventoryNumberNoContent() *DeleteAPIV3InventoriesInventoryNumberNoContent {
	return &DeleteAPIV3InventoriesInventoryNumberNoContent{}
}

/*DeleteAPIV3InventoriesInventoryNumberNoContent handles this case with default header values.

Delete a inventory
*/
type DeleteAPIV3InventoriesInventoryNumberNoContent struct {
}

func (o *DeleteAPIV3InventoriesInventoryNumberNoContent) Error() string {
	return fmt.Sprintf("[DELETE /api/v3/inventories/{inventory_number}][%d] deleteApiV3InventoriesInventoryNumberNoContent ", 204)
}

func (o *DeleteAPIV3InventoriesInventoryNumberNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
