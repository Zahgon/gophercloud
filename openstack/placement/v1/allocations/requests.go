package allocations

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// Get retrieves the allocations for a specific consumer by its UUID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, consumerUUID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// ProviderAllocationsOpts specifies the resources to consume from a single resource
// provider in a write request.
type ProviderAllocationsOpts struct {
	// Resources maps resource class names to the integer amount to consume.
	Resources map[string]int `json:"resources"`
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToAllocationUpdateMap() (map[string]any, error)
}

// UpdateOpts specifies the allocation to be set for a consumer.
//
// This requires microversion 1.28 or later. Write operations on allocations
// using earlier microversions are not safe in concurrent environments.
//
// ConsumerGeneration must be set to nil when creating allocations for a new
// consumer (it serializes as JSON null, which signals to the server that no
// prior allocation is expected). For an existing consumer, set it to the
// generation value returned by a prior Get call. A mismatch causes a 409
// Conflict response, allowing the caller to retry safely.
type UpdateOpts struct {
	// Allocations maps resource provider UUIDs to the resources to consume.
	Allocations map[string]ProviderAllocationsOpts `json:"allocations"`

	// Required from microversion 1.8.
	ProjectID string `json:"project_id"`

	// Required from microversion 1.8.
	UserID string `json:"user_id"`

	// ConsumerGeneration must be nil for new consumers (serializes as null) or
	// the current generation for existing consumers.
	// See the UpdateOpts type documentation for details.
	ConsumerGeneration *int `json:"consumer_generation"`

	// Required from microversion 1.38.
	ConsumerType string `json:"consumer_type,omitempty"`
}

// ToAllocationUpdateMap constructs a request body from UpdateOpts.
func (opts UpdateOpts) ToAllocationUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update replaces all allocations for a consumer. The operation is atomic.
//
// Requires microversion 1.28 or later.
func Update(ctx context.Context, client *gophercloud.ServiceClient, consumerUUID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete removes all allocations for a consumer. Returns 204 on success or
// 404 if the consumer does not exist.
//
// Note: using Update with an empty Allocations map is generally safer because
// it is protected by the consumer generation check, preventing accidental
// deletion under concurrent updates.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, consumerUUID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ManageOptsBuilder allows extensions to add additional parameters to the
// Manage request.
type ManageOptsBuilder interface {
	ToAllocationManageMap() (map[string]any, error)
}

// ManageOpts specifies allocations for multiple consumers in a single atomic
// operation. The map key is the consumer UUID; the value describes the
// allocations and consumer metadata for that consumer.
//
// This requires microversion 1.13 or later. For generation-safe writes, use
// microversion 1.28 or later and set ConsumerGeneration on each entry.
//
// Set ConsumerGeneration to nil for new consumers (serializes as JSON null).
// For existing consumers, set it to the generation returned by a prior Get
// call. A mismatch on any entry causes a 409 Conflict for the entire batch.
type ManageOpts map[string]UpdateOpts

// ToAllocationManageMap constructs a request body from ManageOpts.
func (opts ManageOpts) ToAllocationManageMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Manage atomically sets allocations for one or more consumers in a single
// request.
//
// Requires microversion 1.13 or later. For generation-safe writes, use
// microversion 1.28 or later.
func Manage(ctx context.Context, client *gophercloud.ServiceClient, opts ManageOptsBuilder) (r ManageResult) {
	_ = "STUB: not implemented"
	return *new(ManageResult)
}
