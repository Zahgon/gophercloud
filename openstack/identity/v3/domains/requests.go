package domains

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToDomainListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Enabled filters the response by enabled domains.
	Enabled *bool `q:"enabled"`

	// Name filters the response by domain name.
	Name string `q:"name"`

	// Limit limits the number of projects returned per page.
	Limit int `q:"limit"`
}

// ToDomainListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToDomainListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List enumerates the domains to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// ListAvailable enumerates the domains which are available to a specific user.
func ListAvailable(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details on a single domain, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToDomainCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a domain.
type CreateOpts struct {
	// Name is the name of the new domain.
	Name string `json:"name" required:"true"`

	// Description is a description of the domain.
	Description string `json:"description,omitempty"`

	// Enabled sets the domain status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// ToDomainCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToDomainCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new Domain.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a domain.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, domainID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToDomainUpdateMap() (map[string]any, error)
}

// UpdateOpts represents parameters to update a domain.
type UpdateOpts struct {
	// Name is the name of the domain.
	Name string `json:"name,omitempty"`

	// Description is the description of the domain.
	Description *string `json:"description,omitempty"`

	// Enabled sets the domain status to enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// ToUpdateCreateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToDomainUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update modifies the attributes of a domain.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
