// Code generated by go-swagger; DO NOT EDIT.

package iam_service

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

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-iam/stable/2019-12-10/models"
)

// NewIamServiceGetUserPrincipalsByIDsInOrganizationParams creates a new IamServiceGetUserPrincipalsByIDsInOrganizationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIamServiceGetUserPrincipalsByIDsInOrganizationParams() *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	return &IamServiceGetUserPrincipalsByIDsInOrganizationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIamServiceGetUserPrincipalsByIDsInOrganizationParamsWithTimeout creates a new IamServiceGetUserPrincipalsByIDsInOrganizationParams object
// with the ability to set a timeout on a request.
func NewIamServiceGetUserPrincipalsByIDsInOrganizationParamsWithTimeout(timeout time.Duration) *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	return &IamServiceGetUserPrincipalsByIDsInOrganizationParams{
		timeout: timeout,
	}
}

// NewIamServiceGetUserPrincipalsByIDsInOrganizationParamsWithContext creates a new IamServiceGetUserPrincipalsByIDsInOrganizationParams object
// with the ability to set a context for a request.
func NewIamServiceGetUserPrincipalsByIDsInOrganizationParamsWithContext(ctx context.Context) *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	return &IamServiceGetUserPrincipalsByIDsInOrganizationParams{
		Context: ctx,
	}
}

// NewIamServiceGetUserPrincipalsByIDsInOrganizationParamsWithHTTPClient creates a new IamServiceGetUserPrincipalsByIDsInOrganizationParams object
// with the ability to set a custom HTTPClient for a request.
func NewIamServiceGetUserPrincipalsByIDsInOrganizationParamsWithHTTPClient(client *http.Client) *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	return &IamServiceGetUserPrincipalsByIDsInOrganizationParams{
		HTTPClient: client,
	}
}

/*
IamServiceGetUserPrincipalsByIDsInOrganizationParams contains all the parameters to send to the API endpoint

	for the iam service get user principals by i ds in organization operation.

	Typically these are written to a http.Request.
*/
type IamServiceGetUserPrincipalsByIDsInOrganizationParams struct {

	// Body.
	Body *models.HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest

	/* OrganizationID.

	   organization_id is the ID of organization for users to look up
	*/
	OrganizationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iam service get user principals by i ds in organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) WithDefaults() *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iam service get user principals by i ds in organization params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) WithTimeout(timeout time.Duration) *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) WithContext(ctx context.Context) *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) WithHTTPClient(client *http.Client) *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) WithBody(body *models.HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest) *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) SetBody(body *models.HashicorpCloudIamGetUserPrincipalsByIDsInOrganizationRequest) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) WithOrganizationID(organizationID string) *IamServiceGetUserPrincipalsByIDsInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the iam service get user principals by i ds in organization params
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) SetOrganizationID(organizationID string) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *IamServiceGetUserPrincipalsByIDsInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param organization_id
	if err := r.SetPathParam("organization_id", o.OrganizationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}