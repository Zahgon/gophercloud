package resourcetypes

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// SupportStatus is a type for specifying by which support status to filter the
// list of resource types.
type SupportStatus string

const (
	// SupportStatusUnknown is returned when the resource type does not have a
	// support status.
	SupportStatusUnknown SupportStatus = "UNKNOWN"
	// SupportStatusSupported indicates a resource type that is expected to
	// work.
	SupportStatusSupported SupportStatus = "SUPPORTED"
	// SupportStatusDeprecated indicates a resource type that is in the process
	// being removed, and may or may not be replaced by something else.
	SupportStatusDeprecated SupportStatus = "DEPRECATED"
	// SupportStatusHidden indicates a resource type that has been removed.
	// Existing stacks that contain resources of this type can still be
	// deleted or updated to remove the resources, but they may not actually
	// do anything any more.
	SupportStatusHidden SupportStatus = "HIDDEN"
	// SupportStatusUnsupported indicates a resource type that is provided for
	// preview or other purposes and should not be relied upon.
	SupportStatusUnsupported SupportStatus = "UNSUPPORTED"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToResourceTypeListQuery() (string, error)
}

// ListOpts allows the filtering of collections through the API.
type ListOpts struct {
	// Filters the resource type list by a regex on the name.
	NameRegex string `q:"name"`
	// Filters the resource list by the specified SupportStatus.
	SupportStatus SupportStatus `q:"support_status"`
	// Return descriptions as well as names of resource types
	WithDescription bool `q:"with_description"`
}

// ToResourceTypeListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToResourceTypeListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request against the API to list available resource types.
func List(ctx context.Context, client *gophercloud.ServiceClient, opts ListOptsBuilder) (r ListResult) {
	_ = "STUB: not implemented"
	return *new(ListResult)
}

// GetSchema retreives the schema for a given resource type.
func GetSchema(ctx context.Context, client *gophercloud.ServiceClient, resourceType string) (r GetSchemaResult) {
	_ = "STUB: not implemented"
	return *new(GetSchemaResult)
}

// GenerateTemplateOptsBuilder allows extensions to add additional parameters
// to the GenerateTemplate request.
type GenerateTemplateOptsBuilder interface {
	ToGenerateTemplateQuery() (string, error)
}

type GeneratedTemplateType string

const (
	TemplateTypeHOT GeneratedTemplateType = "hot"
	TemplateTypeCFn GeneratedTemplateType = "cfn"
)

// GenerateTemplateOpts allows the filtering of collections through the API.
type GenerateTemplateOpts struct {
	TemplateType GeneratedTemplateType `q:"template_type"`
}

// ToGenerateTemplateQuery formats a GenerateTemplateOpts into a query string.
func (opts GenerateTemplateOpts) ToGenerateTemplateQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GenerateTemplate retreives an example template for a given resource type.
func GenerateTemplate(ctx context.Context, client *gophercloud.ServiceClient, resourceType string, opts GenerateTemplateOptsBuilder) (r TemplateResult) {
	_ = "STUB: not implemented"
	return *new(TemplateResult)
}
