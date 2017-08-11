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

// NewGetAPIV3MachinesRfqdnAttachmentsParams creates a new GetAPIV3MachinesRfqdnAttachmentsParams object
// with the default values initialized.
func NewGetAPIV3MachinesRfqdnAttachmentsParams() *GetAPIV3MachinesRfqdnAttachmentsParams {
	var ()
	return &GetAPIV3MachinesRfqdnAttachmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3MachinesRfqdnAttachmentsParamsWithTimeout creates a new GetAPIV3MachinesRfqdnAttachmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3MachinesRfqdnAttachmentsParamsWithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnAttachmentsParams {
	var ()
	return &GetAPIV3MachinesRfqdnAttachmentsParams{

		timeout: timeout,
	}
}

// NewGetAPIV3MachinesRfqdnAttachmentsParamsWithContext creates a new GetAPIV3MachinesRfqdnAttachmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3MachinesRfqdnAttachmentsParamsWithContext(ctx context.Context) *GetAPIV3MachinesRfqdnAttachmentsParams {
	var ()
	return &GetAPIV3MachinesRfqdnAttachmentsParams{

		Context: ctx,
	}
}

// NewGetAPIV3MachinesRfqdnAttachmentsParamsWithHTTPClient creates a new GetAPIV3MachinesRfqdnAttachmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3MachinesRfqdnAttachmentsParamsWithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnAttachmentsParams {
	var ()
	return &GetAPIV3MachinesRfqdnAttachmentsParams{
		HTTPClient: client,
	}
}

/*GetAPIV3MachinesRfqdnAttachmentsParams contains all the parameters to send to the API endpoint
for the get Api v3 machines rfqdn attachments operation typically these are written to a http.Request
*/
type GetAPIV3MachinesRfqdnAttachmentsParams struct {

	/*Rfqdn*/
	Rfqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 machines rfqdn attachments params
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) WithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnAttachmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 machines rfqdn attachments params
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 machines rfqdn attachments params
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) WithContext(ctx context.Context) *GetAPIV3MachinesRfqdnAttachmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 machines rfqdn attachments params
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn attachments params
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) WithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnAttachmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn attachments params
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRfqdn adds the rfqdn to the get Api v3 machines rfqdn attachments params
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) WithRfqdn(rfqdn string) *GetAPIV3MachinesRfqdnAttachmentsParams {
	o.SetRfqdn(rfqdn)
	return o
}

// SetRfqdn adds the rfqdn to the get Api v3 machines rfqdn attachments params
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) SetRfqdn(rfqdn string) {
	o.Rfqdn = rfqdn
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3MachinesRfqdnAttachmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
