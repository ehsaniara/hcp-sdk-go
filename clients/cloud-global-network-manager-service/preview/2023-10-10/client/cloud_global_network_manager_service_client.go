// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2023-10-10/client/clusters"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2023-10-10/client/nodes"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2023-10-10/client/service_instances"
	"github.com/hashicorp/hcp-sdk-go/clients/cloud-global-network-manager-service/preview/2023-10-10/client/services"
)

// Default cloud global network manager service HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.cloud.hashicorp.com"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http"}

// NewHTTPClient creates a new cloud global network manager service HTTP client.
func NewHTTPClient(formats strfmt.Registry) *CloudGlobalNetworkManagerService {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new cloud global network manager service HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *CloudGlobalNetworkManagerService {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new cloud global network manager service client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *CloudGlobalNetworkManagerService {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(CloudGlobalNetworkManagerService)
	cli.Transport = transport
	cli.Clusters = clusters.New(transport, formats)
	cli.Nodes = nodes.New(transport, formats)
	cli.ServiceInstances = service_instances.New(transport, formats)
	cli.Services = services.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// CloudGlobalNetworkManagerService is a client for cloud global network manager service
type CloudGlobalNetworkManagerService struct {
	Clusters clusters.ClientService

	Nodes nodes.ClientService

	ServiceInstances service_instances.ClientService

	Services services.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *CloudGlobalNetworkManagerService) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Clusters.SetTransport(transport)
	c.Nodes.SetTransport(transport)
	c.ServiceInstances.SetTransport(transport)
	c.Services.SetTransport(transport)
}
