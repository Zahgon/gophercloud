package groups

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToGroupListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// DomainID filters the response by a domain ID.
	DomainID string `q:"domain_id"`

	// Name filters the response by group name.
	Name string `q:"name"`

	// Filters filters the response by custom filters such as
	// 'name__contains=foo'
	Filters map[string]string `q:"-"`
}

// ToGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToGroupListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List enumerates the Groups to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single group, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToGroupCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a group.
type CreateOpts struct {
	// Name is the name of the new group.
	Name string `json:"name" required:"true"`

	// Description is a description of the group.
	Description string `json:"description,omitempty"`

	// DomainID is the ID of the domain the group belongs to.
	DomainID string `json:"domain_id,omitempty"`

	// Extra is free-form extra key/value pairs to describe the group.
	Extra map[string]any `json:"-"`
}

// ToGroupCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToGroupCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new Group.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToGroupUpdateMap() (map[string]any, error)
}

// UpdateOpts provides options for updating a group.
type UpdateOpts struct {
	// Name is the name of the new group.
	Name string `json:"name,omitempty"`

	// Description is a description of the group.
	Description *string `json:"description,omitempty"`

	// DomainID is the ID of the domain the group belongs to.
	DomainID string `json:"domain_id,omitempty"`

	// Extra is free-form extra key/value pairs to describe the group.
	Extra map[string]any `json:"-"`
}

// ToGroupUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToGroupUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update updates an existing Group.
func Update(ctx context.Context, client *gophercloud.ServiceClient, groupID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete deletes a group.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, groupID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
