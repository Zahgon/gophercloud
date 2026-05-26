package v1

import (
	"context"
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/baremetal/v1/allocations"
	"github.com/gophercloud/gophercloud/v2/openstack/baremetal/v1/nodes"
	"github.com/gophercloud/gophercloud/v2/openstack/baremetal/v1/portgroups"
	"github.com/gophercloud/gophercloud/v2/openstack/baremetal/v1/ports"
)

// CreateNode creates a basic node with a randomly generated name.
func CreateNode(t *testing.T, client *gophercloud.ServiceClient) (*nodes.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteNode deletes a bare metal node via its UUID.
func DeleteNode(t *testing.T, client *gophercloud.ServiceClient, node *nodes.Node) {
	_ = "STUB: not implemented"
	// Force deletion of provisioned nodes requires maintenance mode.
	return
}

// CreateAllocation creates an allocation
func CreateAllocation(t *testing.T, client *gophercloud.ServiceClient) (*allocations.Allocation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteAllocation deletes a bare metal allocation via its UUID.
func DeleteAllocation(t *testing.T, client *gophercloud.ServiceClient, allocation *allocations.Allocation) {
	_ = "STUB: not implemented"
	return
}

// CreatePortGroup creates an allocation
func CreatePortGroup(t *testing.T, client *gophercloud.ServiceClient, node *nodes.Node) (*portgroups.PortGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeletePortGroup deletes a bare metal portgroup via its UUID.
func DeletePortGroup(t *testing.T, client *gophercloud.ServiceClient, portgroup *portgroups.PortGroup) {
	_ = "STUB: not implemented"
	return
}

// CreateFakeNode creates a node with fake-hardware.
func CreateFakeNode(t *testing.T, client *gophercloud.ServiceClient) (*nodes.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ChangeProvisionStateAndWait(ctx context.Context, client *gophercloud.ServiceClient, node *nodes.Node,
	change nodes.ProvisionStateOpts, expectedState nodes.ProvisionState) (*nodes.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeployFakeNode deploys a node that uses fake-hardware.
func DeployFakeNode(t *testing.T, client *gophercloud.ServiceClient, node *nodes.Node) (*nodes.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePort - creates a port for a node with a fixed Address
func CreatePort(t *testing.T, client *gophercloud.ServiceClient, node *nodes.Node) (*ports.Port, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeletePort - deletes a port via its UUID
func DeletePort(t *testing.T, client *gophercloud.ServiceClient, port *ports.Port) {
	_ = "STUB: not implemented"
	return
}
