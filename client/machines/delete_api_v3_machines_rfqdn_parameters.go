// Code generated by go-swagger; DO NOT EDIT.

package machines

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

// NewDeleteAPIV3MachinesRfqdnParams creates a new DeleteAPIV3MachinesRfqdnParams object
// with the default values initialized.
func NewDeleteAPIV3MachinesRfqdnParams() *DeleteAPIV3MachinesRfqdnParams {
	var ()
	return &DeleteAPIV3MachinesRfqdnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV3MachinesRfqdnParamsWithTimeout creates a new DeleteAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV3MachinesRfqdnParamsWithTimeout(timeout time.Duration) *DeleteAPIV3MachinesRfqdnParams {
	var ()
	return &DeleteAPIV3MachinesRfqdnParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV3MachinesRfqdnParamsWithContext creates a new DeleteAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV3MachinesRfqdnParamsWithContext(ctx context.Context) *DeleteAPIV3MachinesRfqdnParams {
	var ()
	return &DeleteAPIV3MachinesRfqdnParams{

		Context: ctx,
	}
}

// NewDeleteAPIV3MachinesRfqdnParamsWithHTTPClient creates a new DeleteAPIV3MachinesRfqdnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV3MachinesRfqdnParamsWithHTTPClient(client *http.Client) *DeleteAPIV3MachinesRfqdnParams {
	var ()
	return &DeleteAPIV3MachinesRfqdnParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV3MachinesRfqdnParams contains all the parameters to send to the API endpoint
for the delete Api v3 machines rfqdn operation typically these are written to a http.Request
*/
type DeleteAPIV3MachinesRfqdnParams struct {

	/*Rfqdn*/
	Rfqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete Api v3 machines rfqdn params
func (o *DeleteAPIV3MachinesRfqdnParams) WithTimeout(timeout time.Duration) *DeleteAPIV3MachinesRfqdnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api v3 machines rfqdn params
func (o *DeleteAPIV3MachinesRfqdnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api v3 machines rfqdn params
func (o *DeleteAPIV3MachinesRfqdnParams) WithContext(ctx context.Context) *DeleteAPIV3MachinesRfqdnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api v3 machines rfqdn params
func (o *DeleteAPIV3MachinesRfqdnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api v3 machines rfqdn params
func (o *DeleteAPIV3MachinesRfqdnParams) WithHTTPClient(client *http.Client) *DeleteAPIV3MachinesRfqdnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api v3 machines rfqdn params
func (o *DeleteAPIV3MachinesRfqdnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRfqdn adds the rfqdn to the delete Api v3 machines rfqdn params
func (o *DeleteAPIV3MachinesRfqdnParams) WithRfqdn(rfqdn string) *DeleteAPIV3MachinesRfqdnParams {
	o.SetRfqdn(rfqdn)
	return o
}

// SetRfqdn adds the rfqdn to the delete Api v3 machines rfqdn params
func (o *DeleteAPIV3MachinesRfqdnParams) SetRfqdn(rfqdn string) {
	o.Rfqdn = rfqdn
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV3MachinesRfqdnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
