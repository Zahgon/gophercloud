package projectendpoints

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type CreateOptsBuilder interface {
	ToEndpointCreateMap() (map[string]any, error)
}

// Create inserts a new Endpoint association to a project.
func Create(ctx context.Context, client *gophercloud.ServiceClient, projectID, endpointID string) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// List enumerates endpoints in a paginated collection, optionally filtered
// by ListOpts criteria.
func List(client *gophercloud.ServiceClient, projectID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Delete removes an endpoint from the service catalog.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, projectID string, endpointID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
