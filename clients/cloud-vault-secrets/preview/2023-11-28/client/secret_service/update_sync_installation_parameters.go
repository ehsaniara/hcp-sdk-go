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

// NewUpdateSyncInstallationParams creates a new UpdateSyncInstallationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSyncInstallationParams() *UpdateSyncInstallationParams {
	return &UpdateSyncInstallationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSyncInstallationParamsWithTimeout creates a new UpdateSyncInstallationParams object
// with the ability to set a timeout on a request.
func NewUpdateSyncInstallationParamsWithTimeout(timeout time.Duration) *UpdateSyncInstallationParams {
	return &UpdateSyncInstallationParams{
		timeout: timeout,
	}
}

// NewUpdateSyncInstallationParamsWithContext creates a new UpdateSyncInstallationParams object
// with the ability to set a context for a request.
func NewUpdateSyncInstallationParamsWithContext(ctx context.Context) *UpdateSyncInstallationParams {
	return &UpdateSyncInstallationParams{
		Context: ctx,
	}
}

// NewUpdateSyncInstallationParamsWithHTTPClient creates a new UpdateSyncInstallationParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSyncInstallationParamsWithHTTPClient(client *http.Client) *UpdateSyncInstallationParams {
	return &UpdateSyncInstallationParams{
		HTTPClient: client,
	}
}

/*
UpdateSyncInstallationParams contains all the parameters to send to the API endpoint

	for the update sync installation operation.

	Typically these are written to a http.Request.
*/
type UpdateSyncInstallationParams struct {

	// Body.
	Body *models.SecretServiceUpdateSyncInstallationBody

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update sync installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSyncInstallationParams) WithDefaults() *UpdateSyncInstallationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update sync installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSyncInstallationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update sync installation params
func (o *UpdateSyncInstallationParams) WithTimeout(timeout time.Duration) *UpdateSyncInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update sync installation params
func (o *UpdateSyncInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update sync installation params
func (o *UpdateSyncInstallationParams) WithContext(ctx context.Context) *UpdateSyncInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update sync installation params
func (o *UpdateSyncInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update sync installation params
func (o *UpdateSyncInstallationParams) WithHTTPClient(client *http.Client) *UpdateSyncInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update sync installation params
func (o *UpdateSyncInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update sync installation params
func (o *UpdateSyncInstallationParams) WithBody(body *models.SecretServiceUpdateSyncInstallationBody) *UpdateSyncInstallationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update sync installation params
func (o *UpdateSyncInstallationParams) SetBody(body *models.SecretServiceUpdateSyncInstallationBody) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the update sync installation params
func (o *UpdateSyncInstallationParams) WithOrganizationID(organizationID string) *UpdateSyncInstallationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the update sync installation params
func (o *UpdateSyncInstallationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the update sync installation params
func (o *UpdateSyncInstallationParams) WithProjectID(projectID string) *UpdateSyncInstallationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the update sync installation params
func (o *UpdateSyncInstallationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSyncInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
