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

// NewWaypointServiceCreateAddOnParams creates a new WaypointServiceCreateAddOnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceCreateAddOnParams() *WaypointServiceCreateAddOnParams {
	return &WaypointServiceCreateAddOnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceCreateAddOnParamsWithTimeout creates a new WaypointServiceCreateAddOnParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceCreateAddOnParamsWithTimeout(timeout time.Duration) *WaypointServiceCreateAddOnParams {
	return &WaypointServiceCreateAddOnParams{
		timeout: timeout,
	}
}

// NewWaypointServiceCreateAddOnParamsWithContext creates a new WaypointServiceCreateAddOnParams object
// with the ability to set a context for a request.
func NewWaypointServiceCreateAddOnParamsWithContext(ctx context.Context) *WaypointServiceCreateAddOnParams {
	return &WaypointServiceCreateAddOnParams{
		Context: ctx,
	}
}

// NewWaypointServiceCreateAddOnParamsWithHTTPClient creates a new WaypointServiceCreateAddOnParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceCreateAddOnParamsWithHTTPClient(client *http.Client) *WaypointServiceCreateAddOnParams {
	return &WaypointServiceCreateAddOnParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceCreateAddOnParams contains all the parameters to send to the API endpoint

	for the waypoint service create add on operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceCreateAddOnParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceCreateAddOnBody

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service create add on params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateAddOnParams) WithDefaults() *WaypointServiceCreateAddOnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service create add on params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceCreateAddOnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithTimeout(timeout time.Duration) *WaypointServiceCreateAddOnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithContext(ctx context.Context) *WaypointServiceCreateAddOnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithHTTPClient(client *http.Client) *WaypointServiceCreateAddOnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithBody(body *models.HashicorpCloudWaypointWaypointServiceCreateAddOnBody) *WaypointServiceCreateAddOnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetBody(body *models.HashicorpCloudWaypointWaypointServiceCreateAddOnBody) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) WithNamespaceID(namespaceID string) *WaypointServiceCreateAddOnParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service create add on params
func (o *WaypointServiceCreateAddOnParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceCreateAddOnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
