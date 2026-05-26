package attachinterfaces

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List makes a request against the nova API to list the server's interfaces.
func List(client *gophercloud.ServiceClient, serverID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get requests details on a single interface attachment by the server and port IDs.
func Get(ctx context.Context, client *gophercloud.ServiceClient, serverID, portID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAttachInterfacesCreateMap() (map[string]any, error)
}

// CreateOpts specifies parameters of a new interface attachment.
type CreateOpts struct {
	// PortID is the ID of the port for which you want to create an interface.
	// The NetworkID and PortID parameters are mutually exclusive.
	// If you do not specify the PortID parameter, the OpenStack Networking API
	// v2.0 allocates a port and creates an interface for it on the network.
	PortID string `json:"port_id,omitempty"`

	// NetworkID is the ID of the network for which you want to create an interface.
	// The NetworkID and PortID parameters are mutually exclusive.
	// If you do not specify the NetworkID parameter, the OpenStack Networking
	// API v2.0 uses the network information cache that is associated with the instance.
	NetworkID string `json:"net_id,omitempty"`

	// Slice of FixedIPs. If you request a specific FixedIP address without a
	// NetworkID, the request returns a Bad Request (400) response code.
	// Note: this uses the FixedIP struct, but only the IPAddress field can be used.
	FixedIPs []FixedIP `json:"fixed_ips,omitempty"`
}

// ToAttachInterfacesCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToAttachInterfacesCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new interface attachment on the server.
func Create(ctx context.Context, client *gophercloud.ServiceClient, serverID string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete makes a request against the nova API to detach a single interface from the server.
// It needs server and port IDs to make a such request.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, serverID, portID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
