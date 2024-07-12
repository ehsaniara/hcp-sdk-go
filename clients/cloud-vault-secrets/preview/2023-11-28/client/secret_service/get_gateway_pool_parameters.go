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

// NewGetGatewayPoolParams creates a new GetGatewayPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGatewayPoolParams() *GetGatewayPoolParams {
	return &GetGatewayPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGatewayPoolParamsWithTimeout creates a new GetGatewayPoolParams object
// with the ability to set a timeout on a request.
func NewGetGatewayPoolParamsWithTimeout(timeout time.Duration) *GetGatewayPoolParams {
	return &GetGatewayPoolParams{
		timeout: timeout,
	}
}

// NewGetGatewayPoolParamsWithContext creates a new GetGatewayPoolParams object
// with the ability to set a context for a request.
func NewGetGatewayPoolParamsWithContext(ctx context.Context) *GetGatewayPoolParams {
	return &GetGatewayPoolParams{
		Context: ctx,
	}
}

// NewGetGatewayPoolParamsWithHTTPClient creates a new GetGatewayPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGatewayPoolParamsWithHTTPClient(client *http.Client) *GetGatewayPoolParams {
	return &GetGatewayPoolParams{
		HTTPClient: client,
	}
}

/*
GetGatewayPoolParams contains all the parameters to send to the API endpoint

	for the get gateway pool operation.

	Typically these are written to a http.Request.
*/
type GetGatewayPoolParams struct {

	// GatewayPoolName.
	GatewayPoolName string

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gateway pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGatewayPoolParams) WithDefaults() *GetGatewayPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gateway pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGatewayPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gateway pool params
func (o *GetGatewayPoolParams) WithTimeout(timeout time.Duration) *GetGatewayPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gateway pool params
func (o *GetGatewayPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gateway pool params
func (o *GetGatewayPoolParams) WithContext(ctx context.Context) *GetGatewayPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gateway pool params
func (o *GetGatewayPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gateway pool params
func (o *GetGatewayPoolParams) WithHTTPClient(client *http.Client) *GetGatewayPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gateway pool params
func (o *GetGatewayPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGatewayPoolName adds the gatewayPoolName to the get gateway pool params
func (o *GetGatewayPoolParams) WithGatewayPoolName(gatewayPoolName string) *GetGatewayPoolParams {
	o.SetGatewayPoolName(gatewayPoolName)
	return o
}

// SetGatewayPoolName adds the gatewayPoolName to the get gateway pool params
func (o *GetGatewayPoolParams) SetGatewayPoolName(gatewayPoolName string) {
	o.GatewayPoolName = gatewayPoolName
}

// WithOrganizationID adds the organizationID to the get gateway pool params
func (o *GetGatewayPoolParams) WithOrganizationID(organizationID string) *GetGatewayPoolParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get gateway pool params
func (o *GetGatewayPoolParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the get gateway pool params
func (o *GetGatewayPoolParams) WithProjectID(projectID string) *GetGatewayPoolParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get gateway pool params
func (o *GetGatewayPoolParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGatewayPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param gateway_pool_name
	if err := r.SetPathParam("gateway_pool_name", o.GatewayPoolName); err != nil {
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
