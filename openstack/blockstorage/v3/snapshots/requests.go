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

// CreateOpts contains options for creating a Snapshot. This object is passed to
// the snapshots.Create function. For more information about these parameters,
// see the Snapshot object.
type CreateOpts struct {
	VolumeID    string            `json:"volume_id" required:"true"`
	Force       bool              `json:"force,omitempty"`
	Name        string            `json:"name,omitempty"`
	Description string            `json:"description,omitempty"`
	Metadata    map[string]string `json:"metadata,omitempty"`
}

// ToSnapshotCreateMap assembles a request body based on the contents of a
// CreateOpts.
func (opts CreateOpts) ToSnapshotCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create will create a new Snapshot based on the values in CreateOpts. To
// extract the Snapshot object from the response, call the Extract method on the
// CreateResult.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete will delete the existing Snapshot with the provided ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Get retrieves the Snapshot with the provided ID. To extract the Snapshot
// object from the response, call the Extract method on the GetResult.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToSnapshotListQuery() (string, error)
}

// ListOpts holds options for listing Snapshots. It is passed to the snapshots.List
// function.
type ListOpts struct {
	// AllTenants will retrieve snapshots of all tenants/projects.
	AllTenants bool `q:"all_tenants"`

	// Name will filter by the specified snapshot name.
	Name string `q:"name"`

	// Status will filter by the specified status.
	Status string `q:"status"`

	// TenantID will filter by a specific tenant/project ID.
	// Setting AllTenants is required to use this.
	TenantID string `q:"project_id"`

	// VolumeID will filter by a specified volume ID.
	VolumeID string `q:"volume_id"`

	// Comma-separated list of sort keys and optional sort directions in the
	// form of <key>[:<direction>].
	Sort string `q:"sort"`

	// Requests a page size of items.
	Limit int `q:"limit"`

	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`

	// The ID of the last-seen item.
	Marker string `q:"marker"`
}

// ToSnapshotListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSnapshotListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns Snapshots optionally limited by the conditions provided in
// ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListDetail returns Snapshots with additional details optionally limited by the conditions provided in ListOpts.
func ListDetail(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSnapshotUpdateMap() (map[string]any, error)
}

// UpdateOpts contain options for updating an existing Snapshot. This object is passed
// to the snapshots.Update function. For more information about the parameters, see
// the Snapshot object.
type UpdateOpts struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
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

// UpdateMetadataOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateMetadataOptsBuilder interface {
	ToSnapshotUpdateMetadataMap() (map[string]any, error)
}

// UpdateMetadataOpts contain options for updating an existing Snapshot. This
// object is passed to the snapshots.Update function. For more information
// about the parameters, see the Snapshot object.
type UpdateMetadataOpts struct {
	Metadata map[string]any `json:"metadata,omitempty"`
}

// ToSnapshotUpdateMetadataMap assembles a request body based on the contents of
// an UpdateMetadataOpts.
func (opts UpdateMetadataOpts) ToSnapshotUpdateMetadataMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateMetadata will update the Snapshot with provided information. To
// extract the updated Snapshot from the response, call the ExtractMetadata
// method on the UpdateMetadataResult.
func UpdateMetadata(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateMetadataOptsBuilder) (r UpdateMetadataResult) {
	_ = "STUB: not implemented"
	return *new(UpdateMetadataResult)
}

// ResetStatusOptsBuilder allows extensions to add additional parameters to the
// ResetStatus request.
type ResetStatusOptsBuilder interface {
	ToSnapshotResetStatusMap() (map[string]any, error)
}

// ResetStatusOpts contains options for resetting a Snapshot status.
// For more information about these parameters, please, refer to the Block Storage API V3,
// Snapshot Actions, ResetStatus snapshot documentation.
type ResetStatusOpts struct {
	// Status is a snapshot status to reset to.
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

// UpdateStatusOptsBuilder allows extensions to add additional parameters to the
// UpdateStatus request.
type UpdateStatusOptsBuilder interface {
	ToSnapshotUpdateStatusMap() (map[string]any, error)
}

// UpdateStatusOpts contains options for resetting a Snapshot status.
// For more information about these parameters, please, refer to the Block Storage API V3,
// Snapshot Actions, UpdateStatus snapshot documentation.
type UpdateStatusOpts struct {
	// Status is a snapshot status to update to.
	Status string `json:"status"`
	// A progress percentage value for snapshot build progress.
	Progress string `json:"progress,omitempty"`
}

// ToSnapshotUpdateStatusMap assembles a request body based on the contents of a
// UpdateStatusOpts.
func (opts UpdateStatusOpts) ToSnapshotUpdateStatusMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateStatus will update the existing snapshot status. UpdateStatusResult contains only the error.
// To extract it, call the ExtractErr method on the UpdateStatusResult.
func UpdateStatus(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateStatusOptsBuilder) (r UpdateStatusResult) {
	_ = "STUB: not implemented"
	return *new(UpdateStatusResult)
}

// ForceDelete will delete the existing snapshot in any state. ForceDeleteResult contains only the error.
// To extract it, call the ExtractErr method on the ForceDeleteResult.
func ForceDelete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r ForceDeleteResult) {
	_ = "STUB: not implemented"
	return *new(ForceDeleteResult)
}
