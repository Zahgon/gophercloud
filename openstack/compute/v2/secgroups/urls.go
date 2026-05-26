package secgroups

import "github.com/gophercloud/gophercloud/v2"

const (
	secgrouppath = "os-security-groups"
	rulepath     = "os-security-group-rules"
)

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func listByServerURL(c *gophercloud.ServiceClient, serverID string) string {
	_ = "STUB: not implemented"
	return ""
}

func rootRuleURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceRuleURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func serverActionURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}
