package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/availabilityzoneprofiles"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/flavorprofiles"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/flavors"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/l7policies"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/listeners"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/loadbalancers"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/monitors"
	"github.com/gophercloud/gophercloud/v2/openstack/loadbalancer/v2/pools"
)

// CreateListener will create a listener for a given load balancer on a random
// port with a random name. An error will be returned if the listener could not
// be created.
func CreateListener(t *testing.T, client *gophercloud.ServiceClient, lb *loadbalancers.LoadBalancer) (*listeners.Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateListenerHTTP will create an HTTP-based listener for a given load
// balancer on a random port with a random name. An error will be returned
// if the listener could not be created.
func CreateListenerHTTP(t *testing.T, client *gophercloud.ServiceClient, lb *loadbalancers.LoadBalancer) (*listeners.Listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// tls_version is only supported in microversion v2.17 introduced in victoria

// CreateLoadBalancer will create a load balancer with a random name on a given
// subnet. An error will be returned if the loadbalancer could not be created.
func CreateLoadBalancer(t *testing.T, client *gophercloud.ServiceClient, subnetID string, tags []string, policyID string, additionalVips []loadbalancers.AdditionalVip) (*loadbalancers.LoadBalancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateLoadBalancerFullyPopulated will create a  fully populated load balancer with a random name on a given
// subnet. It will contain a listener, l7policy, l7rule, pool, member and health monitor.
// An error will be returned if the loadbalancer could not be created.
func CreateLoadBalancerFullyPopulated(t *testing.T, client *gophercloud.ServiceClient, subnetID string, tags []string) (*loadbalancers.LoadBalancer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMember will create a member with a random name, port, address, and
// weight. An error will be returned if the member could not be created.
func CreateMember(t *testing.T, client *gophercloud.ServiceClient, lb *loadbalancers.LoadBalancer, pool *pools.Pool, subnetID, subnetCIDR string) (*pools.Member, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMonitor will create a monitor with a random name for a specific pool.
// An error will be returned if the monitor could not be created.
func CreateMonitor(t *testing.T, client *gophercloud.ServiceClient, lb *loadbalancers.LoadBalancer, pool *pools.Pool) (*monitors.Monitor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePool will create a pool with a random name with a specified listener
// and loadbalancer. An error will be returned if the pool could not be
// created.
func CreatePool(t *testing.T, client *gophercloud.ServiceClient, lb *loadbalancers.LoadBalancer) (*pools.Pool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePoolHTTP will create an HTTP-based pool with a random name with a
// specified listener and loadbalancer. An error will be returned if the pool
// could not be created.
func CreatePoolHTTP(t *testing.T, client *gophercloud.ServiceClient, lb *loadbalancers.LoadBalancer) (*pools.Pool, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateL7Policy will create a l7 policy with a random name with a specified listener
// and loadbalancer. An error will be returned if the l7 policy could not be
// created.
func CreateL7Policy(t *testing.T, client *gophercloud.ServiceClient, listener *listeners.Listener, lb *loadbalancers.LoadBalancer, tags []string) (*l7policies.L7Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateL7Rule creates a l7 rule for specified l7 policy.
func CreateL7Rule(t *testing.T, client *gophercloud.ServiceClient, policyID string, lb *loadbalancers.LoadBalancer, tags []string) (*l7policies.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteL7Policy will delete a specified l7 policy. A fatal error will occur if
// the l7 policy could not be deleted. This works best when used as a deferred
// function.
func DeleteL7Policy(t *testing.T, client *gophercloud.ServiceClient, lbID, policyID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteL7Rule will delete a specified l7 rule. A fatal error will occur if
// the l7 rule could not be deleted. This works best when used as a deferred
// function.
func DeleteL7Rule(t *testing.T, client *gophercloud.ServiceClient, lbID, policyID, ruleID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteListener will delete a specified listener. A fatal error will occur if
// the listener could not be deleted. This works best when used as a deferred
// function.
func DeleteListener(t *testing.T, client *gophercloud.ServiceClient, lbID, listenerID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteMember will delete a specified member. A fatal error will occur if the
// member could not be deleted. This works best when used as a deferred
// function.
func DeleteMember(t *testing.T, client *gophercloud.ServiceClient, lbID, poolID, memberID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteLoadBalancer will delete a specified loadbalancer. A fatal error will
// occur if the loadbalancer could not be deleted. This works best when used
// as a deferred function.
func DeleteLoadBalancer(t *testing.T, client *gophercloud.ServiceClient, lbID string) {
	_ = "STUB: not implemented"
	return
}

// CascadeDeleteLoadBalancer will perform a cascading delete on a loadbalancer.
// A fatal error will occur if the loadbalancer could not be deleted. This works
// best when used as a deferred function.
func CascadeDeleteLoadBalancer(t *testing.T, client *gophercloud.ServiceClient, lbID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteMonitor will delete a specified monitor. A fatal error will occur if
// the monitor could not be deleted. This works best when used as a deferred
// function.
func DeleteMonitor(t *testing.T, client *gophercloud.ServiceClient, lbID, monitorID string) {
	_ = "STUB: not implemented"
	return
}

// DeletePool will delete a specified pool. A fatal error will occur if the
// pool could not be deleted. This works best when used as a deferred function.
func DeletePool(t *testing.T, client *gophercloud.ServiceClient, lbID, poolID string) {
	_ = "STUB: not implemented"
	return
}

// WaitForLoadBalancerState will wait until a loadbalancer reaches a given state.
func WaitForLoadBalancerState(client *gophercloud.ServiceClient, lbID, status string) error {
	_ = "STUB: not implemented"
	return nil
}

func CreateFlavorProfile(t *testing.T, client *gophercloud.ServiceClient) (*flavorprofiles.FlavorProfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteFlavorProfile(t *testing.T, client *gophercloud.ServiceClient, flavorProfile *flavorprofiles.FlavorProfile) {
	_ = "STUB: not implemented"
	return
}

func CreateFlavor(t *testing.T, client *gophercloud.ServiceClient, flavorProfile *flavorprofiles.FlavorProfile) (*flavors.Flavor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteFlavor(t *testing.T, client *gophercloud.ServiceClient, flavor *flavors.Flavor) {
	_ = "STUB: not implemented"
	return
}

func CreateAvailabilityZoneProfile(t *testing.T, client *gophercloud.ServiceClient) (*availabilityzoneprofiles.AvailabilityZoneProfile, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DeleteAvailabilityZoneProfile(t *testing.T, client *gophercloud.ServiceClient, availabilityZoneProfile *availabilityzoneprofiles.AvailabilityZoneProfile) {
	_ = "STUB: not implemented"
	return
}
