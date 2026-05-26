package apiversions

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// List lists all the API versions available to end users.
func List(ctx context.Context, client *gophercloud.ServiceClient) (r ListResult) {
	_ = "STUB: not implemented"
	return *new(ListResult)
}

// Get will get a specific API version, specified by major ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, v string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
