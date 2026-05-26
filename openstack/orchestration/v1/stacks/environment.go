package stacks

// Environment is a structure that represents stack environments
type Environment struct {
	TE
}

// EnvironmentSections is a map containing allowed sections in a stack environment file
var EnvironmentSections = map[string]bool{
	"parameters":         true,
	"parameter_defaults": true,
	"resource_registry":  true,
}

// Validate validates the contents of the Environment
func (e *Environment) Validate() error { _ = "STUB: not implemented"; return nil }

// Parse environment file to resolve the URL's of the resources. This is done by
// reading from the `Resource Registry` section, which is why the function is
// named GetRRFileContents.
func (e *Environment) getRRFileContents(ignoreIf igFunc) error {
	_ = "STUB: not implemented"
	// initialize environment if empty
	return nil
}

// get the resource registry, only process further if it is a map

// search the resource registry for URLs
//
// the resource registry might contain a base URL for the resource. If
// such a field is present, use it. Otherwise, use the default base URL.

// The contents of the resource may be located in a remote file, which
// will be a template. Instantiate a temporary template to manage the
// contents.

// Fetch the contents of remote resource URL's

// check the `resources` section (if it exists) for more URL's. Note that
// the previous call to GetFileContents was (deliberately) not recursive
// as we want more control over where to look for URL's

// if base_url for the resource type is defined, use it

// if the resource registry contained any URL's, store them. This can
// then be passed as parameter to api calls to Heat api.

// In case some element was updated, regenerate the string representation

// function to choose keys whose values are other environment files
func ignoreIfEnvironment(key string, value any) bool {
	_ = "STUB: not implemented"
	// base_url and hooks refer to components which cannot have urls
	return false
}

// if value is not string, it cannot be a URL

// if value contains `::`, it must be a reference to another resource type
// e.g. OS::Nova::Server : Rackspace::Cloud::Server
