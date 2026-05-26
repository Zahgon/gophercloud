package stacks

// Template is a structure that represents OpenStack Heat templates
type Template struct {
	TE
}

// TemplateFormatVersions is a map containing allowed variations of the template format version
// Note that this contains the permitted variations of the _keys_ not the values.
var TemplateFormatVersions = map[string]bool{
	"HeatTemplateFormatVersion": true,
	"heat_template_version":     true,
	"AWSTemplateFormatVersion":  true,
}

// Validate validates the contents of the Template
func (t *Template) Validate() error { _ = "STUB: not implemented"; return nil }

func (t *Template) makeChildTemplate(childURL string, ignoreIf igFunc, recurse bool) (*Template, error) {
	_ = "STUB: not implemented"
	// create a new child template
	return nil, nil
}

// initialize child template

// get the base location of the child template. Child path is relative
// to its parent location so that templates can be composed

// Preserve all elements of the URL but take the directory part of the path

// fetch the contents of the child template or file

// process child template recursively if required. This is
// required if the child template itself contains references to
// other templates

// Applies the transformation for getFileContents() to just one element of a map.
// In case the element requires transforming, the function returns its new value.
func (t *Template) mapElemFileContents(key string, v any, ignoreIf igFunc, recurse bool) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// if the value is not a string, recursively parse that value

// at this point, the k, v pair has a reference to an external template
// or file (for 'get_file' function).
// The assumption of heatclient is that value v is a reference
// to a file in the users environment, so we have to the path

// create a new child template with the referenced contents

// update parent template with current child templates' content.
// At this point, the child template has been parsed recursively.

// Also add child templates' own children (templates or get_file)!

// GetFileContents recursively parses a template to search for urls. These urls
// are assumed to point to other templates (known in OpenStack Heat as child
// templates). The contents of these urls are fetched and stored in the `Files`
// parameter of the template structure. This is the only way that a user can
// use child templates that are located in their filesystem; urls located on the
// web (e.g. on github or swift) can be fetched directly by Heat engine.
func (t *Template) getFileContents(te any, ignoreIf igFunc, recurse bool) error {
	_ = "STUB: not implemented"
	// initialize template if empty
	return nil
}

// if te is a map[string], go check all elements for URLs to replace

// if te is a slice, call the function on each element of the slice.

// if te is anything else, there is nothing to do.

// In case some element was updated, we have to regenerate the string representation

// function to choose keys whose values are other template files
func ignoreIfTemplate(key string, value any) bool {
	_ = "STUB: not implemented"
	// key must be either `get_file` or `type` for value to be a URL
	return false
}

// value must be a string

// `.template` and `.yaml` are allowed suffixes for template URLs when referred to by `type`
