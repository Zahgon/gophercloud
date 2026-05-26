package stacks

import (
	"net/http"
)

// Client is an interface that expects a Get method similar to http.Get. This
// is needed for unit testing, since we can mock an http client. Thus, the
// client will usually be an http.Client EXCEPT in unit tests.
type Client interface {
	Get(string) (*http.Response, error)
}

// TE is a base structure for both Template and Environment
type TE struct {
	// Bin stores the contents of the template or environment.
	Bin []byte
	// URL stores the URL of the template. This is allowed to be a 'file://'
	// for local files.
	URL string
	// Parsed contains a parsed version of Bin. Since there are 2 different
	// fields referring to the same value, you must be careful when accessing
	// this filed.
	Parsed map[string]any
	// Files contains a mapping between the urls in templates to their contents.
	Files map[string]string
	// fileMaps is a map used internally when determining Files.
	fileMaps map[string]string
	// baseURL represents the location of the template or environment file.
	baseURL string
	// client is an interface which allows TE to fetch contents from URLS
	client Client
}

// Fetch fetches the contents of a TE from its URL. Once a TE structure has a
// URL, call the fetch method to fetch the contents.
func (t *TE) Fetch() error {
	_ = "STUB: not implemented"
	// if the baseURL is not provided, use the current directors as the base URL
	return nil
}

// if the contents are already present, do nothing.

// get a fqdn from the URL using the baseURL of the TE. For local files,
// the URL's will have the `file` scheme.

// get an HTTP client if none present

// use the client to fetch the contents of the TE

// get the basepath of the TE
func getBasePath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// get a an HTTP client to retrieve URL's. This client allows the use of `file`
// scheme since we may need to fetch files from users filesystem
func getHTTPClient() Client { _ = "STUB: not implemented"; return *new(Client) }

// Parse will parse the contents and then validate. The contents MUST be either JSON or YAML.
func (t *TE) Parse() error { _ = "STUB: not implemented"; return nil }

// either parse as JSON...

// ... or as YAML (but in this case, take extra care because yaml.Unmarshal()
// might incorrectly decode the `heat_template_version` attribute as a
// time.Time instead of as a string because it looks like "YYYY-MM-DD")

// igfunc is a parameter used by GetFileContents and GetRRFileContents to check
// for valid URL's.
type igFunc func(string, any) bool
