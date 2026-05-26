package attributestags

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// ReplaceAllOptsBuilder allows extensions to add additional parameters to
// the ReplaceAll request.
type ReplaceAllOptsBuilder interface {
	ToAttributeTagsReplaceAllMap() (map[string]any, error)
}

// ReplaceAllOpts provides options used to create Tags on a Resource
type ReplaceAllOpts struct {
	Tags []string `json:"tags" required:"true"`
}

// ToAttributeTagsReplaceAllMap formats a ReplaceAllOpts into the body of the
// replace request
func (opts ReplaceAllOpts) ToAttributeTagsReplaceAllMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReplaceAll updates all tags on a resource, replacing any existing tags
func ReplaceAll(ctx context.Context, client *gophercloud.ServiceClient, resourceType string, resourceID string, opts ReplaceAllOptsBuilder) (r ReplaceAllResult) {
	_ = "STUB: not implemented"
	return *new(ReplaceAllResult)
}

// List all tags on a resource
func List(ctx context.Context, client *gophercloud.ServiceClient, resourceType string, resourceID string) (r ListResult) {
	_ = "STUB: not implemented"
	return *new(ListResult)
}

// DeleteAll deletes all tags on a resource
func DeleteAll(ctx context.Context, client *gophercloud.ServiceClient, resourceType string, resourceID string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Add a tag on a resource
func Add(ctx context.Context, client *gophercloud.ServiceClient, resourceType string, resourceID string, tag string) (r AddResult) {
	_ = "STUB: not implemented"
	return *new(AddResult)
}

// Delete a tag on a resource
func Delete(ctx context.Context, client *gophercloud.ServiceClient, resourceType string, resourceID string, tag string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// Confirm if a tag exists on a resource
func Confirm(ctx context.Context, client *gophercloud.ServiceClient, resourceType string, resourceID string, tag string) (r ConfirmResult) {
	_ = "STUB: not implemented"
	return *new(ConfirmResult)
}
