package instanceactions

import (
	"context"
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToInstanceActionsListQuery() (string, error)
}

// ListOpts represents options used to filter instance action results
// in a List request.
type ListOpts struct {
	// Limit is an integer value to limit the results to return.
	// This requires microversion 2.58 or later.
	Limit int `q:"limit"`

	// Marker is the request ID of the last-seen instance action.
	// This requires microversion 2.58 or later.
	Marker string `q:"marker"`

	// ChangesSince filters the response by actions after the given time.
	// This requires microversion 2.58 or later.
	ChangesSince *time.Time `q:"changes-since"`

	// ChangesBefore filters the response by actions before the given time.
	// This requires microversion 2.66 or later.
	ChangesBefore *time.Time `q:"changes-before"`
}

// ToInstanceActionsListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToInstanceActionsListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request against the API to list the servers actions.
func List(client *gophercloud.ServiceClient, id string, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get makes a request against the API to get a server action.
func Get(ctx context.Context, client *gophercloud.ServiceClient, serverID, requestID string) (r InstanceActionResult) {
	_ = "STUB: not implemented"
	return *new(InstanceActionResult)
}
