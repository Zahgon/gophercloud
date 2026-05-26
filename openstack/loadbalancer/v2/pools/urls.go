package pools

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "lbaas"
	resourcePath = "pools"
	memberPath   = "members"
)

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func memberRootURL(c *gophercloud.ServiceClient, poolId string) string {
	_ = "STUB: not implemented"
	return ""
}

func memberResourceURL(c *gophercloud.ServiceClient, poolID string, memberID string) string {
	_ = "STUB: not implemented"
	return ""
}
