package utils

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

type Status string

const (
	StatusCurrent      Status = "CURRENT"
	StatusSupported    Status = "SUPPORTED"
	StatusDeprecated   Status = "DEPRECATED"
	StatusExperimental Status = "EXPERIMENTAL"
	StatusUnknown      Status = ""
)

// SupportedVersion stores a normalized form of the API version data. It handles APIs that
// support microversions as well as those that do not.
type SupportedVersion struct {
	// Major is the major version number of the API
	Major int
	// Minor is the minor version number of the API
	Minor int
	// Status is the status of the API
	Status Status
	SupportedMicroversions
}

// SupportedMicroversions stores a normalized form of the maximum and minimum API microversions
// supported by a given service.
type SupportedMicroversions struct {
	// MaxMajor is the major version number of the maximum supported API microversion
	MaxMajor int
	// MaxMinor is the minor version number of the maximum supported API microversion
	MaxMinor int
	// MinMajor is the major version number of the minimum supported API microversion
	MinMajor int
	// MinMinor is the minor version number of the minimum supported API microversion
	MinMinor int
}

type version struct {
	ID         string `json:"id"`
	Status     string `json:"status"`
	Version    string `json:"version,omitempty"`
	MaxVersion string `json:"max_version,omitempty"`
	MinVersion string `json:"min_version"`
}

type response struct {
	Versions []version `json:"-"`
}

func (r *response) UnmarshalJSON(in []byte) error {
	_ = "STUB: not implemented"
	// intermediateResponse is an intermediate struct that allows us to offload the difference
	// between a single version document and a multi-version document to the json parser and
	// only focus on differences in the latter
	return nil
}

// case 1: we have a single enveloped version object
//
// this is the approach used by Manila for single version responses

// case 2: we have an singly enveloped array of version objects
//
// this is the approach used by nova, cinder and glance, among others, for multi-version
// responses

// case 3: we have an doubly enveloped array of version objects
//
// this is the approach used by keystone and barbican, among others, for multi-version
// responses

// case 4: we have a single unenveloped version object
//
// this is the approach used by most other services for single version responses

func extractVersion(endpointURL string) (int, int, error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// first, check the nth path element for a version string

// if there are no more parts, quit

// we don't return the error message directly since it might be misleading: at this point
// we might have a *malformed* version identifier rather than *no* version identifier

// the guidelines say we should use the currently scoped project_id from the token, but we
// don't necessarily have a token yet so we speculatively look at the (n-1)th path element
// (but only that) just as keystoneauth does
//
// https://github.com/openstack/keystoneauth/blob/master/keystoneauth1/discover.py#L1534-L1545

// once again, we don't return the error message directly

// GetServiceVersions returns the versions supported by the ServiceClient Endpoint.
// If the endpoint resolves to an unversioned discovery API, this should return one or more supported versions.
// If the endpoint resolves to a versioned discovery API, this should return exactly one supported version.
func GetServiceVersions(ctx context.Context, client *gophercloud.ProviderClient, endpointURL string, discoverVersions bool) ([]SupportedVersion, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// we weren't able to find a discovery document but we have version information from the URL

// Only normalize the microversions if there are microversions to normalize

// GetSupportedMicroversions returns the minimum and maximum microversion that is supported by the ServiceClient Endpoint.
func GetSupportedMicroversions(ctx context.Context, client *gophercloud.ServiceClient) (SupportedMicroversions, error) {
	_ = "STUB: not implemented"
	return *new(SupportedMicroversions), nil
}

// If there are multiple versions then we were handed an unversioned endpoint. These don't
// provide microversion information, so we need to fail. Likewise, if there are no versions
// then something has gone wrong and we also need to fail.

// RequireMicroversion checks that the required microversion is supported and
// returns a ServiceClient with the microversion set.
func RequireMicroversion(ctx context.Context, client gophercloud.ServiceClient, required string) (gophercloud.ServiceClient, error) {
	_ = "STUB: not implemented"
	return *new(gophercloud.ServiceClient), nil
}

// IsSupported checks if a microversion falls in the supported interval.
// It returns true if the version is within the interval and false otherwise.
func (supported SupportedMicroversions) IsSupported(version string) (bool, error) {
	_ = "STUB: not implemented"
	// Parse the version X.Y into X and Y integers that are easier to compare.
	return false, nil
}

// Check that the major version number is supported.

// Check that the minor version number is supported

// ParseVersion parsed the version strings v{MAJOR} and v{MAJOR}.{MINOR} into separate integers
// major and minor.
// For example, "v2.1" becomes 2 and 1, "v3" becomes 3 and 0, and "1" becomes 1 and 0.
func ParseVersion(version string) (major, minor int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// We use the regex indicated by the version discovery guidelines.
//
// https://specs.openstack.org/openstack/api-sig/guidelines/consuming-catalog/version-discovery.html#inferring-version
//
// However, we diverge slightly since not all services include the 'v' prefix (glares at zaqar)

// ParseMicroversion parses the version major.minor into separate integers major and minor.
// For example, "2.53" becomes 2 and 53.
func ParseMicroversion(version string) (major int, minor int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

func ParseStatus(status string) (Status, error) {
	_ = "STUB: not implemented"
	return *new(Status), nil
}

// keystone uses STABLE instead of CURRENT
