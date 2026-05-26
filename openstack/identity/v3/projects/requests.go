package projects

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToProjectListQuery() (string, error)
}

// ListOpts enables filtering of a list request.
type ListOpts struct {
	// DomainID filters the response by a domain ID.
	DomainID string `q:"domain_id"`

	// Enabled filters the response by enabled projects.
	Enabled *bool `q:"enabled"`

	// IsDomain filters the response by projects that are domains.
	// Setting this to true is effectively listing domains.
	IsDomain *bool `q:"is_domain"`

	// Name filters the response by project name.
	Name string `q:"name"`

	// ParentID filters the response by projects of a given parent project.
	ParentID string `q:"parent_id"`

	// Tags filters on specific project tags. All tags must be present for the project.
	Tags string `q:"tags"`

	// TagsAny filters on specific project tags. At least one of the tags must be present for the project.
	TagsAny string `q:"tags-any"`

	// NotTags filters on specific project tags. All tags must be absent for the project.
	NotTags string `q:"not-tags"`

	// NotTagsAny filters on specific project tags. At least one of the tags must be absent for the project.
	NotTagsAny string `q:"not-tags-any"`

	// Limit limits the number of projects returned per page.
	Limit int `q:"limit"`

	// Filters filters the response by custom filters such as
	// 'name__contains=foo'
	Filters map[string]string `q:"-"`
}

// ToProjectListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToProjectListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List enumerates the Projects to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListAvailable enumerates the Projects which are available to a specific user.
func ListAvailable(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single project, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToProjectCreateMap() (map[string]any, error)
}

// CreateOpts represents parameters used to create a project.
type CreateOpts struct {
	// DomainID is the ID this project will belong under.
	DomainID string `json:"domain_id,omitempty"`

	// Enabled sets the project status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// IsDomain indicates if this project is a domain.
	IsDomain *bool `json:"is_domain,omitempty"`

	// Name is the name of the project.
	Name string `json:"name" required:"true"`

	// ParentID specifies the parent project of this new project.
	ParentID string `json:"parent_id,omitempty"`

	// Description is the description of the project.
	Description string `json:"description,omitempty"`

	// Tags is a list of tags to associate with the project.
	Tags []string `json:"tags,omitempty"`

	// Extra is free-form extra key/value pairs to describe the project.
	Extra map[string]any `json:"-"`

	// Options are defined options in the API to enable certain features.
	Options map[Option]any `json:"options,omitempty"`
}

// ToProjectCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToProjectCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new Project.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a project.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToProjectUpdateMap() (map[string]any, error)
}

// UpdateOpts represents parameters to update a project.
type UpdateOpts struct {
	// DomainID is the ID this project will belong under.
	DomainID string `json:"domain_id,omitempty"`

	// Enabled sets the project status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// IsDomain indicates if this project is a domain.
	IsDomain *bool `json:"is_domain,omitempty"`

	// Name is the name of the project.
	Name string `json:"name,omitempty"`

	// ParentID specifies the parent project of this new project.
	ParentID string `json:"parent_id,omitempty"`

	// Description is the description of the project.
	Description *string `json:"description,omitempty"`

	// Tags is a list of tags to associate with the project.
	Tags *[]string `json:"tags,omitempty"`

	// Extra is free-form extra key/value pairs to describe the project.
	Extra map[string]any `json:"-"`

	// Options are defined options in the API to enable certain features.
	Options map[Option]any `json:"options,omitempty"`
}

// ToUpdateCreateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToProjectUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update modifies the attributes of a project.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// CheckTags lists tags for a project.
func ListTags(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r ListTagsResult) {
	_ = "STUB: not implemented"
	return *new(ListTagsResult)
}

// Tags represents a list of Tags object.
type ModifyTagsOpts struct {
	// Tags is the list of tags associated with the project.
	Tags []string `json:"tags,omitempty"`
}

// ModifyTagsOptsBuilder allows extensions to add additional parameters to
// the Modify request.
type ModifyTagsOptsBuilder interface {
	ToModifyTagsCreateMap() (map[string]any, error)
}

// ToModifyTagsCreateMap formats a ModifyTagsOpts into a Modify tags request.
func (opts ModifyTagsOpts) ToModifyTagsCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ModifyTags deletes all tags of a project and adds new ones.
func ModifyTags(ctx context.Context, client *gophercloud.ServiceClient, projectID string, opts ModifyTagsOptsBuilder) (r ModifyTagsResult) {
	_ = "STUB: not implemented"
	return *new(ModifyTagsResult)
}

// DeleteTag deletes a tag from a project.
func DeleteTags(ctx context.Context, client *gophercloud.ServiceClient, projectID string) (r DeleteTagsResult) {
	_ = "STUB: not implemented"
	return *new(DeleteTagsResult)
}
