// Code generated by go-swagger; DO NOT EDIT.

package workflow_library

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

// NewWorkflowLibraryGetFromLibraryParams creates a new WorkflowLibraryGetFromLibraryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewWorkflowLibraryGetFromLibraryParams() *WorkflowLibraryGetFromLibraryParams {
	return &WorkflowLibraryGetFromLibraryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewWorkflowLibraryGetFromLibraryParamsWithTimeout creates a new WorkflowLibraryGetFromLibraryParams object
// with the ability to set a timeout on a request.
func NewWorkflowLibraryGetFromLibraryParamsWithTimeout(timeout time.Duration) *WorkflowLibraryGetFromLibraryParams {
	return &WorkflowLibraryGetFromLibraryParams{
		timeout: timeout,
	}
}

// NewWorkflowLibraryGetFromLibraryParamsWithContext creates a new WorkflowLibraryGetFromLibraryParams object
// with the ability to set a context for a request.
func NewWorkflowLibraryGetFromLibraryParamsWithContext(ctx context.Context) *WorkflowLibraryGetFromLibraryParams {
	return &WorkflowLibraryGetFromLibraryParams{
		Context: ctx,
	}
}

// NewWorkflowLibraryGetFromLibraryParamsWithHTTPClient creates a new WorkflowLibraryGetFromLibraryParams object
// with the ability to set a custom HTTPClient for a request.
func NewWorkflowLibraryGetFromLibraryParamsWithHTTPClient(client *http.Client) *WorkflowLibraryGetFromLibraryParams {
	return &WorkflowLibraryGetFromLibraryParams{
		HTTPClient: client,
	}
}

/*
WorkflowLibraryGetFromLibraryParams contains all the parameters to send to the API endpoint

	for the workflow library get from library operation.

	Typically these are written to a http.Request.
*/
type WorkflowLibraryGetFromLibraryParams struct {

	// ID.
	ID string

	// Name.
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the workflow library get from library params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowLibraryGetFromLibraryParams) WithDefaults() *WorkflowLibraryGetFromLibraryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the workflow library get from library params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *WorkflowLibraryGetFromLibraryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) WithTimeout(timeout time.Duration) *WorkflowLibraryGetFromLibraryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) WithContext(ctx context.Context) *WorkflowLibraryGetFromLibraryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) WithHTTPClient(client *http.Client) *WorkflowLibraryGetFromLibraryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) WithID(id string) *WorkflowLibraryGetFromLibraryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) SetID(id string) {
	o.ID = id
}

// WithName adds the name to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) WithName(name *string) *WorkflowLibraryGetFromLibraryParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the workflow library get from library params
func (o *WorkflowLibraryGetFromLibraryParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *WorkflowLibraryGetFromLibraryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}