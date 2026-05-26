package rules

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/qos/rules"
)

// CreateBandwidthLimitRule will create a QoS BandwidthLimitRule associated with the provided QoS policy.
// An error will be returned if the QoS rule could not be created.
func CreateBandwidthLimitRule(t *testing.T, client *gophercloud.ServiceClient, policyID string) (*rules.BandwidthLimitRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateDSCPMarkingRule will create a QoS DSCPMarkingRule associated with the provided QoS policy.
// An error will be returned if the QoS rule could not be created.
func CreateDSCPMarkingRule(t *testing.T, client *gophercloud.ServiceClient, policyID string) (*rules.DSCPMarkingRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMinimumBandwidthRule will create a QoS MinimumBandwidthRule associated with the provided QoS policy.
// An error will be returned if the QoS rule could not be created.
func CreateMinimumBandwidthRule(t *testing.T, client *gophercloud.ServiceClient, policyID string) (*rules.MinimumBandwidthRule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
