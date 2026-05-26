package queues

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the
// List request.
type ListOptsBuilder interface {
	ToQueueListQuery() (string, error)
}

// ListOpts params to be used with List
type ListOpts struct {
	// Limit instructs List to refrain from sending excessively large lists of queues
	Limit int `q:"limit,omitempty"`

	// Marker and Limit control paging. Marker instructs List where to start listing from.
	Marker string `q:"marker,omitempty"`

	// Specifies if showing the detailed information when querying queues
	Detailed bool `q:"detailed,omitempty"`

	// Specifies if filter the queues by queue’s name when querying queues.
	Name bool `q:"name,omitempty"`

	// Specifies if showing the amount of queues when querying them.
	WithCount bool `q:"with_count,omitempty"`
}

// ToQueueListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToQueueListQuery() (string, error) { _ = "STUB: not implemented"; return "", nil }

// List instructs OpenStack to provide a list of queues.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToQueueCreateMap() (map[string]any, error)
}

// CreateOpts specifies the queue creation parameters.
type CreateOpts struct {
	// The name of the queue to create.
	QueueName string `json:"queue_name" required:"true"`

	// The target incoming messages will be moved to when a message can’t
	// processed successfully after meet the max claim count is met.
	DeadLetterQueue string `json:"_dead_letter_queue,omitempty"`

	// The new TTL setting for messages when moved to dead letter queue.
	DeadLetterQueueMessagesTTL int `json:"_dead_letter_queue_messages_ttl,omitempty"`

	// The delay of messages defined for a queue. When the messages send to
	// the queue, it will be delayed for some times and means it can not be
	// claimed until the delay expired.
	DefaultMessageDelay int `json:"_default_message_delay,omitempty"`

	// The default TTL of messages defined for a queue, which will effect for
	// any messages posted to the queue.
	DefaultMessageTTL int `json:"_default_message_ttl" required:"true"`

	// The flavor name which can tell Zaqar which storage pool will be used
	// to create the queue.
	Flavor string `json:"_flavor,omitempty"`

	// The max number the message can be claimed.
	MaxClaimCount int `json:"_max_claim_count,omitempty"`

	// The max post size of messages defined for a queue, which will effect
	// for any messages posted to the queue.
	MaxMessagesPostSize int `json:"_max_messages_post_size,omitempty"`

	// Does messages should be encrypted
	EnableEncryptMessages *bool `json:"_enable_encrypt_messages,omitempty"`

	// Extra is free-form extra key/value pairs to describe the queue.
	Extra map[string]any `json:"-"`
}

// ToQueueCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToQueueCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new queue.
func Create(ctx context.Context, client *gophercloud.ServiceClient, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// update request.
type UpdateOptsBuilder interface {
	ToQueueUpdateMap() ([]map[string]any, error)
}

// BatchUpdateOpts is an array of UpdateOpts.
type BatchUpdateOpts []UpdateOpts

// UpdateOpts is the struct responsible for updating a property of a queue.
type UpdateOpts struct {
	Op    UpdateOp `json:"op" required:"true"`
	Path  string   `json:"path" required:"true"`
	Value any      `json:"value" required:"true"`
}

type UpdateOp string

const (
	ReplaceOp UpdateOp = "replace"
	AddOp     UpdateOp = "add"
	RemoveOp  UpdateOp = "remove"
)

// ToQueueUpdateMap constructs a request body from UpdateOpts.
func (opts BatchUpdateOpts) ToQueueUpdateMap() ([]map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToMap constructs a request body from UpdateOpts.
func (opts UpdateOpts) ToMap() (map[string]any, error) { _ = "STUB: not implemented"; return nil, nil }

// Update Updates the specified queue.
func Update(ctx context.Context, client *gophercloud.ServiceClient, queueName string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// Get requests details on a single queue, by name.
func Get(ctx context.Context, client *gophercloud.ServiceClient, queueName string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// Delete deletes the specified queue.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, queueName string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// GetStats returns statistics for the specified queue.
func GetStats(ctx context.Context, client *gophercloud.ServiceClient, queueName string) (r StatResult) {
	_ = "STUB: not implemented"
	return *new(StatResult)
}

type SharePath string

const (
	PathMessages      SharePath = "messages"
	PathClaims        SharePath = "claims"
	PathSubscriptions SharePath = "subscriptions"
)

type ShareMethod string

const (
	MethodGet   ShareMethod = "GET"
	MethodPatch ShareMethod = "PATCH"
	MethodPost  ShareMethod = "POST"
	MethodPut   ShareMethod = "PUT"
)

// ShareOpts specifies share creation parameters.
type ShareOpts struct {
	Paths   []SharePath   `json:"paths,omitempty"`
	Methods []ShareMethod `json:"methods,omitempty"`
	Expires string        `json:"expires,omitempty"`
}

// ShareOptsBuilder allows extensions to add additional attributes to the
// Share request.
type ShareOptsBuilder interface {
	ToQueueShareMap() (map[string]any, error)
}

// ToShareQueueMap formats a ShareOpts structure into a request body.
func (opts ShareOpts) ToQueueShareMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Share creates a pre-signed URL for a given queue.
func Share(ctx context.Context, client *gophercloud.ServiceClient, queueName string, opts ShareOptsBuilder) (r ShareResult) {
	_ = "STUB: not implemented"
	return *new(ShareResult)
}

type PurgeResource string

const (
	ResourceMessages      PurgeResource = "messages"
	ResourceSubscriptions PurgeResource = "subscriptions"
)

// PurgeOpts specifies the purge parameters.
type PurgeOpts struct {
	ResourceTypes []PurgeResource `json:"resource_types" required:"true"`
}

// PurgeOptsBuilder allows extensions to add additional attributes to the
// Purge request.
type PurgeOptsBuilder interface {
	ToQueuePurgeMap() (map[string]any, error)
}

// ToPurgeQueueMap formats a PurgeOpts structure into a request body
func (opts PurgeOpts) ToQueuePurgeMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Purge purges particular resource of the queue.
func Purge(ctx context.Context, client *gophercloud.ServiceClient, queueName string, opts PurgeOptsBuilder) (r PurgeResult) {
	_ = "STUB: not implemented"
	return *new(PurgeResult)
}
