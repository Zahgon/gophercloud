package trunks

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToTrunkCreateMap() (map[string]any, error)
}

// CreateOpts represents the attributes used when creating a new trunk.
type CreateOpts struct {
	TenantID     string    `json:"tenant_id,omitempty"`
	ProjectID    string    `json:"project_id,omitempty"`
	PortID       string    `json:"port_id" required:"true"`
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	AdminStateUp *bool     `json:"admin_state_up,omitempty"`
	Subports     []Subport `json:"sub_ports"`
}

// ToTrunkCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToTrunkCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete accepts a unique ID and deletes the trunk associated with it.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToTrunkListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the trunk attributes you want to see returned. SortKey allows you to sort
// by a particular trunk attribute. SortDir sets the direction, and is either
// `asc' or `desc'. Marker and Limit are used for pagination.
type ListOpts struct {
	AdminStateUp *bool  `q:"admin_state_up"`
	Description  string `q:"description"`
	ID           string `q:"id"`
	Name         string `q:"name"`
	PortID       string `q:"port_id"`
	Status       string `q:"status"`
	TenantID     string `q:"tenant_id"`
	ProjectID    string `q:"project_id"`
	SortDir      string `q:"sort_dir"`
	SortKey      string `q:"sort_key"`
	Tags         string `q:"tags"`
	TagsAny      string `q:"tags-any"`
	NotTags      string `q:"not-tags"`
	NotTagsAny   string `q:"not-tags-any"`
	// TODO change type to *int for consistency
	RevisionNumber string `q:"revision_number"`
}

// ToTrunkListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTrunkListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List returns a Pager which allows you to iterate over a collection of
// trunks. It accepts a ListOpts struct, which allows you to filter and sort
// the returned collection for greater efficiency.
//
// Default policy settings return only those trunks that are owned by the tenant
// who submits the request, unless the request is submitted by a user with
// administrative rights.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves a specific trunk based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

type UpdateOptsBuilder interface {
	ToTrunkUpdateMap() (map[string]any, error)
}

type UpdateOpts struct {
	AdminStateUp *bool   `json:"admin_state_up,omitempty"`
	Name         *string `json:"name,omitempty"`
	Description  *string `json:"description,omitempty"`

	// RevisionNumber implements extension:standard-attr-revisions. If != "" it
	// will set revision_number=%s. If the revision number does not match, the
	// update will fail.
	RevisionNumber *int `json:"-" h:"If-Match"`
}

func (opts UpdateOpts) ToTrunkUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

func GetSubports(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetSubportsResult) {
	_ = "STUB: not implemented"
	return *new(GetSubportsResult)
}

type AddSubportsOpts struct {
	Subports []Subport `json:"sub_ports" required:"true"`
}

type AddSubportsOptsBuilder interface {
	ToTrunkAddSubportsMap() (map[string]any, error)
}

func (opts AddSubportsOpts) ToTrunkAddSubportsMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AddSubports(ctx context.Context, c *gophercloud.ServiceClient, id string, opts AddSubportsOptsBuilder) (r UpdateSubportsResult) {
	_ = "STUB: not implemented"
	return *new(UpdateSubportsResult)
}

type RemoveSubport struct {
	PortID string `json:"port_id" required:"true"`
}

type RemoveSubportsOpts struct {
	Subports []RemoveSubport `json:"sub_ports"`
}

type RemoveSubportsOptsBuilder interface {
	ToTrunkRemoveSubportsMap() (map[string]any, error)
}

func (opts RemoveSubportsOpts) ToTrunkRemoveSubportsMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func RemoveSubports(ctx context.Context, c *gophercloud.ServiceClient, id string, opts RemoveSubportsOptsBuilder) (r UpdateSubportsResult) {
	_ = "STUB: not implemented"
	return *new(UpdateSubportsResult)
}
