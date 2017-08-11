// Code generated by go-swagger; DO NOT EDIT.

package switches

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

// NewPostAPIV3SwitchesFqdnPortsParams creates a new PostAPIV3SwitchesFqdnPortsParams object
// with the default values initialized.
func NewPostAPIV3SwitchesFqdnPortsParams() *PostAPIV3SwitchesFqdnPortsParams {
	var ()
	return &PostAPIV3SwitchesFqdnPortsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3SwitchesFqdnPortsParamsWithTimeout creates a new PostAPIV3SwitchesFqdnPortsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3SwitchesFqdnPortsParamsWithTimeout(timeout time.Duration) *PostAPIV3SwitchesFqdnPortsParams {
	var ()
	return &PostAPIV3SwitchesFqdnPortsParams{

		timeout: timeout,
	}
}

// NewPostAPIV3SwitchesFqdnPortsParamsWithContext creates a new PostAPIV3SwitchesFqdnPortsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3SwitchesFqdnPortsParamsWithContext(ctx context.Context) *PostAPIV3SwitchesFqdnPortsParams {
	var ()
	return &PostAPIV3SwitchesFqdnPortsParams{

		Context: ctx,
	}
}

// NewPostAPIV3SwitchesFqdnPortsParamsWithHTTPClient creates a new PostAPIV3SwitchesFqdnPortsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3SwitchesFqdnPortsParamsWithHTTPClient(client *http.Client) *PostAPIV3SwitchesFqdnPortsParams {
	var ()
	return &PostAPIV3SwitchesFqdnPortsParams{
		HTTPClient: client,
	}
}

/*PostAPIV3SwitchesFqdnPortsParams contains all the parameters to send to the API endpoint
for the post Api v3 switches fqdn ports operation typically these are written to a http.Request
*/
type PostAPIV3SwitchesFqdnPortsParams struct {

	/*Fqdn*/
	Fqdn string
	/*Machine
	  Machine nic belongs to

	*/
	Machine string
	/*Nic
	  Nic name

	*/
	Nic string
	/*Number
	  Port number

	*/
	Number int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) WithTimeout(timeout time.Duration) *PostAPIV3SwitchesFqdnPortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) WithContext(ctx context.Context) *PostAPIV3SwitchesFqdnPortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) WithHTTPClient(client *http.Client) *PostAPIV3SwitchesFqdnPortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFqdn adds the fqdn to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) WithFqdn(fqdn string) *PostAPIV3SwitchesFqdnPortsParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) SetFqdn(fqdn string) {
	o.Fqdn = fqdn
}

// WithMachine adds the machine to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) WithMachine(machine string) *PostAPIV3SwitchesFqdnPortsParams {
	o.SetMachine(machine)
	return o
}

// SetMachine adds the machine to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) SetMachine(machine string) {
	o.Machine = machine
}

// WithNic adds the nic to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) WithNic(nic string) *PostAPIV3SwitchesFqdnPortsParams {
	o.SetNic(nic)
	return o
}

// SetNic adds the nic to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) SetNic(nic string) {
	o.Nic = nic
}

// WithNumber adds the number to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) WithNumber(number int32) *PostAPIV3SwitchesFqdnPortsParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the post Api v3 switches fqdn ports params
func (o *PostAPIV3SwitchesFqdnPortsParams) SetNumber(number int32) {
	o.Number = number
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3SwitchesFqdnPortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fqdn
	if err := r.SetPathParam("fqdn", o.Fqdn); err != nil {
		return err
	}

	// form param machine
	frMachine := o.Machine
	fMachine := frMachine
	if fMachine != "" {
		if err := r.SetFormParam("machine", fMachine); err != nil {
			return err
		}
	}

	// form param nic
	frNic := o.Nic
	fNic := frNic
	if fNic != "" {
		if err := r.SetFormParam("nic", fNic); err != nil {
			return err
		}
	}

	// form param number
	frNumber := o.Number
	fNumber := swag.FormatInt32(frNumber)
	if fNumber != "" {
		if err := r.SetFormParam("number", fNumber); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
