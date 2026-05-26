package imageimport

import (
	"context"

	"github.com/gophercloud/gophercloud/v2"
)

// ImportMethod represents valid Import API method.
type ImportMethod string

const (
	// GlanceDirectMethod represents glance-direct Import API method.
	GlanceDirectMethod ImportMethod = "glance-direct"

	// WebDownloadMethod represents web-download Import API method.
	WebDownloadMethod ImportMethod = "web-download"
)

// Get retrieves Import API information data.
func Get(ctx context.Context, c *gophercloud.ServiceClient) (r GetResult) {
	_ = "STUB: not implemented"
	return *new(GetResult)
}

// CreateOptsBuilder allows to add additional parameters to the Create request.
type CreateOptsBuilder interface {
	ToImportCreateMap() (map[string]any, error)
}

// CreateOpts specifies parameters of a new image import.
type CreateOpts struct {
	Name ImportMethod `json:"name"`
	URI  string       `json:"uri"`
}

// ToImportCreateMap constructs a request body from CreateOpts.
func (opts CreateOpts) ToImportCreateMap() (map[string]any, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create requests the creation of a new image import on the server.
func Create(ctx context.Context, client *gophercloud.ServiceClient, imageID string, opts CreateOptsBuilder) (r CreateResult) {
	_ = "STUB: not implemented"
	return *new(CreateResult)
}
