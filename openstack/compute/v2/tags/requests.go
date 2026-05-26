package tags

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// List all tags on a server.
func List(ctx context.Context, client *gophercloud.ServiceClient, serverID string) (r ListResult) {
	_ = "STUB: not implemented"
	return *new(ListResult)
}

// Check if a tag exists on a server.
func Check(ctx context.Context, client *gophercloud.ServiceClient, serverID, tag string) (r CheckResult) {
	_ = "STUB: not implemented"
	return *new(CheckResult)
}

// ReplaceAllOptsBuilder allows to add additional parameters to the ReplaceAll request.
type ReplaceAllOptsBuilder interface {
	ToTagsReplaceAllMap() (map[string]any, error)
}

// ReplaceAllOpts provides options used to replace Tags on a server.
type ReplaceAllOpts struct {
	Tags []string `json:"tags" required:"true"`
}

// ToTagsReplaceAllMap formats a ReplaceALlOpts into the body of the ReplaceAll request.
func (opts ReplaceAllOpts) ToTagsReplaceAllMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReplaceAll replaces all Tags on a server.
func ReplaceAll(ctx context.Context, client *gophercloud.ServiceClient, serverID string, opts ReplaceAllOptsBuilder) (r ReplaceAllResult) {
	_ = "STUB: not implemented"
	return *new(ReplaceAllResult)
}

// Add adds a new Tag on a server.
func Add(ctx context.Context, client *gophercloud.ServiceClient, serverID, tag string) (r AddResult) {
	_ = "STUB: not implemented"
	return *new(AddResult)
}

// Delete removes a tag from a server.
func Delete(ctx context.Context, client *gophercloud.ServiceClient, serverID, tag string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// DeleteAll removes all tag from a server.
func DeleteAll(ctx context.Context, client *gophercloud.ServiceClient, serverID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}
