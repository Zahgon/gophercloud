package availabilityzoneprofiles

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToAvailabilityZoneProfileListQuery() (string, error)
}

// ListOpts allows to manage the output of the request.
type ListOpts struct {
	// The name of the availability zone profile to filter by.
	Name string `q:"name"`
	// The provider name of the availability zone profile to filter by.
	ProviderName string `q:"provider_name"`
	// The fields that you want the server to return
	Fields []string `q:"fields"`
}

// ToAvailabilityZoneProfileListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToAvailabilityZoneProfileListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// AvailabilityZoneProfiles. It accepts a ListOpts struct, which allows you to
// filter and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToAvailabilityZoneProfileCreateMap() (map[string]any, error)
}

// CreateOpts is the common options struct used in this package's create
// operation.
type CreateOpts struct {
	// Human-readable name for the avaialability zone profile.
	// Does not have to be unique.
	Name string `json:"name" required:"true"`

	// Providing the name of the provider supported by the Octavia installation.
	ProviderName string `json:"provider_name" required:"true"`

	// Providing the json string containing the availability zone metadata.
	AvailabilityZoneData string `json:"availability_zone_data" required:"true"`
}

// ToAvailabilityZoneProfileCreateMap builds a request body from CreateOpts.
func (opts CreateOpts) ToAvailabilityZoneProfileCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create is and operation which add a new AvailabilityZoneProfile into the database.
// CreateResult will be returned.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get retrieves a particular AvailabilityZoneProfile based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// update request.
type UpdateOptsBuilder interface {
	ToAvailabiltyZoneProfileUpdateMap() (map[string]any, error)
}

// UpdateOpts is the common options struct used in this package's update
// operation.
type UpdateOpts struct {
	// Human-readable name for the availability zone profile.
	// Does not have to be unique.
	Name *string `json:"name,omitempty"`

	// Providing the name of the provider supported by the Octavia installation.
	ProviderName *string `json:"provider_name,omitempty"`

	// Providing the json string containing the availability zone metadata.
	AvailabiltyZoneData *string `json:"availability_zone_data,omitempty"`
}

// ToAvailabiltyZoneProfileUpdateMap builds a request body from UpdateOpts.
func (opts UpdateOpts) ToAvailabiltyZoneProfileUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update is an operation which modifies the attributes of the specified
// AvailabilityZoneProfile.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete will permanently delete a particular AvailabiltyZoneProfile based on
// its unique ID.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
