package limits

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// GetOptsBuilder allows extensions to add additional parameters to the
// Get request.
type GetOptsBuilder interface {
	ToLimitsQuery() (string, error)
}

// GetOpts enables retrieving limits by a specific tenant.
type GetOpts struct {
	// The tenant ID to retrieve limits for.
	TenantID string `q:"tenant_id"`
}

// ToLimitsQuery formats a GetOpts into a query string.
func (opts GetOpts) ToLimitsQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Get returns the limits about the currently scoped tenant.
func Get(ctx context.Context, client *gophercloud.ServiceClient, opts GetOptsBuilder) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
