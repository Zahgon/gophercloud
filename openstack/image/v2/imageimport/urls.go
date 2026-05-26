package imageimport

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "images"
	infoPath     = "info"
	resourcePath = "import"
)

func infoURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func importURL(c *gophercloud.ServiceClient, imageID string) string {
	_ = "STUB: not implemented"
	return ""
}
