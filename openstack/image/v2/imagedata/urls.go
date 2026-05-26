package imagedata

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath   = "images"
	uploadPath = "file"
	stagePath  = "stage"
)

// `imageDataURL(c,i)` is the URL for the binary image data for the
// image identified by ID `i` in the service `c`.
func uploadURL(c *gophercloud.ServiceClient, imageID string) string {
	_ = "STUB: not implemented"
	return ""
}

func stageURL(c *gophercloud.ServiceClient, imageID string) string {
	_ = "STUB: not implemented"
	return ""
}

func downloadURL(c *gophercloud.ServiceClient, imageID string) string {
	_ = "STUB: not implemented"
	return ""
}
