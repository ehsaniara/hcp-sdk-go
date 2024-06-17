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

// NewWaypointServiceUpdateVariableParams creates a new WaypointServiceUpdateVariableParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceUpdateVariableParams() *WaypointServiceUpdateVariableParams {
	return &WaypointServiceUpdateVariableParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceUpdateVariableParamsWithTimeout creates a new WaypointServiceUpdateVariableParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceUpdateVariableParamsWithTimeout(timeout time.Duration) *WaypointServiceUpdateVariableParams {
	return &WaypointServiceUpdateVariableParams{
		timeout: timeout,
	}
}

// NewWaypointServiceUpdateVariableParamsWithContext creates a new WaypointServiceUpdateVariableParams object
// with the ability to set a context for a request.
func NewWaypointServiceUpdateVariableParamsWithContext(ctx context.Context) *WaypointServiceUpdateVariableParams {
	return &WaypointServiceUpdateVariableParams{
		Context: ctx,
	}
}

// NewWaypointServiceUpdateVariableParamsWithHTTPClient creates a new WaypointServiceUpdateVariableParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceUpdateVariableParamsWithHTTPClient(client *http.Client) *WaypointServiceUpdateVariableParams {
	return &WaypointServiceUpdateVariableParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceUpdateVariableParams contains all the parameters to send to the API endpoint

	for the waypoint service update variable operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceUpdateVariableParams struct {

	// Body.
	Body *models.HashicorpCloudWaypointWaypointServiceUpdateVariableBody

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service update variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateVariableParams) WithDefaults() *WaypointServiceUpdateVariableParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service update variable params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceUpdateVariableParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) WithTimeout(timeout time.Duration) *WaypointServiceUpdateVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) WithContext(ctx context.Context) *WaypointServiceUpdateVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) WithHTTPClient(client *http.Client) *WaypointServiceUpdateVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) WithBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateVariableBody) *WaypointServiceUpdateVariableParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) SetBody(body *models.HashicorpCloudWaypointWaypointServiceUpdateVariableBody) {
	o.Body = body
}

// WithNamespaceID adds the namespaceID to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) WithNamespaceID(namespaceID string) *WaypointServiceUpdateVariableParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service update variable params
func (o *WaypointServiceUpdateVariableParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceUpdateVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
