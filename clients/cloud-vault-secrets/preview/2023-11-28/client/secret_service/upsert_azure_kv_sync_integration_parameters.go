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

// NewUpsertAzureKvSyncIntegrationParams creates a new UpsertAzureKvSyncIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpsertAzureKvSyncIntegrationParams() *UpsertAzureKvSyncIntegrationParams {
	return &UpsertAzureKvSyncIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpsertAzureKvSyncIntegrationParamsWithTimeout creates a new UpsertAzureKvSyncIntegrationParams object
// with the ability to set a timeout on a request.
func NewUpsertAzureKvSyncIntegrationParamsWithTimeout(timeout time.Duration) *UpsertAzureKvSyncIntegrationParams {
	return &UpsertAzureKvSyncIntegrationParams{
		timeout: timeout,
	}
}

// NewUpsertAzureKvSyncIntegrationParamsWithContext creates a new UpsertAzureKvSyncIntegrationParams object
// with the ability to set a context for a request.
func NewUpsertAzureKvSyncIntegrationParamsWithContext(ctx context.Context) *UpsertAzureKvSyncIntegrationParams {
	return &UpsertAzureKvSyncIntegrationParams{
		Context: ctx,
	}
}

// NewUpsertAzureKvSyncIntegrationParamsWithHTTPClient creates a new UpsertAzureKvSyncIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpsertAzureKvSyncIntegrationParamsWithHTTPClient(client *http.Client) *UpsertAzureKvSyncIntegrationParams {
	return &UpsertAzureKvSyncIntegrationParams{
		HTTPClient: client,
	}
}

/*
UpsertAzureKvSyncIntegrationParams contains all the parameters to send to the API endpoint

	for the upsert azure kv sync integration operation.

	Typically these are written to a http.Request.
*/
type UpsertAzureKvSyncIntegrationParams struct {

	// Body.
	Body UpsertAzureKvSyncIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upsert azure kv sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertAzureKvSyncIntegrationParams) WithDefaults() *UpsertAzureKvSyncIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upsert azure kv sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertAzureKvSyncIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) WithTimeout(timeout time.Duration) *UpsertAzureKvSyncIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) WithContext(ctx context.Context) *UpsertAzureKvSyncIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) WithHTTPClient(client *http.Client) *UpsertAzureKvSyncIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) WithBody(body UpsertAzureKvSyncIntegrationBody) *UpsertAzureKvSyncIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) SetBody(body UpsertAzureKvSyncIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) WithOrganizationID(organizationID string) *UpsertAzureKvSyncIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) WithProjectID(projectID string) *UpsertAzureKvSyncIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the upsert azure kv sync integration params
func (o *UpsertAzureKvSyncIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpsertAzureKvSyncIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
