package peers

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List the bgp peers
func List(c *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieve the specific bgp peer by its uuid
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToPeerCreateMap() (map[string]any, error)
}

// CreateOpts represents options used to create a BGP Peer.
type CreateOpts struct {
	AuthType string `json:"auth_type"`
	RemoteAS int    `json:"remote_as"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	PeerIP   string `json:"peer_ip"`
	TenantID string `json:"tenant_id,omitempty"`
}

// ToPeerCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToPeerCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a BGP Peer
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete accepts a unique ID and deletes the bgp Peer associated with it.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, bgpPeerID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToPeerUpdateMap() (map[string]any, error)
}

// UpdateOpts represents options used to update a BGP Peer.
type UpdateOpts struct {
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
}

// ToPeerUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToPeerUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update accept a BGP Peer ID and an UpdateOpts and update the BGP Peer
func Update(ctx context.Context, c *gophercloud.ServiceClient, bgpPeerID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
