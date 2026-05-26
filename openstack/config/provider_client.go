package config

import (
	"context"
	"crypto/tls"
	"net/http"

	"github.com/gophercloud/gophercloud/v2"
)

type options struct {
	httpClient http.Client
	tlsConfig  *tls.Config
}

// WithHTTPClient enables passing a custom http.Client to be used in the
// ProviderClient for authentication and for any further call, for example when
// using a ServiceClient derived from this ProviderClient.
func WithHTTPClient(httpClient http.Client) func(*options) { _ = "STUB: not implemented"; return nil }

// WithTLSConfig replaces the Transport of the default HTTP client (or of the
// HTTP client passed with WithHTTPClient) with a RoundTripper containing the
// given TLS config.
func WithTLSConfig(tlsConfig *tls.Config) func(*options) { _ = "STUB: not implemented"; return nil }

// NewProviderClient logs in to an OpenStack cloud found at the identity
// endpoint specified by the options, acquires a token, and returns a Provider
// Client instance that's ready to operate.
//
// If the full path to a versioned identity endpoint was specified  (example:
// http://example.com:5000/v3), that path will be used as the endpoint to
// query.
//
// If a versionless endpoint was specified (example: http://example.com:5000/),
// the endpoint will be queried to determine which versions of the identity
// service are available, then chooses the most recent or most supported
// version.
func NewProviderClient(ctx context.Context, authOptions gophercloud.AuthOptions, opts ...func(*options)) (*gophercloud.ProviderClient, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
