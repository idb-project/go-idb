// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// Location Return a list of locations, possibly filtered
// swagger:model Location

type Location struct {

	// Child location ids
	Children []int32 `json:"children"`

	// Description
	Description string `json:"description,omitempty"`

	// Id
	ID int32 `json:"id,omitempty"`

	// Location level
	Level string `json:"level,omitempty"`

	// Name
	Name string `json:"name,omitempty"`

	// Parent location id
	Parent int32 `json:"parent,omitempty"`
}

/* polymorph Location children false */

/* polymorph Location description false */

/* polymorph Location id false */

/* polymorph Location level false */

/* polymorph Location name false */

/* polymorph Location parent false */

// Validate validates this location
func (m *Location) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildren(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Location) validateChildren(formats strfmt.Registry) error {

	if swag.IsZero(m.Children) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Location) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Location) UnmarshalBinary(b []byte) error {
	var res Location
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
