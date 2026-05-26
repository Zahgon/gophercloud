package stackresources

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Find retrieves stack resources for the given stack name.
func Find(ctx context.Context, c *gophercloud.ServiceClient, stackName string) (r FindResult) {
	_ = "STUB: not implemented"
	return *new(FindResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToStackResourceListQuery() (string, error)
}

// ListOpts allows the filtering and sorting of paginated collections through
// the API. Marker and Limit are used for pagination.
type ListOpts struct {
	// Include resources from nest stacks up to Depth levels of recursion.
	Depth int `q:"nested_depth"`
}

// ToStackResourceListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToStackResourceListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List makes a request against the API to list resources for the given stack.
func List(client *gophercloud.ServiceClient, stackName, stackID string, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retreives data for the given stack resource.
func Get(ctx context.Context, c *gophercloud.ServiceClient, stackName, stackID, resourceName string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Metadata retreives the metadata for the given stack resource.
func Metadata(ctx context.Context, c *gophercloud.ServiceClient, stackName, stackID, resourceName string) (r MetadataResult) {
	_ = "STUB: not implemented"
	return *new(MetadataResult)
}

// ListTypes makes a request against the API to list resource types.
func ListTypes(client *gophercloud.ServiceClient) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Schema retreives the schema for the given resource type.
func Schema(ctx context.Context, c *gophercloud.ServiceClient, resourceType string) (r SchemaResult) {
	_ = "STUB: not implemented"
	return *new(SchemaResult)
}

// Template retreives the template representation for the given resource type.
func Template(ctx context.Context, c *gophercloud.ServiceClient, resourceType string) (r TemplateResult) {
	_ = "STUB: not implemented"
	return *new(TemplateResult)
}

// MarkUnhealthyOpts contains the common options struct used in this package's
// MarkUnhealthy operations.
type MarkUnhealthyOpts struct {
	// A boolean indicating whether the target resource should be marked as unhealthy.
	MarkUnhealthy bool `json:"mark_unhealthy"`
	// The reason for the current stack resource state.
	ResourceStatusReason string `json:"resource_status_reason,omitempty"`
}

// MarkUnhealthyOptsBuilder is the interface options structs have to satisfy in order
// to be used in the MarkUnhealthy operation in this package
type MarkUnhealthyOptsBuilder interface {
	ToMarkUnhealthyMap() (map[string]any, error)
}

// ToMarkUnhealthyMap validates that a template was supplied and calls
// the ToMarkUnhealthyMap private function.
func (opts MarkUnhealthyOpts) ToMarkUnhealthyMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MarkUnhealthy marks the specified resource in the stack as unhealthy.
func MarkUnhealthy(ctx context.Context, c *gophercloud.ServiceClient, stackName, stackID, resourceName string, opts MarkUnhealthyOptsBuilder) (r MarkUnhealthyResult) {
	_ = "STUB: not implemented"
	return *new(MarkUnhealthyResult)
}
