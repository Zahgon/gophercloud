package loadbalancers

import "github.com/gophercloud/gophercloud/v2"

const (
	rootPath       = "lbaas"
	resourcePath   = "loadbalancers"
	statusPath     = "status"
	statisticsPath = "stats"
	failoverPath   = "failover"
)

func rootURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func statusRootURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func statisticsRootURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func failoverRootURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}
