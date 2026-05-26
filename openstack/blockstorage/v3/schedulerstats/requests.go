package schedulerstats

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToStoragePoolsListQuery() (string, error)
}

// ListOpts controls the view of data returned (e.g globally or per project)
// via tenant_id and the verbosity via detail.
type ListOpts struct {
	// ID of the tenant to look up storage pools for.
	TenantID string `q:"tenant_id"`

	// Whether to list extended details.
	Detail bool `q:"detail"`
}

// ToStoragePoolsListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToStoragePoolsListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request against the API to list storage pool information.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
