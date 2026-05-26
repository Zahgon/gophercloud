package usages

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// GetOptsBuilder allows extensions to add additional parameters to the
// Get request.
type GetOptsBuilder interface {
	ToUsagesGetQuery() (string, error)
}

// GetOpts specifies the query parameters for retrieving total usages.
//
// This requires microversion 1.9 or later.
type GetOpts struct {
	// ProjectID is required: only usages for this project are returned.
	ProjectID string `q:"project_id"`

	// UserID is optional: when set, only usages for this user within the
	// project are returned.
	UserID string `q:"user_id,omitempty"`

	// ConsumerType is optional: when set, results are filtered to this consumer type.
	// Available from microversion 1.38.
	ConsumerType string `q:"consumer_type,omitempty"`
}

// ToUsagesGetQuery formats a GetOpts into a query string.
func (opts GetOpts) ToUsagesGetQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Get retrieves the total resource usages for a project (and optionally a user).
//
// Requires microversion 1.9 or later.
func Get(ctx context.Context, client *gophercloud.ServiceClient, opts GetOptsBuilder) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
