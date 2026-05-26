package services

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToServiceCreateMap() (map[string]any, error)
}

// CreateOpts contains all the values needed to create a new VPN service
type CreateOpts struct {
	// TenantID specifies a tenant to own the VPN service. The caller must have
	// an admin role in order to set this. Otherwise, this field is left unset
	// and the caller will be the owner.
	TenantID string `json:"tenant_id,omitempty"`

	// SubnetID is the ID of the subnet.
	SubnetID string `json:"subnet_id,omitempty"`

	// RouterID is the ID of the router.
	RouterID string `json:"router_id" required:"true"`

	// Description is the human readable description of the service.
	Description string `json:"description,omitempty"`

	// AdminStateUp is the administrative state of the resource, which is up (true) or down (false).
	AdminStateUp *bool `json:"admin_state_up"`

	// Name is the human readable name of the service.
	Name string `json:"name,omitempty"`

	// The ID of the flavor.
	FlavorID string `json:"flavor_id,omitempty"`
}

// ToServiceCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToServiceCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create accepts a CreateOpts struct and uses the values to create a new
// VPN service.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete will permanently delete a particular VPN service based on its
// unique ID.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToServiceUpdateMap() (map[string]any, error)
}

// UpdateOpts contains the values used when updating a VPN service
type UpdateOpts struct {
	// Name is the human readable name of the service.
	Name *string `json:"name,omitempty"`

	// Description is the human readable description of the service.
	Description *string `json:"description,omitempty"`

	// AdminStateUp is the administrative state of the resource, which is up (true) or down (false).
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
}

// ToServiceUpdateMap casts aa UodateOpts struct to a map.
func (opts UpdateOpts) ToServiceUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update allows VPN services to be updated.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToServiceListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the VPN service attributes you want to see returned.
type ListOpts struct {
	TenantID     string `q:"tenant_id"`
	Name         string `q:"name"`
	Description  string `q:"description"`
	AdminStateUp *bool  `q:"admin_state_up"`
	Status       string `q:"status"`
	SubnetID     string `q:"subnet_id"`
	RouterID     string `q:"router_id"`
	ProjectID    string `q:"project_id"`
	ExternalV6IP string `q:"external_v6_ip"`
	ExternalV4IP string `q:"external_v4_ip"`
	FlavorID     string `q:"flavor_id"`
}

// ToServiceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToServiceListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// VPN services. It accepts a ListOpts struct, which allows you to filter
// and sort the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves a particular VPN service based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
