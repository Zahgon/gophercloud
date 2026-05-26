package snapshots

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// WaitForStatus will continually poll the resource, checking for a particular status.
func WaitForStatus(ctx context.Context, c *gophercloud.ServiceClient, id, status string) error {
	_ = "STUB: not implemented"
	return nil
}
