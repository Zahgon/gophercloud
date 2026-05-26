// Package v2 contains common functions for creating compute-based resources
// for use in acceptance tests. See the `*_test.go` files for example usages.
package v2

import (
	"testing"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/openstack/blockstorage/v2/volumes"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/aggregates"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/attachinterfaces"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/flavors"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/keypairs"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/quotasets"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/remoteconsoles"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/secgroups"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servergroups"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/v2/openstack/compute/v2/volumeattach"
)

// AttachInterface will create and attach an interface on a given server.
// An error will returned if the interface could not be created.
func AttachInterface(t *testing.T, client *gophercloud.ServiceClient, serverID string) (*attachinterfaces.Interface, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateAggregate will create an aggregate with random name and available zone.
// An error will be returned if the aggregate could not be created.
func CreateAggregate(t *testing.T, client *gophercloud.ServiceClient) (*aggregates.Aggregate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateBootableVolumeServer works like CreateServer but is configured with
// one or more block devices defined by passing in []servers.BlockDevice.
// An error will be returned if a server was unable to be created.
func CreateBootableVolumeServer(t *testing.T, client *gophercloud.ServiceClient, blockDevices []servers.BlockDevice) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateFlavor will create a flavor with a random name.
// An error will be returned if the flavor could not be created.
func CreateFlavor(t *testing.T, client *gophercloud.ServiceClient) (*flavors.Flavor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Microversion 2.55 is required to add description to flavor

func createKey() (string, error) { _ = "STUB: not implemented"; return "", nil }

// CreateKeyPair will create a KeyPair with a random name. An error will occur
// if the keypair failed to be created. An error will be returned if the
// keypair was unable to be created.
func CreateKeyPair(t *testing.T, client *gophercloud.ServiceClient) (*keypairs.KeyPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMultiEphemeralServer works like CreateServer but is configured with
// one or more block devices defined by passing in []servers.BlockDevice.
// These block devices act like block devices when booting from a volume but
// are actually local ephemeral disks.
// An error will be returned if a server was unable to be created.
func CreateMultiEphemeralServer(t *testing.T, client *gophercloud.ServiceClient, blockDevices []servers.BlockDevice) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreatePrivateFlavor will create a private flavor with a random name.
// An error will be returned if the flavor could not be created.
func CreatePrivateFlavor(t *testing.T, client *gophercloud.ServiceClient) (*flavors.Flavor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecurityGroup will create a security group with a random name.
// An error will be returned if one was failed to be created.
func CreateSecurityGroup(t *testing.T, client *gophercloud.ServiceClient) (*secgroups.SecurityGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateSecurityGroupRule will create a security group rule with a random name
// and a random TCP port range between port 80 and 99. An error will be
// returned if the rule failed to be created.
func CreateSecurityGroupRule(t *testing.T, client *gophercloud.ServiceClient, securityGroupID string) (*secgroups.Rule, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateServer creates a basic instance with a randomly generated name.
// The flavor of the instance will be the value of the OS_FLAVOR_ID environment variable.
// The image will be the value of the OS_IMAGE_ID environment variable.
// The instance will be launched on the network specified in OS_NETWORK_NAME.
// An error will be returned if the instance was unable to be created.
func CreateServer(t *testing.T, client *gophercloud.ServiceClient) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateMicroversionServer creates a basic instance compatible with
// newer microversions with a randomly generated name.
// The flavor of the instance will be the value of the OS_FLAVOR_ID environment variable.
// The image will be the value of the OS_IMAGE_ID environment variable.
// The instance will be launched on the network specified in OS_NETWORK_NAME.
// An error will be returned if the instance was unable to be created.
func CreateMicroversionServer(t *testing.T, client *gophercloud.ServiceClient) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateServerWithoutImageRef creates a basic instance with a randomly generated name.
// The flavor of the instance will be the value of the OS_FLAVOR_ID environment variable.
// The image is intentionally missing to trigger an error.
// The instance will be launched on the network specified in OS_NETWORK_NAME.
// An error will be returned if the instance was unable to be created.
func CreateServerWithoutImageRef(t *testing.T, client *gophercloud.ServiceClient) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateServerWithTags creates a basic instance with a randomly generated name.
// The flavor of the instance will be the value of the OS_FLAVOR_ID environment variable.
// The image will be the value of the OS_IMAGE_ID environment variable.
// The instance will be launched on the network specified in OS_NETWORK_NAME.
// Two tags will be assigned to the server.
// An error will be returned if the instance was unable to be created.
func CreateServerWithTags(t *testing.T, client *gophercloud.ServiceClient, networkID string) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateServerGroup will create a server with a random name. An error will be
// returned if the server group failed to be created.
func CreateServerGroup(t *testing.T, client *gophercloud.ServiceClient, policy string) (*servergroups.ServerGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateServerGroupMicroversion will create a server with a random name using 2.64 microversion. An error will be
// returned if the server group failed to be created.
func CreateServerGroupMicroversion(t *testing.T, client *gophercloud.ServiceClient) (*servergroups.ServerGroup, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateServerInServerGroup works like CreateServer but places the instance in
// a specified Server Group.
func CreateServerInServerGroup(t *testing.T, client *gophercloud.ServiceClient, serverGroup *servergroups.ServerGroup) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateServerWithPublicKey works the same as CreateServer, but additionally
// configures the server with a specified Key Pair name.
func CreateServerWithPublicKey(t *testing.T, client *gophercloud.ServiceClient, keyPairName string) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateVolumeAttachment will attach a volume to a server. An error will be
// returned if the volume failed to attach.
func CreateVolumeAttachment(t *testing.T, client *gophercloud.ServiceClient, blockClient *gophercloud.ServiceClient, server *servers.Server, volume *volumes.Volume) (*volumeattach.VolumeAttachment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// DeleteAggregate will delete a given host aggregate. A fatal error will occur if
// the aggregate deleting is failed. This works best when using it as a
// deferred function.
func DeleteAggregate(t *testing.T, client *gophercloud.ServiceClient, aggregate *aggregates.Aggregate) {
	_ = "STUB: not implemented"
	return
}

// DeleteFlavor will delete a flavor. A fatal error will occur if the flavor
// could not be deleted. This works best when using it as a deferred function.
func DeleteFlavor(t *testing.T, client *gophercloud.ServiceClient, flavor *flavors.Flavor) {
	_ = "STUB: not implemented"
	return
}

// DeleteKeyPair will delete a specified keypair. A fatal error will occur if
// the keypair failed to be deleted. This works best when used as a deferred
// function.
func DeleteKeyPair(t *testing.T, client *gophercloud.ServiceClient, keyPair *keypairs.KeyPair) {
	_ = "STUB: not implemented"
	return
}

// DeleteSecurityGroup will delete a security group. A fatal error will occur
// if the group failed to be deleted. This works best as a deferred function.
func DeleteSecurityGroup(t *testing.T, client *gophercloud.ServiceClient, securityGroupID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteSecurityGroupRule will delete a security group rule. A fatal error
// will occur if the rule failed to be deleted. This works best when used
// as a deferred function.
func DeleteSecurityGroupRule(t *testing.T, client *gophercloud.ServiceClient, ruleID string) {
	_ = "STUB: not implemented"
	return
}

// DeleteServer deletes an instance via its UUID.
// A fatal error will occur if the instance failed to be destroyed. This works
// best when using it as a deferred function.
func DeleteServer(t *testing.T, client *gophercloud.ServiceClient, server *servers.Server) {
	_ = "STUB: not implemented"
	return
}

// If we reach this point, the API returned an actual DELETED status
// which is a very short window of time, but happens occasionally.

// DeleteServerGroup will delete a server group. A fatal error will occur if
// the server group failed to be deleted. This works best when used as a
// deferred function.
func DeleteServerGroup(t *testing.T, client *gophercloud.ServiceClient, serverGroup *servergroups.ServerGroup) {
	_ = "STUB: not implemented"
	return
}

// DeleteVolumeAttachment will disconnect a volume from an instance. A fatal
// error will occur if the volume failed to detach. This works best when used
// as a deferred function.
func DeleteVolumeAttachment(t *testing.T, client *gophercloud.ServiceClient, blockClient *gophercloud.ServiceClient, server *servers.Server, volumeAttachment *volumeattach.VolumeAttachment) {
	_ = "STUB: not implemented"
	return
}

// DetachInterface will detach an interface from a server. A fatal
// error will occur if the interface could not be detached. This works best
// when used as a deferred function.
func DetachInterface(t *testing.T, client *gophercloud.ServiceClient, serverID, portID string) {
	_ = "STUB: not implemented"
	return
}

// GetNetworkIDFromNetworks will return the network UUID for a given network
// name using the Neutron API.
// An error will be returned if the network could not be retrieved.
func GetNetworkIDFromNetworks(t *testing.T, client *gophercloud.ServiceClient, networkName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ImportPublicKey will create a KeyPair with a random name and a specified
// public key. An error will be returned if the keypair failed to be created.
func ImportPublicKey(t *testing.T, client *gophercloud.ServiceClient, publicKey string) (*keypairs.KeyPair, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResizeServer performs a resize action on an instance. An error will be
// returned if the instance failed to resize.
// The new flavor that the instance will be resized to is specified in OS_FLAVOR_ID_RESIZE.
func ResizeServer(t *testing.T, client *gophercloud.ServiceClient, server *servers.Server) error {
	_ = "STUB: not implemented"
	return nil
}

// WaitForComputeStatus will poll an instance's status until it either matches
// the specified status or the status becomes ERROR.
func WaitForComputeStatus(client *gophercloud.ServiceClient, server *servers.Server, status string) error {
	_ = "STUB: not implemented"
	return nil
}

// Success!

// Convenience method to fill an QuotaSet-UpdateOpts-struct from a QuotaSet-struct
func FillUpdateOptsFromQuotaSet(src quotasets.QuotaSet, dest *quotasets.UpdateOpts) {
	_ = "STUB: not implemented"
	return
}

// RescueServer will place the specified server into rescue mode.
func RescueServer(t *testing.T, client *gophercloud.ServiceClient, server *servers.Server) error {
	_ = "STUB: not implemented"
	return nil
}

// UnrescueServer will return server from rescue mode.
func UnrescueServer(t *testing.T, client *gophercloud.ServiceClient, server *servers.Server) error {
	_ = "STUB: not implemented"
	return nil
}

// CreateRemoteConsole will create a remote noVNC console for the specified server.
func CreateRemoteConsole(t *testing.T, client *gophercloud.ServiceClient, serverID string) (*remoteconsoles.RemoteConsole, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CreateNoNetworkServer creates a basic instance with a randomly generated name.
// The flavor of the instance will be the value of the OS_FLAVOR_ID environment variable.
// The image will be the value of the OS_IMAGE_ID environment variable.
// The instance will be launched without network interfaces attached.
// An error will be returned if the instance was unable to be created.
func CreateServerNoNetwork(t *testing.T, client *gophercloud.ServiceClient) (*servers.Server, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
