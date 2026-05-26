package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/vlantransparent"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
)

// VLANTransparentNetwork represents OpenStack V2 Networking Network with the
// "vlan-transparent" extension enabled.
type VLANTransparentNetwork struct {
	networks.Network
	vlantransparent.TransparentExt
}

// ListVLANTransparentNetworks will list networks with the "vlan-transparent"
// extension. An error will be returned networks could not be listed.
func ListVLANTransparentNetworks(t *testing.T, client *gophercloud.ServiceClient) ([]*VLANTransparentNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateVLANTransparentNetwork will create a network with the
// "vlan-transparent" extension. An error will be returned if the network could
// not be created.
func CreateVLANTransparentNetwork(t *testing.T, client *gophercloud.ServiceClient) (*VLANTransparentNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
