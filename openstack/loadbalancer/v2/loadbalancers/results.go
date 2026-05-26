package loadbalancers

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/listeners"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/pools"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// LoadBalancer is the primary load balancing configuration object that
// specifies the virtual IP address on which client traffic is received, as well
// as other details such as the load balancing method to be use, protocol, etc.
type LoadBalancer struct {
	// Human-readable description for the Loadbalancer.
	Description string `json:"description"`

	// The administrative state of the Loadbalancer.
	// A valid value is true (UP) or false (DOWN).
	AdminStateUp bool `json:"admin_state_up"`

	// Owner of the LoadBalancer.
	ProjectID string `json:"project_id"`

	// UpdatedAt and CreatedAt contain ISO-8601 timestamps of when the state of the
	// loadbalancer last changed, and when it was created.
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time `json:"-"`

	// The provisioning status of the LoadBalancer.
	// This value is ACTIVE, PENDING_CREATE or ERROR.
	ProvisioningStatus string `json:"provisioning_status"`

	// The IP address of the Loadbalancer.
	VipAddress string `json:"vip_address"`

	// The UUID of the port associated with the IP address.
	VipPortID string `json:"vip_port_id"`

	// The UUID of the subnet on which to allocate the virtual IP for the
	// Loadbalancer address.
	VipSubnetID string `json:"vip_subnet_id"`

	// The UUID of the network on which to allocate the virtual IP for the
	// Loadbalancer address.
	VipNetworkID string `json:"vip_network_id"`

	// The ID of the QoS Policy which will apply to the Virtual IP
	VipQosPolicyID string `json:"vip_qos_policy_id"`

	// The unique ID for the LoadBalancer.
	ID string `json:"id"`

	// The operating status of the LoadBalancer. This value is ONLINE or OFFLINE.
	OperatingStatus string `json:"operating_status"`

	// Human-readable name for the LoadBalancer. Does not have to be unique.
	Name string `json:"name"`

	// The UUID of a flavor if set.
	FlavorID string `json:"flavor_id"`

	// The name of an Octavia availability zone if set.
	AvailabilityZone string `json:"availability_zone"`

	// The name of the provider.
	Provider string `json:"provider"`

	// Listeners are the listeners related to this Loadbalancer.
	Listeners []listeners.Listener `json:"listeners"`

	// Pools are the pools related to this Loadbalancer.
	Pools []pools.Pool `json:"pools"`

	// Tags is a list of resource tags. Tags are arbitrarily defined strings
	// attached to the resource.
	Tags []string `json:"tags"`

	// The additional ips of the loadbalancer. Subnets must all belong to the same network as the primary VIP.
	// New in version 2.26
	AdditionalVips []AdditionalVip `json:"additional_vips"`
}

// AdditionalVip represent additional ip of a loadbalancer. IpAddress field is optional.
type AdditionalVip struct {
	SubnetID  string `json:"subnet_id"`
	IPAddress string `json:"ip_address,omitempty"`
}

func (r *LoadBalancer) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Support for older neutron time format

// Support for newer neutron time format

// StatusTree represents the status of a loadbalancer.
type StatusTree struct {
	Loadbalancer *LoadBalancer `json:"loadbalancer"`
}

type Stats struct {
	// The currently active connections.
	ActiveConnections int `json:"active_connections"`

	// The total bytes received.
	BytesIn int `json:"bytes_in"`

	// The total bytes sent.
	BytesOut int `json:"bytes_out"`

	// The total requests that were unable to be fulfilled.
	RequestErrors int `json:"request_errors"`

	// The total connections handled.
	TotalConnections int `json:"total_connections"`
}

// LoadBalancerPage is the page returned by a pager when traversing over a
// collection of load balancers.
type LoadBalancerPage struct {
	pagination.LinkedPageBase
}

// NextPageURL is invoked when a paginated collection of load balancers has
// reached the end of a page and the pager seeks to traverse over a new one.
// In order to do this, it needs to construct the next page's URL.
func (r LoadBalancerPage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// IsEmpty checks whether a LoadBalancerPage struct is empty.
func (r LoadBalancerPage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractLoadBalancers accepts a Page struct, specifically a LoadbalancerPage
// struct, and extracts the elements into a slice of LoadBalancer structs. In
// other words, a generic collection is mapped into a relevant slice.
func ExtractLoadBalancers(r pagination.Page) ([]LoadBalancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type commonResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts a loadbalancer.
func (r commonResult) Extract() (*LoadBalancer, error) { _ = "STUB: not implemented"; return nil, nil }

// GetStatusesResult represents the result of a GetStatuses operation.
// Call its Extract method to interpret it as a StatusTree.
type GetStatusesResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts the status of
// a Loadbalancer.
func (r GetStatusesResult) Extract() (*StatusTree, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StatsResult represents the result of a GetStats operation.
// Call its Extract method to interpret it as a Stats.
type StatsResult struct {
	gophercloud.Result
}

// Extract is a function that accepts a result and extracts the status of
// a Loadbalancer.
func (r StatsResult) Extract() (*Stats, error) { _ = "STUB: not implemented"; return nil, nil }

// CreateResult represents the result of a create operation. Call its Extract
// method to interpret it as a LoadBalancer.
type CreateResult struct {
	commonResult
}

// GetResult represents the result of a get operation. Call its Extract
// method to interpret it as a LoadBalancer.
type GetResult struct {
	commonResult
}

// UpdateResult represents the result of an update operation. Call its Extract
// method to interpret it as a LoadBalancer.
type UpdateResult struct {
	commonResult
}

// DeleteResult represents the result of a delete operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// FailoverResult represents the result of a failover operation. Call its
// ExtractErr method to determine if the request succeeded or failed.
type FailoverResult struct {
	gophercloud.ErrResult
}
