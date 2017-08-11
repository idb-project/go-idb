// Code generated by go-swagger; DO NOT EDIT.

package inventories

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

// NewGetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams creates a new GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams object
// with the default values initialized.
func NewGetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams() *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	var ()
	return &GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParamsWithTimeout creates a new GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParamsWithTimeout(timeout time.Duration) *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	var ()
	return &GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams{

		timeout: timeout,
	}
}

// NewGetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParamsWithContext creates a new GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParamsWithContext(ctx context.Context) *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	var ()
	return &GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams{

		Context: ctx,
	}
}

// NewGetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParamsWithHTTPClient creates a new GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParamsWithHTTPClient(client *http.Client) *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	var ()
	return &GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams{
		HTTPClient: client,
	}
}

/*GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams contains all the parameters to send to the API endpoint
for the get Api v3 inventories inventory number attachments fingerprint operation typically these are written to a http.Request
*/
type GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams struct {

	/*Fingerprint*/
	Fingerprint string
	/*InventoryNumber*/
	InventoryNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) WithTimeout(timeout time.Duration) *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) WithContext(ctx context.Context) *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) WithHTTPClient(client *http.Client) *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFingerprint adds the fingerprint to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) WithFingerprint(fingerprint string) *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	o.SetFingerprint(fingerprint)
	return o
}

// SetFingerprint adds the fingerprint to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) SetFingerprint(fingerprint string) {
	o.Fingerprint = fingerprint
}

// WithInventoryNumber adds the inventoryNumber to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) WithInventoryNumber(inventoryNumber string) *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams {
	o.SetInventoryNumber(inventoryNumber)
	return o
}

// SetInventoryNumber adds the inventoryNumber to the get Api v3 inventories inventory number attachments fingerprint params
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) SetInventoryNumber(inventoryNumber string) {
	o.InventoryNumber = inventoryNumber
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3InventoriesInventoryNumberAttachmentsFingerprintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param fingerprint
	if err := r.SetPathParam("fingerprint", o.Fingerprint); err != nil {
		return err
	}

	// path param inventory_number
	if err := r.SetPathParam("inventory_number", o.InventoryNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}