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

// NewDeleteSyncInstallationParams creates a new DeleteSyncInstallationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteSyncInstallationParams() *DeleteSyncInstallationParams {
	return &DeleteSyncInstallationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSyncInstallationParamsWithTimeout creates a new DeleteSyncInstallationParams object
// with the ability to set a timeout on a request.
func NewDeleteSyncInstallationParamsWithTimeout(timeout time.Duration) *DeleteSyncInstallationParams {
	return &DeleteSyncInstallationParams{
		timeout: timeout,
	}
}

// NewDeleteSyncInstallationParamsWithContext creates a new DeleteSyncInstallationParams object
// with the ability to set a context for a request.
func NewDeleteSyncInstallationParamsWithContext(ctx context.Context) *DeleteSyncInstallationParams {
	return &DeleteSyncInstallationParams{
		Context: ctx,
	}
}

// NewDeleteSyncInstallationParamsWithHTTPClient creates a new DeleteSyncInstallationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteSyncInstallationParamsWithHTTPClient(client *http.Client) *DeleteSyncInstallationParams {
	return &DeleteSyncInstallationParams{
		HTTPClient: client,
	}
}

/*
DeleteSyncInstallationParams contains all the parameters to send to the API endpoint

	for the delete sync installation operation.

	Typically these are written to a http.Request.
*/
type DeleteSyncInstallationParams struct {

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

// WithDefaults hydrates default values in the delete sync installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSyncInstallationParams) WithDefaults() *DeleteSyncInstallationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete sync installation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteSyncInstallationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete sync installation params
func (o *DeleteSyncInstallationParams) WithTimeout(timeout time.Duration) *DeleteSyncInstallationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete sync installation params
func (o *DeleteSyncInstallationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete sync installation params
func (o *DeleteSyncInstallationParams) WithContext(ctx context.Context) *DeleteSyncInstallationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete sync installation params
func (o *DeleteSyncInstallationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete sync installation params
func (o *DeleteSyncInstallationParams) WithHTTPClient(client *http.Client) *DeleteSyncInstallationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete sync installation params
func (o *DeleteSyncInstallationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete sync installation params
func (o *DeleteSyncInstallationParams) WithName(name string) *DeleteSyncInstallationParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete sync installation params
func (o *DeleteSyncInstallationParams) SetName(name string) {
	o.Name = name
}

// WithOrganizationID adds the organizationID to the delete sync installation params
func (o *DeleteSyncInstallationParams) WithOrganizationID(organizationID string) *DeleteSyncInstallationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the delete sync installation params
func (o *DeleteSyncInstallationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the delete sync installation params
func (o *DeleteSyncInstallationParams) WithProjectID(projectID string) *DeleteSyncInstallationParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete sync installation params
func (o *DeleteSyncInstallationParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSyncInstallationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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