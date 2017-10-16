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

	"github.com/idb-project/go-idb/models"
)

// NewPostAPIV3InventoriesParams creates a new PostAPIV3InventoriesParams object
// with the default values initialized.
func NewPostAPIV3InventoriesParams() *PostAPIV3InventoriesParams {
	var ()
	return &PostAPIV3InventoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIV3InventoriesParamsWithTimeout creates a new PostAPIV3InventoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIV3InventoriesParamsWithTimeout(timeout time.Duration) *PostAPIV3InventoriesParams {
	var ()
	return &PostAPIV3InventoriesParams{

		timeout: timeout,
	}
}

// NewPostAPIV3InventoriesParamsWithContext creates a new PostAPIV3InventoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIV3InventoriesParamsWithContext(ctx context.Context) *PostAPIV3InventoriesParams {
	var ()
	return &PostAPIV3InventoriesParams{

		Context: ctx,
	}
}

// NewPostAPIV3InventoriesParamsWithHTTPClient creates a new PostAPIV3InventoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIV3InventoriesParamsWithHTTPClient(client *http.Client) *PostAPIV3InventoriesParams {
	var ()
	return &PostAPIV3InventoriesParams{
		HTTPClient: client,
	}
}

/*PostAPIV3InventoriesParams contains all the parameters to send to the API endpoint
for the post Api v3 inventories operation typically these are written to a http.Request
*/
type PostAPIV3InventoriesParams struct {

	/*Inventory*/
	Inventory *models.Inventory

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post Api v3 inventories params
func (o *PostAPIV3InventoriesParams) WithTimeout(timeout time.Duration) *PostAPIV3InventoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post Api v3 inventories params
func (o *PostAPIV3InventoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post Api v3 inventories params
func (o *PostAPIV3InventoriesParams) WithContext(ctx context.Context) *PostAPIV3InventoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post Api v3 inventories params
func (o *PostAPIV3InventoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post Api v3 inventories params
func (o *PostAPIV3InventoriesParams) WithHTTPClient(client *http.Client) *PostAPIV3InventoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post Api v3 inventories params
func (o *PostAPIV3InventoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInventory adds the inventory to the post Api v3 inventories params
func (o *PostAPIV3InventoriesParams) WithInventory(inventory *models.Inventory) *PostAPIV3InventoriesParams {
	o.SetInventory(inventory)
	return o
}

// SetInventory adds the inventory to the post Api v3 inventories params
func (o *PostAPIV3InventoriesParams) SetInventory(inventory *models.Inventory) {
	o.Inventory = inventory
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIV3InventoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Inventory == nil {
		o.Inventory = new(models.Inventory)
	}

	if err := r.SetBodyParam(o.Inventory); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}