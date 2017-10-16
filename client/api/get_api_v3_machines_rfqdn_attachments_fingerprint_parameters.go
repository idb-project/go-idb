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

// NewGetAPIV3MachinesRfqdnAttachmentsFingerprintParams creates a new GetAPIV3MachinesRfqdnAttachmentsFingerprintParams object
// with the default values initialized.
func NewGetAPIV3MachinesRfqdnAttachmentsFingerprintParams() *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	var ()
	return &GetAPIV3MachinesRfqdnAttachmentsFingerprintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3MachinesRfqdnAttachmentsFingerprintParamsWithTimeout creates a new GetAPIV3MachinesRfqdnAttachmentsFingerprintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3MachinesRfqdnAttachmentsFingerprintParamsWithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	var ()
	return &GetAPIV3MachinesRfqdnAttachmentsFingerprintParams{

		timeout: timeout,
	}
}

// NewGetAPIV3MachinesRfqdnAttachmentsFingerprintParamsWithContext creates a new GetAPIV3MachinesRfqdnAttachmentsFingerprintParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3MachinesRfqdnAttachmentsFingerprintParamsWithContext(ctx context.Context) *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	var ()
	return &GetAPIV3MachinesRfqdnAttachmentsFingerprintParams{

		Context: ctx,
	}
}

// NewGetAPIV3MachinesRfqdnAttachmentsFingerprintParamsWithHTTPClient creates a new GetAPIV3MachinesRfqdnAttachmentsFingerprintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3MachinesRfqdnAttachmentsFingerprintParamsWithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	var ()
	return &GetAPIV3MachinesRfqdnAttachmentsFingerprintParams{
		HTTPClient: client,
	}
}

/*GetAPIV3MachinesRfqdnAttachmentsFingerprintParams contains all the parameters to send to the API endpoint
for the get Api v3 machines rfqdn attachments fingerprint operation typically these are written to a http.Request
*/
type GetAPIV3MachinesRfqdnAttachmentsFingerprintParams struct {

	/*Fingerprint*/
	Fingerprint string
	/*Rfqdn*/
	Rfqdn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) WithTimeout(timeout time.Duration) *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) WithContext(ctx context.Context) *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) WithHTTPClient(client *http.Client) *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFingerprint adds the fingerprint to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) WithFingerprint(fingerprint string) *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) SetFingerprint(fingerprint string) {
	o.Fingerprint = fingerprint
}

// WithRfqdn adds the rfqdn to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) WithRfqdn(rfqdn string) *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams {
	o.SetRfqdn(rfqdn)
	return o
}

// SetRfqdn adds the rfqdn to the get Api v3 machines rfqdn attachments fingerprint params
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) SetRfqdn(rfqdn string) {
	o.Rfqdn = rfqdn
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3MachinesRfqdnAttachmentsFingerprintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fingerprint
	if err := r.SetPathParam("fingerprint", o.Fingerprint); err != nil {
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