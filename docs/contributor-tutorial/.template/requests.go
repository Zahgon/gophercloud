package RESOURCE

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToResourceListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
}

// ToResourceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToResourceListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List retrieves a list of RESOURCES.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves details of a RESOURCE.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToResourceCreateMap() (map[string]interface{}, error)
}

// CreateOpts provides options used to create a RESOURCE.
type CreateOpts struct {
}

// ToResourceCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToResourceCreateMap() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new RESOURCE.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Delete deletes a RESOURCE.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToResourceUpdateMap() (map[string]interface{}, error)
}

// UpdateOpts represents parameters to update a RESOURCE.
type UpdateOpts struct {
}

// ToUpdateCreateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToResourceUpdateMap() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update modifies the attributes of a RESOURCE.
func Update(ctx context.Context, client *gophercloud.ServiceClient, id string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
