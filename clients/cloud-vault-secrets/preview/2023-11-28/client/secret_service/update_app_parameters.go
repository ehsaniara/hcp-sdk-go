// Code generated by go-swagger; DO NOT EDIT.

package secret_service

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

// NewUpdateAppParams creates a new UpdateAppParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateAppParams() *UpdateAppParams {
	return &UpdateAppParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppParamsWithTimeout creates a new UpdateAppParams object
// with the ability to set a timeout on a request.
func NewUpdateAppParamsWithTimeout(timeout time.Duration) *UpdateAppParams {
	return &UpdateAppParams{
		timeout: timeout,
	}
}

// NewUpdateAppParamsWithContext creates a new UpdateAppParams object
// with the ability to set a context for a request.
func NewUpdateAppParamsWithContext(ctx context.Context) *UpdateAppParams {
	return &UpdateAppParams{
		Context: ctx,
	}
}

// NewUpdateAppParamsWithHTTPClient creates a new UpdateAppParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateAppParamsWithHTTPClient(client *http.Client) *UpdateAppParams {
	return &UpdateAppParams{
		HTTPClient: client,
	}
}

/*
UpdateAppParams contains all the parameters to send to the API endpoint

	for the update app operation.

	Typically these are written to a http.Request.
*/
type UpdateAppParams struct {

	// Body.
	Body UpdateAppBody

	// Name.
	Name string

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppParams) WithDefaults() *UpdateAppParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update app params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateAppParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update app params
func (o *UpdateAppParams) WithTimeout(timeout time.Duration) *UpdateAppParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update app params
func (o *UpdateAppParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update app params
func (o *UpdateAppParams) WithContext(ctx context.Context) *UpdateAppParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update app params
func (o *UpdateAppParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update app params
func (o *UpdateAppParams) WithHTTPClient(client *http.Client) *UpdateAppParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update app params
func (o *UpdateAppParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update app params
func (o *UpdateAppParams) WithBody(body UpdateAppBody) *UpdateAppParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update app params
func (o *UpdateAppParams) SetBody(body UpdateAppBody) {
	o.Body = body
}

// WithName adds the name to the update app params
func (o *UpdateAppParams) WithName(name string) *UpdateAppParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the update app params
func (o *UpdateAppParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the update app params
func (o *UpdateAppParams) WithOrganizationID(organizationID string) *UpdateAppParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update app params
func (o *UpdateAppParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update app params
func (o *UpdateAppParams) WithProjectID(projectID string) *UpdateAppParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update app params
func (o *UpdateAppParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
