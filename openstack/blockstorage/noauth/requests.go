package noauth

import (
	"github.com/gophercloud/gophercloud/v2"
)

// EndpointOpts specifies a "noauth" Cinder Endpoint.
type EndpointOpts struct {
	// CinderEndpoint [required] is currently only used with "noauth" Cinder.
	// A cinder endpoint with "auth_strategy=noauth" is necessary, for example:
	// http://example.com:8776/v2.
	CinderEndpoint string
}

// NewClient prepares an unauthenticated ProviderClient instance.
func NewClient(options gophercloud.AuthOptions) (*gophercloud.ProviderClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func initClientOpts(client *gophercloud.ProviderClient, eo EndpointOpts, clientType string) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBlockStorageNoAuthV2 creates a ServiceClient that may be used to access "noauth" v2 block storage service.
func NewBlockStorageNoAuthV2(client *gophercloud.ProviderClient, eo EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBlockStorageNoAuthV3 creates a ServiceClient that may be used to access "noauth" v3 block storage service.
func NewBlockStorageNoAuthV3(client *gophercloud.ProviderClient, eo EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
