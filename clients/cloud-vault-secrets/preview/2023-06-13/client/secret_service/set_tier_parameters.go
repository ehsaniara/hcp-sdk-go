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

// NewSetTierParams creates a new SetTierParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSetTierParams() *SetTierParams {
	return &SetTierParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSetTierParamsWithTimeout creates a new SetTierParams object
// with the ability to set a timeout on a request.
func NewSetTierParamsWithTimeout(timeout time.Duration) *SetTierParams {
	return &SetTierParams{
		timeout: timeout,
	}
}

// NewSetTierParamsWithContext creates a new SetTierParams object
// with the ability to set a context for a request.
func NewSetTierParamsWithContext(ctx context.Context) *SetTierParams {
	return &SetTierParams{
		Context: ctx,
	}
}

// NewSetTierParamsWithHTTPClient creates a new SetTierParams object
// with the ability to set a custom HTTPClient for a request.
func NewSetTierParamsWithHTTPClient(client *http.Client) *SetTierParams {
	return &SetTierParams{
		HTTPClient: client,
	}
}

/*
SetTierParams contains all the parameters to send to the API endpoint

	for the set tier operation.

	Typically these are written to a http.Request.
*/
type SetTierParams struct {

	// Body.
	Body SetTierBody

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the set tier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTierParams) WithDefaults() *SetTierParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the set tier params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SetTierParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the set tier params
func (o *SetTierParams) WithTimeout(timeout time.Duration) *SetTierParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set tier params
func (o *SetTierParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set tier params
func (o *SetTierParams) WithContext(ctx context.Context) *SetTierParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set tier params
func (o *SetTierParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set tier params
func (o *SetTierParams) WithHTTPClient(client *http.Client) *SetTierParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set tier params
func (o *SetTierParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the set tier params
func (o *SetTierParams) WithBody(body SetTierBody) *SetTierParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the set tier params
func (o *SetTierParams) SetBody(body SetTierBody) {
	o.Body = body
}

// WithLocationOrganizationID adds the locationOrganizationID to the set tier params
func (o *SetTierParams) WithLocationOrganizationID(locationOrganizationID string) *SetTierParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the set tier params
func (o *SetTierParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WriteToRequest writes these params to a swagger request
func (o *SetTierParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
