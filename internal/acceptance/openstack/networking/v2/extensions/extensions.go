package extensions

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/security/addressgroups"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/security/groups"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/security/rules"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/networks"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/ports"
)

// CreateExternalNetwork will create an external network. An error will be
// returned if the creation failed.
func CreateExternalNetwork(t *testing.T, client *gophercloud.ServiceClient) (*networks.Network, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePortWithSecurityGroup will create a port with a security group
// attached. An error will be returned if the port could not be created.
func CreatePortWithSecurityGroup(t *testing.T, client *gophercloud.ServiceClient, networkID, subnetID, secGroupID string) (*ports.Port, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecurityGroup will create a security group with a random name.
// An error will be returned if one was failed to be created.
func CreateSecurityGroup(t *testing.T, client *gophercloud.ServiceClient) (*groups.SecGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecurityGroupRule will create a security group rule with a random name
// and random port between 80 and 99.
// An error will be returned if one was failed to be created.
func CreateSecurityGroupRule(t *testing.T, client *gophercloud.ServiceClient, secGroupID string) (*rules.SecGroupRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecurityGroupRulesBulk will create 3 security group rules with ports between 1080 and 1099.
// An error will be returned if one was failed to be created.
func CreateSecurityGroupRulesBulk(t *testing.T, client *gophercloud.ServiceClient, secGroupID string) ([]rules.SecGroupRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSecurityGroup will delete a security group of a specified ID.
// A fatal error will occur if the deletion failed. This works best as a
// deferred function
func DeleteSecurityGroup(t *testing.T, client *gophercloud.ServiceClient, secGroupID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteSecurityGroupRule will delete a security group rule of a specified ID.
// A fatal error will occur if the deletion failed. This works best as a
// deferred function
func DeleteSecurityGroupRule(t *testing.T, client *gophercloud.ServiceClient, ruleID string) {
	_ = "STUB: not implemented"
	return
}

// CreateSecurityAddressGroup will create a security address group with a random name.
func CreateSecurityAddressGroup(t *testing.T, client *gophercloud.ServiceClient) (*addressgroups.AddressGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteSecurityAddressGroup will delete a security address group of a specified ID.
// A fatal error will occur if the deletion failed. This works best as a
// deferred function
func DeleteSecurityAddressGroup(t *testing.T, client *gophercloud.ServiceClient, addressGroupID string) {
	_ = "STUB: not implemented"
	return
}
