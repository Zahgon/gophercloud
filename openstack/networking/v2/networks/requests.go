package networks

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToNetworkListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the network attributes you want to see returned. SortKey allows you to sort
// by a particular network attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	Status         string `q:"status"`
	Name           string `q:"name"`
	Description    string `q:"description"`
	AdminStateUp   *bool  `q:"admin_state_up"`
	TenantID       string `q:"tenant_id"`
	ProjectID      string `q:"project_id"`
	Shared         *bool  `q:"shared"`
	ID             string `q:"id"`
	Marker         string `q:"marker"`
	Limit          int    `q:"limit"`
	SortKey        string `q:"sort_key"`
	SortDir        string `q:"sort_dir"`
	Tags           string `q:"tags"`
	TagsAny        string `q:"tags-any"`
	NotTags        string `q:"not-tags"`
	NotTagsAny     string `q:"not-tags-any"`
	RevisionNumber *int   `q:"revision_number"`
}

// ToNetworkListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToNetworkListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// networks. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves a specific network based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToNetworkCreateMap() (map[string]any, error)
}

// CreateOpts represents options used to create a network.
type CreateOpts struct {
	AdminStateUp          *bool    `json:"admin_state_up,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Description           string   `json:"description,omitempty"`
	Shared                *bool    `json:"shared,omitempty"`
	TenantID              string   `json:"tenant_id,omitempty"`
	ProjectID             string   `json:"project_id,omitempty"`
	AvailabilityZoneHints []string `json:"availability_zone_hints,omitempty"`
}

// ToNetworkCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToNetworkCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create accepts a CreateOpts struct and creates a new network using the values
// provided. This operation does not actually require a request body, i.e. the
// CreateOpts struct argument can be empty.
//
// The tenant ID that is contained in the URI is the tenant that creates the
// network. An admin user, however, has the option of specifying another tenant
// ID in the CreateOpts struct.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToNetworkUpdateMap() (map[string]any, error)
}

// UpdateOpts represents options used to update a network.
type UpdateOpts struct {
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`
	Shared       *bool   `json:"shared,omitempty"`

	// RevisionNumber implements extension:standard-attr-revisions. If != "" it
	// will set revision_number=%s. If the revision number does not match, the
	// update will fail.
	RevisionNumber *int `json:"-" h:"If-Match"`
}

// ToNetworkUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToNetworkUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update accepts a UpdateOpts struct and updates an existing network using the
// values provided. For more information, see the Create function.
func Update(ctx context.Context, c *gophercloud.ServiceClient, networkID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete accepts a unique ID and deletes the network associated with it.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, networkID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
