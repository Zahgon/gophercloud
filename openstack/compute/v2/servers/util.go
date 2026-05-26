package servers

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// WaitForStatus will continually poll a server until it successfully
// transitions to a specified status.
func WaitForStatus(ctx context.Context, c *gophercloud.ServiceClient, id, status string) error {
	_ = "STUB: not implemented"
	return nil
}
