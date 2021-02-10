// Code generated by go-swagger; DO NOT EDIT.

package consul_service

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

// NewConsulServiceGetParams creates a new ConsulServiceGetParams object
// with the default values initialized.
func NewConsulServiceGetParams() *ConsulServiceGetParams {
	var ()
	return &ConsulServiceGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConsulServiceGetParamsWithTimeout creates a new ConsulServiceGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConsulServiceGetParamsWithTimeout(timeout time.Duration) *ConsulServiceGetParams {
	var ()
	return &ConsulServiceGetParams{

		timeout: timeout,
	}
}

// NewConsulServiceGetParamsWithContext creates a new ConsulServiceGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewConsulServiceGetParamsWithContext(ctx context.Context) *ConsulServiceGetParams {
	var ()
	return &ConsulServiceGetParams{

		Context: ctx,
	}
}

// NewConsulServiceGetParamsWithHTTPClient creates a new ConsulServiceGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConsulServiceGetParamsWithHTTPClient(client *http.Client) *ConsulServiceGetParams {
	var ()
	return &ConsulServiceGetParams{
		HTTPClient: client,
	}
}

/*ConsulServiceGetParams contains all the parameters to send to the API endpoint
for the consul service get operation typically these are written to a http.Request
*/
type ConsulServiceGetParams struct {

	/*ID
	  id is the unique ID of the HCC cluster to get

	*/
	ID string
	/*LocationOrganizationID
	  organization_id is the id of the organization.

	*/
	LocationOrganizationID string
	/*LocationProjectID
	  project_id is the projects id.

	*/
	LocationProjectID string
	/*LocationRegionProvider
	  provider is the named cloud provider ("aws", "gcp", "azure").

	*/
	LocationRegionProvider *string
	/*LocationRegionRegion
	  region is the cloud region ("us-west1", "us-east1").

	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the consul service get params
func (o *ConsulServiceGetParams) WithTimeout(timeout time.Duration) *ConsulServiceGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the consul service get params
func (o *ConsulServiceGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the consul service get params
func (o *ConsulServiceGetParams) WithContext(ctx context.Context) *ConsulServiceGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the consul service get params
func (o *ConsulServiceGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the consul service get params
func (o *ConsulServiceGetParams) WithHTTPClient(client *http.Client) *ConsulServiceGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the consul service get params
func (o *ConsulServiceGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the consul service get params
func (o *ConsulServiceGetParams) WithID(id string) *ConsulServiceGetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the consul service get params
func (o *ConsulServiceGetParams) SetID(id string) {
	o.ID = id
}

// WithLocationOrganizationID adds the locationOrganizationID to the consul service get params
func (o *ConsulServiceGetParams) WithLocationOrganizationID(locationOrganizationID string) *ConsulServiceGetParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the consul service get params
func (o *ConsulServiceGetParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the consul service get params
func (o *ConsulServiceGetParams) WithLocationProjectID(locationProjectID string) *ConsulServiceGetParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the consul service get params
func (o *ConsulServiceGetParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the consul service get params
func (o *ConsulServiceGetParams) WithLocationRegionProvider(locationRegionProvider *string) *ConsulServiceGetParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the consul service get params
func (o *ConsulServiceGetParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the consul service get params
func (o *ConsulServiceGetParams) WithLocationRegionRegion(locationRegionRegion *string) *ConsulServiceGetParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the consul service get params
func (o *ConsulServiceGetParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *ConsulServiceGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param location.organization_id
	if err := r.SetPathParam("location.organization_id", o.LocationOrganizationID); err != nil {
		return err
	}

	// path param location.project_id
	if err := r.SetPathParam("location.project_id", o.LocationProjectID); err != nil {
		return err
	}

	if o.LocationRegionProvider != nil {

		// query param location.region.provider
		var qrLocationRegionProvider string
		if o.LocationRegionProvider != nil {
			qrLocationRegionProvider = *o.LocationRegionProvider
		}
		qLocationRegionProvider := qrLocationRegionProvider
		if qLocationRegionProvider != "" {
			if err := r.SetQueryParam("location.region.provider", qLocationRegionProvider); err != nil {
				return err
			}
		}

	}

	if o.LocationRegionRegion != nil {

		// query param location.region.region
		var qrLocationRegionRegion string
		if o.LocationRegionRegion != nil {
			qrLocationRegionRegion = *o.LocationRegionRegion
		}
		qLocationRegionRegion := qrLocationRegionRegion
		if qLocationRegionRegion != "" {
			if err := r.SetQueryParam("location.region.region", qLocationRegionRegion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}