package endpointgroups

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// EndpointGroup is an endpoint group.
type EndpointGroup struct {
	// TenantID specifies a tenant to own the endpoint group.
	TenantID string `json:"tenant_id"`

	// TenantID specifies a tenant to own the endpoint group.
	ProjectID string `json:"project_id"`

	// Description is the human readable description of the endpoint group.
	Description string `json:"description"`

	// Name is the human readable name of the endpoint group.
	Name string `json:"name"`

	// Type is the type of the endpoints in the group.
	Type string `json:"type"`

	// Endpoints is a list of endpoints.
	Endpoints []string `json:"endpoints"`

	// ID is the id of the endpoint group
	ID string `json:"id"`
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts an endpoint group.
func (r commonResult) Extract() (*EndpointGroup, error) { _ = "STUB: not implemented"; return nil, nil }

// EndpointGroupPage is the page returned by a pager when traversing over a
// collection of Policies.
type EndpointGroupPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of Endpoint groups has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r EndpointGroupPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsEmpty checks whether an EndpointGroupPage struct is empty.
func (r EndpointGroupPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractEndpointGroups accepts a Page struct, specifically an EndpointGroupPage struct,
// and extracts the elements into a slice of Endpoint group structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractEndpointGroups(r pagination.Page) ([]EndpointGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as an endpoint group.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as an EndpointGroup.
type GetResult struct {
	commonResult
}

// DeleteResult represents the results of a Delete operation. Call its ExtractErr method
// to determine whether the operation succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult represents the result of an update operation. Call its Extract method
// to interpret it as an EndpointGroup.
type UpdateResult struct {
	commonResult
}
