package resourceproviders

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToResourceProviderListQuery() (string, error)
}

// ListOpts allows the filtering resource providers. Filtering is achieved by
// passing in struct field values that map to the resource provider attributes
// you want to see returned.
type ListOpts struct {
	// Name is the name of the resource provider to filter the list
	Name string `q:"name"`

	// UUID is the uuid of the resource provider to filter the list
	UUID string `q:"uuid"`

	// MemberOf is a string representing aggregate uuids to filter or exclude from the list
	MemberOf string `q:"member_of"`

	// Resources is a comma-separated list of string indicating an amount of resource
	// of a specified class that a provider must have the capacity and availability to serve
	Resources string `q:"resources"`

	// InTree is a string that represents a resource provider UUID.  The returned resource
	// providers will be in the same provider tree as the specified provider.
	InTree string `q:"in_tree"`

	// Required is comma-delimited list of string trait names.
	Required string `q:"required"`
}

// ToResourceProviderListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToResourceProviderListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request against the API to list resource providers.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToResourceProviderCreateMap() (map[string]any, error)
}

// CreateOpts represents options used to create a resource provider.
type CreateOpts struct {
	Name string `json:"name"`
	UUID string `json:"uuid,omitempty"`
	// The UUID of the immediate parent of the resource provider.
	// Available in version >= 1.14
	ParentProviderUUID string `json:"parent_provider_uuid,omitempty"`
}

// ToResourceProviderCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToResourceProviderCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create makes a request against the API to create a resource provider
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete accepts a unique ID and deletes the resource provider associated with it.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, resourceProviderID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Get retrieves a specific resource provider based on its unique ID.
func Get(ctx context.Context, c *gophercloud.ServiceClient, resourceProviderID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToResourceProviderUpdateMap() (map[string]any, error)
}

// UpdateOpts represents options used to update a resource provider.
type UpdateOpts struct {
	Name *string `json:"name,omitempty"`
	// Available in version >= 1.37. It can be set to any existing provider UUID
	// except to providers that would cause a loop. Using an empty string
	// transforms the provider to a new root provider.
	ParentProviderUUID *string `json:"parent_provider_uuid,omitempty"`
}

// ToResourceProviderUpdateMap constructs a request body from UpdateOpts.
func (opts UpdateOpts) ToResourceProviderUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// In order to set this to null, use an empty string as a value.

// Update makes a request against the API to create a resource provider
func Update(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

func GetUsages(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string) (r GetUsagesResult) {
	_ = "STUB: not implemented"
	return *new(GetUsagesResult)
}

func GetInventories(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string) (r GetInventoriesResult) {
	_ = "STUB: not implemented"
	return *new(GetInventoriesResult)
}

func GetInventory(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID, resourceClass string) (r GetInventoryResult) {
	_ = "STUB: not implemented"
	return *new(GetInventoryResult)
}

// UpdateInventoriesOptsBuilder allows extensions to add additional parameters to the
// UpdateInventories request.
type UpdateInventoriesOptsBuilder interface {
	ToResourceProviderUpdateInventoriesMap() (map[string]any, error)
}

// UpdateInventoriesOpts represents options used to update all inventories of a resource provider.
type UpdateInventoriesOpts = ResourceProviderInventories

// ToResourceProviderUpdateInventoriesMap constructs a request body from UpdateInventoriesOpts.
func (opts UpdateInventoriesOpts) ToResourceProviderUpdateInventoriesMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateInventories updates all inventories of a resource provider.
func UpdateInventories(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string, opts UpdateInventoriesOptsBuilder) (r GetInventoriesResult) {
	_ = "STUB: not implemented"
	return *new(GetInventoriesResult)
}

// DeleteInventories deletes all inventories from a resource provider.
func DeleteInventories(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateInventoryOptsBuilder allows extensions to add additional parameters to the
// UpdateInventory request.
type UpdateInventoryOptsBuilder interface {
	ToResourceProviderUpdateInventoryMap() (map[string]any, error)
}

// UpdateInventoryOpts represents options used to update one inventory of a resource provider.
type UpdateInventoryOpts = ResourceProviderInventory

// ToResourceProviderUpdateInventoryMap constructs a request body from UpdateInventoryOpts.
func (opts UpdateInventoryOpts) ToResourceProviderUpdateInventoryMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateInventory updates one inventory of a resource provider.
func UpdateInventory(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID, resourceClass string, opts UpdateInventoryOptsBuilder) (r UpdateInventoryResult) {
	_ = "STUB: not implemented"
	return *new(UpdateInventoryResult)
}

// DeleteInventory deletes one inventory from a resource provider.
func DeleteInventory(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID, resourceClass string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

func GetAllocations(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string) (r GetAllocationsResult) {
	_ = "STUB: not implemented"
	return *new(GetAllocationsResult)
}

func GetTraits(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string) (r GetTraitsResult) {
	_ = "STUB: not implemented"
	return *new(GetTraitsResult)
}

func GetAggregates(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string) (r GetAggregatesResult) {
	_ = "STUB: not implemented"
	return *new(GetAggregatesResult)
}

// UpdateAggregatesOptsBuilder allows extensions to add additional parameters to the
// UpdateAggregates request.
type UpdateAggregatesOptsBuilder interface {
	ToResourceProviderUpdateAggregatesMap() (map[string]any, error)
}

// UpdateAggregatesOpts represents options used to update aggregates of a resource provider.
type UpdateAggregatesOpts struct {
	// ResourceProviderGeneration is required from microversion 1.19 and later.
	ResourceProviderGeneration *int     `json:"resource_provider_generation,omitempty"`
	Aggregates                 []string `json:"aggregates"`
}

// ToResourceProviderUpdateAggregatesMap constructs a request body from UpdateAggregatesOpts.
func (opts UpdateAggregatesOpts) ToResourceProviderUpdateAggregatesMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdateAggregates(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string, opts UpdateAggregatesOptsBuilder) (r GetAggregatesResult) {
	_ = "STUB: not implemented"
	return *new(GetAggregatesResult)
}

// In microversions prior to 1.19, the body was not enveloped.

// UpdateTraitsOptsBuilder allows extensions to add additional parameters to the
// UpdateTraits request.
type UpdateTraitsOptsBuilder interface {
	ToResourceProviderUpdateTraitsMap() (map[string]any, error)
}

// UpdateTraitsOpts represents options used to update traits of a resource provider.
type UpdateTraitsOpts = ResourceProviderTraits

// ToResourceProviderUpdateTraitsMap constructs a request body from UpdateTraitsOpts.
func (opts UpdateTraitsOpts) ToResourceProviderUpdateTraitsMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func UpdateTraits(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string, opts UpdateTraitsOptsBuilder) (r GetTraitsResult) {
	_ = "STUB: not implemented"
	return *new(GetTraitsResult)
}

func DeleteTraits(ctx context.Context, client *gophercloud.ServiceClient, resourceProviderID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
