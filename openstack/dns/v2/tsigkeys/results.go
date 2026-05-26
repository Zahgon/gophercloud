package tsigkeys

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract interprets a GetResult, CreateResult or UpdateResult as a TSIGKey.
// An error is returned if the original call or the extraction failed.
func (r commonResult) Extract() (*TSIGKey, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateResult is the result of a Create request. Call its Extract method
// to interpret the result as a TSIGKey.
type CreateResult struct {
	commonResult
}

// GetResult is the result of a Get request. Call its Extract method
// to interpret the result as a TSIGKey.
type GetResult struct {
	commonResult
}

// UpdateResult is the result of an Update request. Call its Extract method
// to interpret the result as a TSIGKey.
type UpdateResult struct {
	commonResult
}

// DeleteResult is the result of a Delete request. Call its ExtractErr method
// to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// TSIGKeyPage is a single page of TSIGKey results.
type TSIGKeyPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if the page contains no results.
func (r TSIGKeyPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractTSIGKeys extracts a slice of TSIGKeys from a List result.
func ExtractTSIGKeys(r pagination.Page) ([]TSIGKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TSIGKey represents a TSIG key for DNS transaction authentication.
type TSIGKey struct {
	// ID uniquely identifies this TSIG key.
	ID string `json:"id"`

	// Name is the name of the TSIG key.
	Name string `json:"name"`

	// Algorithm is the TSIG algorithm used (e.g., hmac-sha256, hmac-sha512).
	Algorithm string `json:"algorithm"`

	// Secret is the base64-encoded secret key.
	Secret string `json:"secret"`

	// Scope defines the scope of the TSIG key (ZONE or POOL).
	Scope string `json:"scope"`

	// ResourceID is the ID of the resource (zone or pool) this key is associated with.
	ResourceID string `json:"resource_id"`

	// CreatedAt is the date when the TSIG key was created.
	CreatedAt time.Time `json:"-"`

	// UpdatedAt is the date when the TSIG key was last updated.
	UpdatedAt time.Time `json:"-"`

	// Links includes HTTP references to the itself.
	Links map[string]any `json:"links"`
}

func (r *TSIGKey) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }
