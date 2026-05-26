package snapshots

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSnapshotCreateMap() (map[string]any, error)
}

// CreateOpts contains the options for create a Snapshot. This object is
// passed to snapshots.Create(). For more information about these parameters,
// please refer to the Snapshot object, or the shared file systems API v2
// documentation
type CreateOpts struct {
	// The UUID of the share from which to create a snapshot
	ShareID string `json:"share_id" required:"true"`
	// Defines the snapshot name
	Name string `json:"name,omitempty"`
	// Defines the snapshot description
	Description string `json:"description,omitempty"`
	// DisplayName is equivalent to Name. The API supports using both
	// This is an inherited attribute from the block storage API
	DisplayName string `json:"display_name,omitempty"`
	// DisplayDescription is equivalent to Description. The API supports using both
	// This is an inherited attribute from the block storage API
	DisplayDescription string `json:"display_description,omitempty"`
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToSnapshotCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create will create a new Snapshot based on the values in CreateOpts. To extract
// the Snapshot object from the response, call the Extract method on the
// CreateResult.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// ListOpts holds options for listing Snapshots. It is passed to the
// snapshots.List function.
type ListOpts struct {
	// (Admin only). Defines whether to list the requested resources for all projects.
	AllTenants bool `q:"all_tenants"`
	// The snapshot name.
	Name string `q:"name"`
	// Filter  by a snapshot description.
	Description string `q:"description"`
	// Filters by a share from which the snapshot was created.
	ShareID string `q:"share_id"`
	// Filters by a snapshot size in GB.
	Size int `q:"size"`
	// Filters by a snapshot status.
	Status string `q:"status"`
	// The maximum number of snapshots to return.
	Limit int `q:"limit"`
	// The offset to define start point of snapshot or snapshot group listing.
	Offset int `q:"offset"`
	// The key to sort a list of snapshots.
	SortKey string `q:"sort_key"`
	// The direction to sort a list of snapshots.
	SortDir string `q:"sort_dir"`
	// The UUID of the project in which the snapshot was created. Useful with all_tenants parameter.
	ProjectID string `q:"project_id"`
	// The name pattern that can be used to filter snapshots, snapshot snapshots, snapshot networks or snapshot groups.
	NamePattern string `q:"name~"`
	// The description pattern that can be used to filter snapshots, snapshot snapshots, snapshot networks or snapshot groups.
	DescriptionPattern string `q:"description~"`
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToSnapshotListQuery() (string, error)
}

// ToSnapshotListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSnapshotListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ListDetail returns []Snapshot optionally limited by the conditions provided in ListOpts.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Delete will delete an existing Snapshot with the given UUID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Get will get a single snapshot with given UUID
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSnapshotUpdateMap() (map[string]any, error)
}

// UpdateOpts contain options for updating an existing Snapshot. This object is passed
// to the snapshot.Update function. For more information about the parameters, see
// the Snapshot object.
type UpdateOpts struct {
	// Snapshot name. Manila snapshot update logic doesn't have a "name" alias.
	DisplayName *string `json:"display_name,omitempty"`
	// Snapshot description. Manila snapshot update logic doesn't have a "description" alias.
	DisplayDescription *string `json:"display_description,omitempty"`
}

// ToSnapshotUpdateMap assembles a request body based on the contents of an
// UpdateOpts.
func (opts UpdateOpts) ToSnapshotUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update will update the Snapshot with provided information. To extract the updated
// Snapshot from the response, call the Extract method on the UpdateResult.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// ResetStatusOptsBuilder allows extensions to add additional parameters to the
// ResetStatus request.
type ResetStatusOptsBuilder interface {
	ToSnapshotResetStatusMap() (map[string]any, error)
}

// ResetStatusOpts contains options for resetting a Snapshot status.
// For more information about these parameters, please, refer to the shared file systems API v2,
// Snapshot Actions, ResetStatus share documentation.
type ResetStatusOpts struct {
	// Status is a snapshot status to reset to. Can be "available", "error",
	// "creating", "deleting", "manage_starting", "manage_error",
	// "unmanage_starting", "unmanage_error" or "error_deleting".
	Status string `json:"status"`
}

// ToSnapshotResetStatusMap assembles a request body based on the contents of a
// ResetStatusOpts.
func (opts ResetStatusOpts) ToSnapshotResetStatusMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResetStatus will reset the existing snapshot status. ResetStatusResult contains only the error.
// To extract it, call the ExtractErr method on the ResetStatusResult.
func ResetStatus(ctx context.Context, client *gophercloud.ServiceClient, id string, opts ResetStatusOptsBuilder) (r ResetStatusResult) {
	_ = "STUB: not implemented"
	return *new(ResetStatusResult)
}

// ForceDelete will delete the existing snapshot in any state. ForceDeleteResult contains only the error.
// To extract it, call the ExtractErr method on the ForceDeleteResult.
func ForceDelete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r ForceDeleteResult) {
	_ = "STUB: not implemented"
	return *new(ForceDeleteResult)
}
