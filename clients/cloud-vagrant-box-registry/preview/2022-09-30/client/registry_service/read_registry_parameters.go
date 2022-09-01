// Code generated by go-swagger; DO NOT EDIT.

package registry_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewReadRegistryParams creates a new ReadRegistryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReadRegistryParams() *ReadRegistryParams {
	return &ReadRegistryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReadRegistryParamsWithTimeout creates a new ReadRegistryParams object
// with the ability to set a timeout on a request.
func NewReadRegistryParamsWithTimeout(timeout time.Duration) *ReadRegistryParams {
	return &ReadRegistryParams{
		timeout: timeout,
	}
}

// NewReadRegistryParamsWithContext creates a new ReadRegistryParams object
// with the ability to set a context for a request.
func NewReadRegistryParamsWithContext(ctx context.Context) *ReadRegistryParams {
	return &ReadRegistryParams{
		Context: ctx,
	}
}

// NewReadRegistryParamsWithHTTPClient creates a new ReadRegistryParams object
// with the ability to set a custom HTTPClient for a request.
func NewReadRegistryParamsWithHTTPClient(client *http.Client) *ReadRegistryParams {
	return &ReadRegistryParams{
		HTTPClient: client,
	}
}

/* ReadRegistryParams contains all the parameters to send to the API endpoint
   for the read registry operation.

   Typically these are written to a http.Request.
*/
type ReadRegistryParams struct {

	/* Registry.

	   The name of the Registry to look up.
	*/
	Registry string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the read registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadRegistryParams) WithDefaults() *ReadRegistryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the read registry params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReadRegistryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the read registry params
func (o *ReadRegistryParams) WithTimeout(timeout time.Duration) *ReadRegistryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read registry params
func (o *ReadRegistryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read registry params
func (o *ReadRegistryParams) WithContext(ctx context.Context) *ReadRegistryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read registry params
func (o *ReadRegistryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read registry params
func (o *ReadRegistryParams) WithHTTPClient(client *http.Client) *ReadRegistryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read registry params
func (o *ReadRegistryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRegistry adds the registry to the read registry params
func (o *ReadRegistryParams) WithRegistry(registry string) *ReadRegistryParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the read registry params
func (o *ReadRegistryParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *ReadRegistryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param registry
	if err := r.SetPathParam("registry", o.Registry); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
