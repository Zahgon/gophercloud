package volumeattach

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List returns a Pager that allows you to iterate over a collection of
// VolumeAttachments.
func List(client *gophercloud.ServiceClient, serverID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add parameters to the Create request.
type CreateOptsBuilder interface {
	ToVolumeAttachmentCreateMap() (map[string]any, error)
}

// CreateOpts specifies volume attachment creation or import parameters.
type CreateOpts struct {
	// Device is the device that the volume will attach to the instance as.
	// Omit for "auto".
	Device string `json:"device,omitempty"`

	// VolumeID is the ID of the volume to attach to the instance.
	VolumeID string `json:"volumeId" required:"true"`

	// Tag is a device role tag that can be applied to a volume when attaching
	// it to the VM. Requires 2.49 microversion
	Tag string `json:"tag,omitempty"`

	// DeleteOnTermination specifies whether or not to delete the volume when the server
	// is destroyed. Requires 2.79 microversion
	DeleteOnTermination bool `json:"delete_on_termination,omitempty"`
}

// ToVolumeAttachmentCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToVolumeAttachmentCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new volume attachment on the server.
func Create(ctx context.Context, client *gophercloud.ServiceClient, serverID string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get returns public data about a previously created VolumeAttachment.
func Get(ctx context.Context, client *gophercloud.ServiceClient, serverID, volumeID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Delete requests the deletion of a previous stored VolumeAttachment from
// the server.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, serverID, volumeID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
