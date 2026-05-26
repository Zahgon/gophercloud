package extensions

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Get retrieves information for a specific extension using its alias.
func Get(ctx context.Context, c *gophercloud.ServiceClient, alias string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
