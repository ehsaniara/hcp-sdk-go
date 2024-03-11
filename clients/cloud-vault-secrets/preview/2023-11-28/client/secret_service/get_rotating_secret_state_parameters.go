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

// NewGetRotatingSecretStateParams creates a new GetRotatingSecretStateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRotatingSecretStateParams() *GetRotatingSecretStateParams {
	return &GetRotatingSecretStateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRotatingSecretStateParamsWithTimeout creates a new GetRotatingSecretStateParams object
// with the ability to set a timeout on a request.
func NewGetRotatingSecretStateParamsWithTimeout(timeout time.Duration) *GetRotatingSecretStateParams {
	return &GetRotatingSecretStateParams{
		timeout: timeout,
	}
}

// NewGetRotatingSecretStateParamsWithContext creates a new GetRotatingSecretStateParams object
// with the ability to set a context for a request.
func NewGetRotatingSecretStateParamsWithContext(ctx context.Context) *GetRotatingSecretStateParams {
	return &GetRotatingSecretStateParams{
		Context: ctx,
	}
}

// NewGetRotatingSecretStateParamsWithHTTPClient creates a new GetRotatingSecretStateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRotatingSecretStateParamsWithHTTPClient(client *http.Client) *GetRotatingSecretStateParams {
	return &GetRotatingSecretStateParams{
		HTTPClient: client,
	}
}

/*
GetRotatingSecretStateParams contains all the parameters to send to the API endpoint

	for the get rotating secret state operation.

	Typically these are written to a http.Request.
*/
type GetRotatingSecretStateParams struct {

	// AppName.
	AppName string

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	// SecretName.
	SecretName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get rotating secret state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRotatingSecretStateParams) WithDefaults() *GetRotatingSecretStateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get rotating secret state params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRotatingSecretStateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get rotating secret state params
func (o *GetRotatingSecretStateParams) WithTimeout(timeout time.Duration) *GetRotatingSecretStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get rotating secret state params
func (o *GetRotatingSecretStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get rotating secret state params
func (o *GetRotatingSecretStateParams) WithContext(ctx context.Context) *GetRotatingSecretStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get rotating secret state params
func (o *GetRotatingSecretStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get rotating secret state params
func (o *GetRotatingSecretStateParams) WithHTTPClient(client *http.Client) *GetRotatingSecretStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get rotating secret state params
func (o *GetRotatingSecretStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the get rotating secret state params
func (o *GetRotatingSecretStateParams) WithAppName(appName string) *GetRotatingSecretStateParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the get rotating secret state params
func (o *GetRotatingSecretStateParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithOrganizationID adds the organizationID to the get rotating secret state params
func (o *GetRotatingSecretStateParams) WithOrganizationID(organizationID string) *GetRotatingSecretStateParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get rotating secret state params
func (o *GetRotatingSecretStateParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the get rotating secret state params
func (o *GetRotatingSecretStateParams) WithProjectID(projectID string) *GetRotatingSecretStateParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get rotating secret state params
func (o *GetRotatingSecretStateParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithSecretName adds the secretName to the get rotating secret state params
func (o *GetRotatingSecretStateParams) WithSecretName(secretName string) *GetRotatingSecretStateParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the get rotating secret state params
func (o *GetRotatingSecretStateParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WriteToRequest writes these params to a swagger request
func (o *GetRotatingSecretStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param secret_name
	if err := r.SetPathParam("secret_name", o.SecretName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
