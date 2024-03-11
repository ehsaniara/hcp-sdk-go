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

// NewForceSyncParams creates a new ForceSyncParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewForceSyncParams() *ForceSyncParams {
	return &ForceSyncParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewForceSyncParamsWithTimeout creates a new ForceSyncParams object
// with the ability to set a timeout on a request.
func NewForceSyncParamsWithTimeout(timeout time.Duration) *ForceSyncParams {
	return &ForceSyncParams{
		timeout: timeout,
	}
}

// NewForceSyncParamsWithContext creates a new ForceSyncParams object
// with the ability to set a context for a request.
func NewForceSyncParamsWithContext(ctx context.Context) *ForceSyncParams {
	return &ForceSyncParams{
		Context: ctx,
	}
}

// NewForceSyncParamsWithHTTPClient creates a new ForceSyncParams object
// with the ability to set a custom HTTPClient for a request.
func NewForceSyncParamsWithHTTPClient(client *http.Client) *ForceSyncParams {
	return &ForceSyncParams{
		HTTPClient: client,
	}
}

/*
ForceSyncParams contains all the parameters to send to the API endpoint

	for the force sync operation.

	Typically these are written to a http.Request.
*/
type ForceSyncParams struct {

	// Body.
	Body ForceSyncBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the force sync params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForceSyncParams) WithDefaults() *ForceSyncParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the force sync params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ForceSyncParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the force sync params
func (o *ForceSyncParams) WithTimeout(timeout time.Duration) *ForceSyncParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the force sync params
func (o *ForceSyncParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the force sync params
func (o *ForceSyncParams) WithContext(ctx context.Context) *ForceSyncParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the force sync params
func (o *ForceSyncParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the force sync params
func (o *ForceSyncParams) WithHTTPClient(client *http.Client) *ForceSyncParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the force sync params
func (o *ForceSyncParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the force sync params
func (o *ForceSyncParams) WithBody(body ForceSyncBody) *ForceSyncParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the force sync params
func (o *ForceSyncParams) SetBody(body ForceSyncBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the force sync params
func (o *ForceSyncParams) WithOrganizationID(organizationID string) *ForceSyncParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the force sync params
func (o *ForceSyncParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the force sync params
func (o *ForceSyncParams) WithProjectID(projectID string) *ForceSyncParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the force sync params
func (o *ForceSyncParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ForceSyncParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
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
