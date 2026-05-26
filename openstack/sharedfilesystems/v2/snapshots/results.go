package snapshots

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

const (
	invalidMarker = "-1"
)

// Snapshot contains all information associated with an OpenStack Snapshot
type Snapshot struct {
	// The UUID of the snapshot
	ID string `json:"id"`
	// The name of the snapshot
	Name string `json:"name,omitempty"`
	// A description of the snapshot
	Description string `json:"description,omitempty"`
	// UUID of the share from which the snapshot was created
	ShareID string `json:"share_id"`
	// The shared file system protocol
	ShareProto string `json:"share_proto"`
	// Size of the snapshot share in GB
	ShareSize int `json:"share_size"`
	// Size of the snapshot in GB
	Size int `json:"size"`
	// The snapshot status
	Status string `json:"status"`
	// The UUID of the project in which the snapshot was created
	ProjectID string `json:"project_id"`
	// Timestamp when the snapshot was created
	CreatedAt time.Time `json:"-"`
	// Snapshot links for pagination
	Links []map[string]string `json:"links"`
}

func (r *Snapshot) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type commonResult struct {
	gophercloud.Result
}

// Extract will get the Snapshot object from the commonResult
func (r commonResult) Extract() (*Snapshot, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// SnapshotPage is a pagination.pager that is returned from a call to the List function.
type SnapshotPage struct {
	pagination.MarkerPageBase
}

// NextPageURL generates the URL for the page of results after this one.
func (r SnapshotPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LastMarker returns the last offset in a ListResult.
func (r SnapshotPage) LastMarker() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Limit is not present, only one page required

// IsEmpty satisifies the IsEmpty method of the Page interface
func (r SnapshotPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractSnapshots extracts and returns a Snapshot slice. It is used while
// iterating over a snapshots.List call.
func ExtractSnapshots(r pagination.Page) ([]Snapshot, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

// ResetStatusResult contains the response error from an ResetStatus request.
type ResetStatusResult struct {
	gophercloud.ErrResult
}

// ForceDeleteResult contains the response error from an ForceDelete request.
type ForceDeleteResult struct {
	gophercloud.ErrResult
}
