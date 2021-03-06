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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/idb-project/go-idb/models"
)

// NewPutAPIV3MachinesRfqdnParams creates a new PutAPIV3MachinesRfqdnParams object
// with the default values initialized.
func NewPutAPIV3MachinesRfqdnParams() *PutAPIV3MachinesRfqdnParams {
	var ()
	return &PutAPIV3MachinesRfqdnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3MachinesRfqdnParamsWithTimeout creates a new PutAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3MachinesRfqdnParamsWithTimeout(timeout time.Duration) *PutAPIV3MachinesRfqdnParams {
	var ()
	return &PutAPIV3MachinesRfqdnParams{

		timeout: timeout,
	}
}

// NewPutAPIV3MachinesRfqdnParamsWithContext creates a new PutAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3MachinesRfqdnParamsWithContext(ctx context.Context) *PutAPIV3MachinesRfqdnParams {
	var ()
	return &PutAPIV3MachinesRfqdnParams{

		Context: ctx,
	}
}

// NewPutAPIV3MachinesRfqdnParamsWithHTTPClient creates a new PutAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3MachinesRfqdnParamsWithHTTPClient(client *http.Client) *PutAPIV3MachinesRfqdnParams {
	var ()
	return &PutAPIV3MachinesRfqdnParams{
		HTTPClient: client,
	}
}

/*PutAPIV3MachinesRfqdnParams contains all the parameters to send to the API endpoint
for the put Api v3 machines rfqdn operation typically these are written to a http.Request
*/
type PutAPIV3MachinesRfqdnParams struct {

	/*Machine*/
	Machine *models.Machine
	/*Rfqdn*/
	Rfqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) WithTimeout(timeout time.Duration) *PutAPIV3MachinesRfqdnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) WithContext(ctx context.Context) *PutAPIV3MachinesRfqdnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) WithHTTPClient(client *http.Client) *PutAPIV3MachinesRfqdnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMachine adds the machine to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) WithMachine(machine *models.Machine) *PutAPIV3MachinesRfqdnParams {
	o.SetMachine(machine)
	return o
}

// SetMachine adds the machine to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) SetMachine(machine *models.Machine) {
	o.Machine = machine
}

// WithRfqdn adds the rfqdn to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) WithRfqdn(rfqdn string) *PutAPIV3MachinesRfqdnParams {
	o.SetRfqdn(rfqdn)
	return o
}

// SetRfqdn adds the rfqdn to the put Api v3 machines rfqdn params
func (o *PutAPIV3MachinesRfqdnParams) SetRfqdn(rfqdn string) {
	o.Rfqdn = rfqdn
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3MachinesRfqdnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Machine == nil {
		o.Machine = new(models.Machine)
	}

	if err := r.SetBodyParam(o.Machine); err != nil {
		return err
	}

	// path param rfqdn
	if err := r.SetPathParam("rfqdn", o.Rfqdn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
