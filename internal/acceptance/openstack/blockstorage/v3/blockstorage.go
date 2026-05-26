// Package v3 contains common functions for creating block storage based
// resources for use in acceptance tests. See the `*_test.go` files for
// example usages.
package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/backups"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/qos"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/snapshots"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumetypes"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
)

// CreateSnapshot will create a snapshot of the specified volume.
// Snapshot will be assigned a random name and description.
func CreateSnapshot(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) (*snapshots.Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateVolume will create a volume with a random name and size of 1GB. An
// error will be returned if the volume was unable to be created.
func CreateVolume(t *testing.T, client *gophercloud.ServiceClient) (*volumes.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateVolumeWithType will create a volume of the given volume type
// with a random name and size of 1GB. An error will be returned if
// the volume was unable to be created.
func CreateVolumeWithType(t *testing.T, client *gophercloud.ServiceClient, vt *volumetypes.VolumeType) (*volumes.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateVolumeType will create a volume type with a random name. An
// error will be returned if the volume was unable to be created.
func CreateVolumeType(t *testing.T, client *gophercloud.ServiceClient) (*volumetypes.VolumeType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO: For some reason returned extra_specs are empty even in API reference: https://developer.openstack.org/api-ref/block-storage/v3/?expanded=create-a-volume-type-detail#volume-types-types
// "extra_specs": {}
// th.AssertEquals(t, vt.ExtraSpecs, createOpts.ExtraSpecs)

// CreateVolumeTypeNoExtraSpecs will create a volume type with a random name and
// no extra specs. This is required to bypass cinder-scheduler filters and be able
// to create a volume with this volumeType. An error will be returned if the volume
// type was unable to be created.
func CreateVolumeTypeNoExtraSpecs(t *testing.T, client *gophercloud.ServiceClient) (*volumetypes.VolumeType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateVolumeTypeMultiAttach will create a volume type with a random name and
// extra specs for multi-attach. An error will be returned if the volume type was
// unable to be created.
func CreateVolumeTypeMultiAttach(t *testing.T, client *gophercloud.ServiceClient) (*volumetypes.VolumeType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePrivateVolumeType will create a private volume type with a random
// name and no extra specs. An error will be returned if the volume type was
// unable to be created.
func CreatePrivateVolumeType(t *testing.T, client *gophercloud.ServiceClient) (*volumetypes.VolumeType, error) {
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
// so block until the snapshoth as been deleted.

// DeleteVolume will delete a volume. A fatal error will occur if the volume
// failed to be deleted. This works best when used as a deferred function.
func DeleteVolume(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) {
	_ = "STUB: not implemented"
	return
}

// VolumeTypes can't be deleted until their volumes have been,
// so block until the volume is deleted.

// DeleteVolumeType will delete a volume type. A fatal error will occur if the
// volume type failed to be deleted. This works best when used as a deferred
// function.
func DeleteVolumeType(t *testing.T, client *gophercloud.ServiceClient, vt *volumetypes.VolumeType) {
	_ = "STUB: not implemented"
	return
}

// CreateQoS will create a QoS with one spec and a random name. An
// error will be returned if the volume was unable to be created.
func CreateQoS(t *testing.T, client *gophercloud.ServiceClient) (*qos.QoS, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteQoS will delete a QoS. A fatal error will occur if the QoS
// failed to be deleted. This works best when used as a deferred function.
func DeleteQoS(t *testing.T, client *gophercloud.ServiceClient, qs *qos.QoS) {
	_ = "STUB: not implemented"
	return
}

// CreateBackup will create a backup based on a volume. An error will be
// will be returned if the backup could not be created.
func CreateBackup(t *testing.T, client *gophercloud.ServiceClient, volumeID string) (*backups.Backup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteBackup will delete a backup. A fatal error will occur if the backup
// could not be deleted. This works best when used as a deferred function.
func DeleteBackup(t *testing.T, client *gophercloud.ServiceClient, backupID string) {
	_ = "STUB: not implemented"
	return
}

// WaitForBackupStatus will continually poll a backup, checking for a particular
// status. It will do this for the amount of seconds defined.
func WaitForBackupStatus(client *gophercloud.ServiceClient, id, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// ResetBackupStatus will reset the status of a backup.
func ResetBackupStatus(t *testing.T, client *gophercloud.ServiceClient, backup *backups.Backup, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateUploadImage will upload volume it as volume-baked image. An name of new image or err will be
// returned
func CreateUploadImage(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) (volumes.VolumeImage, error) {
	_ = "STUB: not implemented"
	return *new(volumes.VolumeImage), nil
}

// DeleteUploadedImage deletes uploaded image. An error will be returned
// if the deletion request failed.
func DeleteUploadedImage(t *testing.T, client *gophercloud.ServiceClient, imageID string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateVolumeAttach will attach a volume to an instance. An error will be
// returned if the attachment failed.
func CreateVolumeAttach(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume, server *servers.Server) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateVolumeReserve creates a volume reservation. An error will be returned
// if the reservation failed.
func CreateVolumeReserve(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteVolumeAttach will detach a volume from an instance. A fatal error will
// occur if the snapshot failed to be deleted. This works best when used as a
// deferred function.
func DeleteVolumeAttach(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) {
	_ = "STUB: not implemented"
	return
}

// DeleteVolumeReserve deletes a volume reservation. A fatal error will occur
// if the deletion request failed. This works best when used as a deferred
// function.
func DeleteVolumeReserve(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) {
	_ = "STUB: not implemented"
	return
}

// ExtendVolumeSize will extend the size of a volume.
func ExtendVolumeSize(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) error {
	_ = "STUB: not implemented"
	return nil
}

// SetImageMetadata will apply the metadata to a volume.
func SetImageMetadata(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) error {
	_ = "STUB: not implemented"
	return nil
}

// SetBootable will set a bootable status to a volume.
func SetBootable(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) error {
	_ = "STUB: not implemented"
	return nil
}

// ChangeVolumeType will extend the size of a volume.
func ChangeVolumeType(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume, vt *volumetypes.VolumeType) error {
	_ = "STUB: not implemented"
	return nil
}

// ResetVolumeStatus will reset the status of a volume.
func ResetVolumeStatus(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// ReImage will re-image a volume
func ReImage(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume, imageID string) error {
	_ = "STUB: not implemented"
	return nil
}

func Unmanage(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) error {
	_ = "STUB: not implemented"
	return nil
}

func ManageExisting(t *testing.T, client *gophercloud.ServiceClient, volume *volumes.Volume) (*volumes.Volume, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
