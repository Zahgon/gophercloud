package groups

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToSecGroupListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the group attributes you want to see returned. SortKey allows you to
// sort by a particular network attribute. SortDir sets the direction, and is
// either `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	ID             string `q:"id"`
	Name           string `q:"name"`
	Description    string `q:"description"`
	Stateful       *bool  `q:"stateful"`
	TenantID       string `q:"tenant_id"`
	ProjectID      string `q:"project_id"`
	Limit          int    `q:"limit"`
	Marker         string `q:"marker"`
	SortKey        string `q:"sort_key"`
	SortDir        string `q:"sort_dir"`
	Tags           string `q:"tags"`
	TagsAny        string `q:"tags-any"`
	NotTags        string `q:"not-tags"`
	NotTagsAny     string `q:"not-tags-any"`
	RevisionNumber *int   `q:"revision_number"`
}

// ToPortListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToSecGroupListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// security groups. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToSecGroupCreateMap() (map[string]any, error)
}

// CreateOpts contains all the values needed to create a new security group.
type CreateOpts struct {
	// Human-readable name for the Security Group. Does not have to be unique.
	Name string `json:"name" required:"true"`

	// TenantID is the UUID of the project who owns the Group.
	// Only administrative users can specify a tenant UUID other than their own.
	TenantID string `json:"tenant_id,omitempty"`

	// ProjectID is the UUID of the project who owns the Group.
	// Only administrative users can specify a tenant UUID other than their own.
	ProjectID string `json:"project_id,omitempty"`

	// Describes the security group.
	Description string `json:"description,omitempty"`

	// Stateful indicates if the security group is stateful or stateless.
	Stateful *bool `json:"stateful,omitempty"`
}

// ToSecGroupCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToSecGroupCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create is an operation which provisions a new security group with default
// security group rules for the IPv4 and IPv6 ether types.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToSecGroupUpdateMap() (map[string]any, error)
}

// UpdateOpts contains all the values needed to update an existing security
// group.
type UpdateOpts struct {
	// Human-readable name for the Security Group. Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// Describes the security group.
	Description *string `json:"description,omitempty"`

	// Stateful indicates if the security group is stateful or stateless.
	Stateful *bool `json:"stateful,omitempty"`

	// RevisionNumber implements extension:standard-attr-revisions. If != "" it
	// will set revision_number=%s. If the revision number does not match, the
	// update will fail.
	RevisionNumber *int `json:"-" h:"If-Match"`
}

// ToSecGroupUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToSecGroupUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update is an operation which updates an existing security group.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Get retrieves a particular security group based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Delete will permanently delete a particular security group based on its
// unique ID.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
