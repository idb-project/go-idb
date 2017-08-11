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

// NewGetAPIV3SwitchesFqdnParams creates a new GetAPIV3SwitchesFqdnParams object
// with the default values initialized.
func NewGetAPIV3SwitchesFqdnParams() *GetAPIV3SwitchesFqdnParams {
	var ()
	return &GetAPIV3SwitchesFqdnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3SwitchesFqdnParamsWithTimeout creates a new GetAPIV3SwitchesFqdnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3SwitchesFqdnParamsWithTimeout(timeout time.Duration) *GetAPIV3SwitchesFqdnParams {
	var ()
	return &GetAPIV3SwitchesFqdnParams{

		timeout: timeout,
	}
}

// NewGetAPIV3SwitchesFqdnParamsWithContext creates a new GetAPIV3SwitchesFqdnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3SwitchesFqdnParamsWithContext(ctx context.Context) *GetAPIV3SwitchesFqdnParams {
	var ()
	return &GetAPIV3SwitchesFqdnParams{

		Context: ctx,
	}
}

// NewGetAPIV3SwitchesFqdnParamsWithHTTPClient creates a new GetAPIV3SwitchesFqdnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3SwitchesFqdnParamsWithHTTPClient(client *http.Client) *GetAPIV3SwitchesFqdnParams {
	var ()
	return &GetAPIV3SwitchesFqdnParams{
		HTTPClient: client,
	}
}

/*GetAPIV3SwitchesFqdnParams contains all the parameters to send to the API endpoint
for the get Api v3 switches fqdn operation typically these are written to a http.Request
*/
type GetAPIV3SwitchesFqdnParams struct {

	/*Fqdn*/
	Fqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 switches fqdn params
func (o *GetAPIV3SwitchesFqdnParams) WithTimeout(timeout time.Duration) *GetAPIV3SwitchesFqdnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 switches fqdn params
func (o *GetAPIV3SwitchesFqdnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 switches fqdn params
func (o *GetAPIV3SwitchesFqdnParams) WithContext(ctx context.Context) *GetAPIV3SwitchesFqdnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 switches fqdn params
func (o *GetAPIV3SwitchesFqdnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 switches fqdn params
func (o *GetAPIV3SwitchesFqdnParams) WithHTTPClient(client *http.Client) *GetAPIV3SwitchesFqdnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 switches fqdn params
func (o *GetAPIV3SwitchesFqdnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFqdn adds the fqdn to the get Api v3 switches fqdn params
func (o *GetAPIV3SwitchesFqdnParams) WithFqdn(fqdn string) *GetAPIV3SwitchesFqdnParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the get Api v3 switches fqdn params
func (o *GetAPIV3SwitchesFqdnParams) SetFqdn(fqdn string) {
	o.Fqdn = fqdn
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3SwitchesFqdnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
