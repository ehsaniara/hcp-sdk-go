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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-vault-secrets/preview/2023-11-28/models"
)

// NewCreateHcpTerraformSyncIntegrationParams creates a new CreateHcpTerraformSyncIntegrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateHcpTerraformSyncIntegrationParams() *CreateHcpTerraformSyncIntegrationParams {
	return &CreateHcpTerraformSyncIntegrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateHcpTerraformSyncIntegrationParamsWithTimeout creates a new CreateHcpTerraformSyncIntegrationParams object
// with the ability to set a timeout on a request.
func NewCreateHcpTerraformSyncIntegrationParamsWithTimeout(timeout time.Duration) *CreateHcpTerraformSyncIntegrationParams {
	return &CreateHcpTerraformSyncIntegrationParams{
		timeout: timeout,
	}
}

// NewCreateHcpTerraformSyncIntegrationParamsWithContext creates a new CreateHcpTerraformSyncIntegrationParams object
// with the ability to set a context for a request.
func NewCreateHcpTerraformSyncIntegrationParamsWithContext(ctx context.Context) *CreateHcpTerraformSyncIntegrationParams {
	return &CreateHcpTerraformSyncIntegrationParams{
		Context: ctx,
	}
}

// NewCreateHcpTerraformSyncIntegrationParamsWithHTTPClient creates a new CreateHcpTerraformSyncIntegrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateHcpTerraformSyncIntegrationParamsWithHTTPClient(client *http.Client) *CreateHcpTerraformSyncIntegrationParams {
	return &CreateHcpTerraformSyncIntegrationParams{
		HTTPClient: client,
	}
}

/*
CreateHcpTerraformSyncIntegrationParams contains all the parameters to send to the API endpoint

	for the create hcp terraform sync integration operation.

	Typically these are written to a http.Request.
*/
type CreateHcpTerraformSyncIntegrationParams struct {

	// Body.
	Body *models.SecretServiceCreateHcpTerraformSyncIntegrationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create hcp terraform sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHcpTerraformSyncIntegrationParams) WithDefaults() *CreateHcpTerraformSyncIntegrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create hcp terraform sync integration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateHcpTerraformSyncIntegrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) WithTimeout(timeout time.Duration) *CreateHcpTerraformSyncIntegrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) WithContext(ctx context.Context) *CreateHcpTerraformSyncIntegrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) WithHTTPClient(client *http.Client) *CreateHcpTerraformSyncIntegrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) WithBody(body *models.SecretServiceCreateHcpTerraformSyncIntegrationBody) *CreateHcpTerraformSyncIntegrationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) SetBody(body *models.SecretServiceCreateHcpTerraformSyncIntegrationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) WithOrganizationID(organizationID string) *CreateHcpTerraformSyncIntegrationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) WithProjectID(projectID string) *CreateHcpTerraformSyncIntegrationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create hcp terraform sync integration params
func (o *CreateHcpTerraformSyncIntegrationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateHcpTerraformSyncIntegrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
