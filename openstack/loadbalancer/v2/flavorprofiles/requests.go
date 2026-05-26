package flavorprofiles

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToFlavorProfileListQuery() (string, error)
}

// ListOpts allows to manage the output of the request.
type ListOpts struct {
	// The name of the flavor profile to filter by.
	Name string `q:"name"`
	// The provider name of the flavor profile to filter by.
	ProviderName string `q:"provider_name"`
	// The fields that you want the server to return
	Fields []string `q:"fields"`
}

// ToFlavorProfileListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToFlavorProfileListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// FlavorProfiles. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToFlavorProfileCreateMap() (map[string]any, error)
}

// CreateOpts is the common options struct used in this package's Create
// operation.
type CreateOpts struct {
	// Human-readable name for the Loadbalancer. Does not have to be unique.
	Name string `json:"name" required:"true"`

	// Providing the name of the provider supported by the Octavia installation.
	ProviderName string `json:"provider_name" required:"true"`

	// Providing the json string containing the flavor metadata.
	FlavorData string `json:"flavor_data" required:"true"`
}

// ToFlavorProfileCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToFlavorProfileCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create is and operation which add a new FlavorProfile into the database.
// CreateResult will be returned.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get retrieves a particular FlavorProfile based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToFlavorProfileUpdateMap() (map[string]any, error)
}

// UpdateOpts is the common options struct used in this package's Update
// operation.
type UpdateOpts struct {
	// Human-readable name for the Loadbalancer. Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// Providing the name of the provider supported by the Octavia installation.
	ProviderName *string `json:"provider_name,omitempty"`

	// Providing the json string containing the flavor metadata.
	FlavorData *string `json:"flavor_data,omitempty"`
}

// ToFlavorProfileUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToFlavorProfileUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update is an operation which modifies the attributes of the specified
// FlavorProfile.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete will permanently delete a particular FlavorProfile based on its
// unique ID.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
