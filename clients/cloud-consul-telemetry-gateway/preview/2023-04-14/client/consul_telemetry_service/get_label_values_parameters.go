// Code generated by go-swagger; DO NOT EDIT.

package consul_telemetry_service

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

// NewGetLabelValuesParams creates a new GetLabelValuesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetLabelValuesParams() *GetLabelValuesParams {
	return &GetLabelValuesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetLabelValuesParamsWithTimeout creates a new GetLabelValuesParams object
// with the ability to set a timeout on a request.
func NewGetLabelValuesParamsWithTimeout(timeout time.Duration) *GetLabelValuesParams {
	return &GetLabelValuesParams{
		timeout: timeout,
	}
}

// NewGetLabelValuesParamsWithContext creates a new GetLabelValuesParams object
// with the ability to set a context for a request.
func NewGetLabelValuesParamsWithContext(ctx context.Context) *GetLabelValuesParams {
	return &GetLabelValuesParams{
		Context: ctx,
	}
}

// NewGetLabelValuesParamsWithHTTPClient creates a new GetLabelValuesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetLabelValuesParamsWithHTTPClient(client *http.Client) *GetLabelValuesParams {
	return &GetLabelValuesParams{
		HTTPClient: client,
	}
}

/*
GetLabelValuesParams contains all the parameters to send to the API endpoint

	for the get label values operation.

	Typically these are written to a http.Request.
*/
type GetLabelValuesParams struct {

	// Body.
	Body GetLabelValuesBody

	// ClusterID.
	ClusterID string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get label values params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelValuesParams) WithDefaults() *GetLabelValuesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get label values params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetLabelValuesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get label values params
func (o *GetLabelValuesParams) WithTimeout(timeout time.Duration) *GetLabelValuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get label values params
func (o *GetLabelValuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get label values params
func (o *GetLabelValuesParams) WithContext(ctx context.Context) *GetLabelValuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get label values params
func (o *GetLabelValuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get label values params
func (o *GetLabelValuesParams) WithHTTPClient(client *http.Client) *GetLabelValuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get label values params
func (o *GetLabelValuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get label values params
func (o *GetLabelValuesParams) WithBody(body GetLabelValuesBody) *GetLabelValuesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get label values params
func (o *GetLabelValuesParams) SetBody(body GetLabelValuesBody) {
	o.Body = body
}

// WithClusterID adds the clusterID to the get label values params
func (o *GetLabelValuesParams) WithClusterID(clusterID string) *GetLabelValuesParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get label values params
func (o *GetLabelValuesParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithLocationOrganizationID adds the locationOrganizationID to the get label values params
func (o *GetLabelValuesParams) WithLocationOrganizationID(locationOrganizationID string) *GetLabelValuesParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the get label values params
func (o *GetLabelValuesParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the get label values params
func (o *GetLabelValuesParams) WithLocationProjectID(locationProjectID string) *GetLabelValuesParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the get label values params
func (o *GetLabelValuesParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLabelValuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}