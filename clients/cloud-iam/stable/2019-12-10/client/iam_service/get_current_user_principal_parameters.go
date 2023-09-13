// Code generated by go-swagger; DO NOT EDIT.

package iam_service

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
	"github.com/go-openapi/swag"
)

// NewGetCurrentUserPrincipalParams creates a new GetCurrentUserPrincipalParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetCurrentUserPrincipalParams() *GetCurrentUserPrincipalParams {
	return &GetCurrentUserPrincipalParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentUserPrincipalParamsWithTimeout creates a new GetCurrentUserPrincipalParams object
// with the ability to set a timeout on a request.
func NewGetCurrentUserPrincipalParamsWithTimeout(timeout time.Duration) *GetCurrentUserPrincipalParams {
	return &GetCurrentUserPrincipalParams{
		timeout: timeout,
	}
}

// NewGetCurrentUserPrincipalParamsWithContext creates a new GetCurrentUserPrincipalParams object
// with the ability to set a context for a request.
func NewGetCurrentUserPrincipalParamsWithContext(ctx context.Context) *GetCurrentUserPrincipalParams {
	return &GetCurrentUserPrincipalParams{
		Context: ctx,
	}
}

// NewGetCurrentUserPrincipalParamsWithHTTPClient creates a new GetCurrentUserPrincipalParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetCurrentUserPrincipalParamsWithHTTPClient(client *http.Client) *GetCurrentUserPrincipalParams {
	return &GetCurrentUserPrincipalParams{
		HTTPClient: client,
	}
}

/*
GetCurrentUserPrincipalParams contains all the parameters to send to the API endpoint

	for the get current user principal operation.

	Typically these are written to a http.Request.
*/
type GetCurrentUserPrincipalParams struct {

	/* WebPortalPreferences.

	     web_portal_preferences is a flag indicating whether user's portal preferences
	need to be included or not. Typically the HCP Web Portal JS application will
	set this option to true, as its the main consumer of this data. Other clients
	(e.g. a CLI, a TF provider) will not need to enable it.
	*/
	WebPortalPreferences *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get current user principal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCurrentUserPrincipalParams) WithDefaults() *GetCurrentUserPrincipalParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get current user principal params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetCurrentUserPrincipalParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get current user principal params
func (o *GetCurrentUserPrincipalParams) WithTimeout(timeout time.Duration) *GetCurrentUserPrincipalParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current user principal params
func (o *GetCurrentUserPrincipalParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current user principal params
func (o *GetCurrentUserPrincipalParams) WithContext(ctx context.Context) *GetCurrentUserPrincipalParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current user principal params
func (o *GetCurrentUserPrincipalParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get current user principal params
func (o *GetCurrentUserPrincipalParams) WithHTTPClient(client *http.Client) *GetCurrentUserPrincipalParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get current user principal params
func (o *GetCurrentUserPrincipalParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWebPortalPreferences adds the webPortalPreferences to the get current user principal params
func (o *GetCurrentUserPrincipalParams) WithWebPortalPreferences(webPortalPreferences *bool) *GetCurrentUserPrincipalParams {
	o.SetWebPortalPreferences(webPortalPreferences)
	return o
}

// SetWebPortalPreferences adds the webPortalPreferences to the get current user principal params
func (o *GetCurrentUserPrincipalParams) SetWebPortalPreferences(webPortalPreferences *bool) {
	o.WebPortalPreferences = webPortalPreferences
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentUserPrincipalParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.WebPortalPreferences != nil {

		// query param web_portal_preferences
		var qrWebPortalPreferences bool

		if o.WebPortalPreferences != nil {
			qrWebPortalPreferences = *o.WebPortalPreferences
		}
		qWebPortalPreferences := swag.FormatBool(qrWebPortalPreferences)
		if qWebPortalPreferences != "" {

			if err := r.SetQueryParam("web_portal_preferences", qWebPortalPreferences); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}