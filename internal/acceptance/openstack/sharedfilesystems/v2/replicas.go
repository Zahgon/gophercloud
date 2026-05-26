package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/replicas"
	"github.com/gophercloud/gophercloud/v2/openstack/sharedfilesystems/v2/shares"
)

// CreateReplica will create a replica from shareID. An error will be returned
// if the replica could not be created.
func CreateReplica(t *testing.T, client *gophercloud.ServiceClient, share *shares.Share) (*replicas.Replica, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteReplica will delete a replica. A fatal error will occur if the replica
// failed to be deleted. This works best when used as a deferred function.
func DeleteReplica(t *testing.T, client *gophercloud.ServiceClient, replica *replicas.Replica) {
	_ = "STUB: not implemented"
	return
}

// ListShareReplicas lists all replicas that belong to shareID.
// An error will be returned if the replicas could not be listed..
func ListShareReplicas(t *testing.T, client *gophercloud.ServiceClient, shareID string) ([]replicas.Replica, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func waitForReplicaStatus(t *testing.T, c *gophercloud.ServiceClient, id, status string) error {
	_ = "STUB: not implemented"
	return nil
}

func waitForReplicaState(t *testing.T, c *gophercloud.ServiceClient, id, state string) error {
	_ = "STUB: not implemented"
	return nil
}
