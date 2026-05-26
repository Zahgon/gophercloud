package quotasets

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get returns public data about a previously created QuotaSet.
func Get(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// GetDefaults returns public data about the project's default block storage quotas.
func GetDefaults(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// GetUsage returns detailed public data about a previously created QuotaSet.
func GetUsage(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r GetUsageResult) {
	_ = "STUB: not implemented"
	return *new(GetUsageResult)
}

// Updates the quotas for the given projectID and returns the new QuotaSet.
func Update(ctx context.Context, client *gophercloud.ServiceClient, projectID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// UpdateOptsBuilder enables extensions to add parameters to the update request.
type UpdateOptsBuilder interface {
	// Extra specific name to prevent collisions with interfaces for other quotas
	// (e.g. neutron)
	ToBlockStorageQuotaUpdateMap() (map[string]any, error)
}

// ToBlockStorageQuotaUpdateMap builds the update options into a serializable
// format.
func (opts UpdateOpts) ToBlockStorageQuotaUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Options for Updating the quotas of a Tenant.
// All int-values are pointers so they can be nil if they are not needed.
// You can use gopercloud.IntToPointer() for convenience
type UpdateOpts struct {
	// Volumes is the number of volumes that are allowed for each project.
	Volumes *int `json:"volumes,omitempty"`

	// Snapshots is the number of snapshots that are allowed for each project.
	Snapshots *int `json:"snapshots,omitempty"`

	// Gigabytes is the size (GB) of volumes and snapshots that are allowed for
	// each project.
	Gigabytes *int `json:"gigabytes,omitempty"`

	// PerVolumeGigabytes is the size (GB) of volumes and snapshots that are
	// allowed for each project and the specifed volume type.
	PerVolumeGigabytes *int `json:"per_volume_gigabytes,omitempty"`

	// Backups is the number of backups that are allowed for each project.
	Backups *int `json:"backups,omitempty"`

	// BackupGigabytes is the size (GB) of backups that are allowed for each
	// project.
	BackupGigabytes *int `json:"backup_gigabytes,omitempty"`

	// Groups is the number of groups that are allowed for each project.
	Groups *int `json:"groups,omitempty"`

	// Force will update the quotaset even if the quota has already been used
	// and the reserved quota exceeds the new quota.
	Force bool `json:"force,omitempty"`

	// Extra is a collection of miscellaneous key/values used to set
	// quota per volume_type
	Extra map[string]any `json:"-"`
}

// Resets the quotas for the given tenant to their default values.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
