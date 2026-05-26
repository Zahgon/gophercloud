package amphorae

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "octavia"
	resourcePath = "amphorae"
	failoverPath = "failover"
)

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func failoverRootURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}
