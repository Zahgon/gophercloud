package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/subnetpools"
)

// CreateSubnetPool will create a subnetpool. An error will be returned if the
// subnetpool could not be created.
func CreateSubnetPool(t *testing.T, client *gophercloud.ServiceClient) (*subnetpools.SubnetPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSubnetPool will delete a subnetpool with a specified ID.
// A fatal error will occur if the delete was not successful.
func DeleteSubnetPool(t *testing.T, client *gophercloud.ServiceClient, subnetPoolID string) {
	_ = "STUB: not implemented"
	return
}
