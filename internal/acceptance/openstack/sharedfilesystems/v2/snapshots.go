package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/snapshots"
)

// CreateSnapshot will create a snapshot from the share ID with a name. An error will
// be returned if the snapshot could not be created
func CreateSnapshot(t *testing.T, client *gophercloud.ServiceClient, shareID string) (*snapshots.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ListSnapshots lists all snapshots that belong to this tenant's project.
// An error will be returned if the snapshots could not be listed..
func ListSnapshots(t *testing.T, client *gophercloud.ServiceClient) ([]snapshots.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSnapshot will delete a snapshot. A fatal error will occur if the snapshot
// failed to be deleted. This works best when used as a deferred function.
func DeleteSnapshot(t *testing.T, client *gophercloud.ServiceClient, snapshot *snapshots.Snapshot) {
	_ = "STUB: not implemented"
	return
}

func waitForSnapshotStatus(t *testing.T, c *gophercloud.ServiceClient, id, status string) error {
	_ = "STUB: not implemented"
	return nil
}
