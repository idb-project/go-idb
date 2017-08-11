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

// NewGetAPIV3SwitchesParams creates a new GetAPIV3SwitchesParams object
// with the default values initialized.
func NewGetAPIV3SwitchesParams() *GetAPIV3SwitchesParams {

	return &GetAPIV3SwitchesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIV3SwitchesParamsWithTimeout creates a new GetAPIV3SwitchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIV3SwitchesParamsWithTimeout(timeout time.Duration) *GetAPIV3SwitchesParams {

	return &GetAPIV3SwitchesParams{

		timeout: timeout,
	}
}

// NewGetAPIV3SwitchesParamsWithContext creates a new GetAPIV3SwitchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIV3SwitchesParamsWithContext(ctx context.Context) *GetAPIV3SwitchesParams {

	return &GetAPIV3SwitchesParams{

		Context: ctx,
	}
}

// NewGetAPIV3SwitchesParamsWithHTTPClient creates a new GetAPIV3SwitchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIV3SwitchesParamsWithHTTPClient(client *http.Client) *GetAPIV3SwitchesParams {

	return &GetAPIV3SwitchesParams{
		HTTPClient: client,
	}
}

/*GetAPIV3SwitchesParams contains all the parameters to send to the API endpoint
for the get Api v3 switches operation typically these are written to a http.Request
*/
type GetAPIV3SwitchesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get Api v3 switches params
func (o *GetAPIV3SwitchesParams) WithTimeout(timeout time.Duration) *GetAPIV3SwitchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Api v3 switches params
func (o *GetAPIV3SwitchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Api v3 switches params
func (o *GetAPIV3SwitchesParams) WithContext(ctx context.Context) *GetAPIV3SwitchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Api v3 switches params
func (o *GetAPIV3SwitchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Api v3 switches params
func (o *GetAPIV3SwitchesParams) WithHTTPClient(client *http.Client) *GetAPIV3SwitchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Api v3 switches params
func (o *GetAPIV3SwitchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIV3SwitchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
