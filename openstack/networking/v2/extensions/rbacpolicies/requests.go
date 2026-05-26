package rbacpolicies

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToRBACPolicyListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the rbac attributes you want to see returned. SortKey allows you to sort
// by a particular rbac attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	TargetTenant string       `q:"target_tenant"`
	ObjectType   string       `q:"object_type"`
	ObjectID     string       `q:"object_id"`
	Action       PolicyAction `q:"action"`
	TenantID     string       `q:"tenant_id"`
	ProjectID    string       `q:"project_id"`
	Marker       string       `q:"marker"`
	Limit        int          `q:"limit"`
	SortKey      string       `q:"sort_key"`
	SortDir      string       `q:"sort_dir"`
	Tags         string       `q:"tags"`
	TagsAny      string       `q:"tags-any"`
	NotTags      string       `q:"not-tags"`
	NotTagsAny   string       `q:"not-tags-any"`
}

// ToRBACPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToRBACPolicyListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// rbac policies. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves a specific rbac policy based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// PolicyAction maps to Action for the RBAC policy.
// Which allows access_as_external or access_as_shared.
type PolicyAction string

const (
	// ActionAccessExternal returns Action for the RBAC policy as access_as_external.
	ActionAccessExternal PolicyAction = "access_as_external"

	// ActionAccessShared returns Action for the RBAC policy as access_as_shared.
	ActionAccessShared PolicyAction = "access_as_shared"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToRBACPolicyCreateMap() (map[string]any, error)
}

// CreateOpts represents options used to create a rbac-policy.
type CreateOpts struct {
	Action       PolicyAction `json:"action" required:"true"`
	ObjectType   string       `json:"object_type" required:"true"`
	TargetTenant string       `json:"target_tenant" required:"true"`
	ObjectID     string       `json:"object_id" required:"true"`
}

// ToRBACPolicyCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToRBACPolicyCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create accepts a CreateOpts struct and creates a new rbac-policy using the values
// provided.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// rbac-policy.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete accepts a unique ID and deletes the rbac-policy associated with it.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, rbacPolicyID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToRBACPolicyUpdateMap() (map[string]any, error)
}

// UpdateOpts represents options used to update a rbac-policy.
type UpdateOpts struct {
	TargetTenant string `json:"target_tenant" required:"true"`
}

// ToRBACPolicyUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToRBACPolicyUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update accepts a UpdateOpts struct and updates an existing rbac-policy using the
// values provided.
func Update(ctx context.Context, c *gophercloud.ServiceClient, rbacPolicyID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
