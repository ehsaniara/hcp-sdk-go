// Code generated by go-swagger; DO NOT EDIT.

package global_network_manager_service

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

// NewListServiceInstancesParams creates a new ListServiceInstancesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServiceInstancesParams() *ListServiceInstancesParams {
	return &ListServiceInstancesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServiceInstancesParamsWithTimeout creates a new ListServiceInstancesParams object
// with the ability to set a timeout on a request.
func NewListServiceInstancesParamsWithTimeout(timeout time.Duration) *ListServiceInstancesParams {
	return &ListServiceInstancesParams{
		timeout: timeout,
	}
}

// NewListServiceInstancesParamsWithContext creates a new ListServiceInstancesParams object
// with the ability to set a context for a request.
func NewListServiceInstancesParamsWithContext(ctx context.Context) *ListServiceInstancesParams {
	return &ListServiceInstancesParams{
		Context: ctx,
	}
}

// NewListServiceInstancesParamsWithHTTPClient creates a new ListServiceInstancesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServiceInstancesParamsWithHTTPClient(client *http.Client) *ListServiceInstancesParams {
	return &ListServiceInstancesParams{
		HTTPClient: client,
	}
}

/*
ListServiceInstancesParams contains all the parameters to send to the API endpoint

	for the list service instances operation.

	Typically these are written to a http.Request.
*/
type ListServiceInstancesParams struct {

	// ClusterResourceName.
	ClusterResourceName string

	// Namespace.
	Namespace *string

	/* PaginationNextPageToken.

	     Specifies a page token to use to retrieve the next page. Set this parameter to the
	`next_page_token` returned by previous list requests to get the next page of
	results. If set, `previous_page_token` must not be set.
	*/
	PaginationNextPageToken *string

	/* PaginationPageSize.

	     The maximum number of results per page to return. If the number
	of available results is larger than `page_size`, a `next_page_token` is
	returned, which you can use to get the next page of results in subsequent
	requests. A value of zero causes `page_size` to be defaulted.

	     Format: int64
	*/
	PaginationPageSize *int64

	/* PaginationPreviousPageToken.

	     Specifies a page token to use to retrieve the previous page. Set this parameter to
	the `previous_page_token` returned by previous list requests to get the
	previous page of results. If set, `next_page_token` must not be set.
	*/
	PaginationPreviousPageToken *string

	// Partition.
	Partition *string

	/* Query.

	   Search query to filter by. Searches across `id` and `tags`.
	*/
	Query *string

	// ServiceName.
	ServiceName string

	/* Status.

	   Query param filter: `status` of the service. This can be combination of `passing`, `warning`, `critical`, `none`
	*/
	Status []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list service instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServiceInstancesParams) WithDefaults() *ListServiceInstancesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list service instances params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServiceInstancesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list service instances params
func (o *ListServiceInstancesParams) WithTimeout(timeout time.Duration) *ListServiceInstancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list service instances params
func (o *ListServiceInstancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list service instances params
func (o *ListServiceInstancesParams) WithContext(ctx context.Context) *ListServiceInstancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list service instances params
func (o *ListServiceInstancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list service instances params
func (o *ListServiceInstancesParams) WithHTTPClient(client *http.Client) *ListServiceInstancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list service instances params
func (o *ListServiceInstancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterResourceName adds the clusterResourceName to the list service instances params
func (o *ListServiceInstancesParams) WithClusterResourceName(clusterResourceName string) *ListServiceInstancesParams {
	o.SetClusterResourceName(clusterResourceName)
	return o
}

// SetClusterResourceName adds the clusterResourceName to the list service instances params
func (o *ListServiceInstancesParams) SetClusterResourceName(clusterResourceName string) {
	o.ClusterResourceName = clusterResourceName
}

// WithNamespace adds the namespace to the list service instances params
func (o *ListServiceInstancesParams) WithNamespace(namespace *string) *ListServiceInstancesParams {
	o.SetNamespace(namespace)
	return o
}

// SetNamespace adds the namespace to the list service instances params
func (o *ListServiceInstancesParams) SetNamespace(namespace *string) {
	o.Namespace = namespace
}

// WithPaginationNextPageToken adds the paginationNextPageToken to the list service instances params
func (o *ListServiceInstancesParams) WithPaginationNextPageToken(paginationNextPageToken *string) *ListServiceInstancesParams {
	o.SetPaginationNextPageToken(paginationNextPageToken)
	return o
}

// SetPaginationNextPageToken adds the paginationNextPageToken to the list service instances params
func (o *ListServiceInstancesParams) SetPaginationNextPageToken(paginationNextPageToken *string) {
	o.PaginationNextPageToken = paginationNextPageToken
}

// WithPaginationPageSize adds the paginationPageSize to the list service instances params
func (o *ListServiceInstancesParams) WithPaginationPageSize(paginationPageSize *int64) *ListServiceInstancesParams {
	o.SetPaginationPageSize(paginationPageSize)
	return o
}

// SetPaginationPageSize adds the paginationPageSize to the list service instances params
func (o *ListServiceInstancesParams) SetPaginationPageSize(paginationPageSize *int64) {
	o.PaginationPageSize = paginationPageSize
}

// WithPaginationPreviousPageToken adds the paginationPreviousPageToken to the list service instances params
func (o *ListServiceInstancesParams) WithPaginationPreviousPageToken(paginationPreviousPageToken *string) *ListServiceInstancesParams {
	o.SetPaginationPreviousPageToken(paginationPreviousPageToken)
	return o
}

// SetPaginationPreviousPageToken adds the paginationPreviousPageToken to the list service instances params
func (o *ListServiceInstancesParams) SetPaginationPreviousPageToken(paginationPreviousPageToken *string) {
	o.PaginationPreviousPageToken = paginationPreviousPageToken
}

// WithPartition adds the partition to the list service instances params
func (o *ListServiceInstancesParams) WithPartition(partition *string) *ListServiceInstancesParams {
	o.SetPartition(partition)
	return o
}

// SetPartition adds the partition to the list service instances params
func (o *ListServiceInstancesParams) SetPartition(partition *string) {
	o.Partition = partition
}

// WithQuery adds the query to the list service instances params
func (o *ListServiceInstancesParams) WithQuery(query *string) *ListServiceInstancesParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the list service instances params
func (o *ListServiceInstancesParams) SetQuery(query *string) {
	o.Query = query
}

// WithServiceName adds the serviceName to the list service instances params
func (o *ListServiceInstancesParams) WithServiceName(serviceName string) *ListServiceInstancesParams {
	o.SetServiceName(serviceName)
	return o
}

// SetServiceName adds the serviceName to the list service instances params
func (o *ListServiceInstancesParams) SetServiceName(serviceName string) {
	o.ServiceName = serviceName
}

// WithStatus adds the status to the list service instances params
func (o *ListServiceInstancesParams) WithStatus(status []string) *ListServiceInstancesParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the list service instances params
func (o *ListServiceInstancesParams) SetStatus(status []string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *ListServiceInstancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_resource_name
	if err := r.SetPathParam("cluster_resource_name", o.ClusterResourceName); err != nil {
		return err
	}

	if o.Namespace != nil {

		// query param namespace
		var qrNamespace string

		if o.Namespace != nil {
			qrNamespace = *o.Namespace
		}
		qNamespace := qrNamespace
		if qNamespace != "" {

			if err := r.SetQueryParam("namespace", qNamespace); err != nil {
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

	if o.Partition != nil {

		// query param partition
		var qrPartition string

		if o.Partition != nil {
			qrPartition = *o.Partition
		}
		qPartition := qrPartition
		if qPartition != "" {

			if err := r.SetQueryParam("partition", qPartition); err != nil {
				return err
			}
		}
	}

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	// path param service_name
	if err := r.SetPathParam("service_name", o.ServiceName); err != nil {
		return err
	}

	if o.Status != nil {

		// binding items for status
		joinedStatus := o.bindParamStatus(reg)

		// query array param status
		if err := r.SetQueryParam("status", joinedStatus...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamListServiceInstances binds the parameter status
func (o *ListServiceInstancesParams) bindParamStatus(formats strfmt.Registry) []string {
	statusIR := o.Status

	var statusIC []string
	for _, statusIIR := range statusIR { // explode []string

		statusIIV := statusIIR // string as string
		statusIC = append(statusIC, statusIIV)
	}

	// items.CollectionFormat: "multi"
	statusIS := swag.JoinByFormat(statusIC, "multi")

	return statusIS
}
