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

// NewGetAPIV3SwitchesFqdnPortsParams creates a new GetAPIV3SwitchesFqdnPortsParams object
// with the default values initialized.
func NewGetAPIV3SwitchesFqdnPortsParams() *GetAPIV3SwitchesFqdnPortsParams {
	var ()
	return &GetAPIV3SwitchesFqdnPortsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3SwitchesFqdnPortsParamsWithTimeout creates a new GetAPIV3SwitchesFqdnPortsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3SwitchesFqdnPortsParamsWithTimeout(timeout time.Duration) *GetAPIV3SwitchesFqdnPortsParams {
	var ()
	return &GetAPIV3SwitchesFqdnPortsParams{

		timeout: timeout,
	}
}

// NewGetAPIV3SwitchesFqdnPortsParamsWithContext creates a new GetAPIV3SwitchesFqdnPortsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3SwitchesFqdnPortsParamsWithContext(ctx context.Context) *GetAPIV3SwitchesFqdnPortsParams {
	var ()
	return &GetAPIV3SwitchesFqdnPortsParams{

		Context: ctx,
	}
}

// NewGetAPIV3SwitchesFqdnPortsParamsWithHTTPClient creates a new GetAPIV3SwitchesFqdnPortsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3SwitchesFqdnPortsParamsWithHTTPClient(client *http.Client) *GetAPIV3SwitchesFqdnPortsParams {
	var ()
	return &GetAPIV3SwitchesFqdnPortsParams{
		HTTPClient: client,
	}
}

/*GetAPIV3SwitchesFqdnPortsParams contains all the parameters to send to the API endpoint
for the get Api v3 switches fqdn ports operation typically these are written to a http.Request
*/
type GetAPIV3SwitchesFqdnPortsParams struct {

	/*Fqdn*/
	Fqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 switches fqdn ports params
func (o *GetAPIV3SwitchesFqdnPortsParams) WithTimeout(timeout time.Duration) *GetAPIV3SwitchesFqdnPortsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 switches fqdn ports params
func (o *GetAPIV3SwitchesFqdnPortsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 switches fqdn ports params
func (o *GetAPIV3SwitchesFqdnPortsParams) WithContext(ctx context.Context) *GetAPIV3SwitchesFqdnPortsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 switches fqdn ports params
func (o *GetAPIV3SwitchesFqdnPortsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 switches fqdn ports params
func (o *GetAPIV3SwitchesFqdnPortsParams) WithHTTPClient(client *http.Client) *GetAPIV3SwitchesFqdnPortsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 switches fqdn ports params
func (o *GetAPIV3SwitchesFqdnPortsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFqdn adds the fqdn to the get Api v3 switches fqdn ports params
func (o *GetAPIV3SwitchesFqdnPortsParams) WithFqdn(fqdn string) *GetAPIV3SwitchesFqdnPortsParams {
	o.SetFqdn(fqdn)
	return o
}

// SetFqdn adds the fqdn to the get Api v3 switches fqdn ports params
func (o *GetAPIV3SwitchesFqdnPortsParams) SetFqdn(fqdn string) {
	o.Fqdn = fqdn
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3SwitchesFqdnPortsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
