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

// NewCreateMongoDBAtlasIntegrationParams creates a new CreateMongoDBAtlasIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateMongoDBAtlasIntegrationParams() *CreateMongoDBAtlasIntegrationParams {
	return &CreateMongoDBAtlasIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMongoDBAtlasIntegrationParamsWithTimeout creates a new CreateMongoDBAtlasIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreateMongoDBAtlasIntegrationParamsWithTimeout(timeout time.Duration) *CreateMongoDBAtlasIntegrationParams {
	return &CreateMongoDBAtlasIntegrationParams{
		timeout: timeout,
	}
}

// NewCreateMongoDBAtlasIntegrationParamsWithContext creates a new CreateMongoDBAtlasIntegrationParams object
// with the ability to set a context for a request.
func NewCreateMongoDBAtlasIntegrationParamsWithContext(ctx context.Context) *CreateMongoDBAtlasIntegrationParams {
	return &CreateMongoDBAtlasIntegrationParams{
		Context: ctx,
	}
}

// NewCreateMongoDBAtlasIntegrationParamsWithHTTPClient creates a new CreateMongoDBAtlasIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateMongoDBAtlasIntegrationParamsWithHTTPClient(client *http.Client) *CreateMongoDBAtlasIntegrationParams {
	return &CreateMongoDBAtlasIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreateMongoDBAtlasIntegrationParams contains all the parameters to send to the API endpoint

	for the create mongo d b atlas integration operation.

	Typically these are written to a http.Request.
*/
type CreateMongoDBAtlasIntegrationParams struct {

	// Body.
	Body CreateMongoDBAtlasIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create mongo d b atlas integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMongoDBAtlasIntegrationParams) WithDefaults() *CreateMongoDBAtlasIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create mongo d b atlas integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateMongoDBAtlasIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) WithTimeout(timeout time.Duration) *CreateMongoDBAtlasIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) WithContext(ctx context.Context) *CreateMongoDBAtlasIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) WithHTTPClient(client *http.Client) *CreateMongoDBAtlasIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) WithBody(body CreateMongoDBAtlasIntegrationBody) *CreateMongoDBAtlasIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) SetBody(body CreateMongoDBAtlasIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) WithOrganizationID(organizationID string) *CreateMongoDBAtlasIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) WithProjectID(projectID string) *CreateMongoDBAtlasIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create mongo d b atlas integration params
func (o *CreateMongoDBAtlasIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMongoDBAtlasIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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