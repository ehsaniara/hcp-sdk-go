// Code generated by go-swagger; DO NOT EDIT.

package packer_service

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

// NewPackerServiceDeleteBucketParams creates a new PackerServiceDeleteBucketParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceDeleteBucketParams() *PackerServiceDeleteBucketParams {
	return &PackerServiceDeleteBucketParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceDeleteBucketParamsWithTimeout creates a new PackerServiceDeleteBucketParams object
// with the ability to set a timeout on a request.
func NewPackerServiceDeleteBucketParamsWithTimeout(timeout time.Duration) *PackerServiceDeleteBucketParams {
	return &PackerServiceDeleteBucketParams{
		timeout: timeout,
	}
}

// NewPackerServiceDeleteBucketParamsWithContext creates a new PackerServiceDeleteBucketParams object
// with the ability to set a context for a request.
func NewPackerServiceDeleteBucketParamsWithContext(ctx context.Context) *PackerServiceDeleteBucketParams {
	return &PackerServiceDeleteBucketParams{
		Context: ctx,
	}
}

// NewPackerServiceDeleteBucketParamsWithHTTPClient creates a new PackerServiceDeleteBucketParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceDeleteBucketParamsWithHTTPClient(client *http.Client) *PackerServiceDeleteBucketParams {
	return &PackerServiceDeleteBucketParams{
		HTTPClient: client,
	}
}

/* PackerServiceDeleteBucketParams contains all the parameters to send to the API endpoint
   for the packer service delete bucket operation.

   Typically these are written to a http.Request.
*/
type PackerServiceDeleteBucketParams struct {

	/* BucketSlug.

	   Human-readable name for the bucket.
	*/
	BucketSlug string

	/* LocationOrganizationID.

	   organization_id is the id of the organization.
	*/
	LocationOrganizationID string

	/* LocationProjectID.

	   project_id is the projects id.
	*/
	LocationProjectID string

	/* LocationRegionProvider.

	   provider is the named cloud provider ("aws", "gcp", "azure").
	*/
	LocationRegionProvider *string

	/* LocationRegionRegion.

	   region is the cloud region ("us-west1", "us-east1").
	*/
	LocationRegionRegion *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packer service delete bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceDeleteBucketParams) WithDefaults() *PackerServiceDeleteBucketParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service delete bucket params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceDeleteBucketParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) WithTimeout(timeout time.Duration) *PackerServiceDeleteBucketParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) WithContext(ctx context.Context) *PackerServiceDeleteBucketParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) WithHTTPClient(client *http.Client) *PackerServiceDeleteBucketParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketSlug adds the bucketSlug to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) WithBucketSlug(bucketSlug string) *PackerServiceDeleteBucketParams {
	o.SetBucketSlug(bucketSlug)
	return o
}

// SetBucketSlug adds the bucketSlug to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) SetBucketSlug(bucketSlug string) {
	o.BucketSlug = bucketSlug
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceDeleteBucketParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) WithLocationProjectID(locationProjectID string) *PackerServiceDeleteBucketParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceDeleteBucketParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceDeleteBucketParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service delete bucket params
func (o *PackerServiceDeleteBucketParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceDeleteBucketParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param bucket_slug
	if err := r.SetPathParam("bucket_slug", o.BucketSlug); err != nil {
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
