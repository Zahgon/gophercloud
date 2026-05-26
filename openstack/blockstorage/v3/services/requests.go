package services

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToServiceListQuery() (string, error)
}

// ListOpts holds options for listing Services.
type ListOpts struct {
	// Filter the service list result by binary name of the service.
	Binary string `q:"binary"`

	// Filter the service list result by host name of the service.
	Host string `q:"host"`
}

// ToServiceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToServiceListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request against the API to list services.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
