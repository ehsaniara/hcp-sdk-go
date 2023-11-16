// Code generated by go-swagger; DO NOT EDIT.

package action

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

// NewActionListActionRunsParams creates a new ActionListActionRunsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActionListActionRunsParams() *ActionListActionRunsParams {
	return &ActionListActionRunsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewActionListActionRunsParamsWithTimeout creates a new ActionListActionRunsParams object
// with the ability to set a timeout on a request.
func NewActionListActionRunsParamsWithTimeout(timeout time.Duration) *ActionListActionRunsParams {
	return &ActionListActionRunsParams{
		timeout: timeout,
	}
}

// NewActionListActionRunsParamsWithContext creates a new ActionListActionRunsParams object
// with the ability to set a context for a request.
func NewActionListActionRunsParamsWithContext(ctx context.Context) *ActionListActionRunsParams {
	return &ActionListActionRunsParams{
		Context: ctx,
	}
}

// NewActionListActionRunsParamsWithHTTPClient creates a new ActionListActionRunsParams object
// with the ability to set a custom HTTPClient for a request.
func NewActionListActionRunsParamsWithHTTPClient(client *http.Client) *ActionListActionRunsParams {
	return &ActionListActionRunsParams{
		HTTPClient: client,
	}
}

/*
ActionListActionRunsParams contains all the parameters to send to the API endpoint

	for the action list action runs operation.

	Typically these are written to a http.Request.
*/
type ActionListActionRunsParams struct {

	/* ActionID.

	   The id of the action config being listed.
	*/
	ActionID *string

	/* ActionName.

	   The name of the action being listed (will only be used if id not provided).
	*/
	ActionName *string

	/* NamespaceID.

	   The namespace the action to be listed in
	*/
	NamespaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the action list action runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionListActionRunsParams) WithDefaults() *ActionListActionRunsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the action list action runs params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActionListActionRunsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the action list action runs params
func (o *ActionListActionRunsParams) WithTimeout(timeout time.Duration) *ActionListActionRunsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the action list action runs params
func (o *ActionListActionRunsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the action list action runs params
func (o *ActionListActionRunsParams) WithContext(ctx context.Context) *ActionListActionRunsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the action list action runs params
func (o *ActionListActionRunsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the action list action runs params
func (o *ActionListActionRunsParams) WithHTTPClient(client *http.Client) *ActionListActionRunsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the action list action runs params
func (o *ActionListActionRunsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionID adds the actionID to the action list action runs params
func (o *ActionListActionRunsParams) WithActionID(actionID *string) *ActionListActionRunsParams {
	o.SetActionID(actionID)
	return o
}

// SetActionID adds the actionId to the action list action runs params
func (o *ActionListActionRunsParams) SetActionID(actionID *string) {
	o.ActionID = actionID
}

// WithActionName adds the actionName to the action list action runs params
func (o *ActionListActionRunsParams) WithActionName(actionName *string) *ActionListActionRunsParams {
	o.SetActionName(actionName)
	return o
}

// SetActionName adds the actionName to the action list action runs params
func (o *ActionListActionRunsParams) SetActionName(actionName *string) {
	o.ActionName = actionName
}

// WithNamespaceID adds the namespaceID to the action list action runs params
func (o *ActionListActionRunsParams) WithNamespaceID(namespaceID string) *ActionListActionRunsParams {
	o.SetNamespaceID(namespaceID)
	return o
}

// SetNamespaceID adds the namespaceId to the action list action runs params
func (o *ActionListActionRunsParams) SetNamespaceID(namespaceID string) {
	o.NamespaceID = namespaceID
}

// WriteToRequest writes these params to a swagger request
func (o *ActionListActionRunsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param namespace_id
	if err := r.SetPathParam("namespace_id", o.NamespaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}