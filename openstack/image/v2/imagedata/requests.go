package imagedata

import (
	"context"
	"io"

	"github.com/gophercloud/gophercloud/v2"
)

// Upload uploads an image file.
func Upload(ctx context.Context, client *gophercloud.ServiceClient, id string, data io.Reader) (r UploadResult) {
	_ = "STUB: not implemented"
	return *new(UploadResult)
}

// Stage performs PUT call on the existing image object in the Image service with
// the provided file.
// Existing image object must be in the "queued" status.
func Stage(ctx context.Context, client *gophercloud.ServiceClient, id string, data io.Reader) (r StageResult) {
	_ = "STUB: not implemented"
	return *new(StageResult)
}

// Download retrieves an image.
func Download(ctx context.Context, client *gophercloud.ServiceClient, id string) (r DownloadResult) {
	_ = "STUB: not implemented"
	return *new(DownloadResult)
}
