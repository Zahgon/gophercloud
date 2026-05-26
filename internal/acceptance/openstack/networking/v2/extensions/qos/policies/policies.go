package policies

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/qos/policies"
)

// CreateQoSPolicy will create a QoS policy. An error will be returned if the
// QoS policy could not be created.
func CreateQoSPolicy(t *testing.T, client *gophercloud.ServiceClient) (*policies.Policy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteQoSPolicy will delete a QoS policy with a specified ID.
// A fatal error will occur if the delete was not successful.
func DeleteQoSPolicy(t *testing.T, client *gophercloud.ServiceClient, policyID string) {
	_ = "STUB: not implemented"
	return
}
