package testhelper

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type FakeServer struct {
	// Mux is a multiplexer that can be used to register handlers.
	Mux *http.ServeMux

	// Server is an in-memory HTTP server for testing.
	Server *httptest.Server
}

func (fakeServer FakeServer) Teardown() { _ = "STUB: not implemented"; return }

func (fakeServer FakeServer) Endpoint() string { _ = "STUB: not implemented"; return "" }

// Serves a static content at baseURL/relPath
func (fakeServer FakeServer) ServeFile(t *testing.T, baseURL, relPath, contentType, content string) string {
	_ = "STUB: not implemented"
	return ""
}

// SetupPersistentPortHTTP prepares the Mux and Server listening specific port.
func SetupPersistentPortHTTP(t *testing.T, port int) FakeServer {
	_ = "STUB: not implemented"
	return *new(FakeServer)
}

// SetupHTTP prepares the Mux and Server.
func SetupHTTP() FakeServer { _ = "STUB: not implemented"; return *new(FakeServer) }

// TestFormValues ensures that all the URL parameters given to the http.Request are the same as values.
func TestFormValues(t *testing.T, r *http.Request, values map[string]string) {
	_ = "STUB: not implemented"
	return
}

// TestMethod checks that the Request has the expected method (e.g. GET, POST).
func TestMethod(t *testing.T, r *http.Request, expected string) { _ = "STUB: not implemented"; return }

// TestHeader checks that the header on the http.Request matches the expected value.
func TestHeader(t *testing.T, r *http.Request, header string, expected string) {
	_ = "STUB: not implemented"
	return
}

// TestHeaderUnset checks that the header on the http.Request doesn't exist.
func TestHeaderUnset(t *testing.T, r *http.Request, header string) {
	_ = "STUB: not implemented"
	return
}

// TestBody verifies that the request body matches an expected body.
func TestBody(t *testing.T, r *http.Request, expected string) { _ = "STUB: not implemented"; return }

// TestJSONRequest verifies that the JSON payload of a request matches an expected structure, without asserting things about
// whitespace or ordering.
func TestJSONRequest(t *testing.T, r *http.Request, expected string) {
	_ = "STUB: not implemented"
	return
}
