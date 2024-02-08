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

// NewWaypointServiceGetActionRunParams creates a new WaypointServiceGetActionRunParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWaypointServiceGetActionRunParams() *WaypointServiceGetActionRunParams {
	return &WaypointServiceGetActionRunParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWaypointServiceGetActionRunParamsWithTimeout creates a new WaypointServiceGetActionRunParams object
// with the ability to set a timeout on a request.
func NewWaypointServiceGetActionRunParamsWithTimeout(timeout time.Duration) *WaypointServiceGetActionRunParams {
	return &WaypointServiceGetActionRunParams{
		timeout: timeout,
	}
}

// NewWaypointServiceGetActionRunParamsWithContext creates a new WaypointServiceGetActionRunParams object
// with the ability to set a context for a request.
func NewWaypointServiceGetActionRunParamsWithContext(ctx context.Context) *WaypointServiceGetActionRunParams {
	return &WaypointServiceGetActionRunParams{
		Context: ctx,
	}
}

// NewWaypointServiceGetActionRunParamsWithHTTPClient creates a new WaypointServiceGetActionRunParams object
// with the ability to set a custom HTTPClient for a request.
func NewWaypointServiceGetActionRunParamsWithHTTPClient(client *http.Client) *WaypointServiceGetActionRunParams {
	return &WaypointServiceGetActionRunParams{
		HTTPClient: client,
	}
}

/*
WaypointServiceGetActionRunParams contains all the parameters to send to the API endpoint

	for the waypoint service get action run operation.

	Typically these are written to a http.Request.
*/
type WaypointServiceGetActionRunParams struct {

	/* ActionID.

	   The id of the action config being requested
	*/
	ActionID *string

	/* ActionName.

	   The name of the action being requested (will only be used if id not provided)
	*/
	ActionName *string

	// NamespaceID.
	NamespaceID string

	/* Sequence.

	   The sequence number of the action run (required)

	   Format: uint64
	*/
	Sequence *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the waypoint service get action run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetActionRunParams) WithDefaults() *WaypointServiceGetActionRunParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the waypoint service get action run params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WaypointServiceGetActionRunParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) WithTimeout(timeout time.Duration) *WaypointServiceGetActionRunParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) WithContext(ctx context.Context) *WaypointServiceGetActionRunParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) WithHTTPClient(client *http.Client) *WaypointServiceGetActionRunParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) WithActionID(actionID *string) *WaypointServiceGetActionRunParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) SetActionID(actionID *string) {
	o.ActionID = actionID
}

// WithActionName adds the actionName to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) WithActionName(actionName *string) *WaypointServiceGetActionRunParams {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) SetActionName(actionName *string) {
	o.ActionName = actionName
}

// WithNamespaceID adds the namespaceID to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) WithNamespaceID(namespaceID string) *WaypointServiceGetActionRunParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WithSequence adds the sequence to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) WithSequence(sequence *string) *WaypointServiceGetActionRunParams {
	o.SetSequence(sequence)
	return o
}

// SetSequence adds the sequence to the waypoint service get action run params
func (o *WaypointServiceGetActionRunParams) SetSequence(sequence *string) {
	o.Sequence = sequence
}

// WriteToRequest writes these params to a swagger request
func (o *WaypointServiceGetActionRunParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionID != nil {

		// query param action_id
		var qrActionID string

		if o.ActionID != nil {
			qrActionID = *o.ActionID
		}
		qActionID := qrActionID
		if qActionID != "" {

			if err := r.SetQueryParam("action_id", qActionID); err != nil {
				return err
			}
		}
	}

	if o.ActionName != nil {

		// query param action_name
		var qrActionName string

		if o.ActionName != nil {
			qrActionName = *o.ActionName
		}
		qActionName := qrActionName
		if qActionName != "" {

			if err := r.SetQueryParam("action_name", qActionName); err != nil {
				return err
			}
		}
	}

	// path param namespace.id
	if err := r.SetPathParam("namespace.id", o.NamespaceID); err != nil {
		return err
	}

	if o.Sequence != nil {

		// query param sequence
		var qrSequence string

		if o.Sequence != nil {
			qrSequence = *o.Sequence
		}
		qSequence := qrSequence
		if qSequence != "" {

			if err := r.SetQueryParam("sequence", qSequence); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
