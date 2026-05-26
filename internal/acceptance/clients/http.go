package clients

import (
	"io"
	"net/http"
)

// List of headers that need to be redacted
var REDACT_HEADERS = []string{"x-auth-token", "x-auth-key", "x-service-token",
	"x-storage-token", "x-account-meta-temp-url-key", "x-account-meta-temp-url-key-2",
	"x-container-meta-temp-url-key", "x-container-meta-temp-url-key-2", "set-cookie",
	"x-subject-token"}

// LogRoundTripper satisfies the http.RoundTripper interface and is used to
// customize the default http client RoundTripper to allow logging.
type LogRoundTripper struct {
	Rt http.RoundTripper
}

// RoundTrip performs a round-trip HTTP request and logs relevant information
// about it.
func (lrt *LogRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// logRequest will log the HTTP Request details.
// If the body is JSON, it will attempt to be pretty-formatted.
func (lrt *LogRoundTripper) logRequest(original io.ReadCloser, contentType string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// Handle request contentType

// logResponse will log the HTTP Response details.
// If the body is JSON, it will attempt to be pretty-formatted.
func (lrt *LogRoundTripper) logResponse(original io.ReadCloser, contentType string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

// formatJSON will try to pretty-format a JSON body.
// It will also mask known fields which contain sensitive information.
func (lrt *LogRoundTripper) formatJSON(raw []byte) string { _ = "STUB: not implemented"; return "" }

// Mask known password fields

// Ignore the catalog

// redactHeaders processes a headers object, returning a redacted list
func redactHeaders(headers http.Header) (processedHeaders []string) {
	_ = "STUB: not implemented"
	return nil
}

// formatHeaders processes a headers object plus a deliminator, returning a string
func formatHeaders(headers http.Header) string { _ = "STUB: not implemented"; return "" }
