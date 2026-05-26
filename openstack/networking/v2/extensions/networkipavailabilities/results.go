package networkipavailabilities

import (
	"encoding/json"
	"math/big"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

type commonResult struct {
	gophercloud.Result
}

// GetResult represents the result of a Get operation. Call its Extract
// method to interpret it as a NetworkIPAvailability.
type GetResult struct {
	commonResult
}

// Extract is a function that accepts a result and extracts a NetworkIPAvailability.
func (r commonResult) Extract() (*NetworkIPAvailability, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NetworkIPAvailability represents availability details for a single network.
type NetworkIPAvailability struct {
	// NetworkID contains an unique identifier of the network.
	NetworkID string `json:"network_id"`

	// NetworkName represents human-readable name of the network.
	NetworkName string `json:"network_name"`

	// ProjectID is the ID of the Identity project.
	ProjectID string `json:"project_id"`

	// TenantID is the ID of the Identity project.
	TenantID string `json:"tenant_id"`

	// SubnetIPAvailabilities contains availability details for every subnet
	// that is associated to the network.
	SubnetIPAvailabilities []SubnetIPAvailability `json:"subnet_ip_availability"`

	// TotalIPs represents a number of IP addresses in the network.
	TotalIPs string `json:"-"`

	// UsedIPs represents a number of used IP addresses in the network.
	UsedIPs string `json:"-"`
}

// Go's encoding/json decodes all JSON numbers into float64 when the target is
// interface{}. For large integers (abs >= 1e21), re-encoding that float64
// produces scientific notation (e.g. "1.1805916207174113e+21"), which
// big.Int.UnmarshalJSON cannot parse. This function handles both plain integer
// and scientific notation forms.
func parseBigIntFromNumber(n json.Number) (*big.Int, error) {
	_ = "STUB: not implemented"

	// Fast path: plain integer notation
	return nil, nil
}

// Slow path: scientific notation from float64 round-trip

func (r *NetworkIPAvailability) UnmarshalJSON(b []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// SubnetIPAvailability represents availability details for a single subnet.
type SubnetIPAvailability struct {
	// SubnetID contains an unique identifier of the subnet.
	SubnetID string `json:"subnet_id"`

	// SubnetName represents human-readable name of the subnet.
	SubnetName string `json:"subnet_name"`

	// CIDR represents prefix in the CIDR format.
	CIDR string `json:"cidr"`

	// IPVersion is the IP protocol version.
	IPVersion int `json:"ip_version"`

	// TotalIPs represents a number of IP addresses in the subnet.
	TotalIPs string `json:"-"`

	// UsedIPs represents a number of used IP addresses in the subnet.
	UsedIPs string `json:"-"`
}

func (r *SubnetIPAvailability) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// NetworkIPAvailabilityPage stores a single page of NetworkIPAvailabilities
// from the List call.
type NetworkIPAvailabilityPage struct {
	pagination.SinglePageBase
}

// IsEmpty determines whether or not a NetworkIPAvailability is empty.
func (r NetworkIPAvailabilityPage) IsEmpty() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// ExtractNetworkIPAvailabilities interprets the results of a single page from
// a List() API call, producing a slice of NetworkIPAvailabilities structures.
func ExtractNetworkIPAvailabilities(r pagination.Page) ([]NetworkIPAvailability, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
