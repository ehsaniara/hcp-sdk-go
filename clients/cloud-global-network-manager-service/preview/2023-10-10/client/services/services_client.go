// Code generated by go-swagger; DO NOT EDIT.

package services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new services API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for services API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetService(params *GetServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceOK, error)

	ListClusterServices(params *ListClusterServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClusterServicesOK, error)

	ListServices(params *ListServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetService get service API
*/
func (a *Client) GetService(params *GetServiceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetServiceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetServiceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Get Service",
		Method:             "GET",
		PathPattern:        "/2023-10-10/{cluster_resource_name}/service/{service_name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetServiceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetServiceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetServiceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListClusterServices list cluster services API
*/
func (a *Client) ListClusterServices(params *ListClusterServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListClusterServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListClusterServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "List Cluster Services",
		Method:             "GET",
		PathPattern:        "/2023-10-10/{cluster_resource_name}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListClusterServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListClusterServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListClusterServicesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListServices services a p is
*/
func (a *Client) ListServices(params *ListServicesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListServicesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListServicesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "List Services",
		Method:             "GET",
		PathPattern:        "/2023-10-10/consul/{project_resource_name}/services",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListServicesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListServicesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListServicesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
