package usage

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// SingleTenantOpts are options for fetching usage of a single tenant.
type SingleTenantOpts struct {
	// The ending time to calculate usage statistics on compute and storage resources.
	End *time.Time `q:"end"`

	// The beginning time to calculate usage statistics on compute and storage resources.
	Start *time.Time `q:"start"`

	// Limit limits the amount of results returned by the API.
	// This requires the client to be set to microversion 2.40 or later.
	Limit int `q:"limit"`

	// Marker instructs the API call where to start listing from.
	// This requires the client to be set to microversion 2.40 or later.
	Marker string `q:"marker"`
}

// SingleTenantOptsBuilder allows extensions to add additional parameters to the
// SingleTenant request.
type SingleTenantOptsBuilder interface {
	ToUsageSingleTenantQuery() (string, error)
}

// ToUsageSingleTenantQuery formats a SingleTenantOpts into a query string.
func (opts SingleTenantOpts) ToUsageSingleTenantQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SingleTenant returns usage data about a single tenant.
func SingleTenant(client *gophercloud.ServiceClient, tenantID string, opts SingleTenantOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// AllTenantsOpts are options for fetching usage of all tenants.
type AllTenantsOpts struct {
	// Detailed will return detailed results.
	Detailed bool

	// The ending time to calculate usage statistics on compute and storage resources.
	End *time.Time `q:"end"`

	// The beginning time to calculate usage statistics on compute and storage resources.
	Start *time.Time `q:"start"`

	// Limit limits the amount of results returned by the API.
	// This requires the client to be set to microversion 2.40 or later.
	Limit int `q:"limit"`

	// Marker instructs the API call where to start listing from.
	// This requires the client to be set to microversion 2.40 or later.
	Marker string `q:"marker"`
}

// AllTenantsOptsBuilder allows extensions to add additional parameters to the
// AllTenants request.
type AllTenantsOptsBuilder interface {
	ToUsageAllTenantsQuery() (string, error)
}

// ToUsageAllTenantsQuery formats a AllTenantsOpts into a query string.
func (opts AllTenantsOpts) ToUsageAllTenantsQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// AllTenants returns usage data about all tenants.
func AllTenants(client *gophercloud.ServiceClient, opts AllTenantsOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}
