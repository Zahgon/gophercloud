package accounts

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// GetOptsBuilder allows extensions to add additional headers to the Get
// request.
type GetOptsBuilder interface {
	ToAccountGetMap() (map[string]string, error)
}

// GetOpts is a structure that contains parameters for getting an account's
// metadata.
type GetOpts struct {
	Newest bool `h:"X-Newest"`
}

// ToAccountGetMap formats a GetOpts into a map[string]string of headers.
func (opts GetOpts) ToAccountGetMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get is a function that retrieves an account's metadata. To extract just the
// custom metadata, call the ExtractMetadata method on the GetResult. To extract
// all the headers that are returned (including the metadata), call the
// Extract method on the GetResult.
func Get(ctx context.Context, c *gophercloud.ServiceClient, opts GetOptsBuilder) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// UpdateOptsBuilder allows extensions to add additional headers to the Update
// request.
type UpdateOptsBuilder interface {
	ToAccountUpdateMap() (map[string]string, error)
}

// UpdateOpts is a structure that contains parameters for updating, creating, or
// deleting an account's metadata.
type UpdateOpts struct {
	Metadata          map[string]string
	RemoveMetadata    []string
	ContentType       *string `h:"Content-Type"`
	DetectContentType *bool   `h:"X-Detect-Content-Type"`
	TempURLKey        string  `h:"X-Account-Meta-Temp-URL-Key"`
	TempURLKey2       string  `h:"X-Account-Meta-Temp-URL-Key-2"`
}

// ToAccountUpdateMap formats an UpdateOpts into a map[string]string of headers.
func (opts UpdateOpts) ToAccountUpdateMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update is a function that creates, updates, or deletes an account's metadata.
// To extract the headers returned, call the Extract method on the UpdateResult.
func Update(ctx context.Context, c *gophercloud.ServiceClient, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}
