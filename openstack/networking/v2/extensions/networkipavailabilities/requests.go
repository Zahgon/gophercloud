package networkipavailabilities

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToNetworkIPAvailabilityListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the Neutron API.
type ListOpts struct {
	// NetworkName allows to filter on the identifier of a network.
	NetworkID string `q:"network_id"`

	// NetworkName allows to filter on the name of a network.
	NetworkName string `q:"network_name"`

	// IPVersion allows to filter on the version of the IP protocol.
	// You can use the well-known IP versions with the gophercloud.IPVersion type.
	IPVersion string `q:"ip_version"`

	// ProjectID allows to filter on the Identity project field.
	ProjectID string `q:"project_id"`

	// TenantID allows to filter on the Identity project field.
	TenantID string `q:"tenant_id"`
}

// ToNetworkIPAvailabilityListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToNetworkIPAvailabilityListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// networkipavailabilities. It accepts a ListOpts struct, which allows you to
// filter the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves a specific NetworkIPAvailability based on its ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
