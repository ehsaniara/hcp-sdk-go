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

// NewCompleteDirectUploadBoxParams creates a new CompleteDirectUploadBoxParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCompleteDirectUploadBoxParams() *CompleteDirectUploadBoxParams {
	return &CompleteDirectUploadBoxParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCompleteDirectUploadBoxParamsWithTimeout creates a new CompleteDirectUploadBoxParams object
// with the ability to set a timeout on a request.
func NewCompleteDirectUploadBoxParamsWithTimeout(timeout time.Duration) *CompleteDirectUploadBoxParams {
	return &CompleteDirectUploadBoxParams{
		timeout: timeout,
	}
}

// NewCompleteDirectUploadBoxParamsWithContext creates a new CompleteDirectUploadBoxParams object
// with the ability to set a context for a request.
func NewCompleteDirectUploadBoxParamsWithContext(ctx context.Context) *CompleteDirectUploadBoxParams {
	return &CompleteDirectUploadBoxParams{
		Context: ctx,
	}
}

// NewCompleteDirectUploadBoxParamsWithHTTPClient creates a new CompleteDirectUploadBoxParams object
// with the ability to set a custom HTTPClient for a request.
func NewCompleteDirectUploadBoxParamsWithHTTPClient(client *http.Client) *CompleteDirectUploadBoxParams {
	return &CompleteDirectUploadBoxParams{
		HTTPClient: client,
	}
}

/*
CompleteDirectUploadBoxParams contains all the parameters to send to the API endpoint

	for the complete direct upload box operation.

	Typically these are written to a http.Request.
*/
type CompleteDirectUploadBoxParams struct {

	/* Architecture.

	   The name of the Architecture.
	*/
	Architecture string

	/* Box.

	     The name segment of the Box. As an example, this field would represent the
	"vagrant" in "hashicorp/vagrant".
	*/
	Box string

	/* Object.

	   Object identifier.
	*/
	Object string

	/* Provider.

	   The name of the Provider.
	*/
	Provider string

	/* Registry.

	     The Registry segment of the Box. As an example, this field would represent
	the "hashicorp" in "hashicorp/vagrant".
	*/
	Registry string

	/* Version.

	   The name of the Version for the Provider.
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the complete direct upload box params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CompleteDirectUploadBoxParams) WithDefaults() *CompleteDirectUploadBoxParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the complete direct upload box params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CompleteDirectUploadBoxParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithTimeout(timeout time.Duration) *CompleteDirectUploadBoxParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithContext(ctx context.Context) *CompleteDirectUploadBoxParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithHTTPClient(client *http.Client) *CompleteDirectUploadBoxParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArchitecture adds the architecture to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithArchitecture(architecture string) *CompleteDirectUploadBoxParams {
	o.SetArchitecture(architecture)
	return o
}

// SetArchitecture adds the architecture to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetArchitecture(architecture string) {
	o.Architecture = architecture
}

// WithBox adds the box to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithBox(box string) *CompleteDirectUploadBoxParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetBox(box string) {
	o.Box = box
}

// WithObject adds the object to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithObject(object string) *CompleteDirectUploadBoxParams {
	o.SetObject(object)
	return o
}

// SetObject adds the object to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetObject(object string) {
	o.Object = object
}

// WithProvider adds the provider to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithProvider(provider string) *CompleteDirectUploadBoxParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithRegistry adds the registry to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithRegistry(registry string) *CompleteDirectUploadBoxParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithVersion adds the version to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) WithVersion(version string) *CompleteDirectUploadBoxParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the complete direct upload box params
func (o *CompleteDirectUploadBoxParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *CompleteDirectUploadBoxParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param architecture
	if err := r.SetPathParam("architecture", o.Architecture); err != nil {
		return err
	}

	// path param box
	if err := r.SetPathParam("box", o.Box); err != nil {
		return err
	}

	// path param object
	if err := r.SetPathParam("object", o.Object); err != nil {
		return err
	}

	// path param provider
	if err := r.SetPathParam("provider", o.Provider); err != nil {
		return err
	}

	// path param registry
	if err := r.SetPathParam("registry", o.Registry); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
