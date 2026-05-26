package gophercloud

import (
	"context"
	"io"
	"net/http"
)

// ServiceClient stores details required to interact with a specific service API implemented by a provider.
// Generally, you'll acquire these by calling the appropriate `New` method on a ProviderClient.
type ServiceClient struct {
	// ProviderClient is a reference to the provider that implements this service.
	*ProviderClient

	// Endpoint is the base URL of the service's API, acquired from a service catalog.
	// It MUST end with a /.
	Endpoint string

	// ResourceBase is the base URL shared by the resources within a service's API. It should include
	// the API version and, like Endpoint, MUST end with a / if set. If not set, the Endpoint is used
	// as-is, instead.
	ResourceBase string

	// This is the service client type (e.g. compute, sharev2).
	// NOTE: FOR INTERNAL USE ONLY. DO NOT SET. GOPHERCLOUD WILL SET THIS.
	// It is only exported because it gets set in a different package.
	Type string

	// The microversion of the service to use. Set this to use a particular microversion.
	Microversion string

	// MoreHeaders allows users (or Gophercloud) to set service-wide headers on requests. Put another way,
	// values set in this field will be set on all the HTTP requests the service client sends.
	MoreHeaders map[string]string
}

// ResourceBaseURL returns the base URL of any resources used by this service. It MUST end with a /.
func (client *ServiceClient) ResourceBaseURL() string { _ = "STUB: not implemented"; return "" }

// ServiceURL constructs a URL for a resource belonging to this provider.
func (client *ServiceClient) ServiceURL(parts ...string) string {
	_ = "STUB: not implemented"
	return ""
}

func (client *ServiceClient) initReqOpts(JSONBody any, JSONResponse any, opts *RequestOpts) {
	_ = "STUB: not implemented"
	return
}

// Get calls `Request` with the "GET" HTTP verb.
func (client *ServiceClient) Get(ctx context.Context, url string, JSONResponse any, opts *RequestOpts) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Post calls `Request` with the "POST" HTTP verb.
func (client *ServiceClient) Post(ctx context.Context, url string, JSONBody any, JSONResponse any, opts *RequestOpts) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Put calls `Request` with the "PUT" HTTP verb.
func (client *ServiceClient) Put(ctx context.Context, url string, JSONBody any, JSONResponse any, opts *RequestOpts) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Patch calls `Request` with the "PATCH" HTTP verb.
func (client *ServiceClient) Patch(ctx context.Context, url string, JSONBody any, JSONResponse any, opts *RequestOpts) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Delete calls `Request` with the "DELETE" HTTP verb.
func (client *ServiceClient) Delete(ctx context.Context, url string, opts *RequestOpts) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Head calls `Request` with the "HEAD" HTTP verb.
func (client *ServiceClient) Head(ctx context.Context, url string, opts *RequestOpts) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (client *ServiceClient) setMicroversionHeader(opts *RequestOpts) {
	_ = "STUB: not implemented"
	return
}

// cinder should accept block-storage but (as of Dalmatian) does not

// magnum should accept container-infrastructure-management but (as of Epoxy) does not

// Request carries out the HTTP operation for the service client
func (client *ServiceClient) Request(ctx context.Context, method, url string, options *RequestOpts) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseResponse is a helper function to parse http.Response to constituents.
func ParseResponse(resp *http.Response, err error) (io.ReadCloser, http.Header, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), *new(http.Header), nil
}
