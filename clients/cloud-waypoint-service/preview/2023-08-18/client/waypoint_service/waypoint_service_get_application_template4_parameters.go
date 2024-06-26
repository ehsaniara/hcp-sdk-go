// Code generated by go-swagger; DO NOT EDIT.

package waypoint_service

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

// NewWaypointServiceGetApplicationTemplate4Params creates a new WaypointServiceGetApplicationTemplate4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetApplicationTemplate4Params() *WaypointServiceGetApplicationTemplate4Params {
	return &WaypointServiceGetApplicationTemplate4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetApplicationTemplate4ParamsWithTimeout creates a new WaypointServiceGetApplicationTemplate4Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetApplicationTemplate4ParamsWithTimeout(timeout time.Duration) *WaypointServiceGetApplicationTemplate4Params {
	return &WaypointServiceGetApplicationTemplate4Params{
		timeout: timeout,
	}
}

// NewWaypointServiceGetApplicationTemplate4ParamsWithContext creates a new WaypointServiceGetApplicationTemplate4Params object
// with the ability to set a context for a request.
func NewWaypointServiceGetApplicationTemplate4ParamsWithContext(ctx context.Context) *WaypointServiceGetApplicationTemplate4Params {
	return &WaypointServiceGetApplicationTemplate4Params{
		Context: ctx,
	}
}

// NewWaypointServiceGetApplicationTemplate4ParamsWithHTTPClient creates a new WaypointServiceGetApplicationTemplate4Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetApplicationTemplate4ParamsWithHTTPClient(client *http.Client) *WaypointServiceGetApplicationTemplate4Params {
	return &WaypointServiceGetApplicationTemplate4Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetApplicationTemplate4Params contains all the parameters to send to the API endpoint

	for the waypoint service get application template4 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetApplicationTemplate4Params struct {

	/* ApplicationTemplateID.

	   ID of the ApplicationTemplate
	*/
	ApplicationTemplateID *string

	/* ApplicationTemplateName.

	   Name of the ApplicationTemplate
	*/
	ApplicationTemplateName string

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service get application template4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetApplicationTemplate4Params) WithDefaults() *WaypointServiceGetApplicationTemplate4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get application template4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetApplicationTemplate4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) WithTimeout(timeout time.Duration) *WaypointServiceGetApplicationTemplate4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) WithContext(ctx context.Context) *WaypointServiceGetApplicationTemplate4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) WithHTTPClient(client *http.Client) *WaypointServiceGetApplicationTemplate4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationTemplateID adds the applicationTemplateID to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) WithApplicationTemplateID(applicationTemplateID *string) *WaypointServiceGetApplicationTemplate4Params {
	o.SetApplicationTemplateID(applicationTemplateID)
	return o
}

// SetApplicationTemplateID adds the applicationTemplateId to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) SetApplicationTemplateID(applicationTemplateID *string) {
	o.ApplicationTemplateID = applicationTemplateID
}

// WithApplicationTemplateName adds the applicationTemplateName to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) WithApplicationTemplateName(applicationTemplateName string) *WaypointServiceGetApplicationTemplate4Params {
	o.SetApplicationTemplateName(applicationTemplateName)
	return o
}

// SetApplicationTemplateName adds the applicationTemplateName to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) SetApplicationTemplateName(applicationTemplateName string) {
	o.ApplicationTemplateName = applicationTemplateName
}

// WithNamespaceID adds the namespaceID to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) WithNamespaceID(namespaceID string) *WaypointServiceGetApplicationTemplate4Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get application template4 params
func (o *WaypointServiceGetApplicationTemplate4Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetApplicationTemplate4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationTemplateID != nil {

		// query param application_template.id
		var qrApplicationTemplateID string

		if o.ApplicationTemplateID != nil {
			qrApplicationTemplateID = *o.ApplicationTemplateID
		}
		qApplicationTemplateID := qrApplicationTemplateID
		if qApplicationTemplateID != "" {

			if err := r.SetQueryParam("application_template.id", qApplicationTemplateID); err != nil {
				return err
			}
		}
	}

	// path param application_template.name
	if err := r.SetPathParam("application_template.name", o.ApplicationTemplateName); err != nil {
		return err
	}

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
