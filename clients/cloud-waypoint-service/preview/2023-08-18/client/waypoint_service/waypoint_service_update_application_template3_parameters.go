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

// NewWaypointServiceUpdateApplicationTemplate3Params creates a new WaypointServiceUpdateApplicationTemplate3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateApplicationTemplate3Params() *WaypointServiceUpdateApplicationTemplate3Params {
	return &WaypointServiceUpdateApplicationTemplate3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateApplicationTemplate3ParamsWithTimeout creates a new WaypointServiceUpdateApplicationTemplate3Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateApplicationTemplate3ParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationTemplate3Params {
	return &WaypointServiceUpdateApplicationTemplate3Params{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateApplicationTemplate3ParamsWithContext creates a new WaypointServiceUpdateApplicationTemplate3Params object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateApplicationTemplate3ParamsWithContext(ctx context.Context) *WaypointServiceUpdateApplicationTemplate3Params {
	return &WaypointServiceUpdateApplicationTemplate3Params{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateApplicationTemplate3ParamsWithHTTPClient creates a new WaypointServiceUpdateApplicationTemplate3Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateApplicationTemplate3ParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationTemplate3Params {
	return &WaypointServiceUpdateApplicationTemplate3Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateApplicationTemplate3Params contains all the parameters to send to the API endpoint

	for the waypoint service update application template3 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateApplicationTemplate3Params struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationTemplateBody

	/* ExistingApplicationTemplateID.

	   ID of the ApplicationTemplate
	*/
	ExistingApplicationTemplateID string

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update application template3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationTemplate3Params) WithDefaults() *WaypointServiceUpdateApplicationTemplate3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update application template3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateApplicationTemplate3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) WithTimeout(timeout time.Duration) *WaypointServiceUpdateApplicationTemplate3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) WithContext(ctx context.Context) *WaypointServiceUpdateApplicationTemplate3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) WithHTTPClient(client *http.Client) *WaypointServiceUpdateApplicationTemplate3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) WithBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationTemplateBody) *WaypointServiceUpdateApplicationTemplate3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) SetBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateApplicationTemplateBody) {
	o.Body = body
}

// WithExistingApplicationTemplateID adds the existingApplicationTemplateID to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) WithExistingApplicationTemplateID(existingApplicationTemplateID string) *WaypointServiceUpdateApplicationTemplate3Params {
	o.SetExistingApplicationTemplateID(existingApplicationTemplateID)
	return o
}

// SetExistingApplicationTemplateID adds the existingApplicationTemplateId to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) SetExistingApplicationTemplateID(existingApplicationTemplateID string) {
	o.ExistingApplicationTemplateID = existingApplicationTemplateID
}

// WithNamespaceID adds the namespaceID to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) WithNamespaceID(namespaceID string) *WaypointServiceUpdateApplicationTemplate3Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service update application template3 params
func (o *WaypointServiceUpdateApplicationTemplate3Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateApplicationTemplate3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param existing_application_template.id
	if err := r.SetPathParam("existing_application_template.id", o.ExistingApplicationTemplateID); err != nil {
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
