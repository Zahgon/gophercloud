package servergroups

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type ListOptsBuilder interface {
	ToServerListQuery() (string, error)
}

type ListOpts struct {
	// AllProjects is a bool to show all projects.
	AllProjects bool `q:"all_projects"`

	// Requests a page size of items.
	Limit int `q:"limit"`

	// Used in conjunction with limit to return a slice of items.
	Offset int `q:"offset"`
}

// ToServerListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToServerListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List returns a Pager that allows you to iterate over a collection of
// ServerGroups.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToServerGroupCreateMap() (map[string]any, error)
}

// CreateOpts specifies Server Group creation parameters.
type CreateOpts struct {
	// Name is the name of the server group.
	Name string `json:"name" required:"true"`

	// Policies are the server group policies.
	Policies []string `json:"policies,omitempty"`

	// Policy specifies the name of a policy.
	// Requires microversion 2.64 or later.
	Policy string `json:"policy,omitempty"`

	// Rules specifies the set of rules.
	// Requires microversion 2.64 or later.
	Rules *Rules `json:"rules,omitempty"`
}

// ToServerGroupCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToServerGroupCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new Server Group.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get returns data about a previously created ServerGroup.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Delete requests the deletion of a previously allocated ServerGroup.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
