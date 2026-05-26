package extensions

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	common "github.com/gophercloud/gophercloud/v2/openstack/common/extensions"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ExtractExtensions interprets a Page as a slice of Extensions.
func ExtractExtensions(page pagination.Page) ([]common.Extension, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get retrieves information for a specific extension using its alias.
func Get(ctx context.Context, c *gophercloud.ServiceClient, alias string) common.GetResult {
	_ = "STUB: not implemented"
	return *new(common.GetResult)
}

// List returns a Pager which allows you to iterate over the full collection of extensions.
// It does not accept query parameters.
func List(c *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
