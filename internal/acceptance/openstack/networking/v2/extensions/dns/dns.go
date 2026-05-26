package dns

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/dns"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
)

// PortWithDNSExt represents a port with the DNS fields
type PortWithDNSExt struct {
	ports.Port
	dns.PortDNSExt
}

// FloatingIPWithDNSExt represents a floating IP with the DNS fields
type FloatingIPWithDNSExt struct {
	floatingips.FloatingIP
	dns.FloatingIPDNSExt
}

// NetworkWithDNSExt represents a network with the DNS fields
type NetworkWithDNSExt struct {
	networks.Network
	dns.NetworkDNSExt
}

// CreatePortDNS will create a port with a DNS name on the specified subnet. An
// error will be returned if the port could not be created.
func CreatePortDNS(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetID, dnsName string) (*PortWithDNSExt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateFloatingIPDNS creates a floating IP with the DNS extension on a given
// network and port. An error will be returned if the creation failed.
func CreateFloatingIPDNS(t *testing.T, client *gophercloud.ServiceClient, networkID, portID, dnsName, dnsDomain string) (*FloatingIPWithDNSExt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateNetworkDNS will create a network with a DNS domain set.
// An error will be returned if the network could not be created.
func CreateNetworkDNS(t *testing.T, client *gophercloud.ServiceClient, dnsDomain string) (*NetworkWithDNSExt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
