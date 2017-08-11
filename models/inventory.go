// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Inventory Create a new inventory
// swagger:model Inventory

type Inventory struct {

	// Additional category description
	Category string `json:"category,omitempty"`

	// Comment field
	Comment string `json:"comment,omitempty"`

	// Creation date RFC3999 formatted
	CreatedAt string `json:"created_at,omitempty"`

	// Deletion date RFC3999 formatted
	DeletedAt string `json:"deleted_at,omitempty"`

	// Id
	ID int32 `json:"id,omitempty"`

	// Installation date as YYYY-MM-DD
	InstallDate string `json:"install_date,omitempty"`

	// Inventory Number
	InventoryNumber string `json:"inventory_number,omitempty"`

	// Inventory status
	InventoryStatus string `json:"inventory_status,omitempty"`

	// Inventory status id
	InventoryStatusID int32 `json:"inventory_status_id,omitempty"`

	// ID of the location
	LocationID int32 `json:"location_id,omitempty"`

	// machines FQDN if this inventoy is a machine
	Machine string `json:"machine,omitempty"`

	// Name
	Name string `json:"name,omitempty"`

	// Factory part number
	PartNumber string `json:"part_number,omitempty"`

	// Additional place description
	Place string `json:"place,omitempty"`

	// Purchase date as YYYY-MM-DD
	PurchaseDate string `json:"purchase_date,omitempty"`

	// Seller
	Seller string `json:"seller,omitempty"`

	// Factory serial number
	Serial string `json:"serial,omitempty"`

	// Update date RFC3999 formatted
	UpdatedAt string `json:"updated_at,omitempty"`

	// user id
	UserID string `json:"user_id,omitempty"`

	// Warranty end date as YYYY-MM-DD
	WarrantyEnd string `json:"warranty_end,omitempty"`
}

/* polymorph Inventory category false */

/* polymorph Inventory comment false */

/* polymorph Inventory created_at false */

/* polymorph Inventory deleted_at false */

/* polymorph Inventory id false */

/* polymorph Inventory install_date false */

/* polymorph Inventory inventory_number false */

/* polymorph Inventory inventory_status false */

/* polymorph Inventory inventory_status_id false */

/* polymorph Inventory location_id false */

/* polymorph Inventory machine false */

/* polymorph Inventory name false */

/* polymorph Inventory part_number false */

/* polymorph Inventory place false */

/* polymorph Inventory purchase_date false */

/* polymorph Inventory seller false */

/* polymorph Inventory serial false */

/* polymorph Inventory updated_at false */

/* polymorph Inventory user_id false */

/* polymorph Inventory warranty_end false */

// Validate validates this inventory
func (m *Inventory) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *Inventory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Inventory) UnmarshalBinary(b []byte) error {
	var res Inventory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}