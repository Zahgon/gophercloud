package portsbinding

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/portsbinding"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
)

// PortWithBindingExt represents a port with the binding fields
type PortWithBindingExt struct {
	ports.Port
	portsbinding.PortsBindingExt
}

// CreatePortsbinding will create a port on the specified subnet. An error will be
// returned if the port could not be created.
func CreatePortsbinding(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetID, hostID string, profile map[string]any) (PortWithBindingExt, error) {
	_ = "STUB: not implemented"
	return *new(PortWithBindingExt), nil
}
