// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new version service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for version service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVersion(params *CreateVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVersionOK, error)

	DeleteVersion(params *DeleteVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVersionOK, error)

	ListVersions(params *ListVersionsParams, opts ...ClientOption) (*ListVersionsOK, error)

	ReadVersion(params *ReadVersionParams, opts ...ClientOption) (*ReadVersionOK, error)

	ReleaseVersion(params *ReleaseVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReleaseVersionOK, error)

	RevokeVersion(params *RevokeVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RevokeVersionOK, error)

	UpdateVersion(params *UpdateVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateVersion creates version creates a new box version
*/
func (a *Client) CreateVersion(params *CreateVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CreateVersion",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVersionReader{formats: a.formats},
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
	success, ok := result.(*CreateVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
	DeleteVersion deletes version deletes a box version

	Deleting a Box Version removes all its Providers as well. This

operation cannot be undone.
*/
func (a *Client) DeleteVersion(params *DeleteVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteVersion",
		Method:             "DELETE",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVersionReader{formats: a.formats},
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
	success, ok := result.(*DeleteVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListVersions lists version lists all of the versions within a particular box
*/
func (a *Client) ListVersions(params *ListVersionsParams, opts ...ClientOption) (*ListVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVersionsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ListVersions",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVersionsReader{formats: a.formats},
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
	success, ok := result.(*ListVersionsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListVersionsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReadVersion reads version reads a box version
*/
func (a *Client) ReadVersion(params *ReadVersionParams, opts ...ClientOption) (*ReadVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReadVersion",
		Method:             "GET",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadVersionReader{formats: a.formats},
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
	success, ok := result.(*ReadVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReadVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ReleaseVersion releases releases the specified version the version must not already be released
*/
func (a *Client) ReleaseVersion(params *ReleaseVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ReleaseVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReleaseVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ReleaseVersion",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/release",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReleaseVersionReader{formats: a.formats},
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
	success, ok := result.(*ReleaseVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ReleaseVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RevokeVersion revokes revokes the specified version the version must be actively released
*/
func (a *Client) RevokeVersion(params *RevokeVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RevokeVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRevokeVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RevokeVersion",
		Method:             "PUT",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RevokeVersionReader{formats: a.formats},
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
	success, ok := result.(*RevokeVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RevokeVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateVersion updates version updates a box version
*/
func (a *Client) UpdateVersion(params *UpdateVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UpdateVersion",
		Method:             "PATCH",
		PathPattern:        "/vagrant/2022-09-30/registry/{registry}/boxes/{box}/versions/{version}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateVersionReader{formats: a.formats},
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
	success, ok := result.(*UpdateVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateVersionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
