package networks

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a network resource.
func (r commonResult) Extract() (*Network, error) { _ = "STUB: not implemented"; return nil, nil }

func (r commonResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a Network.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a Network.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a Network.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// Network represents, well, a network.
type Network struct {
	// UUID for the network
	ID string `json:"id"`

	// Human-readable name for the network. Might not be unique.
	Name string `json:"name"`

	// Description for the network
	Description string `json:"description"`

	// The administrative state of network. If false (down), the network does not
	// forward packets.
	AdminStateUp bool `json:"admin_state_up"`

	// Indicates whether network is currently operational. Possible values include
	// `ACTIVE', `DOWN', `BUILD', or `ERROR'. Plug-ins might define additional
	// values.
	Status string `json:"status"`

	// Subnets associated with this network.
	Subnets []string `json:"subnets"`

	// TenantID is the project owner of the network.
	TenantID string `json:"tenant_id"`

	// UpdatedAt and CreatedAt contain ISO-8601 timestamps of when the state of the
	// network last changed, and when it was created.
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`

	// ProjectID is the project owner of the network.
	ProjectID string `json:"project_id"`

	// Specifies whether the network resource can be accessed by any tenant.
	Shared bool `json:"shared"`

	// Availability zone hints groups network nodes that run services like DHCP, L3, FW, and others.
	// Used to make network resources highly available.
	AvailabilityZoneHints []string `json:"availability_zone_hints"`

	// Tags optionally set via extensions/attributestags
	Tags []string `json:"tags"`

	// RevisionNumber optionally set via extensions/standard-attr-revisions
	RevisionNumber int `json:"revision_number"`
}

func (r *Network) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"

	// Support for older neutron time format
	return nil
}

// Support for newer neutron time format

// NetworkPage is the page returned by a pager when traversing over a
// collection of networks.
type NetworkPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of networks has reached
// the end of a page and the pager seeks to traverse over a new one. In order
// to do this, it needs to construct the next page's URL.
func (r NetworkPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsEmpty checks whether a NetworkPage struct is empty.
func (r NetworkPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractNetworks accepts a Page struct, specifically a NetworkPage struct,
// and extracts the elements into a slice of Network structs. In other words,
// a generic collection is mapped into a relevant slice.
func ExtractNetworks(r pagination.Page) ([]Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractNetworksInto(r pagination.Page, v any) error { _ = "STUB: not implemented"; return nil }
