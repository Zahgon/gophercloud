package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/sharenetworks"
)

// CreateShareNetwork will create a share network with a random name. An
// error will be returned if the share network was unable to be created.
func CreateShareNetwork(t *testing.T, client *gophercloud.ServiceClient) (*sharenetworks.ShareNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteShareNetwork will delete a share network. An error will occur if
// the share network was unable to be deleted.
func DeleteShareNetwork(t *testing.T, client *gophercloud.ServiceClient, shareNetworkID string) {
	_ = "STUB: not implemented"
	return
}
