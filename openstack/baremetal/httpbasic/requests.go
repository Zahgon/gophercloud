package httpbasic

import (
	"github.com/gophercloud/gophercloud/v2"
)

// EndpointOpts specifies a "http_basic" Ironic Endpoint
type EndpointOpts struct {
	IronicEndpoint     string
	IronicUser         string
	IronicUserPassword string
}

func initClientOpts(client *gophercloud.ProviderClient, eo EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBareMetalHTTPBasic creates a ServiceClient that may be used to access a
// "http_basic" bare metal service.
func NewBareMetalHTTPBasic(eo EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
