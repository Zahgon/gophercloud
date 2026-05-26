package subnetpools

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a subnetpool resource.
func (r commonResult) Extract() (*SubnetPool, error) { _ = "STUB: not implemented"; return nil, nil }

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a SubnetPool.
type GetResult struct {
	commonResult
}

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a SubnetPool.
type CreateResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a SubnetPool.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// SubnetPool represents a Neutron subnetpool.
// A subnetpool is a pool of addresses from which subnets can be allocated.
type SubnetPool struct {
	// ID is the id of the subnetpool.
	ID string `json:"id"`

	// Name is the human-readable name of the subnetpool.
	Name string `json:"name"`

	// DefaultQuota is the per-project quota on the prefix space
	// that can be allocated from the subnetpool for project subnets.
	DefaultQuota int `json:"default_quota"`

	// TenantID is the id of the Identity project.
	TenantID string `json:"tenant_id"`

	// ProjectID is the id of the Identity project.
	ProjectID string `json:"project_id"`

	// CreatedAt is the time at which subnetpool has been created.
	CreatedAt time.Time `json:"-"`

	// UpdatedAt is the time at which subnetpool has been created.
	UpdatedAt time.Time `json:"-"`

	// Prefixes is the list of subnet prefixes to assign to the subnetpool.
	// Neutron API merges adjacent prefixes and treats them as a single prefix.
	// Each subnet prefix must be unique among all subnet prefixes in all subnetpools
	// that are associated with the address scope.
	Prefixes []string `json:"prefixes"`

	// DefaultPrefixLen is yhe size of the prefix to allocate when the cidr
	// or prefixlen attributes are omitted when you create the subnet.
	// Defaults to the MinPrefixLen.
	DefaultPrefixLen int `json:"-"`

	// MinPrefixLen is the smallest prefix that can be allocated from a subnetpool.
	// For IPv4 subnetpools, default is 8.
	// For IPv6 subnetpools, default is 64.
	MinPrefixLen int `json:"-"`

	// MaxPrefixLen is the maximum prefix size that can be allocated from the subnetpool.
	// For IPv4 subnetpools, default is 32.
	// For IPv6 subnetpools, default is 128.
	MaxPrefixLen int `json:"-"`

	// AddressScopeID is the Neutron address scope to assign to the subnetpool.
	AddressScopeID string `json:"address_scope_id"`

	// IPversion is the IP protocol version.
	// Valid value is 4 or 6. Default is 4.
	IPversion int `json:"ip_version"`

	// Shared indicates whether this network is shared across all projects.
	Shared bool `json:"shared"`

	// Description is thehuman-readable description for the resource.
	Description string `json:"description"`

	// IsDefault indicates if the subnetpool is default pool or not.
	IsDefault bool `json:"is_default"`

	// RevisionNumber is the revision number of the subnetpool.
	RevisionNumber int `json:"revision_number"`

	// Tags optionally set via extensions/attributestags
	Tags []string `json:"tags"`
}

func (r *SubnetPool) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"

	// Support for older neutron time format
	return nil
}

// Support for newer neutron time format

// SubnetPoolPage stores a single page of SubnetPools from a List() API call.
type SubnetPoolPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of subnetpools has reached
// the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r SubnetPoolPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsEmpty determines whether or not a SubnetPoolPage is empty.
func (r SubnetPoolPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractSubnetPools interprets the results of a single page from a List() API call,
// producing a slice of SubnetPools structs.
func ExtractSubnetPools(r pagination.Page) ([]SubnetPool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
