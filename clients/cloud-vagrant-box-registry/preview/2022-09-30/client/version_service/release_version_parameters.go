// Code generated by go-swagger; DO NOT EDIT.

package version_service

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

// NewReleaseVersionParams creates a new ReleaseVersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewReleaseVersionParams() *ReleaseVersionParams {
	return &ReleaseVersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewReleaseVersionParamsWithTimeout creates a new ReleaseVersionParams object
// with the ability to set a timeout on a request.
func NewReleaseVersionParamsWithTimeout(timeout time.Duration) *ReleaseVersionParams {
	return &ReleaseVersionParams{
		timeout: timeout,
	}
}

// NewReleaseVersionParamsWithContext creates a new ReleaseVersionParams object
// with the ability to set a context for a request.
func NewReleaseVersionParamsWithContext(ctx context.Context) *ReleaseVersionParams {
	return &ReleaseVersionParams{
		Context: ctx,
	}
}

// NewReleaseVersionParamsWithHTTPClient creates a new ReleaseVersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewReleaseVersionParamsWithHTTPClient(client *http.Client) *ReleaseVersionParams {
	return &ReleaseVersionParams{
		HTTPClient: client,
	}
}

/* ReleaseVersionParams contains all the parameters to send to the API endpoint
   for the release version operation.

   Typically these are written to a http.Request.
*/
type ReleaseVersionParams struct {

	// Body.
	Body *models.HashicorpCloudVagrantReleaseVersionRequest

	/* Box.

	     The name segment of the Box. As an example, this field would represent the
	"vagrant" in "hashicorp/vagrant".
	*/
	Box string

	/* Registry.

	     The Registry segment of the Box. As an example, this field would represent
	the "hashicorp" in "hashicorp/vagrant".
	*/
	Registry string

	/* Version.

	   The name of the Version to release.
	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the release version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseVersionParams) WithDefaults() *ReleaseVersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the release version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ReleaseVersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the release version params
func (o *ReleaseVersionParams) WithTimeout(timeout time.Duration) *ReleaseVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the release version params
func (o *ReleaseVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the release version params
func (o *ReleaseVersionParams) WithContext(ctx context.Context) *ReleaseVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the release version params
func (o *ReleaseVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the release version params
func (o *ReleaseVersionParams) WithHTTPClient(client *http.Client) *ReleaseVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the release version params
func (o *ReleaseVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the release version params
func (o *ReleaseVersionParams) WithBody(body *models.HashicorpCloudVagrantReleaseVersionRequest) *ReleaseVersionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the release version params
func (o *ReleaseVersionParams) SetBody(body *models.HashicorpCloudVagrantReleaseVersionRequest) {
	o.Body = body
}

// WithBox adds the box to the release version params
func (o *ReleaseVersionParams) WithBox(box string) *ReleaseVersionParams {
	o.SetBox(box)
	return o
}

// SetBox adds the box to the release version params
func (o *ReleaseVersionParams) SetBox(box string) {
	o.Box = box
}

// WithRegistry adds the registry to the release version params
func (o *ReleaseVersionParams) WithRegistry(registry string) *ReleaseVersionParams {
	o.SetRegistry(registry)
	return o
}

// SetRegistry adds the registry to the release version params
func (o *ReleaseVersionParams) SetRegistry(registry string) {
	o.Registry = registry
}

// WithVersion adds the version to the release version params
func (o *ReleaseVersionParams) WithVersion(version string) *ReleaseVersionParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the release version params
func (o *ReleaseVersionParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ReleaseVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
