package v3

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/domains"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/endpoints"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/groups"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/regions"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/services"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/trusts"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/users"
)

// CreateEndpoint will create an endpoint with a random name.  It only define
// endpoint name and description, and receive from CreateOpts others parameters,
// such as URL, Availability and ServiceID. An error will be returned if the
// endpoint was unabled to be created.
func CreateEndpoint(t *testing.T, client *gophercloud.ServiceClient, c *endpoints.CreateOpts) (*endpoints.Endpoint, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateProject will create a project with a random name.
// It takes an optional createOpts parameter since creating a project
// has so many options. An error will be returned if the project was
// unable to be created.
func CreateProject(t *testing.T, client *gophercloud.ServiceClient, c *projects.CreateOpts) (*projects.Project, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateUser will create a user with a random name.
// It takes an optional createOpts parameter since creating a user
// has so many options. An error will be returned if the user was
// unable to be created.
func CreateUser(t *testing.T, client *gophercloud.ServiceClient, c *users.CreateOpts) (*users.User, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateGroup will create a group with a random name.
// It takes an optional createOpts parameter since creating a group
// has so many options. An error will be returned if the group was
// unable to be created.
func CreateGroup(t *testing.T, client *gophercloud.ServiceClient, c *groups.CreateOpts) (*groups.Group, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateDomain will create a domain with a random name.
// It takes an optional createOpts parameter since creating a domain
// has many options. An error will be returned if the domain was
// unable to be created.
func CreateDomain(t *testing.T, client *gophercloud.ServiceClient, c *domains.CreateOpts) (*domains.Domain, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRole will create a role with a random name.
// It takes an optional createOpts parameter since creating a role
// has so many options. An error will be returned if the role was
// unable to be created.
func CreateRole(t *testing.T, client *gophercloud.ServiceClient, c *roles.CreateOpts) (*roles.Role, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateRegion will create a region with a random name.
// It takes an optional createOpts parameter since creating a region
// has so many options. An error will be returned if the region was
// unable to be created.
func CreateRegion(t *testing.T, client *gophercloud.ServiceClient, c *regions.CreateOpts) (*regions.Region, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateService will create a service with a random name.
// It takes an optional createOpts parameter since creating a service
// has so many options. An error will be returned if the service was
// unable to be created.
func CreateService(t *testing.T, client *gophercloud.ServiceClient, c *services.CreateOpts) (*services.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteEndpoint will delete the specified Endpoint using its ID. A fatal error
// will occur if the endpoint failed to be deleted. This works best when using
// it as a deferred function.
func DeleteEndpoint(t *testing.T, client *gophercloud.ServiceClient, endpointID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteProject will delete a project by ID. A fatal error will occur if
// the project ID failed to be deleted. This works best when using it as
// a deferred function.
func DeleteProject(t *testing.T, client *gophercloud.ServiceClient, projectID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteUser will delete a user by ID. A fatal error will occur if
// the user failed to be deleted. This works best when using it as
// a deferred function.
func DeleteUser(t *testing.T, client *gophercloud.ServiceClient, userID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteGroup will delete a group by ID. A fatal error will occur if
// the group failed to be deleted. This works best when using it as
// a deferred function.
func DeleteGroup(t *testing.T, client *gophercloud.ServiceClient, groupID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteDomain will delete a domain by ID. A fatal error will occur if
// the project ID failed to be deleted. This works best when using it as
// a deferred function.
func DeleteDomain(t *testing.T, client *gophercloud.ServiceClient, domainID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteRole will delete a role by ID. A fatal error will occur if
// the role failed to be deleted. This works best when using it as
// a deferred function.
func DeleteRole(t *testing.T, client *gophercloud.ServiceClient, roleID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteRegion will delete a reg by ID. A fatal error will occur if
// the region failed to be deleted. This works best when using it as
// a deferred function.
func DeleteRegion(t *testing.T, client *gophercloud.ServiceClient, regionID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteService will delete a reg by ID. A fatal error will occur if
// the service failed to be deleted. This works best when using it as
// a deferred function.
func DeleteService(t *testing.T, client *gophercloud.ServiceClient, serviceID string) {
	_ = "STUB: not implemented"
	return
}

// UnassignRole will delete a role assigned to a user/group on a project/domain
// A fatal error will occur if it fails to delete the assignment.
// This works best when using it as a deferred function.
func UnassignRole(t *testing.T, client *gophercloud.ServiceClient, roleID string, opts *roles.UnassignOpts) {
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

// CreateTrust will create a trust with the provided options.
// An error will be returned if the trust was unable to be created.
func CreateTrust(t *testing.T, client *gophercloud.ServiceClient, createOpts trusts.CreateOpts) (*trusts.Trust, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteTrust will delete a trust by ID. A fatal error will occur if
// the trust failed to be deleted. This works best when using it as
// a deferred function.
func DeleteTrust(t *testing.T, client *gophercloud.ServiceClient, trustID string) {
	_ = "STUB: not implemented"
	return
}

// FindTrust finds all trusts that the current authenticated client has access
// to and returns the first one found. An error will be returned if the lookup
// was unsuccessful.
func FindTrust(t *testing.T, client *gophercloud.ServiceClient) (*trusts.Trust, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
