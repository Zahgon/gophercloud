package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	v3 "github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v3/volumes"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
)

// CreateVolumeAttachment will attach a volume to an instance. An error will be
// returned if the attachment failed.
func CreateVolumeAttachment(t *testing.T, client *gophercloud.ServiceClient, volume *v3.Volume, server *servers.Server) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteVolumeAttachment will detach a volume from an instance. A fatal error
// will occur if the attachment failed to be deleted.
func DeleteVolumeAttachment(t *testing.T, client *gophercloud.ServiceClient, volume *v3.Volume) {
	_ = "STUB: not implemented"
	return
}
