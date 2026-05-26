package layer3

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/addressscopes"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/floatingips"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/portforwarding"

	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/layer3/routers"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
)

// CreateFloatingIP creates a floating IP on a given network and port. An error
// will be returned if the creation failed.
func CreateFloatingIP(t *testing.T, client *gophercloud.ServiceClient, networkID, portID string) (*floatingips.FloatingIP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateFloatingIPWithFixedIP creates a floating IP on a given network and port with a
// defined fixed IP. An error will be returned if the creation failed.
func CreateFloatingIPWithFixedIP(t *testing.T, client *gophercloud.ServiceClient, networkID, portID, fixedIP string) (*floatingips.FloatingIP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePortForwarding creates a port forwarding for a given floating IP
// and port. An error will be returned if the creation failed.
func CreatePortForwarding(t *testing.T, client *gophercloud.ServiceClient, fipID string, portID string, portFixedIPs []ports.IP) (*portforwarding.PortForwarding, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeletePortForwarding deletes a Port Forwarding with a given ID and a given floating IP ID.
// A fatal error is returned if the deletion fails. Works best as a deferred function
func DeletePortForwarding(t *testing.T, client *gophercloud.ServiceClient, fipID string, pfID string) {
	_ = "STUB: not implemented"
	return
}

// CreateExternalRouter creates a router on the external network. This requires
// the OS_EXTGW_ID environment variable to be set. An error is returned if the
// creation failed.
func CreateExternalRouter(t *testing.T, client *gophercloud.ServiceClient) (*routers.Router, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRouter creates a router on a specified Network ID. An error will be
// returned if the creation failed.
func CreateRouter(t *testing.T, client *gophercloud.ServiceClient, networkID string) (*routers.Router, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRouterInterface will attach a subnet to a router. An error will be
// returned if the operation fails.
func CreateRouterInterface(t *testing.T, client *gophercloud.ServiceClient, portID, routerID string) (*routers.InterfaceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRouterInterfaceOnSubnet will attach a subnet to a router. An error will be
// returned if the operation fails.
func CreateRouterInterfaceOnSubnet(t *testing.T, client *gophercloud.ServiceClient, subnetID, routerID string) (*routers.InterfaceInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteRouter deletes a router of a specified ID. A fatal error will occur
// if the deletion failed. This works best when used as a deferred function.
func DeleteRouter(t *testing.T, client *gophercloud.ServiceClient, routerID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteRouterInterface will detach a subnet to a router. A fatal error will
// occur if the deletion failed. This works best when used as a deferred
// function.
func DeleteRouterInterface(t *testing.T, client *gophercloud.ServiceClient, portID, routerID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteFloatingIP deletes a floatingIP of a specified ID. A fatal error will
// occur if the deletion failed. This works best when used as a deferred
// function.
func DeleteFloatingIP(t *testing.T, client *gophercloud.ServiceClient, floatingIPID string) {
	_ = "STUB: not implemented"
	return
}

func WaitForRouterToCreate(client *gophercloud.ServiceClient, routerID string) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitForRouterToDelete(client *gophercloud.ServiceClient, routerID string) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitForRouterInterfaceToAttach(client *gophercloud.ServiceClient, routerInterfaceID string) error {
	_ = "STUB: not implemented"
	return nil
}

func WaitForRouterInterfaceToDetach(client *gophercloud.ServiceClient, routerInterfaceID string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateAddressScope will create an address-scope. An error will be returned if
// the address-scope could not be created.
func CreateAddressScope(t *testing.T, client *gophercloud.ServiceClient) (*addressscopes.AddressScope, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteAddressScope will delete an address-scope with the specified ID.
// A fatal error will occur if the delete was not successful.
func DeleteAddressScope(t *testing.T, client *gophercloud.ServiceClient, addressScopeID string) {
	_ = "STUB: not implemented"
	return
}
