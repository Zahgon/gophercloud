package messages

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// Delete will delete the existing Message with the provided ID.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToMessageListQuery() (string, error)
}

// ListOpts holds options for listing Messages. It is passed to the
// messages.List function.
type ListOpts struct {
	// The message ID
	ID string `q:"id"`
	// The ID of the action during which the message was created
	ActionID string `q:"action_id"`
	// The ID of the message detail
	DetailID string `q:"detail_id"`
	// The message level
	MessageLevel string `q:"message_level"`
	// The UUID of the request during which the message was created
	RequestID string `q:"request_id"`
	// The UUID of the resource for which the message was created
	ResourceID string `q:"resource_id"`
	// The type of the resource for which the message was created
	ResourceType string `q:"resource_type"`
	// The key to sort a list of messages
	SortKey string `q:"sort_key"`
	// The key to sort a list of messages
	SortDir string `q:"sort_dir"`
	// The maximum number of messages to return
	Limit int `q:"limit"`
}

// ToMessageListQuery formats a ListOpts into a query string.
func (opts ListOpts) ToMessageListQuery() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List returns Messages optionally limited by the conditions provided in ListOpts.
func List(client *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// Get retrieves the Message with the provided ID. To extract the Message
// object from the response, call the Extract method on the GetResult.
func Get(ctx context.Context, client *gophercloud.ServiceClient, id string) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
