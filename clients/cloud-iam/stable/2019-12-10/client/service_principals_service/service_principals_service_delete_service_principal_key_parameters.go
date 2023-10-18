// Code generated by go-swagger; DO NOT EDIT.

package service_principals_service

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

// NewServicePrincipalsServiceDeleteServicePrincipalKeyParams creates a new ServicePrincipalsServiceDeleteServicePrincipalKeyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServicePrincipalsServiceDeleteServicePrincipalKeyParams() *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	return &ServicePrincipalsServiceDeleteServicePrincipalKeyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipalKeyParamsWithTimeout creates a new ServicePrincipalsServiceDeleteServicePrincipalKeyParams object
// with the ability to set a timeout on a request.
func NewServicePrincipalsServiceDeleteServicePrincipalKeyParamsWithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	return &ServicePrincipalsServiceDeleteServicePrincipalKeyParams{
		timeout: timeout,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipalKeyParamsWithContext creates a new ServicePrincipalsServiceDeleteServicePrincipalKeyParams object
// with the ability to set a context for a request.
func NewServicePrincipalsServiceDeleteServicePrincipalKeyParamsWithContext(ctx context.Context) *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	return &ServicePrincipalsServiceDeleteServicePrincipalKeyParams{
		Context: ctx,
	}
}

// NewServicePrincipalsServiceDeleteServicePrincipalKeyParamsWithHTTPClient creates a new ServicePrincipalsServiceDeleteServicePrincipalKeyParams object
// with the ability to set a custom HTTPClient for a request.
func NewServicePrincipalsServiceDeleteServicePrincipalKeyParamsWithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	return &ServicePrincipalsServiceDeleteServicePrincipalKeyParams{
		HTTPClient: client,
	}
}

/*
ServicePrincipalsServiceDeleteServicePrincipalKeyParams contains all the parameters to send to the API endpoint

	for the service principals service delete service principal key operation.

	Typically these are written to a http.Request.
*/
type ServicePrincipalsServiceDeleteServicePrincipalKeyParams struct {

	/* ResourceName2.

	     resource_name is the resource name of the service principal key to
	delete.
	*/
	ResourceName2 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service principals service delete service principal key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) WithDefaults() *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service principals service delete service principal key params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service principals service delete service principal key params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) WithTimeout(timeout time.Duration) *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service principals service delete service principal key params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service principals service delete service principal key params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) WithContext(ctx context.Context) *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service principals service delete service principal key params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service principals service delete service principal key params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) WithHTTPClient(client *http.Client) *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service principals service delete service principal key params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResourceName2 adds the resourceName2 to the service principals service delete service principal key params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) WithResourceName2(resourceName2 string) *ServicePrincipalsServiceDeleteServicePrincipalKeyParams {
	o.SetResourceName2(resourceName2)
	return o
}

// SetResourceName2 adds the resourceName2 to the service principals service delete service principal key params
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) SetResourceName2(resourceName2 string) {
	o.ResourceName2 = resourceName2
}

// WriteToRequest writes these params to a swagger request
func (o *ServicePrincipalsServiceDeleteServicePrincipalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param resource_name_2
	if err := r.SetPathParam("resource_name_2", o.ResourceName2); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}