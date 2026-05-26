package mtu

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/mtu"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
)

type NetworkMTU struct {
	networks.Network
	mtu.NetworkMTUExt
}

// CreateNetworkWithMTU will create a network with custom MTU. An error will be
// returned if the creation failed.
func CreateNetworkWithMTU(t *testing.T, client *gophercloud.ServiceClient, networkMTU *int) (*NetworkMTU, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
