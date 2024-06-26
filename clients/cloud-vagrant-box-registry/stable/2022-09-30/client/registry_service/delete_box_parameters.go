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

// NewDeleteBoxParams creates a new DeleteBoxParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteBoxParams() *DeleteBoxParams {
	return &DeleteBoxParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBoxParamsWithTimeout creates a new DeleteBoxParams object
// with the ability to set a timeout on a request.
func NewDeleteBoxParamsWithTimeout(timeout time.Duration) *DeleteBoxParams {
	return &DeleteBoxParams{
		timeout: timeout,
	}
}

// NewDeleteBoxParamsWithContext creates a new DeleteBoxParams object
// with the ability to set a context for a request.
func NewDeleteBoxParamsWithContext(ctx context.Context) *DeleteBoxParams {
	return &DeleteBoxParams{
		Context: ctx,
	}
}

// NewDeleteBoxParamsWithHTTPClient creates a new DeleteBoxParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteBoxParamsWithHTTPClient(client *http.Client) *DeleteBoxParams {
	return &DeleteBoxParams{
		HTTPClient: client,
	}
}

/*
DeleteBoxParams contains all the parameters to send to the API endpoint

	for the delete box operation.

	Typically these are written to a http.Request.
*/
type DeleteBoxParams struct {

	/* Box.

	     The name segment of the Box to delete. As an example, this field would
	represent the "vagrant" in "hashicorp/vagrant".
	*/
	Box string

	/* Registry.

	     The Registry segment of the Box. As an example, this field would represent
	the "hashicorp" in "hashicorp/vagrant".
	*/
	Registry string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete box params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBoxParams) WithDefaults() *DeleteBoxParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete box params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteBoxParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete box params
func (o *DeleteBoxParams) WithTimeout(timeout time.Duration) *DeleteBoxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete box params
func (o *DeleteBoxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete box params
func (o *DeleteBoxParams) WithContext(ctx context.Context) *DeleteBoxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete box params
func (o *DeleteBoxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete box params
func (o *DeleteBoxParams) WithHTTPClient(client *http.Client) *DeleteBoxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete box params
func (o *DeleteBoxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBox adds the box to the delete box params
func (o *DeleteBoxParams) WithBox(box string) *DeleteBoxParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the delete box params
func (o *DeleteBoxParams) SetBox(box string) {
	o.Box = box
}

// WithRegistry adds the registry to the delete box params
func (o *DeleteBoxParams) WithRegistry(registry string) *DeleteBoxParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the delete box params
func (o *DeleteBoxParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBoxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param box
	if err := r.SetPathParam("box", o.Box); err != nil {
		return err
	}

	// path param registry
	if err := r.SetPathParam("registry", o.Registry); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
