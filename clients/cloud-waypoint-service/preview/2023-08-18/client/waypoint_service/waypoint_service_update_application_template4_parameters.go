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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-waypoint-service/preview/2023-08-18/models"
)

// NewWaypointServiceUpdateApplicationTemplate4Params creates a new WaypointServiceUpdateApplicationTemplate4Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateApplicationTemplate4Params() *WaypointServiceUpdateApplicationTemplate4Params {
	return &WaypointServiceUpdateApplicationTemplate4Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateApplicationTemplate4ParamsWithTimeout creates a new WaypointServiceUpdateApplicationTemplate4Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateApplicationTemplate4ParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationTemplate4Params {
	return &WaypointServiceUpdateApplicationTemplate4Params{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateApplicationTemplate4ParamsWithContext creates a new WaypointServiceUpdateApplicationTemplate4Params object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateApplicationTemplate4ParamsWithContext(ctx context.Context) *WaypointServiceUpdateApplicationTemplate4Params {
	return &WaypointServiceUpdateApplicationTemplate4Params{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateApplicationTemplate4ParamsWithHTTPClient creates a new WaypointServiceUpdateApplicationTemplate4Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateApplicationTemplate4ParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationTemplate4Params {
	return &WaypointServiceUpdateApplicationTemplate4Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateApplicationTemplate4Params contains all the parameters to send to the API endpoint

	for the waypoint service update application template4 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateApplicationTemplate4Params struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationTemplateBody

	/* ExistingApplicationTemplateName.

	   Name of the ApplicationTemplate
	*/
	ExistingApplicationTemplateName string

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update application template4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationTemplate4Params) WithDefaults() *WaypointServiceUpdateApplicationTemplate4Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update application template4 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationTemplate4Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) WithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationTemplate4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) WithContext(ctx context.Context) *WaypointServiceUpdateApplicationTemplate4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) WithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationTemplate4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) WithBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationTemplateBody) *WaypointServiceUpdateApplicationTemplate4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) SetBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationTemplateBody) {
	o.Body = body
}

// WithExistingApplicationTemplateName adds the existingApplicationTemplateName to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) WithExistingApplicationTemplateName(existingApplicationTemplateName string) *WaypointServiceUpdateApplicationTemplate4Params {
	o.SetExistingApplicationTemplateName(existingApplicationTemplateName)
	return o
}

// SetExistingApplicationTemplateName adds the existingApplicationTemplateName to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) SetExistingApplicationTemplateName(existingApplicationTemplateName string) {
	o.ExistingApplicationTemplateName = existingApplicationTemplateName
}

// WithNamespaceID adds the namespaceID to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) WithNamespaceID(namespaceID string) *WaypointServiceUpdateApplicationTemplate4Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service update application template4 params
func (o *WaypointServiceUpdateApplicationTemplate4Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateApplicationTemplate4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param existing_application_template.name
	if err := r.SetPathParam("existing_application_template.name", o.ExistingApplicationTemplateName); err != nil {
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
