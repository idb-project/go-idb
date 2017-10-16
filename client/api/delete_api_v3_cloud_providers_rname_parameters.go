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

// NewDeleteAPIV3CloudProvidersRnameParams creates a new DeleteAPIV3CloudProvidersRnameParams object
// with the default values initialized.
func NewDeleteAPIV3CloudProvidersRnameParams() *DeleteAPIV3CloudProvidersRnameParams {
	var ()
	return &DeleteAPIV3CloudProvidersRnameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIV3CloudProvidersRnameParamsWithTimeout creates a new DeleteAPIV3CloudProvidersRnameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIV3CloudProvidersRnameParamsWithTimeout(timeout time.Duration) *DeleteAPIV3CloudProvidersRnameParams {
	var ()
	return &DeleteAPIV3CloudProvidersRnameParams{

		timeout: timeout,
	}
}

// NewDeleteAPIV3CloudProvidersRnameParamsWithContext creates a new DeleteAPIV3CloudProvidersRnameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIV3CloudProvidersRnameParamsWithContext(ctx context.Context) *DeleteAPIV3CloudProvidersRnameParams {
	var ()
	return &DeleteAPIV3CloudProvidersRnameParams{

		Context: ctx,
	}
}

// NewDeleteAPIV3CloudProvidersRnameParamsWithHTTPClient creates a new DeleteAPIV3CloudProvidersRnameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIV3CloudProvidersRnameParamsWithHTTPClient(client *http.Client) *DeleteAPIV3CloudProvidersRnameParams {
	var ()
	return &DeleteAPIV3CloudProvidersRnameParams{
		HTTPClient: client,
	}
}

/*DeleteAPIV3CloudProvidersRnameParams contains all the parameters to send to the API endpoint
for the delete Api v3 cloud providers rname operation typically these are written to a http.Request
*/
type DeleteAPIV3CloudProvidersRnameParams struct {

	/*Rname*/
	Rname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete Api v3 cloud providers rname params
func (o *DeleteAPIV3CloudProvidersRnameParams) WithTimeout(timeout time.Duration) *DeleteAPIV3CloudProvidersRnameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete Api v3 cloud providers rname params
func (o *DeleteAPIV3CloudProvidersRnameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete Api v3 cloud providers rname params
func (o *DeleteAPIV3CloudProvidersRnameParams) WithContext(ctx context.Context) *DeleteAPIV3CloudProvidersRnameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete Api v3 cloud providers rname params
func (o *DeleteAPIV3CloudProvidersRnameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete Api v3 cloud providers rname params
func (o *DeleteAPIV3CloudProvidersRnameParams) WithHTTPClient(client *http.Client) *DeleteAPIV3CloudProvidersRnameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete Api v3 cloud providers rname params
func (o *DeleteAPIV3CloudProvidersRnameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRname adds the rname to the delete Api v3 cloud providers rname params
func (o *DeleteAPIV3CloudProvidersRnameParams) WithRname(rname string) *DeleteAPIV3CloudProvidersRnameParams {
	o.SetRname(rname)
	return o
}

// SetRname adds the rname to the delete Api v3 cloud providers rname params
func (o *DeleteAPIV3CloudProvidersRnameParams) SetRname(rname string) {
	o.Rname = rname
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIV3CloudProvidersRnameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param rname
	if err := r.SetPathParam("rname", o.Rname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}