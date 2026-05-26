package apiversions

import (
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// APIVersion represents an API version for Neutron. It contains the status of
// the API, and its unique ID.
type APIVersion struct {
	Status string `json:"status"`
	ID     string `json:"id"`
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

// APIVersionResource represents a generic API resource. It contains the name
// of the resource and its plural collection name.
type APIVersionResource struct {
	Name       string `json:"name"`
	Collection string `json:"collection"`
}

// APIVersionResourcePage is a concrete type which embeds the common
// SinglePageBase struct, and is used when traversing API versions collections.
type APIVersionResourcePage struct {
	pagination.SinglePageBase
}

// IsEmpty is a concrete function which indicates whether an
// APIVersionResourcePage is empty or not.
func (r APIVersionResourcePage) IsEmpty() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ExtractVersionResources accepts a Page struct, specifically a
// APIVersionResourcePage struct, and extracts the elements into a slice of
// APIVersionResource structs. In other words, the collection is mapped into
// a relevant slice.
func ExtractVersionResources(r pagination.Page) ([]APIVersionResource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
