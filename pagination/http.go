package pagination

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gophercloud/gophercloud/v2"
)

// PageResult stores the HTTP response that returned the current page of results.
type PageResult struct {
	gophercloud.Result
	url.URL
}

// PageResultFrom parses an HTTP response as JSON and returns a PageResult containing the
// results, interpreting it as JSON if the content type indicates.
func PageResultFrom(resp *http.Response) (PageResult, error) {
	_ = "STUB: not implemented"
	return *new(PageResult), nil
}

// PageResultFromParsed constructs a PageResult from an HTTP response that has already had its
// body parsed as JSON (and closed).
func PageResultFromParsed(resp *http.Response, body any) PageResult {
	_ = "STUB: not implemented"
	return *new(PageResult)
}

// Request performs an HTTP request and extracts the http.Response from the result.
func Request(ctx context.Context, client *gophercloud.ServiceClient, headers map[string]string, url string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
