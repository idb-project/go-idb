// Code generated by go-swagger; DO NOT EDIT.

package inventories

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/idb-project/go-idb/models"
)

// GetAPIV3InventoriesInventoryNumberReader is a Reader for the GetAPIV3InventoriesInventoryNumber structure.
type GetAPIV3InventoriesInventoryNumberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIV3InventoriesInventoryNumberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIV3InventoriesInventoryNumberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIV3InventoriesInventoryNumberOK creates a GetAPIV3InventoriesInventoryNumberOK with default headers values
func NewGetAPIV3InventoriesInventoryNumberOK() *GetAPIV3InventoriesInventoryNumberOK {
	return &GetAPIV3InventoriesInventoryNumberOK{}
}

/*GetAPIV3InventoriesInventoryNumberOK handles this case with default header values.

Get a inventory by inventory number
*/
type GetAPIV3InventoriesInventoryNumberOK struct {
	Payload *models.Inventory
}

func (o *GetAPIV3InventoriesInventoryNumberOK) Error() string {
	return fmt.Sprintf("[GET /api/v3/inventories/{inventory_number}][%d] getApiV3InventoriesInventoryNumberOK  %+v", 200, o.Payload)
}

func (o *GetAPIV3InventoriesInventoryNumberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Inventory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
