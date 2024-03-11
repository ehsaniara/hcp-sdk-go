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

// NewListOpenAppSecretVersionsParams creates a new ListOpenAppSecretVersionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListOpenAppSecretVersionsParams() *ListOpenAppSecretVersionsParams {
	return &ListOpenAppSecretVersionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListOpenAppSecretVersionsParamsWithTimeout creates a new ListOpenAppSecretVersionsParams object
// with the ability to set a timeout on a request.
func NewListOpenAppSecretVersionsParamsWithTimeout(timeout time.Duration) *ListOpenAppSecretVersionsParams {
	return &ListOpenAppSecretVersionsParams{
		timeout: timeout,
	}
}

// NewListOpenAppSecretVersionsParamsWithContext creates a new ListOpenAppSecretVersionsParams object
// with the ability to set a context for a request.
func NewListOpenAppSecretVersionsParamsWithContext(ctx context.Context) *ListOpenAppSecretVersionsParams {
	return &ListOpenAppSecretVersionsParams{
		Context: ctx,
	}
}

// NewListOpenAppSecretVersionsParamsWithHTTPClient creates a new ListOpenAppSecretVersionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListOpenAppSecretVersionsParamsWithHTTPClient(client *http.Client) *ListOpenAppSecretVersionsParams {
	return &ListOpenAppSecretVersionsParams{
		HTTPClient: client,
	}
}

/*
ListOpenAppSecretVersionsParams contains all the parameters to send to the API endpoint

	for the list open app secret versions operation.

	Typically these are written to a http.Request.
*/
type ListOpenAppSecretVersionsParams struct {

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

// WithDefaults hydrates default values in the list open app secret versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenAppSecretVersionsParams) WithDefaults() *ListOpenAppSecretVersionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list open app secret versions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListOpenAppSecretVersionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithTimeout(timeout time.Duration) *ListOpenAppSecretVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithContext(ctx context.Context) *ListOpenAppSecretVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithHTTPClient(client *http.Client) *ListOpenAppSecretVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAppName adds the appName to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithAppName(appName string) *ListOpenAppSecretVersionsParams {
	o.SetAppName(appName)
	return o
}

// SetAppName adds the appName to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetAppName(appName string) {
	o.AppName = appName
}

// WithOrganizationID adds the organizationID to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithOrganizationID(organizationID string) *ListOpenAppSecretVersionsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithProjectID(projectID string) *ListOpenAppSecretVersionsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithSecretName adds the secretName to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) WithSecretName(secretName string) *ListOpenAppSecretVersionsParams {
	o.SetSecretName(secretName)
	return o
}

// SetSecretName adds the secretName to the list open app secret versions params
func (o *ListOpenAppSecretVersionsParams) SetSecretName(secretName string) {
	o.SecretName = secretName
}

// WriteToRequest writes these params to a swagger request
func (o *ListOpenAppSecretVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
