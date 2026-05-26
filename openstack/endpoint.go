package openstack

import (
	"context"
	"regexp"

	"github.com/gophercloud/gophercloud/v2"
	tokens2 "github.com/gophercloud/gophercloud/v2/openstack/identity/v2/tokens"
	tokens3 "github.com/gophercloud/gophercloud/v2/openstack/identity/v3/tokens"
)

var versionedServiceTypeAliasRegexp = regexp.MustCompile(`^.*v(\d)$`)

func extractServiceTypeVersion(serviceType string) int { _ = "STUB: not implemented"; return 0 }

// no point converting to an int

func endpointSupportsVersion(ctx context.Context, client *gophercloud.ProviderClient, serviceType, endpointURL string, expectedVersion int) (bool, error) {
	_ = "STUB: not implemented"
	// Swift doesn't support version discovery :(
	return false, nil
}

// Repeating verbatim from keystoneauth1 [1]:
//
// > The sins of our fathers become the blood on our hands.
// > If a user requests an old-style service type such as volumev2, then they
// > are inherently requesting the major API version 2. It's not a good
// > interface, but it's the one that was imposed on the world years ago
// > because the client libraries all hid the version discovery document.
// > In order to be able to ensure that a user who requests volumev2 does not
// > get a block-storage endpoint that only provides v3 of the block-storage
// > service, we need to pull the version out of the service_type. The
// > service-types-authority will prevent the growth of new monstrosities such
// > as this, but in order to move forward without breaking people, we have
// > to just cry in the corner while striking ourselves with thorned branches.
// > That said, for sure only do this hack for officially known service_types.
//
// So yeah, what mordred said.
//
// https://github.com/openstack/keystoneauth/blob/5.10.0/keystoneauth1/discover.py#L270-L290

// NOTE(stephenfin) In addition to the above, keystoneauth also supports a URL
// hack whereby it will extract the version from the URL. We may wish to
// implement this too.

/*
V2Endpoint discovers the endpoint URL for a specific service from a
ServiceCatalog acquired during the v2 identity service.

The specified EndpointOpts are used to identify a unique, unambiguous endpoint
to return. It's an error both when multiple endpoints match the provided
criteria and when none do. The minimum that can be specified is a Type, but you
will also often need to specify a Name and/or a Region depending on what's
available on your OpenStack deployment.
*/
func V2Endpoint(ctx context.Context, client *gophercloud.ProviderClient, catalog *tokens2.ServiceCatalog, opts gophercloud.EndpointOpts) (string, error) {
	_ = "STUB: not implemented"
	// Extract Endpoints from the catalog entries that match the requested Type, Name if provided, and Region if provided.
	//
	// If multiple endpoints are found, we return the first result and disregard the rest.
	// This behavior matches the Python library. See GH-1764.
	return "", nil
}

// Report an error if there were no matching endpoints.

/*
V3Endpoint discovers the endpoint URL for a specific service from a Catalog
acquired during the v3 identity service.

The specified EndpointOpts are used to identify a unique, unambiguous endpoint
to return. It's an error both when multiple endpoints match the provided
criteria and when none do. The minimum that can be specified is a Type, but you
will also often need to specify a Name and/or a Region depending on what's
available on your OpenStack deployment.
*/
func V3Endpoint(ctx context.Context, client *gophercloud.ProviderClient, catalog *tokens3.ServiceCatalog, opts gophercloud.EndpointOpts) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Extract Endpoints from the catalog entries that match the requested Type, Interface,
// Name if provided, and Region if provided.
//
// If multiple endpoints are found, we return the first result and disregard the rest.
// This behavior matches the Python library. See GH-1764.

// Report an error if there were no matching endpoints.
