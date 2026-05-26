package shareaccessrules

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get retrieves details about a share access rule.
func Get(ctx context.Context, client *gophercloud.ServiceClient, accessID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// List gets all access rules of a share.
func List(ctx context.Context, client *gophercloud.ServiceClient, shareID string) (r ListResult) {
	_ = "STUB: not implemented"
	return *new(ListResult)
}
