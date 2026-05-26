package tags

import "github.com/gophercloud/gophercloud/v2"

const (
	rootResourcePath = "servers"
	resourcePath     = "tags"
)

func rootURL(c *gophercloud.ServiceClient, serverID string) string {
	_ = "STUB: not implemented"
	return ""
}

func resourceURL(c *gophercloud.ServiceClient, serverID, tag string) string {
	_ = "STUB: not implemented"
	return ""
}

func listURL(c *gophercloud.ServiceClient, serverID string) string {
	_ = "STUB: not implemented"
	return ""
}

func checkURL(c *gophercloud.ServiceClient, serverID, tag string) string {
	_ = "STUB: not implemented"
	return ""
}

func replaceAllURL(c *gophercloud.ServiceClient, serverID string) string {
	_ = "STUB: not implemented"
	return ""
}

func addURL(c *gophercloud.ServiceClient, serverID, tag string) string {
	_ = "STUB: not implemented"
	return ""
}

func deleteURL(c *gophercloud.ServiceClient, serverID, tag string) string {
	_ = "STUB: not implemented"
	return ""
}

func deleteAllURL(c *gophercloud.ServiceClient, serverID string) string {
	_ = "STUB: not implemented"
	return ""
}
