package l7policies

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath     = "lbaas"
	resourcePath = "l7policies"
	rulePath     = "rules"
)

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func ruleRootURL(c *gophercloud.ServiceClient, policyID string) string {
	_ = "STUB: not implemented"
	return ""
}

func ruleResourceURL(c *gophercloud.ServiceClient, policyID string, ruleID string) string {
	_ = "STUB: not implemented"
	return ""
}
