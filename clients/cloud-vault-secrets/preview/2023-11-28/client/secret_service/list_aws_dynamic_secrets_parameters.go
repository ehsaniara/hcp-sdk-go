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

// NewListAwsDynamicSecretsParams creates a new ListAwsDynamicSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListAwsDynamicSecretsParams() *ListAwsDynamicSecretsParams {
	return &ListAwsDynamicSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListAwsDynamicSecretsParamsWithTimeout creates a new ListAwsDynamicSecretsParams object
// with the ability to set a timeout on a request.
func NewListAwsDynamicSecretsParamsWithTimeout(timeout time.Duration) *ListAwsDynamicSecretsParams {
	return &ListAwsDynamicSecretsParams{
		timeout: timeout,
	}
}

// NewListAwsDynamicSecretsParamsWithContext creates a new ListAwsDynamicSecretsParams object
// with the ability to set a context for a request.
func NewListAwsDynamicSecretsParamsWithContext(ctx context.Context) *ListAwsDynamicSecretsParams {
	return &ListAwsDynamicSecretsParams{
		Context: ctx,
	}
}

// NewListAwsDynamicSecretsParamsWithHTTPClient creates a new ListAwsDynamicSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListAwsDynamicSecretsParamsWithHTTPClient(client *http.Client) *ListAwsDynamicSecretsParams {
	return &ListAwsDynamicSecretsParams{
		HTTPClient: client,
	}
}

/*
ListAwsDynamicSecretsParams contains all the parameters to send to the API endpoint

	for the list aws dynamic secrets operation.

	Typically these are written to a http.Request.
*/
type ListAwsDynamicSecretsParams struct {

	// AppName.
	AppName string

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list aws dynamic secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAwsDynamicSecretsParams) WithDefaults() *ListAwsDynamicSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list aws dynamic secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListAwsDynamicSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) WithTimeout(timeout time.Duration) *ListAwsDynamicSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) WithContext(ctx context.Context) *ListAwsDynamicSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) WithHTTPClient(client *http.Client) *ListAwsDynamicSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) WithAppName(appName string) *ListAwsDynamicSecretsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithOrganizationID adds the organizationID to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) WithOrganizationID(organizationID string) *ListAwsDynamicSecretsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) WithProjectID(projectID string) *ListAwsDynamicSecretsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list aws dynamic secrets params
func (o *ListAwsDynamicSecretsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAwsDynamicSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param app_name
	if err := r.SetPathParam("app_name", o.AppName); err != nil {
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
