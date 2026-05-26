package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/extradhcpopts"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/subnets"
)

// PortWithExtraDHCPOpts represents a port with extra DHCP options configuration.
type PortWithExtraDHCPOpts struct {
	ports.Port
	extradhcpopts.ExtraDHCPOptsExt
}

// CreateNetwork will create basic network. An error will be returned if the
// network could not be created.
func CreateNetwork(t *testing.T, client *gophercloud.ServiceClient) (*networks.Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateNetworkWithoutPortSecurity will create a network without port security.
// An error will be returned if the network could not be created.
func CreateNetworkWithoutPortSecurity(t *testing.T, client *gophercloud.ServiceClient) (*networks.Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePort will create a port on the specified subnet. An error will be
// returned if the port could not be created.
func CreatePort(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetID string) (*ports.Port, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePortWithNoSecurityGroup will create a port with no security group
// attached. An error will be returned if the port could not be created.
func CreatePortWithNoSecurityGroup(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetID string) (*ports.Port, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePortWithoutPortSecurity will create a port without port security on the
// specified subnet. An error will be returned if the port could not be created.
func CreatePortWithoutPortSecurity(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetID string) (*ports.Port, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePortWithExtraDHCPOpts will create a port with DHCP options on the
// specified subnet. An error will be returned if the port could not be created.
func CreatePortWithExtraDHCPOpts(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetID string) (*PortWithExtraDHCPOpts, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePortWithMultipleFixedIPs will create a port with two FixedIPs on the
// specified subnet. An error will be returned if the port could not be created.
func CreatePortWithMultipleFixedIPs(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetID string) (*ports.Port, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSubnet will create a subnet on the specified Network ID. An error
// will be returned if the subnet could not be created.
func CreateSubnet(t *testing.T, client *gophercloud.ServiceClient, networkID string) (*subnets.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSubnetWithCIDR will create a subnet on the specified Network ID and CIDR. An error
// will be returned if the subnet could not be created.
func CreateSubnetWithCIDR(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetCIDR, subnetGateway string) (*subnets.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSubnet will create a subnet on the specified Network ID and service types.
//
//	An error will be returned if the subnet could not be created.
func CreateSubnetWithServiceTypes(t *testing.T, client *gophercloud.ServiceClient, networkID string) (*subnets.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSubnetWithDefaultGateway will create a subnet on the specified Network
// ID and have Neutron set the gateway by default An error will be returned if
// the subnet could not be created.
func CreateSubnetWithDefaultGateway(t *testing.T, client *gophercloud.ServiceClient, networkID string) (*subnets.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSubnetWithNoGateway will create a subnet with no gateway on the
// specified Network ID.  An error will be returned if the subnet could not be
// created.
func CreateSubnetWithNoGateway(t *testing.T, client *gophercloud.ServiceClient, networkID string) (*subnets.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSubnetWithSubnetPool will create a subnet associated with the provided subnetpool on the specified Network ID.
// An error will be returned if the subnet or the subnetpool could not be created.
func CreateSubnetWithSubnetPool(t *testing.T, client *gophercloud.ServiceClient, networkID string, subnetPoolID string) (*subnets.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSubnetWithSubnetPoolNoCIDR will create a subnet associated with the
// provided subnetpool on the specified Network ID.
// An error will be returned if the subnet or the subnetpool could not be created.
func CreateSubnetWithSubnetPoolNoCIDR(t *testing.T, client *gophercloud.ServiceClient, networkID string, subnetPoolID string) (*subnets.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSubnetWithSubnetPoolPrefixlen will create a subnet associated with the
// provided subnetpool on the specified Network ID and with overwritten
// prefixlen instead of the default subnetpool prefixlen.
// An error will be returned if the subnet or the subnetpool could not be created.
func CreateSubnetWithSubnetPoolPrefixlen(t *testing.T, client *gophercloud.ServiceClient, networkID string, subnetPoolID string) (*subnets.Subnet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteNetwork will delete a network with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeleteNetwork(t *testing.T, client *gophercloud.ServiceClient, networkID string) {
	_ = "STUB: not implemented"
	return
}

// DeletePort will delete a port with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeletePort(t *testing.T, client *gophercloud.ServiceClient, portID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteSubnet will delete a subnet with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeleteSubnet(t *testing.T, client *gophercloud.ServiceClient, subnetID string) {
	_ = "STUB: not implemented"
	return
}

func WaitForPortToCreate(client *gophercloud.ServiceClient, portID string) error {
	_ = "STUB: not implemented"
	return nil
}

// This is duplicated from https://github.com/gophercloud/utils
// so that Gophercloud "core" doesn't have a dependency on the
// complementary utils repository.
func IDFromName(client *gophercloud.ServiceClient, name string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
