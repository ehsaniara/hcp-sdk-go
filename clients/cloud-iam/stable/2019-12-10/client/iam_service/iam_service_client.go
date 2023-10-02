// Code generated by go-swagger; DO NOT EDIT.

package iam_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new iam service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iam service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	IamServiceCreateUserPrincipal(params *IamServiceCreateUserPrincipalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceCreateUserPrincipalOK, error)

	IamServiceDeleteOrganizationMembership(params *IamServiceDeleteOrganizationMembershipParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceDeleteOrganizationMembershipOK, error)

	IamServiceGetCallerIdentity(params *IamServiceGetCallerIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetCallerIdentityOK, error)

	IamServiceGetCurrentUserPrincipal(params *IamServiceGetCurrentUserPrincipalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetCurrentUserPrincipalOK, error)

	IamServiceGetOrganizationAuthMetadata(params *IamServiceGetOrganizationAuthMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetOrganizationAuthMetadataOK, error)

	IamServiceGetUserPrincipalByIDInOrganization(params *IamServiceGetUserPrincipalByIDInOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetUserPrincipalByIDInOrganizationOK, error)

	IamServiceGetUserPrincipalsByIDsInOrganization(params *IamServiceGetUserPrincipalsByIDsInOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetUserPrincipalsByIDsInOrganizationOK, error)

	IamServiceListUserPrincipalsByOrganization(params *IamServiceListUserPrincipalsByOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceListUserPrincipalsByOrganizationOK, error)

	IamServicePing(params *IamServicePingParams, opts ...ClientOption) (*IamServicePingOK, error)

	IamServiceSearchPrincipals(params *IamServiceSearchPrincipalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceSearchPrincipalsOK, error)

	IamServiceUpdateWebConsolePreferences(params *IamServiceUpdateWebConsolePreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceUpdateWebConsolePreferencesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
IamServiceCreateUserPrincipal creates user principal creates a new user principal
*/
func (a *Client) IamServiceCreateUserPrincipal(params *IamServiceCreateUserPrincipalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceCreateUserPrincipalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceCreateUserPrincipalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_CreateUserPrincipal",
		Method:             "POST",
		PathPattern:        "/iam/2019-12-10/user-principals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceCreateUserPrincipalReader{formats: a.formats},
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
	success, ok := result.(*IamServiceCreateUserPrincipalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceCreateUserPrincipalDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceDeleteOrganizationMembership deletes organization membership deletes a user principal s organization membership
*/
func (a *Client) IamServiceDeleteOrganizationMembership(params *IamServiceDeleteOrganizationMembershipParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceDeleteOrganizationMembershipOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceDeleteOrganizationMembershipParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_DeleteOrganizationMembership",
		Method:             "DELETE",
		PathPattern:        "/iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceDeleteOrganizationMembershipReader{formats: a.formats},
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
	success, ok := result.(*IamServiceDeleteOrganizationMembershipOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceDeleteOrganizationMembershipDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceGetCallerIdentity gets caller identity returns the identity of the current caller
*/
func (a *Client) IamServiceGetCallerIdentity(params *IamServiceGetCallerIdentityParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetCallerIdentityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceGetCallerIdentityParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_GetCallerIdentity",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/caller-identity",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceGetCallerIdentityReader{formats: a.formats},
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
	success, ok := result.(*IamServiceGetCallerIdentityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceGetCallerIdentityDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceGetCurrentUserPrincipal gets current user principal retrieves information about the current user principal this endpoint it meant to be used by external clients over an HTTP API it supports retrieving the basic user principal data useful for any client and optionally the user preferences for the h c p web portal j s application
*/
func (a *Client) IamServiceGetCurrentUserPrincipal(params *IamServiceGetCurrentUserPrincipalParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetCurrentUserPrincipalOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceGetCurrentUserPrincipalParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_GetCurrentUserPrincipal",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceGetCurrentUserPrincipalReader{formats: a.formats},
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
	success, ok := result.(*IamServiceGetCurrentUserPrincipalOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceGetCurrentUserPrincipalDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceGetOrganizationAuthMetadata gets organization auth metadata returns metadata about the organization s configured authentication methods
*/
func (a *Client) IamServiceGetOrganizationAuthMetadata(params *IamServiceGetOrganizationAuthMetadataParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetOrganizationAuthMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceGetOrganizationAuthMetadataParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_GetOrganizationAuthMetadata",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/organizations/{organization_id}/auth-metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceGetOrganizationAuthMetadataReader{formats: a.formats},
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
	success, ok := result.(*IamServiceGetOrganizationAuthMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceGetOrganizationAuthMetadataDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceGetUserPrincipalByIDInOrganization gets user principal retrieves a user principal
*/
func (a *Client) IamServiceGetUserPrincipalByIDInOrganization(params *IamServiceGetUserPrincipalByIDInOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetUserPrincipalByIDInOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceGetUserPrincipalByIDInOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_GetUserPrincipalByIdInOrganization",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/organizations/{organization_id}/user-principals/{user_principal_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceGetUserPrincipalByIDInOrganizationReader{formats: a.formats},
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
	success, ok := result.(*IamServiceGetUserPrincipalByIDInOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceGetUserPrincipalByIDInOrganizationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceGetUserPrincipalsByIDsInOrganization gets user principals by i ds in organization is a batch method to fetch users by ID for a given organization if some of the requested users don t exist or aren t members of the given organization then they will be omitted in response we re using p o s t because g e t has a length limitation for URL which is given that user id is UUID would limit us to fetching up to 48 users at a time which is less than ideal
*/
func (a *Client) IamServiceGetUserPrincipalsByIDsInOrganization(params *IamServiceGetUserPrincipalsByIDsInOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceGetUserPrincipalsByIDsInOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceGetUserPrincipalsByIDsInOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_GetUserPrincipalsByIDsInOrganization",
		Method:             "POST",
		PathPattern:        "/iam/2019-12-10/organizations/{organization_id}/user-principals/batch-fetch",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceGetUserPrincipalsByIDsInOrganizationReader{formats: a.formats},
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
	success, ok := result.(*IamServiceGetUserPrincipalsByIDsInOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceGetUserPrincipalsByIDsInOrganizationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceListUserPrincipalsByOrganization lists user principals by organization returns a list of principals that are members of an organization
*/
func (a *Client) IamServiceListUserPrincipalsByOrganization(params *IamServiceListUserPrincipalsByOrganizationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceListUserPrincipalsByOrganizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceListUserPrincipalsByOrganizationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_ListUserPrincipalsByOrganization",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/organizations/{organization_id}/user-principals",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceListUserPrincipalsByOrganizationReader{formats: a.formats},
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
	success, ok := result.(*IamServiceListUserPrincipalsByOrganizationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceListUserPrincipalsByOrganizationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServicePing pings pings the healthcheck endpoint exposed for HTTP healthchecks via datadog synthetic monitoring
*/
func (a *Client) IamServicePing(params *IamServicePingParams, opts ...ClientOption) (*IamServicePingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServicePingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_Ping",
		Method:             "GET",
		PathPattern:        "/iam/2019-12-10/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServicePingReader{formats: a.formats},
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
	success, ok := result.(*IamServicePingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServicePingDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceSearchPrincipals searches principals returns principal details for principals within the supplied organization optional filters may be specified to filter the result set
*/
func (a *Client) IamServiceSearchPrincipals(params *IamServiceSearchPrincipalsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceSearchPrincipalsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceSearchPrincipalsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_SearchPrincipals",
		Method:             "POST",
		PathPattern:        "/iam/2019-12-10/organizations/{organization_id}/principals/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceSearchPrincipalsReader{formats: a.formats},
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
	success, ok := result.(*IamServiceSearchPrincipalsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceSearchPrincipalsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
IamServiceUpdateWebConsolePreferences updates web console preferences updates a user principal s web portal fka web console preferences
*/
func (a *Client) IamServiceUpdateWebConsolePreferences(params *IamServiceUpdateWebConsolePreferencesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*IamServiceUpdateWebConsolePreferencesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewIamServiceUpdateWebConsolePreferencesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "IamService_UpdateWebConsolePreferences",
		Method:             "PUT",
		PathPattern:        "/iam/2019-12-10/me/web-portal-preferences",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &IamServiceUpdateWebConsolePreferencesReader{formats: a.formats},
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
	success, ok := result.(*IamServiceUpdateWebConsolePreferencesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*IamServiceUpdateWebConsolePreferencesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
