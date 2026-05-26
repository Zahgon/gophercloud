package apiversions

import (
	"time"

	"github.com/gophercloud/gophercloud/v2/pagination"
)

// APIVersion represents an API version for Cinder.
type APIVersion struct {
	// ID is the unique identifier of the API version.
	ID string `json:"id"`

	// MinVersion is the minimum microversion supported.
	MinVersion string `json:"min_version"`

	// Status represents the status of the API version.
	Status string `json:"status"`

	// Updated is the date the API version was updated.
	Updated time.Time `json:"updated"`

	// Version is the current version and microversion.
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

// ExtractAPIVersion takes a List result and extracts a single requested
// version, which is returned as an APIVersion
func ExtractAPIVersion(r pagination.Page, v string) (*APIVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
