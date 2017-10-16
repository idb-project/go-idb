// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PutAPIV3MachinesRfqdnNics Update a nic
// swagger:model putApiV3MachinesRfqdnNics

type PutAPIV3MachinesRfqdnNics struct {

	// IP address of the nic
	IPAddress *IPAddress `json:"ip_address,omitempty"`

	// MAC address
	Mac string `json:"mac,omitempty"`

	// Machine FQDN this nic belongs to
	Machine string `json:"machine,omitempty"`

	// Name
	Name string `json:"name,omitempty"`
}

/* polymorph putApiV3MachinesRfqdnNics ip_address false */

/* polymorph putApiV3MachinesRfqdnNics mac false */

/* polymorph putApiV3MachinesRfqdnNics machine false */

/* polymorph putApiV3MachinesRfqdnNics name false */

// Validate validates this put Api v3 machines rfqdn nics
func (m *PutAPIV3MachinesRfqdnNics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPAddress(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PutAPIV3MachinesRfqdnNics) validateIPAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.IPAddress) { // not required
		return nil
	}

	if m.IPAddress != nil {

		if err := m.IPAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip_address")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PutAPIV3MachinesRfqdnNics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PutAPIV3MachinesRfqdnNics) UnmarshalBinary(b []byte) error {
	var res PutAPIV3MachinesRfqdnNics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}