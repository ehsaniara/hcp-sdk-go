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

// NewWaypointServiceGetAddOnInputVariablesParams creates a new WaypointServiceGetAddOnInputVariablesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetAddOnInputVariablesParams() *WaypointServiceGetAddOnInputVariablesParams {
	return &WaypointServiceGetAddOnInputVariablesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetAddOnInputVariablesParamsWithTimeout creates a new WaypointServiceGetAddOnInputVariablesParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetAddOnInputVariablesParamsWithTimeout(timeout time.Duration) *WaypointServiceGetAddOnInputVariablesParams {
	return &WaypointServiceGetAddOnInputVariablesParams{
		timeout: timeout,
	}
}

// NewWaypointServiceGetAddOnInputVariablesParamsWithContext creates a new WaypointServiceGetAddOnInputVariablesParams object
// with the ability to set a context for a request.
func NewWaypointServiceGetAddOnInputVariablesParamsWithContext(ctx context.Context) *WaypointServiceGetAddOnInputVariablesParams {
	return &WaypointServiceGetAddOnInputVariablesParams{
		Context: ctx,
	}
}

// NewWaypointServiceGetAddOnInputVariablesParamsWithHTTPClient creates a new WaypointServiceGetAddOnInputVariablesParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetAddOnInputVariablesParamsWithHTTPClient(client *http.Client) *WaypointServiceGetAddOnInputVariablesParams {
	return &WaypointServiceGetAddOnInputVariablesParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetAddOnInputVariablesParams contains all the parameters to send to the API endpoint

	for the waypoint service get add on input variables operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetAddOnInputVariablesParams struct {

	// AddOnID.
	AddOnID string

	// AddOnName.
	AddOnName *string

	// NamespaceID.
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service get add on input variables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetAddOnInputVariablesParams) WithDefaults() *WaypointServiceGetAddOnInputVariablesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get add on input variables params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetAddOnInputVariablesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) WithTimeout(timeout time.Duration) *WaypointServiceGetAddOnInputVariablesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) WithContext(ctx context.Context) *WaypointServiceGetAddOnInputVariablesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) WithHTTPClient(client *http.Client) *WaypointServiceGetAddOnInputVariablesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddOnID adds the addOnID to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) WithAddOnID(addOnID string) *WaypointServiceGetAddOnInputVariablesParams {
	o.SetAddOnID(addOnID)
	return o
}

// SetAddOnID adds the addOnId to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) SetAddOnID(addOnID string) {
	o.AddOnID = addOnID
}

// WithAddOnName adds the addOnName to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) WithAddOnName(addOnName *string) *WaypointServiceGetAddOnInputVariablesParams {
	o.SetAddOnName(addOnName)
	return o
}

// SetAddOnName adds the addOnName to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) SetAddOnName(addOnName *string) {
	o.AddOnName = addOnName
}

// WithNamespaceID adds the namespaceID to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) WithNamespaceID(namespaceID string) *WaypointServiceGetAddOnInputVariablesParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get add on input variables params
func (o *WaypointServiceGetAddOnInputVariablesParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetAddOnInputVariablesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param add_on.id
	if err := r.SetPathParam("add_on.id", o.AddOnID); err != nil {
		return err
	}

	if o.AddOnName != nil {

		// query param add_on.name
		var qrAddOnName string

		if o.AddOnName != nil {
			qrAddOnName = *o.AddOnName
		}
		qAddOnName := qrAddOnName
		if qAddOnName != "" {

			if err := r.SetQueryParam("add_on.name", qAddOnName); err != nil {
				return err
			}
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
