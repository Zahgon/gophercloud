package capsules

import "github.com/gophercloud/gophercloud/v2"

func getURL(client *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}

func createURL(client *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

// `listURL` is a pure function. `listURL(c)` is a URL for which a GET
// request will respond with a list of capsules in the service `c`.
func listURL(c *gophercloud.ServiceClient) string { _ = "STUB: not implemented"; return "" }

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	_ = "STUB: not implemented"
	return ""
}
