package regions

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToRegionListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// ParentRegionID filters the response by a parent region ID.
	ParentRegionID string `q:"parent_region_id"`
}

// ToRegionListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRegionListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List enumerates the Regions to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single region, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToRegionCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a region.
type CreateOpts struct {
	// ID is the ID of the new region.
	ID string `json:"id,omitempty"`

	// Description is a description of the region.
	Description string `json:"description,omitempty"`

	// ParentRegionID is the ID of the parent the region to add this region under.
	ParentRegionID string `json:"parent_region_id,omitempty"`

	// Extra is free-form extra key/value pairs to describe the region.
	Extra map[string]any `json:"-"`
}

// ToRegionCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToRegionCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new Region.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToRegionUpdateMap() (map[string]any, error)
}

// UpdateOpts provides options for updating a region.
type UpdateOpts struct {
	// Description is a description of the region.
	Description *string `json:"description,omitempty"`

	// ParentRegionID is the ID of the parent region.
	ParentRegionID string `json:"parent_region_id,omitempty"`

	/*
		// Due to a bug in Keystone, the Extra column of the Region table
		// is not updatable, see: https://bugs.launchpad.net/keystone/+bug/1729933
		// The following lines should be uncommented once the fix is merged.

		// Extra is free-form extra key/value pairs to describe the region.
		Extra map[string]any `json:"-"`
	*/
}

// ToRegionUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToRegionUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/*
	// Due to a bug in Keystone, the Extra column of the Region table
	// is not updatable, see: https://bugs.launchpad.net/keystone/+bug/1729933
	// The following lines should be uncommented once the fix is merged.

	if opts.Extra != nil {
		if v, ok := b["region"].(map[string]any); ok {
			for key, value := range opts.Extra {
				v[key] = value
			}
		}
	}
*/

// Update updates an existing Region.
func Update(ctx context.Context, client *gophercloud.ServiceClient, regionID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete deletes a region.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, regionID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
