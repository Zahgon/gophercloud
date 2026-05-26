package claims

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToClaimCreateRequest() (map[string]any, string, error)
}

// CreateOpts params to be used with Create.
type CreateOpts struct {
	// Sets the TTL for the claim. When the claim expires un-deleted messages will be able to be claimed again.
	TTL int `json:"ttl,omitempty"`

	// Sets the Grace period for the claimed messages. The server extends the lifetime of claimed messages
	// to be at least as long as the lifetime of the claim itself, plus the specified grace period.
	Grace int `json:"grace,omitempty"`

	// Set the limit of messages returned by create.
	Limit int `q:"limit" json:"-"`
}

// ToClaimCreateRequest assembles a body and URL for a Create request based on
// the contents of a CreateOpts.
func (opts CreateOpts) ToClaimCreateRequest() (map[string]any, string, error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

// Create creates a Claim that claims messages on a specified queue.
func Create(ctx context.Context, client *gophercloud.ServiceClient, queueName string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// Get queries the specified claim for the specified queue.
func Get(ctx context.Context, client *gophercloud.ServiceClient, queueName string, claimID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToClaimUpdateMap() (map[string]any, error)
}

// UpdateOpts implements UpdateOpts.
type UpdateOpts struct {
	// Update the TTL for the specified Claim.
	TTL int `json:"ttl,omitempty"`

	// Update the grace period for Messages in a specified Claim.
	Grace int `json:"grace,omitempty"`
}

// ToClaimUpdateMap assembles a request body based on the contents of
// UpdateOpts.
func (opts UpdateOpts) ToClaimUpdateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update will update the options for a specified claim.
func Update(ctx context.Context, client *gophercloud.ServiceClient, queueName string, claimID string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Delete will delete a Claim for a specified Queue.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, queueName string, claimID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
