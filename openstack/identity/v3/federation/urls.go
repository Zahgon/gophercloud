package federation

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "OS-FEDERATION"
	mappingsPath = "mappings"
)

func mappingsRootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func mappingsResourceURL(c *gophercloud.ServiceClient, mappingID string) string {
	_ = "STUB: not implemented"
	return ""
}
