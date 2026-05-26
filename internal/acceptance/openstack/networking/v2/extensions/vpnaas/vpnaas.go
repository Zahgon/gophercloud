package vpnaas

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/vpnaas/endpointgroups"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/vpnaas/ikepolicies"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/vpnaas/ipsecpolicies"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/vpnaas/services"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/vpnaas/siteconnections"
)

// CreateService will create a Service with a random name and a specified router ID
// An error will be returned if the service could not be created.
func CreateService(t *testing.T, client *gophercloud.ServiceClient, routerID string) (*services.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteService will delete a service with a specified ID. A fatal error
// will occur if the delete was not successful. This works best when used as
// a deferred function.
func DeleteService(t *testing.T, client *gophercloud.ServiceClient, serviceID string) {
	_ = "STUB: not implemented"
	return
}

// CreateIPSecPolicy will create an IPSec Policy with a random name and given
// rule. An error will be returned if the rule could not be created.
func CreateIPSecPolicy(t *testing.T, client *gophercloud.ServiceClient) (*ipsecpolicies.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateIKEPolicy will create an IKE Policy with a random name and given
// rule. An error will be returned if the policy could not be created.
func CreateIKEPolicy(t *testing.T, client *gophercloud.ServiceClient) (*ikepolicies.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteIPSecPolicy will delete an IPSec policy with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeleteIPSecPolicy(t *testing.T, client *gophercloud.ServiceClient, policyID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteIKEPolicy will delete an IKE policy with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeleteIKEPolicy(t *testing.T, client *gophercloud.ServiceClient, policyID string) {
	_ = "STUB: not implemented"
	return
}

// CreateEndpointGroup will create an endpoint group with a random name.
// An error will be returned if the group could not be created.
func CreateEndpointGroup(t *testing.T, client *gophercloud.ServiceClient) (*endpointgroups.EndpointGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateEndpointGroupWithCIDR will create an endpoint group with a random name and a specified CIDR.
// An error will be returned if the group could not be created.
func CreateEndpointGroupWithCIDR(t *testing.T, client *gophercloud.ServiceClient, cidr string) (*endpointgroups.EndpointGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteEndpointGroup will delete an Endpoint group with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeleteEndpointGroup(t *testing.T, client *gophercloud.ServiceClient, epGroupID string) {
	_ = "STUB: not implemented"
	return
}

// CreateEndpointGroupWithSubnet will create an endpoint group with a random name.
// An error will be returned if the group could not be created.
func CreateEndpointGroupWithSubnet(t *testing.T, client *gophercloud.ServiceClient, subnetID string) (*endpointgroups.EndpointGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSiteConnection will create an IPSec site connection with a random name and specified
// IKE policy, IPSec policy, service, peer EP group and local EP Group.
// An error will be returned if the connection could not be created.
func CreateSiteConnection(t *testing.T, client *gophercloud.ServiceClient, ikepolicyID string, ipsecpolicyID string, serviceID string, peerEPGroupID string, localEPGroupID string) (*siteconnections.Connection, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSiteConnection will delete an IPSec site connection with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeleteSiteConnection(t *testing.T, client *gophercloud.ServiceClient, siteConnectionID string) {
	_ = "STUB: not implemented"
	return
}
