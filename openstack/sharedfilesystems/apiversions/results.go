package apiversions

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// APIVersion represents an API version for the Shared File System service.
type APIVersion struct {
	// ID is the unique identifier of the API version.
	ID string `json:"id"`

	// MinVersion is the minimum microversion supported.
	MinVersion string `json:"min_version"`

	// Status is the API versions status.
	Status string `json:"status"`

	// Updated is the date when the API was last updated.
	Updated time.Time `json:"updated"`

	// Version is the maximum microversion supported.
	Version string `json:"version"`
}

// APIVersionPage is the page returned by a pager when traversing over a
// collection of API versions.
type APIVersionPage struct {
	pagination.SinglePageBase
}

// IsEmpty checks whether an APIVersionPage struct is empty.
func (r APIVersionPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractAPIVersions takes a collection page, extracts all of the elements,
// and returns them a slice of APIVersion structs. It is effectively a cast.
func ExtractAPIVersions(r pagination.Page) ([]APIVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetResult represents the result of a get operation.
type GetResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts an API version resource.
func (r GetResult) Extract() (*APIVersion, error) { _ = "STUB: not implemented"; return nil, nil }
