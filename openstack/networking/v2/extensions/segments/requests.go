package segments

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOpts allows filtering when listing segments.
type ListOpts struct {
	Name            string `q:"name"`
	Description     string `q:"description"`
	NetworkID       string `q:"network_id"`
	PhysicalNetwork string `q:"physical_network"`
	NetworkType     string `q:"network_type"`
	SegmentationID  int    `q:"segmentation_id"`
	RevisionNumber  int    `q:"revision_number"`
	SortDir         string `q:"sort_dir"`
	SortKey         string `q:"sort_key"`
	Fields          string `q:"fields"`
}

// ListOptsBuilder interface for listing.
type ListOptsBuilder interface {
	ToSegmentListQuery() (string, error)
}

func (opts ListOpts) ToSegmentListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List all segments.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters.
type CreateOptsBuilder interface {
	ToSegmentCreateMap() (map[string]any, error)
}

// CreateOpts contains the fields needed for creating a segment.
type CreateOpts struct {
	Name            string `json:"name,omitempty"`
	Description     string `json:"description,omitempty"`
	NetworkID       string `json:"network_id" required:"true"`
	NetworkType     string `json:"network_type" required:"true"`
	PhysicalNetwork string `json:"physical_network,omitempty"`
	SegmentationID  int    `json:"segmentation_id,omitempty"`
}

func (opts CreateOpts) ToSegmentCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create a new segment.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get retrieves a segment by ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Delete removes a segment by ID.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOpts contains fields to update a segment.
type UpdateOpts struct {
	Name           *string `json:"name,omitempty"`
	Description    *string `json:"description,omitempty"`
	SegmentationID *int    `json:"segmentation_id,omitempty"`
}

// UpdateOptsBuilder is the interface for update options.
type UpdateOptsBuilder interface {
	ToSegmentUpdateMap() (map[string]any, error)
}

func (opts UpdateOpts) ToSegmentUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update a segment.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
