package tenants

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Tenant is a grouping of users in the identity service.
type Tenant struct {
	// ID is a unique identifier for this tenant.
	ID string `json:"id"`

	// Name is a friendlier user-facing name for this tenant.
	Name string `json:"name"`

	// Description is a human-readable explanation of this Tenant's purpose.
	Description string `json:"description"`

	// Enabled indicates whether or not a tenant is active.
	Enabled bool `json:"enabled"`
}

// TenantPage is a single page of Tenant results.
type TenantPage struct {
	pagination.LinkedPageBase
}

// IsEmpty determines whether or not a page of Tenants contains any results.
func (r TenantPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL extracts the "next" link from the tenants_links section of the result.
func (r TenantPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ExtractTenants returns a slice of Tenants contained in a single page of
// results.
func ExtractTenants(r pagination.Page) ([]Tenant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type tenantResult struct {
	gophercloud.Result
}

// Extract interprets any tenantResults as a Tenant.
func (r tenantResult) Extract() (*Tenant, error) { _ = "STUB: not implemented"; return nil, nil }

// GetResult is the response from a Get request. Call its Extract method to
// interpret it as a Tenant.
type GetResult struct {
	tenantResult
}

// CreateResult is the response from a Create request. Call its Extract method
// to interpret it as a Tenant.
type CreateResult struct {
	tenantResult
}

// DeleteResult is the response from a Get request. Call its ExtractErr method
// to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// UpdateResult is the response from a Update request. Call its Extract method
// to interpret it as a Tenant.
type UpdateResult struct {
	tenantResult
}
