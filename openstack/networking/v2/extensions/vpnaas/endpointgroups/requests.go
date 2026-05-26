package endpointgroups

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type EndpointType string

const (
	TypeSubnet  EndpointType = "subnet"
	TypeCIDR    EndpointType = "cidr"
	TypeVLAN    EndpointType = "vlan"
	TypeNetwork EndpointType = "network"
	TypeRouter  EndpointType = "router"
)

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToEndpointGroupCreateMap() (map[string]any, error)
}

// CreateOpts contains all the values needed to create a new endpoint group
type CreateOpts struct {
	// TenantID specifies a tenant to own the endpoint group. The caller must have
	// an admin role in order to set this. Otherwise, this field is left unset
	// and the caller will be the owner.
	TenantID string `json:"tenant_id,omitempty"`

	// Description is the human readable description of the endpoint group.
	Description string `json:"description,omitempty"`

	// Name is the human readable name of the endpoint group.
	Name string `json:"name,omitempty"`

	// The type of the endpoints in the group.
	// A valid value is subnet, cidr, network, router, or vlan.
	Type EndpointType `json:"type,omitempty"`

	// List of endpoints of the same type, for the endpoint group.
	// The values will depend on the type.
	Endpoints []string `json:"endpoints"`
}

// ToEndpointGroupCreateMap casts a CreateOpts struct to a map.
func (opts CreateOpts) ToEndpointGroupCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create accepts a CreateOpts struct and uses the values to create a new
// endpoint group.
func Create(ctx context.Context, c *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get retrieves a particular endpoint group based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToEndpointGroupListQuery() (string, error)
}

// ListOpts allows the filtering of paginated collections through
// the API. Filtering is achieved by passing in struct field values that map to
// the Endpoint group attributes you want to see returned.
type ListOpts struct {
	TenantID    string `q:"tenant_id"`
	ProjectID   string `q:"project_id"`
	Description string `q:"description"`
	Name        string `q:"name"`
	Type        string `q:"type"`
}

// ToEndpointGroupListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToEndpointGroupListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns a Pager which allows you to iterate over a collection of
// Endpoint groups. It accepts a ListOpts struct, which allows you to filter
// the returned collection for greater efficiency.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Delete will permanently delete a particular endpoint group based on its
// unique ID.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToEndpointGroupUpdateMap() (map[string]any, error)
}

// UpdateOpts contains the values used when updating an endpoint group.
type UpdateOpts struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

// ToEndpointGroupUpdateMap casts an UpdateOpts struct to a map.
func (opts UpdateOpts) ToEndpointGroupUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update allows endpoint groups to be updated.
func Update(ctx context.Context, c *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
