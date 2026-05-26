package volumeattach

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// VolumeAttachment contains attachment information between a volume
// and server.
type VolumeAttachment struct {
	// ID is a unique id of the attachment.
	ID string `json:"id"`

	// Device is what device the volume is attached as.
	Device string `json:"device"`

	// VolumeID is the ID of the attached volume.
	VolumeID string `json:"volumeId"`

	// ServerID is the ID of the instance that has the volume attached.
	ServerID string `json:"serverId"`

	// Tag is a device role tag that can be applied to a volume when attaching
	// it to the VM. Requires 2.70 microversion
	Tag *string `json:"tag"`

	// DeleteOnTermination specifies whether or not to delete the volume when the server
	// is destroyed. Requires 2.79 microversion
	DeleteOnTermination *bool `json:"delete_on_termination"`
}

// VolumeAttachmentPage stores a single page all of VolumeAttachment
// results from a List call.
type VolumeAttachmentPage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a VolumeAttachmentPage is empty.
func (page VolumeAttachmentPage) IsEmpty() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ExtractVolumeAttachments interprets a page of results as a slice of
// VolumeAttachment.
func ExtractVolumeAttachments(r pagination.Page) ([]VolumeAttachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VolumeAttachmentResult is the result from a volume attachment operation.
type VolumeAttachmentResult struct {
	gophercloud.Result
}

// Extract is a method that attempts to interpret any VolumeAttachment resource
// response as a VolumeAttachment struct.
func (r VolumeAttachmentResult) Extract() (*VolumeAttachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a VolumeAttachment.
type CreateResult struct {
	VolumeAttachmentResult
}

// GetResult is the response from a Get operation. Call its Extract method to
// interpret it as a VolumeAttachment.
type GetResult struct {
	VolumeAttachmentResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}
