package quotas

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get retrieves the details of quotas for a specified tenant.
func Get(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
