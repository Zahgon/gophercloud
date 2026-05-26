// Package v2 contains common functions for creating db resources for use
// in acceptance tests. See the `*_test.go` files for example usages.
package v1

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/db/v1/instances"
)

// CreateDatabase will create a database with a randomly generated name.
// An error will be returned if the database was unable to be created.
func CreateDatabase(t *testing.T, client *gophercloud.ServiceClient, instanceID string) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateInstance will create an instance with a randomly generated name.
// The flavor of the instance will be the value of the OS_FLAVOR_ID
// environment variable. The Datastore will be pulled from the
// OS_DATASTORE_TYPE_ID environment variable.
// An error will be returned if the instance was unable to be created.
func CreateInstance(t *testing.T, client *gophercloud.ServiceClient) (*instances.Instance, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateUser will create a user with a randomly generated name.
// An error will be returned if the user was unable to be created.
func CreateUser(t *testing.T, client *gophercloud.ServiceClient, instanceID string) error {
	_ = "STUB: not implemented"
	return nil
}

// DeleteDatabase deletes a database. A fatal error will occur if the database
// failed to delete. This works best when used as a deferred function.
func DeleteDatabase(t *testing.T, client *gophercloud.ServiceClient, instanceID, name string) {
	_ = "STUB: not implemented"
	return
}

// DeleteInstance deletes an instance. A fatal error will occur if the instance
// failed to delete. This works best when used as a deferred function.
func DeleteInstance(t *testing.T, client *gophercloud.ServiceClient, id string) {
	_ = "STUB: not implemented"
	return
}

// DeleteUser deletes a user. A fatal error will occur if the user
// failed to delete. This works best when used as a deferred function.
func DeleteUser(t *testing.T, client *gophercloud.ServiceClient, instanceID, name string) {
	_ = "STUB: not implemented"
	return
}

// WaitForInstanceState will poll an instance's status until it either matches
// the specified status or the status becomes ERROR.
func WaitForInstanceStatus(
	client *gophercloud.ServiceClient, instance *instances.Instance, status string) error {
	_ = "STUB: not implemented"
	return nil
}
