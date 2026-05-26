// Package noauth contains common functions for creating block storage based
// resources for use in acceptance tests. See the `*_test.go` files for
// example usages.
package noauth

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/snapshots"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
)

// CreateVolume will create a volume with a random name and size of 1GB. An
// error will be returned if the volume was unable to be created.
func CreateVolume(t *testing.T, client *gophercloud.ServiceClient) (*volumes.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateVolumeFromImage will create a volume from with a random name and size of
// 1GB. An error will be returned if the volume was unable to be created.
func CreateVolumeFromImage(t *testing.T, client *gophercloud.ServiceClient) (*volumes.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteVolume will delete a volume. A fatal error will occur if the volume
// failed to be deleted. This works best when used as a deferred function.
func DeleteVolume(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) {
	_ = "STUB: not implemented"
	return
}

// CreateSnapshot will create a snapshot of the specified volume.
// Snapshot will be assigned a random name and description.
func CreateSnapshot(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) (*snapshots.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSnapshot will delete a snapshot. A fatal error will occur if the
// snapshot failed to be deleted.
func DeleteSnapshot(t *testing.T, client *gophercloud.ServiceClient, snapshot *snapshots.Snapshot) {
	_ = "STUB: not implemented"
	return
}

// Volumes can't be deleted until their snapshots have been,
// so block up to 120 seconds for the snapshot to delete.
