package policies

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

const policyTypeMaxLength = 255

// ListOptsBuilder allows extensions to add additional parameters to
// the List request
type ListOptsBuilder interface {
	ToPolicyListQuery() (string, error)
}

// ListOpts provides options to filter the List results.
type ListOpts struct {
	// Type filters the response by MIME media type
	// of the serialized policy blob.
	Type string `q:"type"`

	// Filters filters the response by custom filters such as
	// 'type__contains=foo'
	Filters map[string]string `q:"-"`
}

// ToPolicyListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToPolicyListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List enumerates the policies to which the current token has access.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to
// the Create request.
type CreateOptsBuilder interface {
	ToPolicyCreateMap() (map[string]any, error)
}

// CreateOpts provides options used to create a policy.
type CreateOpts struct {
	// Type is the MIME media type of the serialized policy blob.
	Type string `json:"type" required:"true"`

	// Blob is the policy rule as a serialized blob.
	Blob []byte `json:"-" required:"true"`

	// Extra is free-form extra key/value pairs to describe the policy.
	Extra map[string]any `json:"-"`
}

// ToPolicyCreateMap formats a CreateOpts into a create request.
func (opts CreateOpts) ToPolicyCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create creates a new Policy.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get retrieves details on a single policy, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, policyID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to
// the Update request.
type UpdateOptsBuilder interface {
	ToPolicyUpdateMap() (map[string]any, error)
}

// UpdateOpts provides options for updating a policy.
type UpdateOpts struct {
	// Type is the MIME media type of the serialized policy blob.
	Type string `json:"type,omitempty"`

	// Blob is the policy rule as a serialized blob.
	Blob []byte `json:"-"`

	// Extra is free-form extra key/value pairs to describe the policy.
	Extra map[string]any `json:"-"`
}

// ToPolicyUpdateMap formats a UpdateOpts into an update request.
func (opts UpdateOpts) ToPolicyUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update updates an existing Role.
func Update(ctx context.Context, client *gophercloud.ServiceClient, policyID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete deletes a policy.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, policyID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
