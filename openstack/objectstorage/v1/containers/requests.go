package containers

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
	"github.com/gophercloud/gophercloud/v2/pagination"
)

// ListOptsBuilder allows extensions to add additional parameters to the List
// request.
type ListOptsBuilder interface {
	ToContainerListParams() (string, error)
}

// ListOpts is a structure that holds options for listing containers.
type ListOpts struct {
	// Full has been removed from the Gophercloud API. Gophercloud will now
	// always request the "full" (json) listing, because simplified listing
	// (plaintext) returns false results when names contain end-of-line
	// characters.

	Limit     int    `q:"limit"`
	Marker    string `q:"marker"`
	EndMarker string `q:"end_marker"`
	Format    string `q:"format"`
	Prefix    string `q:"prefix"`
	Delimiter string `q:"delimiter"`
}

// ToContainerListParams formats a ListOpts into a query string.
func (opts ListOpts) ToContainerListParams() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// List is a function that retrieves containers associated with the account as
// well as account metadata. It returns a pager which can be iterated with the
// EachPage function.
func List(c *gophercloud.ServiceClient, opts ListOptsBuilder) pagination.Pager {
	_ = "STUB: not implemented"
	return *new(pagination.Pager)
}

// CreateOptsBuilder allows extensions to add additional parameters to the
// Create request.
type CreateOptsBuilder interface {
	ToContainerCreateMap() (map[string]string, error)
}

// CreateOpts is a structure that holds parameters for creating a container.
type CreateOpts struct {
	Metadata          map[string]string
	ContainerRead     string `h:"X-Container-Read"`
	ContainerSyncTo   string `h:"X-Container-Sync-To"`
	ContainerSyncKey  string `h:"X-Container-Sync-Key"`
	ContainerWrite    string `h:"X-Container-Write"`
	ContentType       string `h:"Content-Type"`
	DetectContentType bool   `h:"X-Detect-Content-Type"`
	IfNoneMatch       string `h:"If-None-Match"`
	VersionsLocation  string `h:"X-Versions-Location"`
	HistoryLocation   string `h:"X-History-Location"`
	TempURLKey        string `h:"X-Container-Meta-Temp-URL-Key"`
	TempURLKey2       string `h:"X-Container-Meta-Temp-URL-Key-2"`
	StoragePolicy     string `h:"X-Storage-Policy"`
	VersionsEnabled   bool   `h:"X-Versions-Enabled"`
}

// ToContainerCreateMap formats a CreateOpts into a map of headers.
func (opts CreateOpts) ToContainerCreateMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create is a function that creates a new container.
func Create(ctx context.Context, c *gophercloud.ServiceClient, containerName string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}

// BulkDelete is a function that bulk deletes containers.
func BulkDelete(ctx context.Context, c *gophercloud.ServiceClient, containers []string) (r BulkDeleteResult) {
	_ = "STUB: not implemented"
	return *new(BulkDeleteResult)
}

// Delete is a function that deletes a container.
func Delete(ctx context.Context, c *gophercloud.ServiceClient, containerName string) (r DeleteResult) {
	_ = "STUB: not implemented"
	return *new(DeleteResult)
}

// UpdateOptsBuilder allows extensions to add additional parameters to the
// Update request.
type UpdateOptsBuilder interface {
	ToContainerUpdateMap() (map[string]string, error)
}

// UpdateOpts is a structure that holds parameters for updating, creating, or
// deleting a container's metadata.
type UpdateOpts struct {
	Metadata               map[string]string
	RemoveMetadata         []string
	ContainerRead          *string `h:"X-Container-Read"`
	ContainerSyncTo        *string `h:"X-Container-Sync-To"`
	ContainerSyncKey       *string `h:"X-Container-Sync-Key"`
	ContainerWrite         *string `h:"X-Container-Write"`
	ContentType            *string `h:"Content-Type"`
	DetectContentType      *bool   `h:"X-Detect-Content-Type"`
	RemoveVersionsLocation string  `h:"X-Remove-Versions-Location"`
	VersionsLocation       string  `h:"X-Versions-Location"`
	RemoveHistoryLocation  string  `h:"X-Remove-History-Location"`
	HistoryLocation        string  `h:"X-History-Location"`
	TempURLKey             string  `h:"X-Container-Meta-Temp-URL-Key"`
	TempURLKey2            string  `h:"X-Container-Meta-Temp-URL-Key-2"`
	VersionsEnabled        *bool   `h:"X-Versions-Enabled"`
}

// ToContainerUpdateMap formats a UpdateOpts into a map of headers.
func (opts UpdateOpts) ToContainerUpdateMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Update is a function that creates, updates, or deletes a container's
// metadata.
func Update(ctx context.Context, c *gophercloud.ServiceClient, containerName string, opts UpdateOptsBuilder) (r UpdateResult) {
	_ = "STUB: not implemented"
	return *new(UpdateResult)
}

// GetOptsBuilder allows extensions to add additional parameters to the Get
// request.
type GetOptsBuilder interface {
	ToContainerGetMap() (map[string]string, error)
}

// GetOpts is a structure that holds options for listing containers.
type GetOpts struct {
	Newest bool `h:"X-Newest"`
}

// ToContainerGetMap formats a GetOpts into a map of headers.
func (opts GetOpts) ToContainerGetMap() (map[string]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Get is a function that retrieves the metadata of a container. To extract just
// the custom metadata, pass the GetResult response to the ExtractMetadata
// function.
func Get(ctx context.Context, c *gophercloud.ServiceClient, containerName string, opts GetOptsBuilder) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}
