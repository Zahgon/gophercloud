package rbacpolicies

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/networking/v2/extensions/rbacpolicies"
)

// CreateRBACPolicy will create a rbac-policy. An error will be returned if the
// rbac-policy could not be created.
func CreateRBACPolicy(t *testing.T, client *gophercloud.ServiceClient, tenantID, networkID string) (*rbacpolicies.RBACPolicy, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteRBACPolicy will delete a rbac-policy with a specified ID. A fatal error will
// occur if the delete was not successful. This works best when used as a
// deferred function.
func DeleteRBACPolicy(t *testing.T, client *gophercloud.ServiceClient, rbacPolicyID string) {
	_ = "STUB: not implemented"
	return
}
