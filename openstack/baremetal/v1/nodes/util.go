package nodes

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// WaitForProvisionState will continually poll a node until it successfully
// transitions to a specified state. It will do this for at most the number
// of seconds specified.
func WaitForProvisionState(ctx context.Context, c *gophercloud.ServiceClient, id string, state ProvisionState) error {
	_ = "STUB: not implemented"
	return nil
}
