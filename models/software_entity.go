// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SoftwareEntity software entity
// swagger:model SoftwareEntity

type SoftwareEntity struct {

	// Software name
	Name string `json:"name,omitempty"`

	// Software version
	Version string `json:"version,omitempty"`
}

/* polymorph SoftwareEntity name false */

/* polymorph SoftwareEntity version false */

// Validate validates this software entity
func (m *SoftwareEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *SoftwareEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SoftwareEntity) UnmarshalBinary(b []byte) error {
	var res SoftwareEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
