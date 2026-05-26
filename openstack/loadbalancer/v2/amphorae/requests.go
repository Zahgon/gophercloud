package amphorae

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToAmphoraListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Amphorae attributes you want to see returned. SortKey allows you to
// sort by a particular attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	LoadbalancerID string `q:"loadbalancer_id"`
	ImageID        string `q:"image_id"`
	Role           string `q:"role"`
	Status         string `q:"status"`
	HAPortID       string `q:"ha_port_id"`
	VRRPPortID     string `q:"vrrp_port_id"`
	Limit          int    `q:"limit"`
	Marker         string `q:"marker"`
	SortKey        string `q:"sort_key"`
	SortDir        string `q:"sort_dir"`
}

// ToAmphoraListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToAmphoraListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// amphorae. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves a particular amphora based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Failover performs a failover of an amphora.
func Failover(ctx context.Context, c *gophercloud.ServiceClient, id string) (r FailoverResult) {
	_ = "STUB: not implemented"
	return *new(FailoverResult)
}
