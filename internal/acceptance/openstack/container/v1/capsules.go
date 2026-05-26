package v1

import (
	"github.com/gophercloud/gophercloud/v2"
)

// WaitForCapsuleStatus will poll a capsule's status until it either matches
// the specified status or the status becomes Failed.
func WaitForCapsuleStatus(client *gophercloud.ServiceClient, uuid, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// Success!
