package gophercloud

import (
	"context"
)

// NormalizePathURL is used to convert rawPath to a fqdn, using basePath as
// a reference in the filesystem, if necessary. basePath is assumed to contain
// either '.' when first used, or the file:// type fqdn of the parent resource.
// e.g. myFavScript.yaml => file://opt/lib/myFavScript.yaml
func NormalizePathURL(basePath, rawPath string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// if a scheme is defined, it must be a fqdn already

// if basePath is a url, then child resources are assumed to be relative to it

// NormalizeURL is an internal function to be used by provider clients.
//
// It ensures that each endpoint URL has a closing `/`, as expected by
// ServiceClient's methods.
func NormalizeURL(url string) string { _ = "STUB: not implemented"; return "" }

// RemainingKeys will inspect a struct and compare it to a map. Any struct
// field that does not have a JSON tag that matches a key in the map or
// a matching lower-case field in the map will be returned as an extra.
//
// This is useful for determining the extra fields returned in response bodies
// for resources that can contain an arbitrary or dynamic number of fields.
func RemainingKeys(s any, m map[string]any) (extras map[string]any) {
	_ = "STUB: not implemented"
	return nil
}

// WaitFor polls a predicate function, once per second, up to a context cancellation.
// This is useful to wait for a resource to transition to a certain state.
// Resource packages will wrap this in a more convenient function that's
// specific to a certain resource, but it can also be useful on its own.
func WaitFor(ctx context.Context, predicate func(context.Context) (bool, error)) error {
	_ = "STUB: not implemented"
	return nil
}
