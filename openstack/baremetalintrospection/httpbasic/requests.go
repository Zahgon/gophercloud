package httpbasic

import (
	"github.com/gophercloud/gophercloud/v2"
)

// EndpointOpts specifies a "http_basic" Ironic Inspector Endpoint.
type EndpointOpts struct {
	IronicInspectorEndpoint     string
	IronicInspectorUser         string
	IronicInspectorUserPassword string
}

func initClientOpts(client *gophercloud.ProviderClient, eo EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBareMetalIntrospectionHTTPBasic creates a ServiceClient that may be used to access a
// "http_basic" bare metal introspection service.
func NewBareMetalIntrospectionHTTPBasic(eo EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
