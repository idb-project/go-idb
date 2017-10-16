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
)

// NewDeleteAPIV3MachinesRfqdnNicsRnicParams creates a new DeleteAPIV3MachinesRfqdnNicsRnicParams object
// with the default values initialized.
func NewDeleteAPIV3MachinesRfqdnNicsRnicParams() *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	var ()
	return &DeleteAPIV3MachinesRfqdnNicsRnicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV3MachinesRfqdnNicsRnicParamsWithTimeout creates a new DeleteAPIV3MachinesRfqdnNicsRnicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV3MachinesRfqdnNicsRnicParamsWithTimeout(timeout time.Duration) *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	var ()
	return &DeleteAPIV3MachinesRfqdnNicsRnicParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV3MachinesRfqdnNicsRnicParamsWithContext creates a new DeleteAPIV3MachinesRfqdnNicsRnicParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV3MachinesRfqdnNicsRnicParamsWithContext(ctx context.Context) *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	var ()
	return &DeleteAPIV3MachinesRfqdnNicsRnicParams{

		Context: ctx,
	}
}

// NewDeleteAPIV3MachinesRfqdnNicsRnicParamsWithHTTPClient creates a new DeleteAPIV3MachinesRfqdnNicsRnicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV3MachinesRfqdnNicsRnicParamsWithHTTPClient(client *http.Client) *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	var ()
	return &DeleteAPIV3MachinesRfqdnNicsRnicParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV3MachinesRfqdnNicsRnicParams contains all the parameters to send to the API endpoint
for the delete Api v3 machines rfqdn nics rnic operation typically these are written to a http.Request
*/
type DeleteAPIV3MachinesRfqdnNicsRnicParams struct {

	/*Rfqdn*/
	Rfqdn string
	/*Rnic*/
	Rnic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) WithTimeout(timeout time.Duration) *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) WithContext(ctx context.Context) *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) WithHTTPClient(client *http.Client) *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRfqdn adds the rfqdn to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) WithRfqdn(rfqdn string) *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	o.SetRfqdn(rfqdn)
	return o
}

// SetRfqdn adds the rfqdn to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) SetRfqdn(rfqdn string) {
	o.Rfqdn = rfqdn
}

// WithRnic adds the rnic to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) WithRnic(rnic string) *DeleteAPIV3MachinesRfqdnNicsRnicParams {
	o.SetRnic(rnic)
	return o
}

// SetRnic adds the rnic to the delete Api v3 machines rfqdn nics rnic params
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) SetRnic(rnic string) {
	o.Rnic = rnic
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV3MachinesRfqdnNicsRnicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rfqdn
	if err := r.SetPathParam("rfqdn", o.Rfqdn); err != nil {
		return err
	}

	// path param rnic
	if err := r.SetPathParam("rnic", o.Rnic); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
