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
	"github.com/go-openapi/swag"
)

// NewPackerServiceListBucketsParams creates a new PackerServiceListBucketsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPackerServiceListBucketsParams() *PackerServiceListBucketsParams {
	return &PackerServiceListBucketsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPackerServiceListBucketsParamsWithTimeout creates a new PackerServiceListBucketsParams object
// with the ability to set a timeout on a request.
func NewPackerServiceListBucketsParamsWithTimeout(timeout time.Duration) *PackerServiceListBucketsParams {
	return &PackerServiceListBucketsParams{
		timeout: timeout,
	}
}

// NewPackerServiceListBucketsParamsWithContext creates a new PackerServiceListBucketsParams object
// with the ability to set a context for a request.
func NewPackerServiceListBucketsParamsWithContext(ctx context.Context) *PackerServiceListBucketsParams {
	return &PackerServiceListBucketsParams{
		Context: ctx,
	}
}

// NewPackerServiceListBucketsParamsWithHTTPClient creates a new PackerServiceListBucketsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPackerServiceListBucketsParamsWithHTTPClient(client *http.Client) *PackerServiceListBucketsParams {
	return &PackerServiceListBucketsParams{
		HTTPClient: client,
	}
}

/* PackerServiceListBucketsParams contains all the parameters to send to the API endpoint
   for the packer service list buckets operation.

   Typically these are written to a http.Request.
*/
type PackerServiceListBucketsParams struct {

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

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The max number of results per page that should be returned. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned which can be used to get the next page of results in subsequent
	requests. A value of zero will cause `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	/* SortingOrderBy.

	     Specifies the list of per field ordering that should be used for sorting.
	The order matters as rows are sorted in order by fields and when the field
	matches, the next field is used to tie break the ordering.
	The per field default ordering is ascending.

	The fields should be immutabile, unique, and orderable. If the field is
	not unique, more than one sort fields should be passed.

	Example: oder_by=name,age desc,created_at asc
	In that case, 'name' will get the default 'ascending' order.
	*/
	SortingOrderBy []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the packer service list buckets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceListBucketsParams) WithDefaults() *PackerServiceListBucketsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the packer service list buckets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PackerServiceListBucketsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithTimeout(timeout time.Duration) *PackerServiceListBucketsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithContext(ctx context.Context) *PackerServiceListBucketsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithHTTPClient(client *http.Client) *PackerServiceListBucketsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLocationOrganizationID adds the locationOrganizationID to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithLocationOrganizationID(locationOrganizationID string) *PackerServiceListBucketsParams {
	o.SetLocationOrganizationID(locationOrganizationID)
	return o
}

// SetLocationOrganizationID adds the locationOrganizationId to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetLocationOrganizationID(locationOrganizationID string) {
	o.LocationOrganizationID = locationOrganizationID
}

// WithLocationProjectID adds the locationProjectID to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithLocationProjectID(locationProjectID string) *PackerServiceListBucketsParams {
	o.SetLocationProjectID(locationProjectID)
	return o
}

// SetLocationProjectID adds the locationProjectId to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetLocationProjectID(locationProjectID string) {
	o.LocationProjectID = locationProjectID
}

// WithLocationRegionProvider adds the locationRegionProvider to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithLocationRegionProvider(locationRegionProvider *string) *PackerServiceListBucketsParams {
	o.SetLocationRegionProvider(locationRegionProvider)
	return o
}

// SetLocationRegionProvider adds the locationRegionProvider to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetLocationRegionProvider(locationRegionProvider *string) {
	o.LocationRegionProvider = locationRegionProvider
}

// WithLocationRegionRegion adds the locationRegionRegion to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithLocationRegionRegion(locationRegionRegion *string) *PackerServiceListBucketsParams {
	o.SetLocationRegionRegion(locationRegionRegion)
	return o
}

// SetLocationRegionRegion adds the locationRegionRegion to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetLocationRegionRegion(locationRegionRegion *string) {
	o.LocationRegionRegion = locationRegionRegion
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithPaginationNextPageToken(paginationNextPageToken *string) *PackerServiceListBucketsParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithPaginationPageSize(paginationPageSize *int64) *PackerServiceListBucketsParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *PackerServiceListBucketsParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithSortingOrderBy adds the sortingOrderBy to the packer service list buckets params
func (o *PackerServiceListBucketsParams) WithSortingOrderBy(sortingOrderBy []string) *PackerServiceListBucketsParams {
	o.SetSortingOrderBy(sortingOrderBy)
	return o
}

// SetSortingOrderBy adds the sortingOrderBy to the packer service list buckets params
func (o *PackerServiceListBucketsParams) SetSortingOrderBy(sortingOrderBy []string) {
	o.SortingOrderBy = sortingOrderBy
}

// WriteToRequest writes these params to a swagger request
func (o *PackerServiceListBucketsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.PaginationNextPageToken != nil {

		// query param pagination.next_page_token
		var qrPaginationNextPageToken string

		if o.PaginationNextPageToken != nil {
			qrPaginationNextPageToken = *o.PaginationNextPageToken
		}
		qPaginationNextPageToken := qrPaginationNextPageToken
		if qPaginationNextPageToken != "" {

			if err := r.SetQueryParam("pagination.next_page_token", qPaginationNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.PaginationPageSize != nil {

		// query param pagination.page_size
		var qrPaginationPageSize int64

		if o.PaginationPageSize != nil {
			qrPaginationPageSize = *o.PaginationPageSize
		}
		qPaginationPageSize := swag.FormatInt64(qrPaginationPageSize)
		if qPaginationPageSize != "" {

			if err := r.SetQueryParam("pagination.page_size", qPaginationPageSize); err != nil {
				return err
			}
		}
	}

	if o.PaginationPreviousPageToken != nil {

		// query param pagination.previous_page_token
		var qrPaginationPreviousPageToken string

		if o.PaginationPreviousPageToken != nil {
			qrPaginationPreviousPageToken = *o.PaginationPreviousPageToken
		}
		qPaginationPreviousPageToken := qrPaginationPreviousPageToken
		if qPaginationPreviousPageToken != "" {

			if err := r.SetQueryParam("pagination.previous_page_token", qPaginationPreviousPageToken); err != nil {
				return err
			}
		}
	}

	if o.SortingOrderBy != nil {

		// binding items for sorting.order_by
		joinedSortingOrderBy := o.bindParamSortingOrderBy(reg)

		// query array param sorting.order_by
		if err := r.SetQueryParam("sorting.order_by", joinedSortingOrderBy...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPackerServiceListBuckets binds the parameter sorting.order_by
func (o *PackerServiceListBucketsParams) bindParamSortingOrderBy(formats strfmt.Registry) []string {
	sortingOrderByIR := o.SortingOrderBy

	var sortingOrderByIC []string
	for _, sortingOrderByIIR := range sortingOrderByIR { // explode []string

		sortingOrderByIIV := sortingOrderByIIR // string as string
		sortingOrderByIC = append(sortingOrderByIC, sortingOrderByIIV)
	}

	// items.CollectionFormat: "multi"
	sortingOrderByIS := swag.JoinByFormat(sortingOrderByIC, "multi")

	return sortingOrderByIS
}
