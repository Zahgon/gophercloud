package groups

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Group helps manage related users.
type Group struct {
	// Description describes the group purpose.
	Description string `json:"description"`

	// DomainID is the domain ID the group belongs to.
	DomainID string `json:"domain_id"`

	// ID is the unique ID of the group.
	ID string `json:"id"`

	// Extra is a collection of miscellaneous key/values.
	Extra map[string]any `json:"-"`

	// Links contains referencing links to the group.
	Links map[string]any `json:"links"`

	// Name is the name of the group.
	Name string `json:"name"`
}

func (r *Group) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Collect other fields and bundle them into Extra
// but only if a field titled "extra" wasn't sent.

type groupResult struct {
	gophercloud.Result
}

// GetResult is the response from a Get operation. Call its Extract method
// to interpret it as a Group.
type GetResult struct {
	groupResult
}

// CreateResult is the response from a Create operation. Call its Extract method
// to interpret it as a Group.
type CreateResult struct {
	groupResult
}

// UpdateResult is the response from an Update operation. Call its Extract
// method to interpret it as a Group.
type UpdateResult struct {
	groupResult
}

// DeleteResult is the response from a Delete operation. Call its ExtractErr to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// GroupPage is a single page of Group results.
type GroupPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Groups contains any results.
func (r GroupPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r GroupPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractGroups returns a slice of Groups contained in a single page of results.
func ExtractGroups(r pagination.Page) ([]Group, error) { _ = "STUB: not implemented"; return nil, nil }

// Extract interprets any group results as a Group.
func (r groupResult) Extract() (*Group, error) { _ = "STUB: not implemented"; return nil, nil }
