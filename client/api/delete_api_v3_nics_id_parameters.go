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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteAPIV3NicsIDParams creates a new DeleteAPIV3NicsIDParams object
// with the default values initialized.
func NewDeleteAPIV3NicsIDParams() *DeleteAPIV3NicsIDParams {
	var ()
	return &DeleteAPIV3NicsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV3NicsIDParamsWithTimeout creates a new DeleteAPIV3NicsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV3NicsIDParamsWithTimeout(timeout time.Duration) *DeleteAPIV3NicsIDParams {
	var ()
	return &DeleteAPIV3NicsIDParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV3NicsIDParamsWithContext creates a new DeleteAPIV3NicsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV3NicsIDParamsWithContext(ctx context.Context) *DeleteAPIV3NicsIDParams {
	var ()
	return &DeleteAPIV3NicsIDParams{

		Context: ctx,
	}
}

// NewDeleteAPIV3NicsIDParamsWithHTTPClient creates a new DeleteAPIV3NicsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV3NicsIDParamsWithHTTPClient(client *http.Client) *DeleteAPIV3NicsIDParams {
	var ()
	return &DeleteAPIV3NicsIDParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV3NicsIDParams contains all the parameters to send to the API endpoint
for the delete Api v3 nics Id operation typically these are written to a http.Request
*/
type DeleteAPIV3NicsIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete Api v3 nics Id params
func (o *DeleteAPIV3NicsIDParams) WithTimeout(timeout time.Duration) *DeleteAPIV3NicsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api v3 nics Id params
func (o *DeleteAPIV3NicsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api v3 nics Id params
func (o *DeleteAPIV3NicsIDParams) WithContext(ctx context.Context) *DeleteAPIV3NicsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api v3 nics Id params
func (o *DeleteAPIV3NicsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api v3 nics Id params
func (o *DeleteAPIV3NicsIDParams) WithHTTPClient(client *http.Client) *DeleteAPIV3NicsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api v3 nics Id params
func (o *DeleteAPIV3NicsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete Api v3 nics Id params
func (o *DeleteAPIV3NicsIDParams) WithID(id int32) *DeleteAPIV3NicsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete Api v3 nics Id params
func (o *DeleteAPIV3NicsIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV3NicsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
