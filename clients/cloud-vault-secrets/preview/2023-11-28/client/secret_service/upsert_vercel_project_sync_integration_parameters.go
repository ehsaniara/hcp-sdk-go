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

// NewUpsertVercelProjectSyncIntegrationParams creates a new UpsertVercelProjectSyncIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpsertVercelProjectSyncIntegrationParams() *UpsertVercelProjectSyncIntegrationParams {
	return &UpsertVercelProjectSyncIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpsertVercelProjectSyncIntegrationParamsWithTimeout creates a new UpsertVercelProjectSyncIntegrationParams object
// with the ability to set a timeout on a request.
func NewUpsertVercelProjectSyncIntegrationParamsWithTimeout(timeout time.Duration) *UpsertVercelProjectSyncIntegrationParams {
	return &UpsertVercelProjectSyncIntegrationParams{
		timeout: timeout,
	}
}

// NewUpsertVercelProjectSyncIntegrationParamsWithContext creates a new UpsertVercelProjectSyncIntegrationParams object
// with the ability to set a context for a request.
func NewUpsertVercelProjectSyncIntegrationParamsWithContext(ctx context.Context) *UpsertVercelProjectSyncIntegrationParams {
	return &UpsertVercelProjectSyncIntegrationParams{
		Context: ctx,
	}
}

// NewUpsertVercelProjectSyncIntegrationParamsWithHTTPClient creates a new UpsertVercelProjectSyncIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpsertVercelProjectSyncIntegrationParamsWithHTTPClient(client *http.Client) *UpsertVercelProjectSyncIntegrationParams {
	return &UpsertVercelProjectSyncIntegrationParams{
		HTTPClient: client,
	}
}

/*
UpsertVercelProjectSyncIntegrationParams contains all the parameters to send to the API endpoint

	for the upsert vercel project sync integration operation.

	Typically these are written to a http.Request.
*/
type UpsertVercelProjectSyncIntegrationParams struct {

	// Body.
	Body UpsertVercelProjectSyncIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upsert vercel project sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertVercelProjectSyncIntegrationParams) WithDefaults() *UpsertVercelProjectSyncIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upsert vercel project sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertVercelProjectSyncIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) WithTimeout(timeout time.Duration) *UpsertVercelProjectSyncIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) WithContext(ctx context.Context) *UpsertVercelProjectSyncIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) WithHTTPClient(client *http.Client) *UpsertVercelProjectSyncIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) WithBody(body UpsertVercelProjectSyncIntegrationBody) *UpsertVercelProjectSyncIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) SetBody(body UpsertVercelProjectSyncIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) WithOrganizationID(organizationID string) *UpsertVercelProjectSyncIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) WithProjectID(projectID string) *UpsertVercelProjectSyncIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the upsert vercel project sync integration params
func (o *UpsertVercelProjectSyncIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpsertVercelProjectSyncIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
