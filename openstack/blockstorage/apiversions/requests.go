package apiversions

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List lists all the Cinder API versions available to end-users.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
