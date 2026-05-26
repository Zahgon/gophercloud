// package clouds provides a parser for OpenStack credentials stored in a clouds.yaml file.
//
// Example use:
//
//	ctx := context.Background()
//	ao, eo, tlsConfig, err := clouds.Parse()
//	if err != nil {
//		panic(err)
//	}
//
//	providerClient, err := config.NewProviderClient(ctx, ao, config.WithTLSConfig(tlsConfig))
//	if err != nil {
//		panic(err)
//	}
//
//	networkClient, err := openstack.NewNetworkV2(ctx, providerClient, eo)
//	if err != nil {
//		panic(err)
//	}
package clouds

import (
	"crypto/tls"

	"github.com/gophercloud/gophercloud/v2"
)

// Parse fetches a clouds.yaml file from disk and returns the parsed
// credentials.
//
// By default this function mimics the behaviour of python-openstackclient, which is:
//
//   - if the environment variable `OS_CLIENT_CONFIG_FILE` is set and points to a
//     clouds.yaml, use that location as the only search location for `clouds.yaml` and `secure.yaml`;
//   - otherwise, the search locations for `clouds.yaml` and `secure.yaml` are:
//     1. the current working directory (on Linux: `./`)
//     2. the directory `openstack` under the standard user config location for
//     the operating system (on Linux: `${XDG_CONFIG_HOME:-$HOME/.config}/openstack/`)
//     3. on Linux, `/etc/openstack/`
//
// Once `clouds.yaml` is found in a search location, the same location is used to search for `secure.yaml`.
//
// Like in python-openstackclient, relative paths in the `clouds.yaml` section
// `cacert` are interpreted as relative the the current directory, and not to
// the `clouds.yaml` location.
//
// Search locations, as well as individual `clouds.yaml` properties, can be
// overwritten with functional options.
func Parse(opts ...ParseOption) (gophercloud.AuthOptions, gophercloud.EndpointOpts, *tls.Config, error) {
	_ = "STUB: not implemented"
	return *new(gophercloud.AuthOptions), *new(gophercloud.EndpointOpts), nil, nil
}

// Set the defaults and open the files for reading. This code only runs
// if no override has been set, because it is fallible.

// Use XDG_CONFIG_HOME or fall back to ~/.config, matching the
// OpenStack convention for clouds.yaml location on all platforms.

// Parse the YAML payloads.

// If secureCloud has content and it differs from the cloud entry,
// merge the two together.

// computeAvailability is a helper method to determine the endpoint type
// requested by the user.
func computeAvailability(endpointType string) gophercloud.Availability {
	_ = "STUB: not implemented"
	return *new(gophercloud.Availability)
}

// coalesce returns the first argument that is not the zero value for its type,
// or the zero value for its type.
func coalesce[T comparable](items ...T) T { _ = "STUB: not implemented"; return *new(T) }

// mergeClouds merges two Clouds recursively (the AuthInfo also gets merged).
// In case both Clouds define a value, the value in the 'override' cloud takes precedence
func mergeClouds(override, cloud Cloud) (Cloud, error) {
	_ = "STUB: not implemented"
	return *new(Cloud), nil
}

// merges two interfaces. In cases where a value is defined for both 'overridingInterface' and
// 'inferiorInterface' the value in 'overridingInterface' will take precedence.
func mergeInterfaces(overridingInterface, inferiorInterface any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

// mergeClouds(nil, map[string]interface{...}) -> map[string]interface{...}

// We don't want to override with empty values
