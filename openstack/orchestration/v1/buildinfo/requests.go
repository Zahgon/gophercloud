package buildinfo

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get retreives data for the given stack template.
func Get(ctx context.Context, c *gophercloud.ServiceClient) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
