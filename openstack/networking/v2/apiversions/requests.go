package apiversions

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListVersions lists all the Neutron API versions available to end-users.
func ListVersions(c *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListVersionResources lists all of the different API resources for a
// particular API versions. Typical resources for Neutron might be: networks,
// subnets, etc.
func ListVersionResources(c *gophercloud.ServiceClient, v string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
