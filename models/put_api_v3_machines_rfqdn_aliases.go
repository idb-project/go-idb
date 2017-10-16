// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PutAPIV3MachinesRfqdnAliases Update an alias
// swagger:model putApiV3MachinesRfqdnAliases

type PutAPIV3MachinesRfqdnAliases struct {

	// Aliased machine FQDN
	Machine string `json:"machine,omitempty"`

	// Name
	Name string `json:"name,omitempty"`
}

/* polymorph putApiV3MachinesRfqdnAliases machine false */

/* polymorph putApiV3MachinesRfqdnAliases name false */

// Validate validates this put Api v3 machines rfqdn aliases
func (m *PutAPIV3MachinesRfqdnAliases) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *PutAPIV3MachinesRfqdnAliases) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutAPIV3MachinesRfqdnAliases) UnmarshalBinary(b []byte) error {
	var res PutAPIV3MachinesRfqdnAliases
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}