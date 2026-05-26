package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/placement/v1/resourceproviders"
)

func CreateResourceProvider(t *testing.T, client *gophercloud.ServiceClient) (*resourceproviders.ResourceProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func CreateResourceProviderWithParent(t *testing.T, client *gophercloud.ServiceClient, parentUUID string) (*resourceproviders.ResourceProvider, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteResourceProvider will delete a resource provider with a specified ID.
// A fatal error will occur if the delete was not successful. This works best when
// used as a deferred function.
func DeleteResourceProvider(t *testing.T, client *gophercloud.ServiceClient, resourceProviderID string) {
	_ = "STUB: not implemented"
	return
}

// CreateResourceProviderWithVCPUInventory creates a resource provider and seeds it
// with a VCPU inventory, returning the provider and the inventory generation.
// This is used by acceptance tests that need a resource provider with available
// capacity before setting allocations against it.
func CreateResourceProviderWithVCPUInventory(t *testing.T, client *gophercloud.ServiceClient) (*resourceproviders.ResourceProvider, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}
