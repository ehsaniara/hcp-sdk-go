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

// NewGetGitHubInstallLinksParams creates a new GetGitHubInstallLinksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGitHubInstallLinksParams() *GetGitHubInstallLinksParams {
	return &GetGitHubInstallLinksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGitHubInstallLinksParamsWithTimeout creates a new GetGitHubInstallLinksParams object
// with the ability to set a timeout on a request.
func NewGetGitHubInstallLinksParamsWithTimeout(timeout time.Duration) *GetGitHubInstallLinksParams {
	return &GetGitHubInstallLinksParams{
		timeout: timeout,
	}
}

// NewGetGitHubInstallLinksParamsWithContext creates a new GetGitHubInstallLinksParams object
// with the ability to set a context for a request.
func NewGetGitHubInstallLinksParamsWithContext(ctx context.Context) *GetGitHubInstallLinksParams {
	return &GetGitHubInstallLinksParams{
		Context: ctx,
	}
}

// NewGetGitHubInstallLinksParamsWithHTTPClient creates a new GetGitHubInstallLinksParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGitHubInstallLinksParamsWithHTTPClient(client *http.Client) *GetGitHubInstallLinksParams {
	return &GetGitHubInstallLinksParams{
		HTTPClient: client,
	}
}

/*
GetGitHubInstallLinksParams contains all the parameters to send to the API endpoint

	for the get git hub install links operation.

	Typically these are written to a http.Request.
*/
type GetGitHubInstallLinksParams struct {

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get git hub install links params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitHubInstallLinksParams) WithDefaults() *GetGitHubInstallLinksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get git hub install links params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGitHubInstallLinksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get git hub install links params
func (o *GetGitHubInstallLinksParams) WithTimeout(timeout time.Duration) *GetGitHubInstallLinksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get git hub install links params
func (o *GetGitHubInstallLinksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get git hub install links params
func (o *GetGitHubInstallLinksParams) WithContext(ctx context.Context) *GetGitHubInstallLinksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get git hub install links params
func (o *GetGitHubInstallLinksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get git hub install links params
func (o *GetGitHubInstallLinksParams) WithHTTPClient(client *http.Client) *GetGitHubInstallLinksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get git hub install links params
func (o *GetGitHubInstallLinksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get git hub install links params
func (o *GetGitHubInstallLinksParams) WithOrganizationID(organizationID string) *GetGitHubInstallLinksParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get git hub install links params
func (o *GetGitHubInstallLinksParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the get git hub install links params
func (o *GetGitHubInstallLinksParams) WithProjectID(projectID string) *GetGitHubInstallLinksParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get git hub install links params
func (o *GetGitHubInstallLinksParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGitHubInstallLinksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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