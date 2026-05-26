package providers

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToProviderListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Provider attributes you want to see returned.
type ListOpts struct {
	Fields []string `q:"fields"`
}

// ToProviderListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToProviderListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// providers.
//
// Default policy settings return only those providers that are owned by
// the project who submits the request, unless an admin user submits the request.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
