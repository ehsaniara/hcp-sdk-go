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

// NewConnectGitHubInstallationParams creates a new ConnectGitHubInstallationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewConnectGitHubInstallationParams() *ConnectGitHubInstallationParams {
	return &ConnectGitHubInstallationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewConnectGitHubInstallationParamsWithTimeout creates a new ConnectGitHubInstallationParams object
// with the ability to set a timeout on a request.
func NewConnectGitHubInstallationParamsWithTimeout(timeout time.Duration) *ConnectGitHubInstallationParams {
	return &ConnectGitHubInstallationParams{
		timeout: timeout,
	}
}

// NewConnectGitHubInstallationParamsWithContext creates a new ConnectGitHubInstallationParams object
// with the ability to set a context for a request.
func NewConnectGitHubInstallationParamsWithContext(ctx context.Context) *ConnectGitHubInstallationParams {
	return &ConnectGitHubInstallationParams{
		Context: ctx,
	}
}

// NewConnectGitHubInstallationParamsWithHTTPClient creates a new ConnectGitHubInstallationParams object
// with the ability to set a custom HTTPClient for a request.
func NewConnectGitHubInstallationParamsWithHTTPClient(client *http.Client) *ConnectGitHubInstallationParams {
	return &ConnectGitHubInstallationParams{
		HTTPClient: client,
	}
}

/*
ConnectGitHubInstallationParams contains all the parameters to send to the API endpoint

	for the connect git hub installation operation.

	Typically these are written to a http.Request.
*/
type ConnectGitHubInstallationParams struct {

	// Body.
	Body *models.SecretServiceConnectGitHubInstallationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the connect git hub installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectGitHubInstallationParams) WithDefaults() *ConnectGitHubInstallationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the connect git hub installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ConnectGitHubInstallationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) WithTimeout(timeout time.Duration) *ConnectGitHubInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) WithContext(ctx context.Context) *ConnectGitHubInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) WithHTTPClient(client *http.Client) *ConnectGitHubInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) WithBody(body *models.SecretServiceConnectGitHubInstallationBody) *ConnectGitHubInstallationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) SetBody(body *models.SecretServiceConnectGitHubInstallationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) WithOrganizationID(organizationID string) *ConnectGitHubInstallationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) WithProjectID(projectID string) *ConnectGitHubInstallationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the connect git hub installation params
func (o *ConnectGitHubInstallationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectGitHubInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
