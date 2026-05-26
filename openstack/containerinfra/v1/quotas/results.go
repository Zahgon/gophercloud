package quotas

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
)

type commonResult struct {
	gophercloud.Result
}

// CreateResult is the response of a Create operations.
type CreateResult struct {
	commonResult
}

// Extract is a function that accepts a result and extracts a quota resource.
func (r commonResult) Extract() (*Quotas, error) { _ = "STUB: not implemented"; return nil, nil }

type Quotas struct {
	Resource  string    `json:"resource"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	HardLimit int       `json:"hard_limit"`
	ProjectID string    `json:"project_id"`
	ID        string    `json:"-"`
}

func (r *Quotas) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
