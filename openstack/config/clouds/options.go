package clouds

import (
	"io"

	"github.com/gophercloud/gophercloud/v2"
)

type cloudOpts struct {
	cloudName        string
	locations        []string
	cloudsyamlReader io.Reader
	secureyamlReader io.Reader

	applicationCredentialID     string
	applicationCredentialName   string
	applicationCredentialSecret string
	authURL                     string
	domainID                    string
	domainName                  string
	endpointType                string
	password                    string
	projectID                   string
	projectName                 string
	region                      string
	scope                       *gophercloud.AuthScope
	token                       string
	userID                      string
	username                    string

	caCertPath     string
	clientCertPath string
	clientKeyPath  string
	insecure       *bool
}

// ParseOption one of parse configuration returned by With* modifier
type ParseOption = func(*cloudOpts)

// WithCloudName allows to override the environment variable `OS_CLOUD`.
func WithCloudName(osCloud string) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }

// WithLocations is a functional option that sets the search locations for the
// clouds.yaml file (and its optional companion secure.yaml). Each location is
// a file path pointing to a possible `clouds.yaml`.
func WithLocations(locations ...string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

// WithCloudsYAML is a functional option that lets you pass a clouds.yaml file
// as an io.Reader interface. When this option is passed, FromCloudsYaml will
// not attempt to fetch any file from the file system. To add a secure.yaml,
// use in conjunction with WithSecureYAML.
func WithCloudsYAML(clouds io.Reader) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

// WithSecureYAML is a functional option that lets you pass a secure.yaml file
// as an io.Reader interface, to complement the clouds.yaml that is either
// fetched from the filesystem, or passed with WithCloudsYAML.
func WithSecureYAML(secure io.Reader) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithApplicationCredentialID(applicationCredentialID string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithApplicationCredentialName(applicationCredentialName string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithApplicationCredentialSecret(applicationCredentialSecret string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithIdentityEndpoint(authURL string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithDomainID(domainID string) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }

func WithDomainName(domainName string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

// WithRegion allows to override the endpoint type set in clouds.yaml or in the
// environment variable `OS_INTERFACE`.
func WithEndpointType(endpointType string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithPassword(password string) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }

func WithProjectID(projectID string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithProjectName(projectName string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

// WithRegion allows to override the region set in clouds.yaml or in the
// environment variable `OS_REGION_NAME`
func WithRegion(region string) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }

func WithScope(scope *gophercloud.AuthScope) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithToken(token string) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }

func WithUserID(userID string) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }

func WithUsername(username string) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }

func WithCACertPath(caCertPath string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithClientCertPath(clientCertPath string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithClientKeyPath(clientKeyPath string) ParseOption {
	_ = "STUB: not implemented"
	return *new(ParseOption)
}

func WithInsecure(insecure bool) ParseOption { _ = "STUB: not implemented"; return *new(ParseOption) }
