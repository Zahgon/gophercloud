package resourceclasses

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type Link struct {
	Href string `json:"href"`
	Rel  string `json:"rel"`
}

type ResourceClass struct {
	Name string `json:"name"`

	// Links is a list of links associated with the resource class.
	Links []Link `json:"links"`
}

// resourceClassResult is the response of a base ResourceClass result.
type resourceClassResult struct {
	gophercloud.Result
}

// Extract interprets any resourceClassResult-base result as a ResourceClass.
func (r resourceClassResult) Extract() (*ResourceClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetResult represents the result of a Get operation. Call its Extract
// method to interpret it as a ResourceClass.
type GetResult struct {
	resourceClassResult
}

// CreateResult is the response from a Create operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type CreateResult struct {
	gophercloud.ErrResult
}

// UpdateResult is the response from an Update operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type UpdateResult struct {
	gophercloud.ErrResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ResourceClassesPage contains a single page of all resource classes from a List call.
type ResourceClassesPage struct {
	pagination.SinglePageBase
}

// IsEmpty satisfies the IsEmpty method of the Page interface. It returns true
// if a List contains no results.
func (r ResourceClassesPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractResourceClasses takes a List result and extracts the collection of resource classes
// returned by the API.
func ExtractResourceClasses(p pagination.Page) ([]ResourceClass, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
