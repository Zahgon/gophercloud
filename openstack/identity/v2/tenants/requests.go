package tenants

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToTenantListQuery() (string, error)
}

// ListOpts filters the Tenants that are returned by the List call.
type ListOpts struct {
	// Marker is the ID of the last Tenant on the previous page.
	Marker string `q:"marker"`

	// Limit specifies the page size.
	Limit int `q:"limit"`
}

// ToTenantListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToTenantListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List enumerates the Tenants to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOpts represents the options needed when creating new tenant.
type CreateOpts struct {
	// Name is the name of the tenant.
	Name string `json:"name" required:"true"`

	// Description is the description of the tenant.
	Description string `json:"description,omitempty"`

	// Enabled sets the tenant status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// CreateOptsBuilder enables extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToTenantCreateMap() (map[string]any, error)
}

// ToTenantCreateMap assembles a request body based on the contents of
// a CreateOpts.
func (opts CreateOpts) ToTenantCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create is the operation responsible for creating new tenant.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get requests details on a single tenant by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToTenantUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies the base attributes that may be updated on an existing
// tenant.
type UpdateOpts struct {
	// Name is the name of the tenant.
	Name string `json:"name,omitempty"`

	// Description is the description of the tenant.
	Description *string `json:"description,omitempty"`

	// Enabled sets the tenant status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// ToTenantUpdateMap formats an UpdateOpts structure into a request body.
func (opts UpdateOpts) ToTenantUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update is the operation responsible for updating exist tenants by their TenantID.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete is the operation responsible for permanently deleting a tenant.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
