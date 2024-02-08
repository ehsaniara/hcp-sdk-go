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

// NewWaypointServiceUIGetApplicationBundle2Params creates a new WaypointServiceUIGetApplicationBundle2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUIGetApplicationBundle2Params() *WaypointServiceUIGetApplicationBundle2Params {
	return &WaypointServiceUIGetApplicationBundle2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUIGetApplicationBundle2ParamsWithTimeout creates a new WaypointServiceUIGetApplicationBundle2Params object
// with the ability to set a timeout on a request.
func NewWaypointServiceUIGetApplicationBundle2ParamsWithTimeout(timeout time.Duration) *WaypointServiceUIGetApplicationBundle2Params {
	return &WaypointServiceUIGetApplicationBundle2Params{
		timeout: timeout,
	}
}

// NewWaypointServiceUIGetApplicationBundle2ParamsWithContext creates a new WaypointServiceUIGetApplicationBundle2Params object
// with the ability to set a context for a request.
func NewWaypointServiceUIGetApplicationBundle2ParamsWithContext(ctx context.Context) *WaypointServiceUIGetApplicationBundle2Params {
	return &WaypointServiceUIGetApplicationBundle2Params{
		Context: ctx,
	}
}

// NewWaypointServiceUIGetApplicationBundle2ParamsWithHTTPClient creates a new WaypointServiceUIGetApplicationBundle2Params object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUIGetApplicationBundle2ParamsWithHTTPClient(client *http.Client) *WaypointServiceUIGetApplicationBundle2Params {
	return &WaypointServiceUIGetApplicationBundle2Params{
		HTTPClient: client,
	}
}

/*
WaypointServiceUIGetApplicationBundle2Params contains all the parameters to send to the API endpoint

	for the waypoint service UI get application bundle2 operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUIGetApplicationBundle2Params struct {

	// ApplicationID.
	ApplicationID *string

	// ApplicationName.
	ApplicationName string

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service UI get application bundle2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUIGetApplicationBundle2Params) WithDefaults() *WaypointServiceUIGetApplicationBundle2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service UI get application bundle2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUIGetApplicationBundle2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) WithTimeout(timeout time.Duration) *WaypointServiceUIGetApplicationBundle2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) WithContext(ctx context.Context) *WaypointServiceUIGetApplicationBundle2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) WithHTTPClient(client *http.Client) *WaypointServiceUIGetApplicationBundle2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) WithApplicationID(applicationID *string) *WaypointServiceUIGetApplicationBundle2Params {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) SetApplicationID(applicationID *string) {
	o.ApplicationID = applicationID
}

// WithApplicationName adds the applicationName to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) WithApplicationName(applicationName string) *WaypointServiceUIGetApplicationBundle2Params {
	o.SetApplicationName(applicationName)
	return o
}

// SetApplicationName adds the applicationName to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) SetApplicationName(applicationName string) {
	o.ApplicationName = applicationName
}

// WithNamespaceID adds the namespaceID to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) WithNamespaceID(namespaceID string) *WaypointServiceUIGetApplicationBundle2Params {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service UI get application bundle2 params
func (o *WaypointServiceUIGetApplicationBundle2Params) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUIGetApplicationBundle2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationID != nil {

		// query param application.id
		var qrApplicationID string

		if o.ApplicationID != nil {
			qrApplicationID = *o.ApplicationID
		}
		qApplicationID := qrApplicationID
		if qApplicationID != "" {

			if err := r.SetQueryParam("application.id", qApplicationID); err != nil {
				return err
			}
		}
	}

	// path param application.name
	if err := r.SetPathParam("application.name", o.ApplicationName); err != nil {
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
