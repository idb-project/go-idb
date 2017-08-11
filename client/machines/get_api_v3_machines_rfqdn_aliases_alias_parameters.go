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

// NewGetAPIV3MachinesRfqdnAliasesAliasParams creates a new GetAPIV3MachinesRfqdnAliasesAliasParams object
// with the default values initialized.
func NewGetAPIV3MachinesRfqdnAliasesAliasParams() *GetAPIV3MachinesRfqdnAliasesAliasParams {
	var ()
	return &GetAPIV3MachinesRfqdnAliasesAliasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3MachinesRfqdnAliasesAliasParamsWithTimeout creates a new GetAPIV3MachinesRfqdnAliasesAliasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3MachinesRfqdnAliasesAliasParamsWithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnAliasesAliasParams {
	var ()
	return &GetAPIV3MachinesRfqdnAliasesAliasParams{

		timeout: timeout,
	}
}

// NewGetAPIV3MachinesRfqdnAliasesAliasParamsWithContext creates a new GetAPIV3MachinesRfqdnAliasesAliasParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3MachinesRfqdnAliasesAliasParamsWithContext(ctx context.Context) *GetAPIV3MachinesRfqdnAliasesAliasParams {
	var ()
	return &GetAPIV3MachinesRfqdnAliasesAliasParams{

		Context: ctx,
	}
}

// NewGetAPIV3MachinesRfqdnAliasesAliasParamsWithHTTPClient creates a new GetAPIV3MachinesRfqdnAliasesAliasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3MachinesRfqdnAliasesAliasParamsWithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnAliasesAliasParams {
	var ()
	return &GetAPIV3MachinesRfqdnAliasesAliasParams{
		HTTPClient: client,
	}
}

/*GetAPIV3MachinesRfqdnAliasesAliasParams contains all the parameters to send to the API endpoint
for the get Api v3 machines rfqdn aliases alias operation typically these are written to a http.Request
*/
type GetAPIV3MachinesRfqdnAliasesAliasParams struct {

	/*Alias*/
	Alias string
	/*Rfqdn*/
	Rfqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) WithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnAliasesAliasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) WithContext(ctx context.Context) *GetAPIV3MachinesRfqdnAliasesAliasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) WithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnAliasesAliasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAlias adds the alias to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) WithAlias(alias string) *GetAPIV3MachinesRfqdnAliasesAliasParams {
	o.SetAlias(alias)
	return o
}

// SetAlias adds the alias to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) SetAlias(alias string) {
	o.Alias = alias
}

// WithRfqdn adds the rfqdn to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) WithRfqdn(rfqdn string) *GetAPIV3MachinesRfqdnAliasesAliasParams {
	o.SetRfqdn(rfqdn)
	return o
}

// SetRfqdn adds the rfqdn to the get Api v3 machines rfqdn aliases alias params
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) SetRfqdn(rfqdn string) {
	o.Rfqdn = rfqdn
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3MachinesRfqdnAliasesAliasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param alias
	if err := r.SetPathParam("alias", o.Alias); err != nil {
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
