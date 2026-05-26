package accept

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add parameters to the List request.
type ListOptsBuilder interface {
	ToTransferAcceptListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the server attributes you want to see returned.
// https://developer.openstack.org/api-ref/dns/
type ListOpts struct {
	Status string `q:"status"`
}

// ToTransferAcceptListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTransferAcceptListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List implements a transfer accept List request.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get returns information about a transfer accept, given its ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, transferAcceptID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional attributes to the
// Create request.
type CreateOptsBuilder interface {
	ToTransferAcceptCreateMap() (map[string]any, error)
}

// CreateOpts specifies the attributes used to create a transfer accept.
type CreateOpts struct {
	// Key is used as part of the zone transfer accept process.
	// This is only shown to the creator, and must be communicated out of band.
	Key string `json:"key" required:"true"`

	// ZoneTransferRequestID is ID for this zone transfer request
	ZoneTransferRequestID string `json:"zone_transfer_request_id" required:"true"`
}

// ToTransferAcceptCreateMap formats an CreateOpts structure into a request body.
func (opts CreateOpts) ToTransferAcceptCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create implements a transfer accept create request.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}
