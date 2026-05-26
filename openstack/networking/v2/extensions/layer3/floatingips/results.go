package floatingips

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// FloatingIP represents a floating IP resource. A floating IP is an external
// IP address that is mapped to an internal port and, optionally, a specific
// IP address on a private network. In other words, it enables access to an
// instance on a private network from an external network. For this reason,
// floating IPs can only be defined on networks where the `router:external'
// attribute (provided by the external network extension) is set to True.
type FloatingIP struct {
	// ID is the unique identifier for the floating IP instance.
	ID string `json:"id"`

	// Description for the floating IP instance.
	Description string `json:"description"`

	// FloatingNetworkID is the UUID of the external network where the floating
	// IP is to be created.
	FloatingNetworkID string `json:"floating_network_id"`

	// FloatingIP is the address of the floating IP on the external network.
	FloatingIP string `json:"floating_ip_address"`

	// PortID is the UUID of the port on an internal network that is associated
	// with the floating IP.
	PortID string `json:"port_id"`

	// FixedIP is the specific IP address of the internal port which should be
	// associated with the floating IP.
	FixedIP string `json:"fixed_ip_address"`

	// TenantID is the project owner of the floating IP. Only admin users can
	// specify a project identifier other than its own.
	TenantID string `json:"tenant_id"`

	// UpdatedAt and CreatedAt contain ISO-8601 timestamps of when the state of
	// the floating ip last changed, and when it was created.
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`

	// ProjectID is the project owner of the floating IP.
	ProjectID string `json:"project_id"`

	// Status is the condition of the API resource.
	Status string `json:"status"`

	// RouterID is the ID of the router used for this floating IP.
	RouterID string `json:"router_id"`

	// Tags optionally set via extensions/attributestags
	Tags []string `json:"tags"`

	// RevisionNumber optionally set via extensions/standard-attr-revisions
	RevisionNumber int `json:"revision_number"`
}

func (r *FloatingIP) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"

	// Support for older neutron time format
	return nil
}

// Support for newer neutron time format

type commonResult struct {
	gophercloud.Result
}

// Extract will extract a FloatingIP resource from a result.
func (r commonResult) Extract() (*FloatingIP, error) { _ = "STUB: not implemented"; return nil, nil }

func (r commonResult) ExtractInto(v any) error { _ = "STUB: not implemented"; return nil }

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a FloatingIP.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a FloatingIP.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a FloatingIP.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of an update operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// FloatingIPPage is the page returned by a pager when traversing over a
// collection of floating IPs.
type FloatingIPPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of floating IPs has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r FloatingIPPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsEmpty checks whether a FloatingIPPage struct is empty.
func (r FloatingIPPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractFloatingIPs accepts a Page struct, specifically a FloatingIPPage
// struct, and extracts the elements into a slice of FloatingIP structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractFloatingIPs(r pagination.Page) ([]FloatingIP, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ExtractFloatingIPsInto(r pagination.Page, v any) error { _ = "STUB: not implemented"; return nil }
