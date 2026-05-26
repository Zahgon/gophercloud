package messages

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToMessageListQuery() (string, error)
}

// ListOpts params to be used with List.
type ListOpts struct {
	// Limit instructs List to refrain from sending excessively large lists of queues
	Limit int `q:"limit,omitempty"`

	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Marker string `q:"marker,omitempty"`

	// Indicate if the messages can be echoed back to the client that posted them.
	Echo bool `q:"echo,omitempty"`

	// Indicate if the messages list should include the claimed messages.
	IncludeClaimed bool `q:"include_claimed,omitempty"`

	//Indicate if the messages list should include the delayed messages.
	IncludeDelayed bool `q:"include_delayed,omitempty"`
}

func (opts ListOpts) ToMessageListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// ListMessages lists messages on a specific queue based off queue name.
func List(client *gophercloud.ServiceClient, queueName string, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder Builder.
type CreateOptsBuilder interface {
	ToMessageCreateMap() (map[string]any, error)
}

// BatchCreateOpts is an array of CreateOpts.
type BatchCreateOpts []CreateOpts

// CreateOpts params to be used with Create.
type CreateOpts struct {
	// TTL specifies how long the server waits before marking the message
	// as expired and removing it from the queue.
	TTL int `json:"ttl,omitempty"`

	// Delay specifies how long the message can be claimed.
	Delay int `json:"delay,omitempty"`

	// Body specifies an arbitrary document that constitutes the body of the message being sent.
	Body map[string]any `json:"body" required:"true"`
}

// ToMessageCreateMap constructs a request body from BatchCreateOpts.
func (opts BatchCreateOpts) ToMessageCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToMap constructs a request body from UpdateOpts.
func (opts CreateOpts) ToMap() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

// Create creates a message on a specific queue based of off queue name.
func Create(ctx context.Context, client *gophercloud.ServiceClient, queueName string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// DeleteMessagesOptsBuilder allows extensions to add additional parameters to the
// DeleteMessages request.
type DeleteMessagesOptsBuilder interface {
	ToMessagesDeleteQuery() (string, error)
}

// DeleteMessagesOpts params to be used with DeleteMessages.
type DeleteMessagesOpts struct {
	IDs []string `q:"ids,omitempty"`
}

// ToMessagesDeleteQuery formats a DeleteMessagesOpts structure into a query string.
func (opts DeleteMessagesOpts) ToMessagesDeleteQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// DeleteMessages deletes multiple messages based off of ID.
func DeleteMessages(ctx context.Context, client *gophercloud.ServiceClient, queueName string, opts DeleteMessagesOptsBuilder) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// PopMessagesOptsBuilder allows extensions to add additional parameters to the
// DeleteMessages request.
type PopMessagesOptsBuilder interface {
	ToMessagesPopQuery() (string, error)
}

// PopMessagesOpts params to be used with PopMessages.
type PopMessagesOpts struct {
	Pop int `q:"pop,omitempty"`
}

// ToMessagesPopQuery formats a PopMessagesOpts structure into a query string.
func (opts PopMessagesOpts) ToMessagesPopQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// PopMessages deletes and returns multiple messages based off of number of messages.
func PopMessages(ctx context.Context, client *gophercloud.ServiceClient, queueName string, opts PopMessagesOptsBuilder) (r PopResult) {
	_ = "STUB: not implemented"
	return *new(PopResult)
}

// GetMessagesOptsBuilder allows extensions to add additional parameters to the
// GetMessages request.
type GetMessagesOptsBuilder interface {
	ToGetMessagesListQuery() (string, error)
}

// GetMessagesOpts params to be used with GetMessages.
type GetMessagesOpts struct {
	IDs []string `q:"ids,omitempty"`
}

// ToGetMessagesListQuery formats a GetMessagesOpts structure into a query string.
func (opts GetMessagesOpts) ToGetMessagesListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetMessages requests details on a multiple messages, by IDs.
func GetMessages(ctx context.Context, client *gophercloud.ServiceClient, queueName string, opts GetMessagesOptsBuilder) (r GetMessagesResult) {
	_ = "STUB: not implemented"
	return *new(GetMessagesResult)
}

// Get requests details on a single message, by ID.
func Get(ctx context.Context, client *gophercloud.ServiceClient, queueName string, messageID string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// DeleteOptsBuilder allows extensions to add additional parameters to the
// delete request.
type DeleteOptsBuilder interface {
	ToMessageDeleteQuery() (string, error)
}

// DeleteOpts params to be used with Delete.
type DeleteOpts struct {
	// ClaimID instructs Delete to delete a message that is associated with a claim ID
	ClaimID string `q:"claim_id,omitempty"`
}

// ToMessageDeleteQuery formats a DeleteOpts structure into a query string.
func (opts DeleteOpts) ToMessageDeleteQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// Delete deletes a specific message from the queue.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, queueName string, messageID string, opts DeleteOptsBuilder) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
