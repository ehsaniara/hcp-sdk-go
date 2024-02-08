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

// NewWaypointServiceUpdateTFCConfigParams creates a new WaypointServiceUpdateTFCConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateTFCConfigParams() *WaypointServiceUpdateTFCConfigParams {
	return &WaypointServiceUpdateTFCConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateTFCConfigParamsWithTimeout creates a new WaypointServiceUpdateTFCConfigParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateTFCConfigParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateTFCConfigParams {
	return &WaypointServiceUpdateTFCConfigParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateTFCConfigParamsWithContext creates a new WaypointServiceUpdateTFCConfigParams object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateTFCConfigParamsWithContext(ctx context.Context) *WaypointServiceUpdateTFCConfigParams {
	return &WaypointServiceUpdateTFCConfigParams{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateTFCConfigParamsWithHTTPClient creates a new WaypointServiceUpdateTFCConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateTFCConfigParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateTFCConfigParams {
	return &WaypointServiceUpdateTFCConfigParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateTFCConfigParams contains all the parameters to send to the API endpoint

	for the waypoint service update t f c config operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateTFCConfigParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceUpdateTFCConfigBody

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update t f c config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateTFCConfigParams) WithDefaults() *WaypointServiceUpdateTFCConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update t f c config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateTFCConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) WithTimeout(timeout time.Duration) *WaypointServiceUpdateTFCConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) WithContext(ctx context.Context) *WaypointServiceUpdateTFCConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) WithHTTPClient(client *http.Client) *WaypointServiceUpdateTFCConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) WithBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateTFCConfigBody) *WaypointServiceUpdateTFCConfigParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) SetBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateTFCConfigBody) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) WithNamespaceID(namespaceID string) *WaypointServiceUpdateTFCConfigParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service update t f c config params
func (o *WaypointServiceUpdateTFCConfigParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateTFCConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
