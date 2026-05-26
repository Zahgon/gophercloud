package snapshots

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Snapshot contains all the information associated with a Cinder Snapshot.
type Snapshot struct {
	// Unique identifier.
	ID string `json:"id"`

	// Date created.
	CreatedAt time.Time `json:"-"`

	// Date updated.
	UpdatedAt time.Time `json:"-"`

	// Display name.
	Name string `json:"name"`

	// Display description.
	Description string `json:"description"`

	// ID of the Volume from which this Snapshot was created.
	VolumeID string `json:"volume_id"`

	// Currect status of the Snapshot.
	Status string `json:"status"`

	// Size of the Snapshot, in GB.
	Size int `json:"size"`

	// User-defined key-value pairs.
	Metadata map[string]string `json:"metadata"`

	// Progress of the snapshot creation.
	Progress string `json:"os-extended-snapshot-attributes:progress"`

	// Project ID that owns the snapshot.
	ProjectID string `json:"os-extended-snapshot-attributes:project_id"`

	// ID of the group snapshot, if applicable.
	GroupSnapshotID string `json:"group_snapshot_id"`

	// User ID that created the snapshot.
	UserID string `json:"user_id"`

	// Indicates whether the snapshot consumes quota.
	ConsumesQuota bool `json:"consumes_quota"`
}

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

// SnapshotPage is a pagination.Pager that is returned from a call to the List function.
type SnapshotPage struct {
	pagination.LinkedPageBase
}

// UnmarshalJSON converts our JSON API response into our snapshot struct
func (r *Snapshot) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// IsEmpty returns true if a SnapshotPage contains no Snapshots.
func (r SnapshotPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r SnapshotPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractSnapshots extracts and returns Snapshots. It is used while iterating over a snapshots.List call.
func ExtractSnapshots(r pagination.Page) ([]Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateMetadataResult contains the response body and error from an UpdateMetadata request.
type UpdateMetadataResult struct {
	commonResult
}

// ExtractMetadata returns the metadata from a response from snapshots.UpdateMetadata.
func (r UpdateMetadataResult) ExtractMetadata() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Snapshot object out of the commonResult object.
func (r commonResult) Extract() (*Snapshot, error) { _ = "STUB: not implemented"; return nil, nil }

// ResetStatusResult contains the response error from a ResetStatus request.
type ResetStatusResult struct {
	gophercloud.ErrResult
}

// UpdateStatusResult contains the response error from an UpdateStatus request.
type UpdateStatusResult struct {
	gophercloud.ErrResult
}

// ForceDeleteResult contains the response error from a ForceDelete request.
type ForceDeleteResult struct {
	gophercloud.ErrResult
}
