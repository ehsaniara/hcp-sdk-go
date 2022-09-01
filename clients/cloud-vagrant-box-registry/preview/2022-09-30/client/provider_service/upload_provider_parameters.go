// Code generated by go-swagger; DO NOT EDIT.

package provider_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vagrant-box-registry/preview/2022-09-30/models"
)

// NewUploadProviderParams creates a new UploadProviderParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUploadProviderParams() *UploadProviderParams {
	return &UploadProviderParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUploadProviderParamsWithTimeout creates a new UploadProviderParams object
// with the ability to set a timeout on a request.
func NewUploadProviderParamsWithTimeout(timeout time.Duration) *UploadProviderParams {
	return &UploadProviderParams{
		timeout: timeout,
	}
}

// NewUploadProviderParamsWithContext creates a new UploadProviderParams object
// with the ability to set a context for a request.
func NewUploadProviderParamsWithContext(ctx context.Context) *UploadProviderParams {
	return &UploadProviderParams{
		Context: ctx,
	}
}

// NewUploadProviderParamsWithHTTPClient creates a new UploadProviderParams object
// with the ability to set a custom HTTPClient for a request.
func NewUploadProviderParamsWithHTTPClient(client *http.Client) *UploadProviderParams {
	return &UploadProviderParams{
		HTTPClient: client,
	}
}

/* UploadProviderParams contains all the parameters to send to the API endpoint
   for the upload provider operation.

   Typically these are written to a http.Request.
*/
type UploadProviderParams struct {

	// Body.
	Body *models.HashicorpCloudVagrantUploadProviderRequest

	/* Box.

	     The name segment of the Box. As an example, this field would represent the
	"vagrant" in "hashicorp/vagrant".
	*/
	Box string

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

// WithDefaults hydrates default values in the upload provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadProviderParams) WithDefaults() *UploadProviderParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upload provider params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UploadProviderParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upload provider params
func (o *UploadProviderParams) WithTimeout(timeout time.Duration) *UploadProviderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload provider params
func (o *UploadProviderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload provider params
func (o *UploadProviderParams) WithContext(ctx context.Context) *UploadProviderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload provider params
func (o *UploadProviderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload provider params
func (o *UploadProviderParams) WithHTTPClient(client *http.Client) *UploadProviderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload provider params
func (o *UploadProviderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upload provider params
func (o *UploadProviderParams) WithBody(body *models.HashicorpCloudVagrantUploadProviderRequest) *UploadProviderParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upload provider params
func (o *UploadProviderParams) SetBody(body *models.HashicorpCloudVagrantUploadProviderRequest) {
	o.Body = body
}

// WithBox adds the box to the upload provider params
func (o *UploadProviderParams) WithBox(box string) *UploadProviderParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the upload provider params
func (o *UploadProviderParams) SetBox(box string) {
	o.Box = box
}

// WithProvider adds the provider to the upload provider params
func (o *UploadProviderParams) WithProvider(provider string) *UploadProviderParams {
	o.SetProvider(provider)
	return o
}

// SetProvider adds the provider to the upload provider params
func (o *UploadProviderParams) SetProvider(provider string) {
	o.Provider = provider
}

// WithRegistry adds the registry to the upload provider params
func (o *UploadProviderParams) WithRegistry(registry string) *UploadProviderParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the upload provider params
func (o *UploadProviderParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithVersion adds the version to the upload provider params
func (o *UploadProviderParams) WithVersion(version string) *UploadProviderParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the upload provider params
func (o *UploadProviderParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *UploadProviderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param box
	if err := r.SetPathParam("box", o.Box); err != nil {
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
