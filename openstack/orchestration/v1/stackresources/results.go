package stackresources

import (
	"time"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Resource represents a stack resource.
type Resource struct {
	Attributes     map[string]any     `json:"attributes"`
	CreationTime   time.Time          `json:"-"`
	Description    string             `json:"description"`
	Links          []gophercloud.Link `json:"links"`
	LogicalID      string             `json:"logical_resource_id"`
	Name           string             `json:"resource_name"`
	ParentResource string             `json:"parent_resource"`
	PhysicalID     string             `json:"physical_resource_id"`
	RequiredBy     []any              `json:"required_by"`
	Status         string             `json:"resource_status"`
	StatusReason   string             `json:"resource_status_reason"`
	Type           string             `json:"resource_type"`
	UpdatedTime    time.Time          `json:"-"`
}

func (r *Resource) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// FindResult represents the result of a Find operation.
type FindResult struct {
	gophercloud.Result
}

// Extract returns a slice of Resource objects and is called after a
// Find operation.
func (r FindResult) Extract() ([]Resource, error) { _ = "STUB: not implemented"; return nil, nil }

// ResourcePage abstracts the raw results of making a List() request against the API.
// As OpenStack extensions may freely alter the response bodies of structures returned to the client, you may only safely access the
// data provided through the ExtractResources call.
type ResourcePage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a page contains no Server results.
func (r ResourcePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ExtractResources interprets the results of a single page from a List() call, producing a slice of Resource entities.
func ExtractResources(r pagination.Page) ([]Resource, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetResult represents the result of a Get operation.
type GetResult struct {
	gophercloud.Result
}

// Extract returns a pointer to a Resource object and is called after a
// Get operation.
func (r GetResult) Extract() (*Resource, error) { _ = "STUB: not implemented"; return nil, nil }

// MetadataResult represents the result of a Metadata operation.
type MetadataResult struct {
	gophercloud.Result
}

// Extract returns a map object and is called after a
// Metadata operation.
func (r MetadataResult) Extract() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ResourceTypePage abstracts the raw results of making a ListTypes() request against the API.
// As OpenStack extensions may freely alter the response bodies of structures returned to the client, you may only safely access the
// data provided through the ExtractResourceTypes call.
type ResourceTypePage struct {
	pagination.SinglePageBase
}

// IsEmpty returns true if a ResourceTypePage contains no resource types.
func (r ResourceTypePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// ResourceTypes represents the type that holds the result of ExtractResourceTypes.
// We define methods on this type to sort it before output
type ResourceTypes []string

func (r ResourceTypes) Len() int { _ = "STUB: not implemented"; return 0 }

func (r ResourceTypes) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (r ResourceTypes) Less(i, j int) bool {
	_ = "STUB: not implemented"

	// ExtractResourceTypes extracts and returns resource types.
	return false
}

func ExtractResourceTypes(r pagination.Page) (ResourceTypes, error) {
	_ = "STUB: not implemented"
	return *new(ResourceTypes), nil
}

// TypeSchema represents a stack resource schema.
type TypeSchema struct {
	Attributes    map[string]any `json:"attributes"`
	Properties    map[string]any `json:"properties"`
	ResourceType  string         `json:"resource_type"`
	SupportStatus map[string]any `json:"support_status"`
}

// SchemaResult represents the result of a Schema operation.
type SchemaResult struct {
	gophercloud.Result
}

// Extract returns a pointer to a TypeSchema object and is called after a
// Schema operation.
func (r SchemaResult) Extract() (*TypeSchema, error) { _ = "STUB: not implemented"; return nil, nil }

// TemplateResult represents the result of a Template operation.
type TemplateResult struct {
	gophercloud.Result
}

// Extract returns the template and is called after a
// Template operation.
func (r TemplateResult) Extract() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// MarkUnhealthyResult represents the result of a mark unhealthy operation.
type MarkUnhealthyResult struct {
	gophercloud.ErrResult
}
