// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutAPIV3SwitchesFqdnPortsNumberParams creates a new PutAPIV3SwitchesFqdnPortsNumberParams object
// with the default values initialized.
func NewPutAPIV3SwitchesFqdnPortsNumberParams() *PutAPIV3SwitchesFqdnPortsNumberParams {
	var ()
	return &PutAPIV3SwitchesFqdnPortsNumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3SwitchesFqdnPortsNumberParamsWithTimeout creates a new PutAPIV3SwitchesFqdnPortsNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3SwitchesFqdnPortsNumberParamsWithTimeout(timeout time.Duration) *PutAPIV3SwitchesFqdnPortsNumberParams {
	var ()
	return &PutAPIV3SwitchesFqdnPortsNumberParams{

		timeout: timeout,
	}
}

// NewPutAPIV3SwitchesFqdnPortsNumberParamsWithContext creates a new PutAPIV3SwitchesFqdnPortsNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3SwitchesFqdnPortsNumberParamsWithContext(ctx context.Context) *PutAPIV3SwitchesFqdnPortsNumberParams {
	var ()
	return &PutAPIV3SwitchesFqdnPortsNumberParams{

		Context: ctx,
	}
}

// NewPutAPIV3SwitchesFqdnPortsNumberParamsWithHTTPClient creates a new PutAPIV3SwitchesFqdnPortsNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3SwitchesFqdnPortsNumberParamsWithHTTPClient(client *http.Client) *PutAPIV3SwitchesFqdnPortsNumberParams {
	var ()
	return &PutAPIV3SwitchesFqdnPortsNumberParams{
		HTTPClient: client,
	}
}

/*PutAPIV3SwitchesFqdnPortsNumberParams contains all the parameters to send to the API endpoint
for the put Api v3 switches fqdn ports number operation typically these are written to a http.Request
*/
type PutAPIV3SwitchesFqdnPortsNumberParams struct {

	/*Fqdn*/
	Fqdn string
	/*Machine
	  Machine nic belongs to

	*/
	Machine *string
	/*Nic
	  Nic name

	*/
	Nic *string
	/*Number
	  Port number

	*/
	Number int32
	/*Switch
	  Switch FQDN the port belongs to

	*/
	Switch *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WithTimeout(timeout time.Duration) *PutAPIV3SwitchesFqdnPortsNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WithContext(ctx context.Context) *PutAPIV3SwitchesFqdnPortsNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WithHTTPClient(client *http.Client) *PutAPIV3SwitchesFqdnPortsNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFqdn adds the fqdn to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WithFqdn(fqdn string) *PutAPIV3SwitchesFqdnPortsNumberParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) SetFqdn(fqdn string) {
	o.Fqdn = fqdn
}

// WithMachine adds the machine to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WithMachine(machine *string) *PutAPIV3SwitchesFqdnPortsNumberParams {
	o.SetMachine(machine)
	return o
}

// SetMachine adds the machine to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) SetMachine(machine *string) {
	o.Machine = machine
}

// WithNic adds the nic to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WithNic(nic *string) *PutAPIV3SwitchesFqdnPortsNumberParams {
	o.SetNic(nic)
	return o
}

// SetNic adds the nic to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) SetNic(nic *string) {
	o.Nic = nic
}

// WithNumber adds the number to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WithNumber(number int32) *PutAPIV3SwitchesFqdnPortsNumberParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) SetNumber(number int32) {
	o.Number = number
}

// WithSwitch adds the switchVar to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WithSwitch(switchVar *string) *PutAPIV3SwitchesFqdnPortsNumberParams {
	o.SetSwitch(switchVar)
	return o
}

// SetSwitch adds the switch to the put Api v3 switches fqdn ports number params
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) SetSwitch(switchVar *string) {
	o.Switch = switchVar
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3SwitchesFqdnPortsNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fqdn
	if err := r.SetPathParam("fqdn", o.Fqdn); err != nil {
		return err
	}

	if o.Machine != nil {

		// form param machine
		var frMachine string
		if o.Machine != nil {
			frMachine = *o.Machine
		}
		fMachine := frMachine
		if fMachine != "" {
			if err := r.SetFormParam("machine", fMachine); err != nil {
				return err
			}
		}

	}

	if o.Nic != nil {

		// form param nic
		var frNic string
		if o.Nic != nil {
			frNic = *o.Nic
		}
		fNic := frNic
		if fNic != "" {
			if err := r.SetFormParam("nic", fNic); err != nil {
				return err
			}
		}

	}

	// path param number
	if err := r.SetPathParam("number", swag.FormatInt32(o.Number)); err != nil {
		return err
	}

	if o.Switch != nil {

		// form param switch
		var frSwitch string
		if o.Switch != nil {
			frSwitch = *o.Switch
		}
		fSwitch := frSwitch
		if fSwitch != "" {
			if err := r.SetFormParam("switch", fSwitch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
