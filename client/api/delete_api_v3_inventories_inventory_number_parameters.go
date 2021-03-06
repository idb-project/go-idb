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

// NewDeleteAPIV3InventoriesInventoryNumberParams creates a new DeleteAPIV3InventoriesInventoryNumberParams object
// with the default values initialized.
func NewDeleteAPIV3InventoriesInventoryNumberParams() *DeleteAPIV3InventoriesInventoryNumberParams {
	var ()
	return &DeleteAPIV3InventoriesInventoryNumberParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV3InventoriesInventoryNumberParamsWithTimeout creates a new DeleteAPIV3InventoriesInventoryNumberParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV3InventoriesInventoryNumberParamsWithTimeout(timeout time.Duration) *DeleteAPIV3InventoriesInventoryNumberParams {
	var ()
	return &DeleteAPIV3InventoriesInventoryNumberParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV3InventoriesInventoryNumberParamsWithContext creates a new DeleteAPIV3InventoriesInventoryNumberParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV3InventoriesInventoryNumberParamsWithContext(ctx context.Context) *DeleteAPIV3InventoriesInventoryNumberParams {
	var ()
	return &DeleteAPIV3InventoriesInventoryNumberParams{

		Context: ctx,
	}
}

// NewDeleteAPIV3InventoriesInventoryNumberParamsWithHTTPClient creates a new DeleteAPIV3InventoriesInventoryNumberParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV3InventoriesInventoryNumberParamsWithHTTPClient(client *http.Client) *DeleteAPIV3InventoriesInventoryNumberParams {
	var ()
	return &DeleteAPIV3InventoriesInventoryNumberParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV3InventoriesInventoryNumberParams contains all the parameters to send to the API endpoint
for the delete Api v3 inventories inventory number operation typically these are written to a http.Request
*/
type DeleteAPIV3InventoriesInventoryNumberParams struct {

	/*InventoryNumber*/
	InventoryNumber string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete Api v3 inventories inventory number params
func (o *DeleteAPIV3InventoriesInventoryNumberParams) WithTimeout(timeout time.Duration) *DeleteAPIV3InventoriesInventoryNumberParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api v3 inventories inventory number params
func (o *DeleteAPIV3InventoriesInventoryNumberParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api v3 inventories inventory number params
func (o *DeleteAPIV3InventoriesInventoryNumberParams) WithContext(ctx context.Context) *DeleteAPIV3InventoriesInventoryNumberParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api v3 inventories inventory number params
func (o *DeleteAPIV3InventoriesInventoryNumberParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api v3 inventories inventory number params
func (o *DeleteAPIV3InventoriesInventoryNumberParams) WithHTTPClient(client *http.Client) *DeleteAPIV3InventoriesInventoryNumberParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api v3 inventories inventory number params
func (o *DeleteAPIV3InventoriesInventoryNumberParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInventoryNumber adds the inventoryNumber to the delete Api v3 inventories inventory number params
func (o *DeleteAPIV3InventoriesInventoryNumberParams) WithInventoryNumber(inventoryNumber string) *DeleteAPIV3InventoriesInventoryNumberParams {
	o.SetInventoryNumber(inventoryNumber)
	return o
}

// SetInventoryNumber adds the inventoryNumber to the delete Api v3 inventories inventory number params
func (o *DeleteAPIV3InventoriesInventoryNumberParams) SetInventoryNumber(inventoryNumber string) {
	o.InventoryNumber = inventoryNumber
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV3InventoriesInventoryNumberParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param inventory_number
	if err := r.SetPathParam("inventory_number", o.InventoryNumber); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
