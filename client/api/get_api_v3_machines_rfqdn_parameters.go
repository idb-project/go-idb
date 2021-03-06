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

// NewGetAPIV3MachinesRfqdnParams creates a new GetAPIV3MachinesRfqdnParams object
// with the default values initialized.
func NewGetAPIV3MachinesRfqdnParams() *GetAPIV3MachinesRfqdnParams {
	var ()
	return &GetAPIV3MachinesRfqdnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3MachinesRfqdnParamsWithTimeout creates a new GetAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3MachinesRfqdnParamsWithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnParams {
	var ()
	return &GetAPIV3MachinesRfqdnParams{

		timeout: timeout,
	}
}

// NewGetAPIV3MachinesRfqdnParamsWithContext creates a new GetAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3MachinesRfqdnParamsWithContext(ctx context.Context) *GetAPIV3MachinesRfqdnParams {
	var ()
	return &GetAPIV3MachinesRfqdnParams{

		Context: ctx,
	}
}

// NewGetAPIV3MachinesRfqdnParamsWithHTTPClient creates a new GetAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3MachinesRfqdnParamsWithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnParams {
	var ()
	return &GetAPIV3MachinesRfqdnParams{
		HTTPClient: client,
	}
}

/*GetAPIV3MachinesRfqdnParams contains all the parameters to send to the API endpoint
for the get Api v3 machines rfqdn operation typically these are written to a http.Request
*/
type GetAPIV3MachinesRfqdnParams struct {

	/*Rfqdn*/
	Rfqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 machines rfqdn params
func (o *GetAPIV3MachinesRfqdnParams) WithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 machines rfqdn params
func (o *GetAPIV3MachinesRfqdnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 machines rfqdn params
func (o *GetAPIV3MachinesRfqdnParams) WithContext(ctx context.Context) *GetAPIV3MachinesRfqdnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 machines rfqdn params
func (o *GetAPIV3MachinesRfqdnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn params
func (o *GetAPIV3MachinesRfqdnParams) WithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn params
func (o *GetAPIV3MachinesRfqdnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRfqdn adds the rfqdn to the get Api v3 machines rfqdn params
func (o *GetAPIV3MachinesRfqdnParams) WithRfqdn(rfqdn string) *GetAPIV3MachinesRfqdnParams {
	o.SetRfqdn(rfqdn)
	return o
}

// SetRfqdn adds the rfqdn to the get Api v3 machines rfqdn params
func (o *GetAPIV3MachinesRfqdnParams) SetRfqdn(rfqdn string) {
	o.Rfqdn = rfqdn
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3MachinesRfqdnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rfqdn
	if err := r.SetPathParam("rfqdn", o.Rfqdn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
