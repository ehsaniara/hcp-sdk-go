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

// NewWaypointServiceGetTFModuleDetailsParams creates a new WaypointServiceGetTFModuleDetailsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetTFModuleDetailsParams() *WaypointServiceGetTFModuleDetailsParams {
	return &WaypointServiceGetTFModuleDetailsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetTFModuleDetailsParamsWithTimeout creates a new WaypointServiceGetTFModuleDetailsParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetTFModuleDetailsParamsWithTimeout(timeout time.Duration) *WaypointServiceGetTFModuleDetailsParams {
	return &WaypointServiceGetTFModuleDetailsParams{
		timeout: timeout,
	}
}

// NewWaypointServiceGetTFModuleDetailsParamsWithContext creates a new WaypointServiceGetTFModuleDetailsParams object
// with the ability to set a context for a request.
func NewWaypointServiceGetTFModuleDetailsParamsWithContext(ctx context.Context) *WaypointServiceGetTFModuleDetailsParams {
	return &WaypointServiceGetTFModuleDetailsParams{
		Context: ctx,
	}
}

// NewWaypointServiceGetTFModuleDetailsParamsWithHTTPClient creates a new WaypointServiceGetTFModuleDetailsParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetTFModuleDetailsParamsWithHTTPClient(client *http.Client) *WaypointServiceGetTFModuleDetailsParams {
	return &WaypointServiceGetTFModuleDetailsParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetTFModuleDetailsParams contains all the parameters to send to the API endpoint

	for the waypoint service get t f module details operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetTFModuleDetailsParams struct {

	// NamespaceID.
	NamespaceID string

	// Source.
	Source string

	// Version.
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service get t f module details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFModuleDetailsParams) WithDefaults() *WaypointServiceGetTFModuleDetailsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get t f module details params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetTFModuleDetailsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) WithTimeout(timeout time.Duration) *WaypointServiceGetTFModuleDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) WithContext(ctx context.Context) *WaypointServiceGetTFModuleDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) WithHTTPClient(client *http.Client) *WaypointServiceGetTFModuleDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamespaceID adds the namespaceID to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) WithNamespaceID(namespaceID string) *WaypointServiceGetTFModuleDetailsParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithSource adds the source to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) WithSource(source string) *WaypointServiceGetTFModuleDetailsParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) SetSource(source string) {
	o.Source = source
}

// WithVersion adds the version to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) WithVersion(version string) *WaypointServiceGetTFModuleDetailsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the waypoint service get t f module details params
func (o *WaypointServiceGetTFModuleDetailsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetTFModuleDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	// path param source
	if err := r.SetPathParam("source", o.Source); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}