package monitors

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "lbaas"
	resourcePath = "healthmonitors"
)

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}
