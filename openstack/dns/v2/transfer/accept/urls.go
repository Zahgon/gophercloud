package accept

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "zones"
	tasksPath    = "tasks"
	resourcePath = "transfer_accepts"
)

func baseURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, transferAcceptID string) string {
	_ = "STUB: not implemented"
	return ""
}
