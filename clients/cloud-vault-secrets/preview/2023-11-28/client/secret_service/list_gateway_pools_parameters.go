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

// NewListGatewayPoolsParams creates a new ListGatewayPoolsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListGatewayPoolsParams() *ListGatewayPoolsParams {
	return &ListGatewayPoolsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListGatewayPoolsParamsWithTimeout creates a new ListGatewayPoolsParams object
// with the ability to set a timeout on a request.
func NewListGatewayPoolsParamsWithTimeout(timeout time.Duration) *ListGatewayPoolsParams {
	return &ListGatewayPoolsParams{
		timeout: timeout,
	}
}

// NewListGatewayPoolsParamsWithContext creates a new ListGatewayPoolsParams object
// with the ability to set a context for a request.
func NewListGatewayPoolsParamsWithContext(ctx context.Context) *ListGatewayPoolsParams {
	return &ListGatewayPoolsParams{
		Context: ctx,
	}
}

// NewListGatewayPoolsParamsWithHTTPClient creates a new ListGatewayPoolsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListGatewayPoolsParamsWithHTTPClient(client *http.Client) *ListGatewayPoolsParams {
	return &ListGatewayPoolsParams{
		HTTPClient: client,
	}
}

/*
ListGatewayPoolsParams contains all the parameters to send to the API endpoint

	for the list gateway pools operation.

	Typically these are written to a http.Request.
*/
type ListGatewayPoolsParams struct {

	// OrganizationID.
	OrganizationID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list gateway pools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGatewayPoolsParams) WithDefaults() *ListGatewayPoolsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list gateway pools params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListGatewayPoolsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list gateway pools params
func (o *ListGatewayPoolsParams) WithTimeout(timeout time.Duration) *ListGatewayPoolsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list gateway pools params
func (o *ListGatewayPoolsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list gateway pools params
func (o *ListGatewayPoolsParams) WithContext(ctx context.Context) *ListGatewayPoolsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list gateway pools params
func (o *ListGatewayPoolsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list gateway pools params
func (o *ListGatewayPoolsParams) WithHTTPClient(client *http.Client) *ListGatewayPoolsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list gateway pools params
func (o *ListGatewayPoolsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the list gateway pools params
func (o *ListGatewayPoolsParams) WithOrganizationID(organizationID string) *ListGatewayPoolsParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the list gateway pools params
func (o *ListGatewayPoolsParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WithProjectID adds the projectID to the list gateway pools params
func (o *ListGatewayPoolsParams) WithProjectID(projectID string) *ListGatewayPoolsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list gateway pools params
func (o *ListGatewayPoolsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListGatewayPoolsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
