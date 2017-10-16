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

// NewGetAPIV3MachinesRfqdnNicsRnicParams creates a new GetAPIV3MachinesRfqdnNicsRnicParams object
// with the default values initialized.
func NewGetAPIV3MachinesRfqdnNicsRnicParams() *GetAPIV3MachinesRfqdnNicsRnicParams {
	var ()
	return &GetAPIV3MachinesRfqdnNicsRnicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3MachinesRfqdnNicsRnicParamsWithTimeout creates a new GetAPIV3MachinesRfqdnNicsRnicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3MachinesRfqdnNicsRnicParamsWithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnNicsRnicParams {
	var ()
	return &GetAPIV3MachinesRfqdnNicsRnicParams{

		timeout: timeout,
	}
}

// NewGetAPIV3MachinesRfqdnNicsRnicParamsWithContext creates a new GetAPIV3MachinesRfqdnNicsRnicParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3MachinesRfqdnNicsRnicParamsWithContext(ctx context.Context) *GetAPIV3MachinesRfqdnNicsRnicParams {
	var ()
	return &GetAPIV3MachinesRfqdnNicsRnicParams{

		Context: ctx,
	}
}

// NewGetAPIV3MachinesRfqdnNicsRnicParamsWithHTTPClient creates a new GetAPIV3MachinesRfqdnNicsRnicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3MachinesRfqdnNicsRnicParamsWithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnNicsRnicParams {
	var ()
	return &GetAPIV3MachinesRfqdnNicsRnicParams{
		HTTPClient: client,
	}
}

/*GetAPIV3MachinesRfqdnNicsRnicParams contains all the parameters to send to the API endpoint
for the get Api v3 machines rfqdn nics rnic operation typically these are written to a http.Request
*/
type GetAPIV3MachinesRfqdnNicsRnicParams struct {

	/*Rfqdn*/
	Rfqdn string
	/*Rnic*/
	Rnic string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) WithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnNicsRnicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) WithContext(ctx context.Context) *GetAPIV3MachinesRfqdnNicsRnicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) WithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnNicsRnicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRfqdn adds the rfqdn to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) WithRfqdn(rfqdn string) *GetAPIV3MachinesRfqdnNicsRnicParams {
	o.SetRfqdn(rfqdn)
	return o
}

// SetRfqdn adds the rfqdn to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) SetRfqdn(rfqdn string) {
	o.Rfqdn = rfqdn
}

// WithRnic adds the rnic to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) WithRnic(rnic string) *GetAPIV3MachinesRfqdnNicsRnicParams {
	o.SetRnic(rnic)
	return o
}

// SetRnic adds the rnic to the get Api v3 machines rfqdn nics rnic params
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) SetRnic(rnic string) {
	o.Rnic = rnic
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3MachinesRfqdnNicsRnicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
