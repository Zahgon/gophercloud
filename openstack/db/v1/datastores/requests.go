package datastores

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// List will list all available datastore types that instances can use.
func List(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get will retrieve the details of a specified datastore type.
func Get(ctx context.Context, client *gophercloud.ServiceClient, datastoreID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ListVersions will list all of the available versions for a specified
// datastore type.
func ListVersions(client *gophercloud.ServiceClient, datastoreID string) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetVersion will retrieve the details of a specified datastore version.
func GetVersion(ctx context.Context, client *gophercloud.ServiceClient, datastoreID, versionID string) (r GetVersionResult) {
	_ = "STUB: not implemented"
	return *new(GetVersionResult)
}
