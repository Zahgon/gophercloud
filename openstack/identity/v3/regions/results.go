package regions

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Region helps manage related users.
type Region struct {
	// Description describes the region purpose.
	Description string `json:"description"`

	// ID is the unique ID of the region.
	ID string `json:"id"`

	// Extra is a collection of miscellaneous key/values.
	Extra map[string]any `json:"-"`

	// Links contains referencing links to the region.
	Links map[string]any `json:"links"`

	// ParentRegionID is the ID of the parent region.
	ParentRegionID string `json:"parent_region_id"`
}

func (r *Region) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Collect other fields and bundle them into Extra
// but only if a field titled "extra" wasn't sent.

type regionResult struct {
	gophercloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Region.
type GetResult struct {
	regionResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Region.
type CreateResult struct {
	regionResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a Region.
type UpdateResult struct {
	regionResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// RegionPage is a single page of Region results.
type RegionPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Regions contains any results.
func (r RegionPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r RegionPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractRegions returns a slice of Regions contained in a single page of results.
func ExtractRegions(r pagination.Page) ([]Region, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract interprets any region results as a Region.
func (r regionResult) Extract() (*Region, error) { _ = "STUB: not implemented"; return nil, nil }
