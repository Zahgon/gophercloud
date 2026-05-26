package fwaas_v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/fwaas_v2/groups"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/fwaas_v2/policies"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/fwaas_v2/rules"
)

// RemoveRule will remove a rule from the  policy.
func RemoveRule(t *testing.T, client *gophercloud.ServiceClient, policyID string, ruleID string) {
	_ = "STUB: not implemented"
	return
}

// AddRule will add a rule to to a policy.
func AddRule(t *testing.T, client *gophercloud.ServiceClient, policyID string, ruleID string, beforeRuleID string) {
	_ = "STUB: not implemented"
	return
}

// CreatePolicy will create a Firewall Policy with a random name and given
// rule. An error will be returned if the rule could not be created.
func CreatePolicy(t *testing.T, client *gophercloud.ServiceClient, ruleID string) (*policies.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRule will create a Firewall Rule with a random source address and
// source port, destination address and port. An error will be returned if
// the rule could not be created.
func CreateRule(t *testing.T, client *gophercloud.ServiceClient) (*rules.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeletePolicy will delete a policy with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeletePolicy(t *testing.T, client *gophercloud.ServiceClient, policyID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteRule will delete a rule with a specified ID. A fatal error will occur
// if the delete was not successful. This works best when used as a deferred
// function.
func DeleteRule(t *testing.T, client *gophercloud.ServiceClient, ruleID string) {
	_ = "STUB: not implemented"
	return
}

// CreateGroup will create a Firewall Group. An error will be returned if the
// firewall group could not be created.
func CreateGroup(t *testing.T, client *gophercloud.ServiceClient) (*groups.Group, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteGroup will delete a group with a specified ID. A fatal error will occur
// if the delete was not successful. This works best when used as a deferred
// function.
func DeleteGroup(t *testing.T, client *gophercloud.ServiceClient, groupId string) {
	_ = "STUB: not implemented"
	return
}
