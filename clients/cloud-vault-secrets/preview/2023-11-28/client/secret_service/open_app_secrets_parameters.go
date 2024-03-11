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

// NewOpenAppSecretsParams creates a new OpenAppSecretsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOpenAppSecretsParams() *OpenAppSecretsParams {
	return &OpenAppSecretsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOpenAppSecretsParamsWithTimeout creates a new OpenAppSecretsParams object
// with the ability to set a timeout on a request.
func NewOpenAppSecretsParamsWithTimeout(timeout time.Duration) *OpenAppSecretsParams {
	return &OpenAppSecretsParams{
		timeout: timeout,
	}
}

// NewOpenAppSecretsParamsWithContext creates a new OpenAppSecretsParams object
// with the ability to set a context for a request.
func NewOpenAppSecretsParamsWithContext(ctx context.Context) *OpenAppSecretsParams {
	return &OpenAppSecretsParams{
		Context: ctx,
	}
}

// NewOpenAppSecretsParamsWithHTTPClient creates a new OpenAppSecretsParams object
// with the ability to set a custom HTTPClient for a request.
func NewOpenAppSecretsParamsWithHTTPClient(client *http.Client) *OpenAppSecretsParams {
	return &OpenAppSecretsParams{
		HTTPClient: client,
	}
}

/*
OpenAppSecretsParams contains all the parameters to send to the API endpoint

	for the open app secrets operation.

	Typically these are written to a http.Request.
*/
type OpenAppSecretsParams struct {

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

// WithDefaults hydrates default values in the open app secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAppSecretsParams) WithDefaults() *OpenAppSecretsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the open app secrets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OpenAppSecretsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the open app secrets params
func (o *OpenAppSecretsParams) WithTimeout(timeout time.Duration) *OpenAppSecretsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the open app secrets params
func (o *OpenAppSecretsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the open app secrets params
func (o *OpenAppSecretsParams) WithContext(ctx context.Context) *OpenAppSecretsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the open app secrets params
func (o *OpenAppSecretsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the open app secrets params
func (o *OpenAppSecretsParams) WithHTTPClient(client *http.Client) *OpenAppSecretsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the open app secrets params
func (o *OpenAppSecretsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the open app secrets params
func (o *OpenAppSecretsParams) WithAppName(appName string) *OpenAppSecretsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the open app secrets params
func (o *OpenAppSecretsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithOrganizationID adds the organizationID to the open app secrets params
func (o *OpenAppSecretsParams) WithOrganizationID(organizationID string) *OpenAppSecretsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the open app secrets params
func (o *OpenAppSecretsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the open app secrets params
func (o *OpenAppSecretsParams) WithProjectID(projectID string) *OpenAppSecretsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the open app secrets params
func (o *OpenAppSecretsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *OpenAppSecretsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
