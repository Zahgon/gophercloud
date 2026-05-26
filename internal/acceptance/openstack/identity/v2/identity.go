// Package v2 contains common functions for creating identity-based resources
// for use in acceptance tests. See the `*_test.go` files for example usages.
package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v2/roles"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v2/tenants"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v2/users"
)

// AddUserRole will grant a role to a user in a tenant. An error will be
// returned if the grant was unsuccessful.
func AddUserRole(t *testing.T, client *gophercloud.ServiceClient, tenant *tenants.Tenant, user *users.User, role *roles.Role) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateTenant will create a project with a random name.
// It takes an optional createOpts parameter since creating a project
// has so many options. An error will be returned if the project was
// unable to be created.
func CreateTenant(t *testing.T, client *gophercloud.ServiceClient, c *tenants.CreateOpts) (*tenants.Tenant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateUser will create a user with a random name and adds them to the given
// tenant. An error will be returned if the user was unable to be created.
func CreateUser(t *testing.T, client *gophercloud.ServiceClient, tenant *tenants.Tenant) (*users.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteTenant will delete a tenant by ID. A fatal error will occur if
// the tenant ID failed to be deleted. This works best when using it as
// a deferred function.
func DeleteTenant(t *testing.T, client *gophercloud.ServiceClient, tenantID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteUser will delete a user. A fatal error will occur if the delete was
// unsuccessful. This works best when used as a deferred function.
func DeleteUser(t *testing.T, client *gophercloud.ServiceClient, user *users.User) {
	_ = "STUB: not implemented"
	return
}

// DeleteUserRole will revoke a role of a user in a tenant. A fatal error will
// occur if the revoke was unsuccessful. This works best when used as a
// deferred function.
func DeleteUserRole(t *testing.T, client *gophercloud.ServiceClient, tenant *tenants.Tenant, user *users.User, role *roles.Role) {
	_ = "STUB: not implemented"
	return
}

// FindRole finds all roles that the current authenticated client has access
// to and returns the first one found. An error will be returned if the lookup
// was unsuccessful.
func FindRole(t *testing.T, client *gophercloud.ServiceClient) (*roles.Role, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FindTenant finds all tenants that the current authenticated client has access
// to and returns the first one found. An error will be returned if the lookup
// was unsuccessful.
func FindTenant(t *testing.T, client *gophercloud.ServiceClient) (*tenants.Tenant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// UpdateUser will update an existing user with a new randomly generated name.
// An error will be returned if the update was unsuccessful.
func UpdateUser(t *testing.T, client *gophercloud.ServiceClient, user *users.User) (*users.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
