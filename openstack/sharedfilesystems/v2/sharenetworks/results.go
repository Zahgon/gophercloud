package sharenetworks

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ShareNetwork contains all the information associated with an OpenStack
// ShareNetwork.
type ShareNetwork struct {
	// The Share Network ID
	ID string `json:"id"`
	// The UUID of the project where the share network was created
	ProjectID string `json:"project_id"`
	// The neutron network ID
	NeutronNetID string `json:"neutron_net_id"`
	// The neutron subnet ID
	NeutronSubnetID string `json:"neutron_subnet_id"`
	// The nova network ID
	NovaNetID string `json:"nova_net_id"`
	// The network type. A valid value is VLAN, VXLAN, GRE or flat
	NetworkType string `json:"network_type"`
	// The segmentation ID
	SegmentationID int `json:"segmentation_id"`
	// The IP block from which to allocate the network, in CIDR notation
	CIDR string `json:"cidr"`
	// The IP version of the network. A valid value is 4 or 6
	IPVersion int `json:"ip_version"`
	// The Share Network name
	Name string `json:"name"`
	// The Share Network description
	Description string `json:"description"`
	// The date and time stamp when the Share Network was created
	CreatedAt time.Time `json:"-"`
	// The date and time stamp when the Share Network was updated
	UpdatedAt time.Time `json:"-"`
}

func (r *ShareNetwork) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

type commonResult struct {
	gophercloud.Result
}

// ShareNetworkPage is a pagination.pager that is returned from a call to the List function.
type ShareNetworkPage struct {
	pagination.MarkerPageBase
}

// NextPageURL generates the URL for the page of results after this one.
func (r ShareNetworkPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// LastMarker returns the last offset in a ListResult.
func (r ShareNetworkPage) LastMarker() (string, error) { _ = "STUB: not implemented"; return "", nil }

// Limit is not present, only one page required

// IsEmpty satisifies the IsEmpty method of the Page interface
func (r ShareNetworkPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractShareNetworks extracts and returns ShareNetworks. It is used while
// iterating over a sharenetworks.List call.
func ExtractShareNetworks(r pagination.Page) ([]ShareNetwork, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract will get the ShareNetwork object out of the commonResult object.
func (r commonResult) Extract() (*ShareNetwork, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateResult contains the response body and error from a Create request.
type CreateResult struct {
	commonResult
}

// DeleteResult contains the response body and error from a Delete request.
type DeleteResult struct {
	gophercloud.ErrResult
}

// GetResult contains the response body and error from a Get request.
type GetResult struct {
	commonResult
}

// UpdateResult contains the response body and error from an Update request.
type UpdateResult struct {
	commonResult
}

// AddSecurityServiceResult contains the response body and error from a security
// service addition request.
type AddSecurityServiceResult struct {
	commonResult
}

// RemoveSecurityServiceResult contains the response body and error from a security
// service removal request.
type RemoveSecurityServiceResult struct {
	commonResult
}
