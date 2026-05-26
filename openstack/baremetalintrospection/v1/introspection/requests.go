package introspection

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListIntrospectionsOptsBuilder allows extensions to add additional parameters to the
// ListIntrospections request.
type ListIntrospectionsOptsBuilder interface {
	ToIntrospectionsListQuery() (string, error)
}

// ListIntrospectionsOpts allows the filtering and sorting of paginated collections through
// the Introspection API. Filtering is achieved by passing in struct field values that map to
// the node attributes you want to see returned. Marker and Limit are used
// for pagination.
type ListIntrospectionsOpts struct {
	// Requests a page size of items.
	Limit int `q:"limit"`

	// The ID of the last-seen item.
	Marker string `q:"marker"`
}

// ToIntrospectionsListQuery formats a ListIntrospectionsOpts into a query string.
func (opts ListIntrospectionsOpts) ToIntrospectionsListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ListIntrospections makes a request against the Inspector API to list the current introspections.
func ListIntrospections(client *gophercloud.ServiceClient, opts ListIntrospectionsOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// GetIntrospectionStatus makes a request against the Inspector API to get the
// status of a single introspection.
func GetIntrospectionStatus(ctx context.Context, client *gophercloud.ServiceClient, nodeID string) (r GetIntrospectionStatusResult) {
	_ = "STUB: not implemented"
	return *new(GetIntrospectionStatusResult)
}

// StartOptsBuilder allows extensions to add additional parameters to the
// Start request.
type StartOptsBuilder interface {
	ToStartIntrospectionQuery() (string, error)
}

// StartOpts represents options to start an introspection.
type StartOpts struct {
	// Whether the current installation of ironic-inspector can manage PXE booting of nodes.
	ManageBoot *bool `q:"manage_boot"`
}

// ToStartIntrospectionQuery converts a StartOpts into a request.
func (opts StartOpts) ToStartIntrospectionQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// StartIntrospection initiate hardware introspection for node NodeID .
// All power management configuration for this node needs to be done prior to calling the endpoint.
func StartIntrospection(ctx context.Context, client *gophercloud.ServiceClient, nodeID string, opts StartOptsBuilder) (r StartResult) {
	_ = "STUB: not implemented"
	return *new(StartResult)
}

// AbortIntrospection abort running introspection.
func AbortIntrospection(ctx context.Context, client *gophercloud.ServiceClient, nodeID string) (r AbortResult) {
	_ = "STUB: not implemented"
	return *new(AbortResult)
}

// GetIntrospectionData return stored data from successful introspection.
func GetIntrospectionData(ctx context.Context, client *gophercloud.ServiceClient, nodeID string) (r DataResult) {
	_ = "STUB: not implemented"
	return *new(DataResult)
}

// ReApplyIntrospection triggers introspection on stored unprocessed data.
// No data is allowed to be sent along with the request.
func ReApplyIntrospection(ctx context.Context, client *gophercloud.ServiceClient, nodeID string) (r ApplyDataResult) {
	_ = "STUB: not implemented"
	return *new(ApplyDataResult)
}
