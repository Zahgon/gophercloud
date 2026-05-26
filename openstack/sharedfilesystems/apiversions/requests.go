package apiversions

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List lists all the API versions available to end-users.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get will get a specific API version, specified by major ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, v string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
