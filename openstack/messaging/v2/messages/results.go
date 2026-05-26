package messages

import (
	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// CreateResult is the response of a Create operations.
type CreateResult struct {
	gophercloud.Result
}

// MessagePage contains a single page of all clusters from a ListDetails call.
type MessagePage struct {
	pagination.LinkedPageBase
}

// DeleteResult is the result from a Delete operation. Call its ExtractErr
// method to determine if the call succeeded or failed.
type DeleteResult struct {
	gophercloud.ErrResult
}

// CreateResult is the response of a Create operations.
type PopResult struct {
	gophercloud.Result
}

// GetMessagesResult is the response of a GetMessages operations.
type GetMessagesResult struct {
	gophercloud.Result
}

// GetResult is the response of a Get operations.
type GetResult struct {
	gophercloud.Result
}

// Message represents a message on a queue.
type Message struct {
	Body     map[string]any `json:"body"`
	Age      int            `json:"age"`
	Href     string         `json:"href"`
	ID       string         `json:"id"`
	TTL      int            `json:"ttl"`
	Checksum string         `json:"checksum"`
}

// PopMessage represents a message returned from PopMessages.
type PopMessage struct {
	Body       map[string]any `json:"body"`
	Age        int            `json:"age"`
	ID         string         `json:"id"`
	TTL        int            `json:"ttl"`
	ClaimCount int            `json:"claim_count"`
	ClaimID    string         `json:"claim_id"`
}

// ResourceList represents the result of creating a message.
type ResourceList struct {
	Resources []string `json:"resources"`
}

// Extract interprets any CreateResult as a ResourceList.
func (r CreateResult) Extract() (ResourceList, error) {
	_ = "STUB: not implemented"
	return *new(ResourceList), nil
}

// Extract interprets any PopResult as a list of PopMessage.
func (r PopResult) Extract() ([]PopMessage, error) { _ = "STUB: not implemented"; return nil, nil }

// Extract interprets any GetMessagesResult as a list of Message.
func (r GetMessagesResult) Extract() ([]Message, error) { _ = "STUB: not implemented"; return nil, nil }

// Extract interprets any GetResult as a Message.
func (r GetResult) Extract() (Message, error) { _ = "STUB: not implemented"; return *new(Message), nil }

// ExtractMessage extracts message into a  list of Message.
func ExtractMessages(r pagination.Page) ([]Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// IsEmpty determines if a MessagePage contains any results.
func (r MessagePage) IsEmpty() (bool, error) { _ = "STUB: not implemented"; return false, nil }

// NextPageURL uses the response's embedded link reference to navigate to the
// next page of results.
func (r MessagePage) NextPageURL(endpointURL string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
