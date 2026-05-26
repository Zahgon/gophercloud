package apiversions

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// APIVersion represents an API version for Neutron. It contains the status of
// the API, and its unique ID.
type APIVersion struct {
	Status string             `json:"status"`
	ID     string             `json:"id"`
	Links  []gophercloud.Link `json:"links"`
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
