// Package clients contains functions for creating OpenStack service clients
// for use in acceptance tests. It also manages the required environment
// variables to run the tests.
package clients

import (
	"github.com/gophercloud/gophercloud/v2"
)

// AcceptanceTestChoices contains image and flavor selections for use by the acceptance tests.
type AcceptanceTestChoices struct {
	// ImageID contains the ID of a valid image.
	ImageID string

	// FlavorID contains the ID of a valid flavor.
	FlavorID string

	// FlavorIDResize contains the ID of a different flavor available on the same OpenStack installation, that is distinct
	// from FlavorID.
	FlavorIDResize string

	// FloatingIPPool contains the name of the pool from where to obtain floating IPs.
	FloatingIPPoolName string

	// MagnumKeypair contains the ID of a valid key pair.
	MagnumKeypair string

	// MagnumImageID contains the ID of a valid magnum image.
	MagnumImageID string

	// NetworkName is the name of a network to launch the instance on.
	NetworkName string

	// NetworkID is the ID of a network to launch the instance on.
	NetworkID string

	// SubnetID is the ID of a subnet to launch the instance on.
	SubnetID string

	// ExternalNetworkID is the network ID of the external network.
	ExternalNetworkID string

	// DBDatastoreType is the datastore type for DB tests.
	DBDatastoreType string

	// DBDatastoreTypeID is the datastore type version for DB tests.
	DBDatastoreVersion string
}

// AcceptanceTestChoicesFromEnv populates a ComputeChoices struct from environment variables.
// If any required state is missing, an `error` will be returned that enumerates the missing properties.
func AcceptanceTestChoicesFromEnv() (*AcceptanceTestChoices, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

/* // Temporarily disabled, see https://github.com/gophercloud/gophercloud/issues/1345
if networkID == "" {
	missing = append(missing, "OS_NETWORK_ID")
}
if subnetID == "" {
	missing = append(missing, "OS_SUBNET_ID")
}
*/

// NewBlockStorageV1Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBlockStorageV2Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBlockStorageV3Client returns a *ServiceClient for making calls
// to the OpenStack Block Storage v3 API. An error will be returned
// if authentication or client creation was not possible.
func NewBlockStorageV3Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBlockStorageV2NoAuthClient returns a noauth *ServiceClient for
// making calls to the OpenStack Block Storage v2 API. An error will be
// returned if client creation was not possible.
func NewBlockStorageV2NoAuthClient() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBlockStorageV3NoAuthClient returns a noauth *ServiceClient for
// making calls to the OpenStack Block Storage v3 API. An error will be
// returned if client creation was not possible.
func NewBlockStorageV3NoAuthClient() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewComputeV2Client returns a *ServiceClient for making calls
// to the OpenStack Compute v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewComputeV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBareMetalV1Client returns a *ServiceClient for making calls
// to the OpenStack Bare Metal v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewBareMetalV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBareMetalV1NoAuthClient returns a *ServiceClient for making calls
// to the OpenStack Bare Metal v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewBareMetalV1NoAuthClient() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBareMetalV1HTTPBasic returns a *ServiceClient for making calls
// to the OpenStack Bare Metal v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewBareMetalV1HTTPBasic() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBareMetalIntrospectionV1Client returns a *ServiceClient for making calls
// to the OpenStack Bare Metal Introspection v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewBareMetalIntrospectionV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDBV1Client returns a *ServiceClient for making calls
// to the OpenStack Database v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewDBV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDNSV2Client returns a *ServiceClient for making calls
// to the OpenStack Compute v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewDNSV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIdentityV2Client returns a *ServiceClient for making calls
// to the OpenStack Identity v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewIdentityV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIdentityV2AdminClient returns a *ServiceClient for making calls
// to the Admin Endpoint of the OpenStack Identity v2 API. An error
// will be returned if authentication or client creation was not possible.
func NewIdentityV2AdminClient() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIdentityV2UnauthenticatedClient returns an unauthenticated *ServiceClient
// for the OpenStack Identity v2 API. An error  will be returned if
// authentication or client creation was not possible.
func NewIdentityV2UnauthenticatedClient() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIdentityV3Client returns a *ServiceClient for making calls
// to the OpenStack Identity v3 API. An error will be returned
// if authentication or client creation was not possible.
func NewIdentityV3Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIdentityV3UnauthenticatedClient returns an unauthenticated *ServiceClient
// for the OpenStack Identity v3 API. An error  will be returned if
// authentication or client creation was not possible.
func NewIdentityV3UnauthenticatedClient() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewImageV2Client returns a *ServiceClient for making calls to the
// OpenStack Image v2 API. An error will be returned if authentication or
// client creation was not possible.
func NewImageV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewNetworkV2Client returns a *ServiceClient for making calls to the
// OpenStack Networking v2 API. An error will be returned if authentication
// or client creation was not possible.
func NewNetworkV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewObjectStorageV1Client returns a *ServiceClient for making calls to the
// OpenStack Object Storage v1 API. An error will be returned if authentication
// or client creation was not possible.
func NewObjectStorageV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSharedFileSystemV2Client returns a *ServiceClient for making calls
// to the OpenStack Shared File System v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewSharedFileSystemV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewLoadBalancerV2Client returns a *ServiceClient for making calls to the
// OpenStack Octavia v2 API. An error will be returned if authentication
// or client creation was not possible.
func NewLoadBalancerV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMessagingV2Client returns a *ServiceClient for making calls
// to the OpenStack Messaging (Zaqar) v2 API. An error will be returned
// if authentication or client creation was not possible.
func NewMessagingV2Client(clientID string) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMetricV1Client returns a *ServiceClient for making calls
// to the OpenStack Metric v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewMetricV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewContainerV1Client returns a *ServiceClient for making calls
// to the OpenStack Container V1 API. An error will be returned
// if authentication or client creation was not possible.
func NewContainerV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewKeyManagerV1Client returns a *ServiceClient for making calls
// to the OpenStack Key Manager (Barbican) v1 API. An error will be
// returned if authentication or client creation was not possible.
func NewKeyManagerV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// configureDebug will configure the provider client to print the API
// requests and responses if OS_DEBUG is enabled.
func configureDebug(client *gophercloud.ProviderClient) *gophercloud.ProviderClient {
	_ = "STUB: not implemented"
	return nil
}

// NewContainerInfraV1Client returns a *ServiceClient for making calls
// to the OpenStack Container Infra Management v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewContainerInfraV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWorkflowV2Client returns a *ServiceClient for making calls
// to the OpenStack Workflow v2 API (Mistral). An error will be returned if
// authentication or client creation failed.
func NewWorkflowV2Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewOrchestrationV1Client returns a *ServiceClient for making calls
// to the OpenStack Orchestration v1 API. An error will be returned
// if authentication or client creation was not possible.
func NewOrchestrationV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewPlacementV1Client returns a *ServiceClient for making calls
// to the OpenStack Placement API. An error will be returned
// if authentication or client creation was not possible.
func NewPlacementV1Client() (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
