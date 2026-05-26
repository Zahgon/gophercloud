package resourceclasses

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List retrieves a list of resource classes.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves the resource class with the provided name.
func Get(ctx context.Context, client *gophercloud.ServiceClient, name string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToResourceClassCreateMap() (map[string]any, error)
}

// CreateOpts represents the attributes of a new resource class.
type CreateOpts struct {
	Name string `json:"name" required:"true"`
}

// ToResourceClassCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToResourceClassCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new resource class.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Update ensures the existence of a custom resource class with the
// provided name (can be safely called multiple times).
func Update(ctx context.Context, client *gophercloud.ServiceClient, name string) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete deletes the resource class with the provided name.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, name string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
