package openstack

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	tokens2 "github.com/gophercloud/gophercloud/v2/openstack/identity/v2/tokens"
	tokens3 "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
)

const (
	// v2 represents Keystone v2.
	// It should never increase beyond 2.0.
	v2 = "v2.0"

	// v3 represents Keystone v3.
	// The version can be anything from v3 to v3.x.
	v3 = "v3"
)

// NewClient prepares an unauthenticated ProviderClient instance.
// Most users will probably prefer using the AuthenticatedClient function
// instead.
//
// This is useful if you wish to explicitly control the version of the identity
// service that's used for authentication explicitly, for example.
//
// A basic example of using this would be:
//
//	ao, err := openstack.AuthOptionsFromEnv()
//	provider, err := openstack.NewClient(ao.IdentityEndpoint)
//	client, err := openstack.NewIdentityV3(ctx, provider, gophercloud.EndpointOpts{})
func NewClient(endpoint string) (*gophercloud.ProviderClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// AuthenticatedClient logs in to an OpenStack cloud found at the identity endpoint
// specified by the options, acquires a token, and returns a Provider Client
// instance that's ready to operate.
//
// If the full path to a versioned identity endpoint was specified  (example:
// http://example.com:5000/v3), that path will be used as the endpoint to query.
//
// If a versionless endpoint was specified (example: http://example.com:5000/),
// the endpoint will be queried to determine which versions of the identity service
// are available, then chooses the most recent or most supported version.
//
// Example:
//
//	ao, err := openstack.AuthOptionsFromEnv()
//	provider, err := openstack.AuthenticatedClient(ctx, ao)
//	client, err := openstack.NewNetworkV2(ctx, provider, gophercloud.EndpointOpts{
//		Region: os.Getenv("OS_REGION_NAME"),
//	})
func AuthenticatedClient(ctx context.Context, options gophercloud.AuthOptions) (*gophercloud.ProviderClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Authenticate authenticates or re-authenticates against the most
// recent identity service supported at the provided endpoint.
func Authenticate(ctx context.Context, client *gophercloud.ProviderClient, options gophercloud.AuthOptions) error {
	_ = "STUB: not implemented"
	return nil
}

// The switch statement must be out of date from the versions list.

// AuthenticateV2 explicitly authenticates against the identity v2 endpoint.
func AuthenticateV2(ctx context.Context, client *gophercloud.ProviderClient, options tokens2.AuthOptionsBuilder, eo gophercloud.EndpointOpts) error {
	_ = "STUB: not implemented"
	return nil
}

type v2TokenNoReauth struct {
	tokens2.AuthOptionsBuilder
}

func (v2TokenNoReauth) CanReauth() bool { _ = "STUB: not implemented"; return false }

func v2auth(ctx context.Context, client *gophercloud.ProviderClient, endpoint string, options tokens2.AuthOptionsBuilder, eo gophercloud.EndpointOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
// this should retry authentication only once

// AuthenticateV3 explicitly authenticates against the identity v3 service.
func AuthenticateV3(ctx context.Context, client *gophercloud.ProviderClient, options tokens3.AuthOptionsBuilder, eo gophercloud.EndpointOpts) error {
	_ = "STUB: not implemented"
	return nil
}

func v3auth(ctx context.Context, client *gophercloud.ProviderClient, endpoint string, opts tokens3.AuthOptionsBuilder, eo gophercloud.EndpointOpts) error {
	_ = "STUB: not implemented"
	// Override the generated service endpoint with the one returned by the version endpoint.
	return nil
}

// passthroughToken allows to passthrough the token without a scope

// passing through the token ID without requesting a new scope

// here we're creating a throw-away client (tac). it's a copy of the user's provider client, but
// with the token and reauth func zeroed out. combined with setting `AllowReauth` to `false`,
// this should retry authentication only once

// NewIdentityV2 creates a ServiceClient that may be used to interact with the
// v2 identity service.
func NewIdentityV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewIdentityV3 creates a ServiceClient that may be used to access the v3
// identity service.
func NewIdentityV3(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Ensure endpoint still has a suffix of v3.
// This is because EndpointLocator might have found a versionless
// endpoint or the published endpoint is still /v2.0. In both
// cases, we need to fix the endpoint to point to /v3.

// TODO(stephenfin): Allow passing aliases to all New${SERVICE}V${VERSION} methods in v3
func initClientOpts(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts, clientType string, version int) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBareMetalV1 creates a ServiceClient that may be used with the v1
// bare metal package.
func NewBareMetalV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBareMetalIntrospectionV1 creates a ServiceClient that may be used with the v1
// bare metal introspection package.
func NewBareMetalIntrospectionV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewObjectStorageV1 creates a ServiceClient that may be used with the v1
// object storage package.
func NewObjectStorageV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewComputeV2 creates a ServiceClient that may be used with the v2 compute
// package.
func NewComputeV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewNetworkV2 creates a ServiceClient that may be used with the v2 network
// package.
func NewNetworkV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TODO(stephenfin): Remove this in v3. We no longer support the V1 Block Storage service.
// NewBlockStorageV1 creates a ServiceClient that may be used to access the v1
// block storage service.
func NewBlockStorageV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBlockStorageV2 creates a ServiceClient that may be used to access the v2
// block storage service.
func NewBlockStorageV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewBlockStorageV3 creates a ServiceClient that may be used to access the v3 block storage service.
func NewBlockStorageV3(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewSharedFileSystemV2 creates a ServiceClient that may be used to access the v2 shared file system service.
func NewSharedFileSystemV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewOrchestrationV1 creates a ServiceClient that may be used to access the v1
// orchestration service.
func NewOrchestrationV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDBV1 creates a ServiceClient that may be used to access the v1 DB service.
func NewDBV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewDNSV2 creates a ServiceClient that may be used to access the v2 DNS
// service.
func NewDNSV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewImageV2 creates a ServiceClient that may be used to access the v2 image
// service.
func NewImageV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewLoadBalancerV2 creates a ServiceClient that may be used to access the v2
// load balancer service.
func NewLoadBalancerV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Fixes edge case having an OpenStack lb endpoint with trailing version number.

// NewMetricV1 creates a ServiceClient that may be used with the v1 metric-storage
// service (Aetos).
func NewMetricV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewMessagingV2 creates a ServiceClient that may be used with the v2 messaging
// service.
func NewMessagingV2(ctx context.Context, client *gophercloud.ProviderClient, clientID string, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewContainerV1 creates a ServiceClient that may be used with v1 container package
func NewContainerV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewKeyManagerV1 creates a ServiceClient that may be used with the v1 key
// manager service.
func NewKeyManagerV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewContainerInfraV1 creates a ServiceClient that may be used with the v1 container infra management
// package.
func NewContainerInfraV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewWorkflowV2 creates a ServiceClient that may be used with the v2 workflow management package.
func NewWorkflowV2(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewPlacementV1 creates a ServiceClient that may be used with the placement package.
func NewPlacementV1(ctx context.Context, client *gophercloud.ProviderClient, eo gophercloud.EndpointOpts) (*gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
