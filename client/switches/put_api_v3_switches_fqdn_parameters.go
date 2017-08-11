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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPutAPIV3SwitchesFqdnParams creates a new PutAPIV3SwitchesFqdnParams object
// with the default values initialized.
func NewPutAPIV3SwitchesFqdnParams() *PutAPIV3SwitchesFqdnParams {
	var ()
	return &PutAPIV3SwitchesFqdnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIV3SwitchesFqdnParamsWithTimeout creates a new PutAPIV3SwitchesFqdnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIV3SwitchesFqdnParamsWithTimeout(timeout time.Duration) *PutAPIV3SwitchesFqdnParams {
	var ()
	return &PutAPIV3SwitchesFqdnParams{

		timeout: timeout,
	}
}

// NewPutAPIV3SwitchesFqdnParamsWithContext creates a new PutAPIV3SwitchesFqdnParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIV3SwitchesFqdnParamsWithContext(ctx context.Context) *PutAPIV3SwitchesFqdnParams {
	var ()
	return &PutAPIV3SwitchesFqdnParams{

		Context: ctx,
	}
}

// NewPutAPIV3SwitchesFqdnParamsWithHTTPClient creates a new PutAPIV3SwitchesFqdnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIV3SwitchesFqdnParamsWithHTTPClient(client *http.Client) *PutAPIV3SwitchesFqdnParams {
	var ()
	return &PutAPIV3SwitchesFqdnParams{
		HTTPClient: client,
	}
}

/*PutAPIV3SwitchesFqdnParams contains all the parameters to send to the API endpoint
for the put Api v3 switches fqdn operation typically these are written to a http.Request
*/
type PutAPIV3SwitchesFqdnParams struct {

	/*Fqdn*/
	Fqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put Api v3 switches fqdn params
func (o *PutAPIV3SwitchesFqdnParams) WithTimeout(timeout time.Duration) *PutAPIV3SwitchesFqdnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put Api v3 switches fqdn params
func (o *PutAPIV3SwitchesFqdnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put Api v3 switches fqdn params
func (o *PutAPIV3SwitchesFqdnParams) WithContext(ctx context.Context) *PutAPIV3SwitchesFqdnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put Api v3 switches fqdn params
func (o *PutAPIV3SwitchesFqdnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put Api v3 switches fqdn params
func (o *PutAPIV3SwitchesFqdnParams) WithHTTPClient(client *http.Client) *PutAPIV3SwitchesFqdnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put Api v3 switches fqdn params
func (o *PutAPIV3SwitchesFqdnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFqdn adds the fqdn to the put Api v3 switches fqdn params
func (o *PutAPIV3SwitchesFqdnParams) WithFqdn(fqdn string) *PutAPIV3SwitchesFqdnParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the put Api v3 switches fqdn params
func (o *PutAPIV3SwitchesFqdnParams) SetFqdn(fqdn string) {
	o.Fqdn = fqdn
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIV3SwitchesFqdnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fqdn
	if err := r.SetPathParam("fqdn", o.Fqdn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}