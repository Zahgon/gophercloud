package policies

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "fwaas"
	resourcePath = "firewall_policies"
	insertPath   = "insert_rule"
	removePath   = "remove_rule"
)

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func insertURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func removeURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}
