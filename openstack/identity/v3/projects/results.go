package projects

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Option is a specific option defined at the API to enable features
// on a project.
type Option string

const (
	Immutable Option = "immutable"
)

type projectResult struct {
	gophercloud.Result
}

// GetResult is the result of a Get request. Call its Extract method to
// interpret it as a Project.
type GetResult struct {
	projectResult
}

// CreateResult is the result of a Create request. Call its Extract method to
// interpret it as a Project.
type CreateResult struct {
	projectResult
}

// DeleteResult is the result of a Delete request. Call its ExtractErr method to
// determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult is the result of an Update request. Call its Extract method to
// interpret it as a Project.
type UpdateResult struct {
	projectResult
}

// Project represents an OpenStack Identity Project.
type Project struct {
	// IsDomain indicates whether the project is a domain.
	IsDomain bool `json:"is_domain"`

	// Description is the description of the project.
	Description string `json:"description"`

	// DomainID is the domain ID the project belongs to.
	DomainID string `json:"domain_id"`

	// Enabled is whether or not the project is enabled.
	Enabled bool `json:"enabled"`

	// ID is the unique ID of the project.
	ID string `json:"id"`

	// Name is the name of the project.
	Name string `json:"name"`

	// ParentID is the parent_id of the project.
	ParentID string `json:"parent_id"`

	// Tags is the list of tags associated with the project.
	Tags []string `json:"tags,omitempty"`

	// Extra is free-form extra key/value pairs to describe the project.
	Extra map[string]any `json:"-"`

	// Options are defined options in the API to enable certain features.
	Options map[Option]any `json:"options,omitempty"`
}

func (r *Project) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Collect other fields and bundle them into Extra
// but only if a field titled "extra" wasn't sent.

// ProjectPage is a single page of Project results.
type ProjectPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Projects contains any results.
func (r ProjectPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the links section of the result.
func (r ProjectPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractProjects returns a slice of Projects contained in a single page of
// results.
func ExtractProjects(r pagination.Page) ([]Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract interprets any projectResults as a Project.
func (r projectResult) Extract() (*Project, error) { _ = "STUB: not implemented"; return nil, nil }

// Tags represents a list of Tags object.
type Tags struct {
	// Tags is the list of tags associated with the project.
	Tags []string `json:"tags,omitempty"`
}

// ListTagsResult is the result of a List Tags request. Call its Extract method to
// interpret it as a list of tags.
type ListTagsResult struct {
	gophercloud.Result
}

// Extract interprets any ListTagsResult as a Tags Object.
func (r ListTagsResult) Extract() (*Tags, error) { _ = "STUB: not implemented"; return nil, nil }

// ProjectTags represents a list of Tags object.
type ProjectTags struct {
	// Tags is the list of tags associated with the project.
	Projects []Project `json:"projects,omitempty"`
	// Links contains referencing links to the implied_role.
	Links map[string]any `json:"links"`
}

// ModifyTagsResLinksult is the result of a  Tags request. Call its Extract method to
// interpret it as a project of tags.
type ModifyTagsResult struct {
	gophercloud.Result
}

// Extract interprets any ModifyTags as a Tags Object.
func (r ModifyTagsResult) Extract() (*ProjectTags, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteTagsResult is the result of a Delete Tags request. Call its ExtractErr method to
// determine if the request succeeded or failed.
type DeleteTagsResult struct {
	gophercloud.ErrResult
}
