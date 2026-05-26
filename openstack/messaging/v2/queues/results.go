package queues

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// commonResult is the response of a base result.
type commonResult struct {
	gophercloud.Result
}

// QueuePage contains a single page of all queues from a List operation.
type QueuePage struct {
	pagination.LinkedPageBase
}

// CreateResult is the response of a Create operation.
type CreateResult struct {
	gophercloud.ErrResult
}

// UpdateResult is the response of a Update operation.
type UpdateResult struct {
	commonResult
}

// GetResult is the response of a Get operation.
type GetResult struct {
	commonResult
}

// StatResult contains the result of a Share operation.
type StatResult struct {
	gophercloud.Result
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// ShareResult contains the result of a Share operation.
type ShareResult struct {
	gophercloud.Result
}

// PurgeResult is the response of a Purge operation.
type PurgeResult struct {
	gophercloud.ErrResult
}

// Queue represents a messaging queue.
type Queue struct {
	Href          string       `json:"href"`
	Methods       []string     `json:"methods"`
	Name          string       `json:"name"`
	Paths         []string     `json:"paths"`
	ResourceTypes []string     `json:"resource_types"`
	Metadata      QueueDetails `json:"metadata"`
}

// QueueDetails represents the metadata of a queue.
type QueueDetails struct {
	// The queue the message will be moved to when the message can’t
	// be processed successfully after the max claim count is met.
	DeadLetterQueue string `json:"_dead_letter_queue"`

	// The TTL setting for messages when moved to dead letter queue.
	DeadLetterQueueMessageTTL int `json:"_dead_letter_queue_messages_ttl"`

	// The delay of messages defined for the queue.
	DefaultMessageDelay int `json:"_default_message_delay"`

	// The default TTL of messages defined for the queue.
	DefaultMessageTTL int `json:"_default_message_ttl"`

	// Extra is a collection of miscellaneous key/values.
	Extra map[string]any `json:"-"`

	// The max number the message can be claimed from the queue.
	MaxClaimCount int `json:"_max_claim_count"`

	// The max post size of messages defined for the queue.
	MaxMessagesPostSize int `json:"_max_messages_post_size"`

	// Is message encryption enabled
	EnableEncryptMessages bool `json:"_enable_encrypt_messages"`

	// The flavor defined for the queue.
	Flavor string `json:"flavor"`
}

// Stats represents a stats response.
type Stats struct {
	// Number of Claimed messages for a queue
	Claimed int `json:"claimed"`

	// Total Messages for a queue
	Total int `json:"total"`

	// Number of free messages
	Free int `json:"free"`
}

// QueueShare represents a share response.
type QueueShare struct {
	Project   string   `json:"project"`
	Paths     []string `json:"paths"`
	Expires   string   `json:"expires"`
	Methods   []string `json:"methods"`
	Signature string   `json:"signature"`
}

// Extract interprets any commonResult as a Queue.
func (r commonResult) Extract() (QueueDetails, error) {
	_ = "STUB: not implemented"
	return *new(QueueDetails), nil
}

// Extract interprets any StatResult as a Stats.
func (r StatResult) Extract() (Stats, error) { _ = "STUB: not implemented"; return *new(Stats), nil }

// Extract interprets any ShareResult as a QueueShare.
func (r ShareResult) Extract() (QueueShare, error) {
	_ = "STUB: not implemented"
	return *new(QueueShare), nil
}

// ExtractQueues interprets the results of a single page from a
// List() call, producing a map of queues.
func ExtractQueues(r pagination.Page) ([]Queue, error) { _ = "STUB: not implemented"; return nil, nil }

// IsEmpty determines if a QueuesPage contains any results.
func (r QueuePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r QueuePage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetCount value if it request was supplied `WithCount` param
func (r QueuePage) GetCount() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (r *QueueDetails) UnmarshalJSON(b []byte) error { _ = "STUB: not implemented"; return nil }

// Collect other fields and bundle them into Extra
// but only if a field titled "extra" wasn't sent.
