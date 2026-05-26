package limits

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get returns the limits about the currently scoped tenant.
func Get(ctx context.Context, client *gophercloud.ServiceClient) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
